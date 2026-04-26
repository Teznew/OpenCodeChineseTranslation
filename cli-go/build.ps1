# OpenCode CLI 多目标编译脚本 (PowerShell)

$ErrorActionPreference = "Stop"

$APP_NAME = "opencode-cli"
$VERSION = "8.6.1"
$OUTPUT_DIR = "dist"

# 创建输出目录
New-Item -ItemType Directory -Force -Path $OUTPUT_DIR | Out-Null

Write-Host "`n📦 构建 $APP_NAME v$VERSION" -ForegroundColor Cyan
Write-Host ""

function Build {
    param (
        [string]$GOOS,
        [string]$GOARCH,
        [string]$EXT
    )
    
    $OUTPUT = "$OUTPUT_DIR/$APP_NAME-$GOOS-$GOARCH$EXT"
    Write-Host "  → 构建 $GOOS/$GOARCH..." -ForegroundColor Yellow
    
    $env:GOOS = $GOOS
    $env:GOARCH = $GOARCH
    
    go build -ldflags="-s -w" -o $OUTPUT .
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "    ✓ $OUTPUT" -ForegroundColor Green
    } else {
        Write-Host "    ✗ 构建失败" -ForegroundColor Red
    }
}

# Windows x64
Build -GOOS "windows" -GOARCH "amd64" -EXT ".exe"

# macOS
Build -GOOS "darwin" -GOARCH "amd64" -EXT ""
Build -GOOS "darwin" -GOARCH "arm64" -EXT ""

# Linux x64
Build -GOOS "linux" -GOARCH "amd64" -EXT ""

# 清理环境变量
Remove-Item Env:\GOOS -ErrorAction SilentlyContinue
Remove-Item Env:\GOARCH -ErrorAction SilentlyContinue

Write-Host "`n✓ 构建完成!" -ForegroundColor Green
Write-Host ""

# 显示构建产物
Get-ChildItem $OUTPUT_DIR | Format-Table Name, @{Label="Size"; Expression={"{0:N0} KB" -f ($_.Length / 1KB)}}
