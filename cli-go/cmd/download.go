package cmd

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"opencode-cli/internal/core"

	"github.com/spf13/cobra"
)

const (
	GiteaRepo = "Mirror/OpenCodeChineseTranslation"
)

var giteaHosts = []string{
	"http://192.168.2.7:23000",
	"http://10.1.1.7:23000",
	"http://10.10.10.7:23000",
	"https://gitea.re-v0.com",
}

func getGiteaCredentials() (username, token string, ok bool) {
	username = os.Getenv("GITEA_USER")
	token = os.Getenv("GITEA_TOKEN")
	return username, token, username != "" && token != ""
}

func buildGiteaAPIURL(host string) string {
	return host + "/api/v1/repos/" + GiteaRepo + "/releases/latest"
}

func buildGiteaReleaseURL(host, tag string) string {
	return host + "/" + GiteaRepo + "/releases/tag/" + tag
}

func buildGiteaReleasesURL(host string) string {
	return host + "/" + GiteaRepo + "/releases"
}

func buildAssetDownloadURL(host, tag, name string) string {
	return host + "/" + GiteaRepo + "/releases/download/" + tag + "/" + name
}

type GiteaRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
	Assets  []struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
	} `json:"assets"`
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "下载预编译汉化版 (无需编译环境)",
	Long:  "Download prebuilt OpenCode Chinese version from Gitea mirror (no compilation required)",
	Run: func(cmd *cobra.Command, args []string) {
		runDownload()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func runDownload() {
	fmt.Println("")
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Println("  下载预编译版 OpenCode 汉化版")
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Println("")
	fmt.Println("  无需本地编译环境，直接从 Gitea 镜像下载")
	fmt.Println("  适用于无法安装 Bun/Node.js 或想快速体验的用户")
	fmt.Println("")

	username, token, hasCreds := getGiteaCredentials()
	if !hasCreds {
		fmt.Println("  ⚠ 未检测到 Gitea 认证信息 (GITEA_USER / GITEA_TOKEN)")
		fmt.Println("  将尝试匿名访问，部分节点可能需要认证才能下载")
		fmt.Println("")
		fmt.Println("  配置方法 (添加到 ~/.bashrc 或 ~/.zshrc):")
		fmt.Println("    export GITEA_USER=<你的用户名>")
		fmt.Println("    export GITEA_TOKEN=<你的访问令牌>")
		fmt.Println("")
		fmt.Println("  获取令牌: Gitea → 设置 → 应用 → 管理 Access Token")
		fmt.Println("")
		fmt.Printf("  示例 (带认证下载):\n")
		fmt.Printf("    curl -u '$GITEA_USER:$GITEA_TOKEN' -LO %s/%s/releases/download/nightly/opencode-zh-CN-linux-x64.zip\n",
			giteaHosts[len(giteaHosts)-1], GiteaRepo)
		fmt.Println("")
	} else {
		fmt.Printf("  认证用户: %s\n", username)
		fmt.Println("")
	}

	fmt.Println("▶ 正在获取最新版本信息...")

	var release *GiteaRelease
	var activeHost string
	for _, host := range giteaHosts {
		r, err := getLatestRelease(host, username, token)
		if err != nil {
			fmt.Printf("  ✗ %s — %v\n", host, err)
			continue
		}
		release = r
		activeHost = host
		fmt.Printf("  ✓ 已连接: %s\n", host)
		break
	}

	if release == nil {
		fmt.Println("")
		fmt.Println("✗ 所有 Gitea 节点均无法连接")
		fmt.Println("")
		fmt.Println("  Gitea 节点（按优先级）:")
		for _, h := range giteaHosts {
			fmt.Printf("    - %s\n", h)
		}
		fmt.Println("")
		fmt.Println("  请检查:")
		fmt.Println("    1. 网络是否可达以上任一节点")
		fmt.Println("    2. GITEA_USER / GITEA_TOKEN 是否正确")
		fmt.Println("    3. 令牌是否有仓库读取权限")
		return
	}

	fmt.Printf("✓ 最新版本: %s\n", release.TagName)
	fmt.Println("")

	platform := core.DetectPlatform()

	var fileSize int64
	var assetName string

	for _, asset := range release.Assets {
		name := asset.Name
		if strings.HasPrefix(name, "opencode-zh-CN") &&
			strings.HasSuffix(name, ".zip") &&
			strings.Contains(name, platform) {
			fileSize = asset.Size
			assetName = name
			break
		}
	}

	if assetName == "" {
		fmt.Printf("✗ 未找到适用于当前平台的预编译包: %s\n", platform)
		fmt.Println("")
		fmt.Println("  可用的预编译包:")
		for _, asset := range release.Assets {
			if strings.HasSuffix(asset.Name, ".zip") && strings.HasPrefix(asset.Name, "opencode-") {
				fmt.Printf("    - %s\n", asset.Name)
			}
		}
		fmt.Println("")
		fmt.Printf("  手动下载页面: %s\n", buildGiteaReleaseURL(activeHost, release.TagName))
		return
	}

	fmt.Printf("  平台: %s\n", platform)
	fmt.Printf("  文件: %s\n", assetName)
	fmt.Printf("  大小: %.2f MB\n", float64(fileSize)/(1024*1024))
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("是否下载并安装? [Y/n]: ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer == "n" || answer == "no" {
		fmt.Println("操作已取消")
		return
	}

	tempDir := filepath.Join(os.TempDir(), "opencode-download")
	os.RemoveAll(tempDir)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fmt.Printf("✗ 创建临时目录失败: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	zipPath := filepath.Join(tempDir, assetName)

	fmt.Println("")
	fmt.Println("▶ 正在下载...")

	downloaded := false
	for _, host := range giteaHosts {
		url := buildAssetDownloadURL(host, release.TagName, assetName)
		fmt.Printf("  尝试: %s\n", host)
		if err := downloadFileWithAuth(url, zipPath, username, token); err != nil {
			fmt.Printf("  ✗ 失败: %v\n", err)
			continue
		}
		downloaded = true
		break
	}

	if !downloaded {
		fmt.Println("")
		fmt.Println("✗ 所有节点下载均失败")
		fmt.Println("")
		fmt.Println("  可能的解决方案:")
		fmt.Println("    1. 检查网络连接")
		fmt.Println("    2. 确认 GITEA_TOKEN 有下载权限")
		fmt.Printf("    3. 手动下载页面: %s\n", buildGiteaReleasesURL(activeHost))
		return
	}

	fmt.Println("✓ 下载完成")

	fmt.Println("")
	fmt.Println("▶ 正在解压...")

	extractDir := filepath.Join(tempDir, "extracted")
	if err := unzip(zipPath, extractDir); err != nil {
		fmt.Printf("✗ 解压失败: %v\n", err)
		return
	}

	fmt.Println("✓ 解压完成")

	exeName := "opencode"
	if runtime.GOOS == "windows" {
		exeName = "opencode.exe"
	}

	var exePath string
	filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && info.Name() == exeName {
			exePath = path
			return filepath.SkipAll
		}
		return nil
	})

	if exePath == "" {
		fmt.Printf("✗ 未在压缩包中找到 %s\n", exeName)
		return
	}

	fmt.Println("")
	fmt.Println("▶ 正在部署...")

	binDir, err := getDeployDir()
	if err != nil {
		fmt.Printf("✗ 获取部署目录失败: %v\n", err)
		return
	}

	if err := os.MkdirAll(binDir, 0755); err != nil {
		fmt.Printf("✗ 创建目录失败: %v\n", err)
		return
	}

	targetPath := filepath.Join(binDir, exeName)

	if err := copyFileWithProgress(exePath, targetPath); err != nil {
		fmt.Printf("✗ 复制文件失败: %v\n", err)
		return
	}

	if runtime.GOOS != "windows" {
		os.Chmod(targetPath, 0755)
	}

	fmt.Printf("✓ 已部署到: %s\n", targetPath)

	fmt.Println("")
	fmt.Println("▶ 正在配置系统 PATH...")

	if err := configurePathForDownload(binDir); err != nil {
		fmt.Printf("⚠ PATH 配置失败: %v\n", err)
		fmt.Println("  请手动将以下目录添加到 PATH:")
		fmt.Printf("    %s\n", binDir)
	} else {
		fmt.Println("✓ PATH 配置完成")
	}

	fmt.Println("")
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Println("  ✓ OpenCode 汉化版安装完成!")
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Println("")
	fmt.Printf("  版本: %s\n", release.TagName)
	fmt.Printf("  位置: %s\n", targetPath)
	fmt.Println("")
	fmt.Println("  下一步:")
	fmt.Println("    1. 重新打开终端使 PATH 生效")
	fmt.Println("    2. 运行 'opencode' 启动程序")
	fmt.Println("    3. 输入 /connect 配置 AI 模型")
}

