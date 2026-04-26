# OpenCode 汉化管理工具 (Go 版本)

这是 OpenCode 汉化管理工具的 Go 语言重写版本，提供与 Node.js 版本相同的功能，但无需任何运行时依赖。

## ✨ 特性

- **无依赖运行** - 单文件二进制，无需 Node.js/npm
- **跨平台支持** - Windows x64、Linux x64、macOS (Intel/Apple Silicon)
- **完整功能** - 复刻原 JS 版本的所有命令
- **交互式菜单** - 美观的 TUI 界面，支持鼠标操作
- **自动更新** - 支持从 GitHub Releases 自动下载和更新

## 📦 安装

### 一键安装（推荐）

**Windows PowerShell:**
```powershell
powershell -c "irm https://gh-proxy.com/https://raw.githubusercontent.com/Teznew/OpenCodeChineseTranslation/refs/heads/main/install.ps1 | iex"
```

**macOS / Linux:**
```bash
curl -fsSL https://gh-proxy.com/https://raw.githubusercontent.com/Teznew/OpenCodeChineseTranslation/refs/heads/main/install.sh | bash
```

默认使用 `gh-proxy.com`。如需切换，可设置环境变量 `OPENCODE_GITHUB_PROXY`，或在 Unix 安装命令后追加 `bash -s -- --no-proxy`；CLI 下载命令也支持 `opencode-cli download --proxy=...` 与 `opencode-cli download --no-proxy`。

### 手动安装

从 GitHub Releases 页面下载对应平台的二进制文件，解压后将其放入系统 PATH 中即可。

推荐按以下步骤操作：

1. 下载与你平台匹配的 `opencode-cli` 二进制文件。
2. 将文件移动到：
   - Windows：`%USERPROFILE%\.opencode-i18n\bin\opencode-cli.exe`
   - macOS / Linux：`~/.opencode-i18n/bin/opencode-cli`
3. macOS / Linux 执行：
   ```bash
   chmod +x ~/.opencode-i18n/bin/opencode-cli
   ```
4. 把目录加入 PATH：
   - Windows（PowerShell）：
     ```powershell
     [Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", "User") + ";$env:USERPROFILE\.opencode-i18n\bin", "User")
     ```
   - macOS / Linux：
     ```bash
     echo 'export PATH="$HOME/.opencode-i18n/bin:$PATH"' >> ~/.bashrc
     ```
5. 重新打开终端后验证：
   ```bash
   opencode-cli --help
   ```

如果你还想手动安装预编译 `opencode` 本体，请额外执行：

1. 下载与你平台匹配的 ZIP 包，例如：
   - `opencode-zh-CN-<version>-windows-x64.zip`
   - `opencode-zh-CN-<version>-darwin-arm64.zip`
   - `opencode-zh-CN-<version>-darwin-x64.zip`
   - `opencode-zh-CN-<version>-linux-x64.zip`
2. 解压后取出 `opencode` 或 `opencode.exe`。
3. 将文件移动到：
   - Windows：`%USERPROFILE%\.opencode-i18n\bin\opencode.exe`
   - macOS / Linux：`~/.opencode-i18n/bin/opencode`
4. macOS / Linux 执行：
   ```bash
   chmod +x ~/.opencode-i18n/bin/opencode
   ```
5. 重新打开终端后验证：
   ```bash
   opencode --version
   ```

## 🚀 使用

### 交互式菜单（推荐）

```bash
opencode-cli interactive
```

或者直接运行 `opencode-cli` (如果没有其他参数，默认进入交互式菜单)。

### 命令行模式

```bash
# 查看帮助
opencode-cli --help

# 更新源码
opencode-cli update

# 应用汉化 (自动备份)
opencode-cli apply

# 验证配置
opencode-cli verify --detailed

# 编译构建
opencode-cli build

# 打包发布
opencode-cli package

# 部署到系统 (同时配置 opencode-cli 和 opencode)
opencode-cli deploy --shortcut

# 回滚备份
opencode-cli rollback --list
opencode-cli rollback [backup-id]

# 配置 Antigravity
opencode-cli antigravity

# 安装 Oh-My-OpenCode
opencode-cli ohmyopencode

# 校准 Bun 版本
opencode-cli fix-bun
```

## 📋 可用命令

| 命令 | 说明 |
|------|------|
| `interactive` | 启动交互式菜单 (默认) |
| `update` | 更新 OpenCode 源码 |
| `apply` | 应用汉化配置到源码 |
| `verify` | 验证汉化配置完整性 |
| `build` | 编译构建 OpenCode |
| `package` | 打包 Windows x64 / Linux x64 / macOS Intel/Apple Silicon 发布版 |
| `deploy` | 部署到系统 PATH，可选创建桌面快捷方式 |
| `rollback` | 回滚到之前的备份 |
| `antigravity` | 配置 Antigravity AI 代理 |
| `ohmyopencode` | 安装 Oh-My-OpenCode 插件 |
| `helper` | 安装智谱编码助手 |
| `fix-bun` | 校准 Bun 版本 |

## ⌨️ 快捷键

交互式菜单支持以下快捷键：

| 按键 | 功能 |
|------|------|
| ↑↓←→ | 导航菜单 |
| h/j/k/l | Vim 风格导航 |
| Enter | 确认选择 |
| 1-9 | 快速选择前 9 项 |
| Tab | 切换教程标签页 |
| Q/Esc | 退出 |
| 鼠标点击 | 选择菜单项 |
| 滚轮 | 上下移动 |

## 📁 项目结构

```
cli-go/
├── main.go                    # 入口
├── go.mod                     # Go 模块定义
├── build.ps1                  # Windows 编译脚本
├── build.sh                   # Unix 编译脚本
├── cmd/                       # Cobra 命令
│   ├── root.go               # 根命令
│   ├── menu.go               # 交互式菜单
│   ├── apply.go              # 应用汉化
│   ├── build.go              # 编译构建
│   ├── update.go             # 更新源码
│   ├── verify.go             # 验证配置
│   ├── package.go            # 打包发布
│   ├── deploy.go             # 部署命令
│   ├── rollback.go           # 回滚备份
│   ├── antigravity.go        # Antigravity 配置
│   └── extras.go             # 其他命令
└── internal/
    ├── core/                 # 核心逻辑
    │   ├── i18n.go          # 汉化处理
    │   ├── git.go           # Git 操作
    │   ├── build.go         # 编译逻辑
    │   ├── backup.go        # 备份管理
    │   ├── utils.go         # 工具函数
    │   └── version.go       # 版本管理
    └── tui/                  # TUI 界面
        ├── menu.go          # BubbleTea 菜单
        ├── items.go         # 菜单项数据
        └── theme.go         # 主题样式
```

## 📄 许可证

MIT License