func setAuthIfPresent(req *http.Request, username, token string) {
	if username != "" && token != "" {
		req.SetBasicAuth(username, token)
	}
}

func getLatestRelease(host, username, token string) (*GiteaRelease, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", buildGiteaAPIURL(host), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "OpenCode-CLI")
	setAuthIfPresent(req, username, token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	var release GiteaRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func downloadFileWithAuth(url, dest, username, token string) error {
	client := &http.Client{Timeout: 5 * time.Minute}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	setAuthIfPresent(req, username, token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	totalSize := resp.ContentLength
	downloaded := int64(0)
	buf := make([]byte, 32*1024)

	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			out.Write(buf[:n])
			downloaded += int64(n)

			if totalSize > 0 {
				percent := float64(downloaded) / float64(totalSize) * 100
				fmt.Printf("\r  进度: %.1f%% (%.2f / %.2f MB)", percent,
					float64(downloaded)/(1024*1024),
					float64(totalSize)/(1024*1024))
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	fmt.Println()
	return nil
}

// unzip 解压 ZIP 文件
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		// 安全检查：防止 Zip Slip 漏洞
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("非法文件路径: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}

		// 确保二进制文件有执行权限 (Unix)
		// ZIP 可能在 Windows 上创建，权限信息可能丢失
		if runtime.GOOS != "windows" {
			name := strings.ToLower(f.Name)
			if strings.Contains(name, "opencode") && !strings.HasSuffix(name, ".json") && !strings.HasSuffix(name, ".txt") {
				os.Chmod(fpath, 0755)
			}
		}
	}

	return nil
}

// copyFileWithProgress 复制文件
func copyFileWithProgress(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// configurePathForDownload 配置 PATH 环境变量
func configurePathForDownload(binDir string) error {
	switch runtime.GOOS {
	case "windows":
		// 检查是否已在 PATH 中
		currentPath := os.Getenv("PATH")
		if strings.Contains(strings.ToLower(currentPath), strings.ToLower(binDir)) {
			return nil // 已在 PATH 中
		}

		// 使用 PowerShell 添加到用户 PATH
		script := fmt.Sprintf(`
$userPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
if (-not $userPath.ToLower().Contains('%s'.ToLower())) {
    $newPath = '%s;' + $userPath
    [Environment]::SetEnvironmentVariable('PATH', $newPath, 'User')
}
`, binDir, binDir)

		return core.ExecLive("powershell", "-NoProfile", "-Command", script)

	default:
		// Unix: 提示用户手动配置
		homeDir, _ := os.UserHomeDir()
		shellRC := filepath.Join(homeDir, ".bashrc")
		if _, err := os.Stat(filepath.Join(homeDir, ".zshrc")); err == nil {
			shellRC = filepath.Join(homeDir, ".zshrc")
		}

		// 检查是否已配置
		if data, err := os.ReadFile(shellRC); err == nil {
			if strings.Contains(string(data), binDir) {
				return nil
			}
		}

		// 追加到配置文件
		exportLine := fmt.Sprintf("\n# OpenCode\nexport PATH=\"%s:$PATH\"\n", binDir)
		f, err := os.OpenFile(shellRC, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(exportLine)
		return err
	}
}
