# OpenCode 汉化项目 - AI 维护指南

> 本文档为 AI 助手（如 Claude Code、Cursor 等）提供维护此汉化项目的完整指南。

[![Version](https://img.shields.io/badge/i18n-v7.0-green.svg)](../opencode-i18n/config.json)
[![Coverage](https://img.shields.io/badge/汉化覆盖率-100%25-brightgreen.svg)]()

---

## 📋 项目概述

| 项目 | 说明 |
|------|------|
| **项目名称** | OpenCode 中文汉化版 |
| **原项目** | https://github.com/opencode-ai/opencode |
| **汉化仓库** | https://github.com/Teznew/OpenCodeChineseTranslation |
| **管理工具** | `opencodenpm` (npm 包) |
| **当前版本** | v7.0 |

---

## 📂 目录结构

```
OpenCodeChineseTranslation/
├── scripts/                     # 管理脚本目录
│   ├── commands/                # 命令模块
│   │   ├── update.js            # 更新源码
│   │   ├── apply.js             # 应用汉化
│   │   ├── build.js             # 编译构建
│   │   ├── verify.js            # 验证汉化
│   │   ├── full.js              # 完整工作流
│   │   ├── launch.js            # 启动程序
│   │   ├── helper.js            # 智谱助手
│   │   ├── package.js           # 打包发布
│   │   └── deploy.js            # 部署命令
│   ├── core/                    # 核心模块
│   │   ├── cli.js               # CLI 入口
│   │   ├── menu.js              # 交互菜单
│   │   ├── utils.js             # 工具函数
│   │   ├── git.js               # Git 操作
│   │   ├── i18n.js              # 汉化应用核心 ⭐
│   │   ├── build.js             # 编译逻辑
│   │   ├── env.js               # 环境检查
│   │   ├── colors.js            # 输出样式
│   │   └── version.js           # 版本检测
│   ├── bin/                     # CLI 入口
│   │   └── opencodenpm          # 命令行工具
│   └── package.json             # 依赖配置
├── opencode-i18n/               # 汉化配置目录 ⭐
│   ├── config.json              # 主配置文件（版本、模块列表）
│   ├── dialogs/                 # 对话框翻译配置 (20个)
│   ├── routes/                  # 路由翻译配置 (6个)
│   ├── components/              # 组件翻译配置 (6个)
│   ├── common/                  # 通用翻译配置 (6个)
│   └── app.json                 # 应用根配置
├── opencode-zh-CN/              # OpenCode 源码（自动克隆）
├── bin/                         # 编译输出目录
├── releases/                    # 打包发布目录
└── docs/                        # 项目文档
    ├── AI_MAINTENANCE.md        # 本文档
    └── ANTIGRAVITY_INTEGRATION.md  # Antigravity 集成指南
```

---

## 🚀 快速开始

### 1. 安装管理工具

```bash
# 全局安装
npm install -g opencodenpm

# 或从本地安装
cd scripts && npm install -g .
```

### 2. 检查编译环境

```bash
opencodenpm env
```

**环境要求：**

| 工具 | 版本要求 | 说明 |
|------|----------|------|
| Node.js | >= 18.0.0 | JavaScript 运行时 |
| Bun | >= 1.3.8 | 快速 JavaScript 运行时 (需与上游 OpenCode 匹配) |
| Git | latest | 版本控制 |

### 3. 完整工作流

```bash
# 交互式菜单（推荐）
opencodenpm

# 或直接执行完整流程
opencodenpm full
```

---

## 🛠️ opencodenpm 命令参考

| 命令 | 别名 | 说明 |
|------|------|------|
| `opencodenpm` | `ui` | 交互式菜单 |
| `opencodenpm update` | - | 更新 OpenCode 源码 |
| `opencodenpm apply` | - | 应用汉化配置 |
| `opencodenpm build` | - | 编译构建 OpenCode |
| `opencodenpm verify` | - | 验证汉化覆盖率 |
| `opencodenpm full` | - | 完整工作流（更新→汉化→编译） |
| `opencodenpm launch` | `start` | 启动已编译的 OpenCode |
| `opencodenpm package` | `pack` | 打包 Releases |
| `opencodenpm deploy` | - | 部署全局命令 |
| `opencodenpm helper` | - | 智谱助手 |
| `opencodenpm env` | - | 检查编译环境 |
| `opencodenpm config` | - | 显示当前配置 |

### 常用命令示例

```bash
# 更新源码
opencodenpm update              # 更新到最新版本
opencodenpm update --force      # 强制重新克隆

# 应用汉化
opencodenpm apply               # 应用汉化配置
opencodenpm apply --silent      # 静默模式

# 编译构建
opencodenpm build               # 编译当前平台
opencodenpm build -p linux-x64  # 编译指定平台
opencodenpm build --no-deploy   # 不部署到 bin 目录

# 验证汉化
opencodenpm verify              # 验证汉化
opencodenpm verify -d           # 显示详细信息

# 打包发布
opencodenpm package -p windows-x64   # 打包指定平台
opencodenpm package -a               # 打包所有平台
```

---

## 🔧 汉化配置详解

### 配置文件结构

主配置文件 `opencode-i18n/config.json`:

```json
{
  "version": "6.0",
  "description": "OpenCode 中文汉化配置文件（模块化结构）",
  "lastUpdate": "2026-01-16",
  "testPassRate": "100%",
  "supportedCommit": "99a1e73fa1bd5c92c02abd8a20b0e274d5b0d214",
  "maintainer": {
    "name": "CodeCreator",
    "github": "https://github.com/Teznew/OpenCodeChineseTranslation"
  },
  "modules": {
    "dialogs": ["dialogs/dialog-agent.json", ...],
    "routes": ["routes/route-footer.json", ...],
    "components": ["components/autocomplete.json", ...],
    "common": ["common/app-messages.json", ...],
    "root": ["app.json"]
  }
}
```

### 翻译配置文件格式

每个翻译配置文件格式如下：

```json
{
  "file": "src/cli/cmd/tui/dialogs/xxx.tsx",
  "description": "文件描述",
  "note": "翻译注意事项",
  "replacements": {
    "Original Text": "翻译文本",
    "Another Text": "另一个翻译"
  }
}
```

### 模块分类

| 模块 | 目录 | 文件数 | 说明 |
|------|------|--------|------|
| **dialogs** | `dialogs/` | 20 | 对话框组件翻译 |
| **routes** | `routes/` | 6 | 路由页面翻译 |
| **components** | `components/` | 6 | UI 组件翻译 |
| **common** | `common/` | 6 | 通用文本翻译 |
| **root** | `/` | 1 | 应用根配置 |

---

## 📝 翻译规范

### 命名规范

| 类型 | 文件名格式 | 示例 |
|------|------------|------|
| 对话框 | `dialog-{name}.json` | `dialog-status.json` |
| 路由 | `route-{name}.json` | `route-sidebar.json` |
| 组件 | `component-{name}.json` | `component-question.json` |
| 通用 | `{category}-{name}.json` | `app-messages.json` |

### 翻译原则

1. **只翻译用户可见文本**
   - ✅ UI 文本、按钮、提示信息
   - ❌ 函数名、变量名、类型名
   - ❌ 日志输出（除非面向用户）

2. **保持技术术语一致性**

   | 英文 | 中文 |
   |------|------|
   | MCP Server | MCP 服务器 |
   | LSP Server | LSP 服务器 |
   | Plugin | 插件 |
   | Formatter | 格式化器 |
   | Session | 会话 |
   | Agent | 智能体 |
   | Provider | 提供商 |
   | Model | 模型 |
   | Context | 上下文 |
   | Prompt | 提示词 |

3. **匹配完整上下文**
   - 包含必要的 HTML/JSX 标签
   - 示例: `</text>` 而非单独的 `text`

---

## 🔄 更新流程

### 场景一：OpenCode 发布了新版本

```bash
# 1. 更新源码
opencodenpm update

# 2. 应用汉化
opencodenpm apply

# 3. 验证结果
opencodenpm verify

# 4. 如有失败，检查并修复配置
opencodenpm verify -d  # 查看详细信息

# 5. 编译测试
opencodenpm build
opencodenpm launch
```

### 场景二：新增/修改翻译配置

1. **编辑配置文件**
   ```bash
   # 位置: opencode-i18n/ 下对应目录
   # 格式: JSON
   ```

2. **测试配置**
   ```bash
   opencodenpm apply
   opencodenpm verify
   ```

3. **更新版本号**
   ```bash
   # 编辑 opencode-i18n/config.json
   # 更新 version 和 lastUpdate
   ```

4. **提交更改**
   ```bash
   git add opencode-i18n/
   git commit -m "chore(i18n): 更新汉化配置到 vX.X"
   git push
   ```

---

## 🐛 常见问题排查

| 问题 | 原因 | 解决方案 |
|------|------|----------|
| `[原文不存在]` | 源文件已更新，模式不匹配 | 检查源文件，更新翻译配置 |
| `验证失败` | 配置模式与源文件不符 | `opencodenpm verify -d` 查看详情 |
| `路径错误` | 源码路径配置错误 | 检查配置文件中的 `file` 字段 |
| `编译失败` | 环境问题 | `opencodenpm env` 检查环境 |
| `汉化未生效` | 未应用汉化 | `opencodenpm apply` 重新应用 |

---

## 🤖 自动化构建 (CI/CD)

本项目使用 GitHub Actions 实现自动化构建和发布。

### 工作流概览

| 工作流 | 文件 | 触发条件 | 说明 |
|--------|------|----------|------|
| **Release** | `release.yml` | Tag 推送 / 手动触发 | 正式版发布，经过测试的稳定版本 |
| **Nightly** | `nightly.yml` | 每日定时 / 手动触发 | 每日构建，跟进上游最新代码 |

### Nightly Build (自动跟进构建)

Nightly Build 会**每小时**自动检查上游仓库更新，当累计有 **≥5 个新 commit** 时自动触发构建。

**工作原理：**

1. **检查上游更新**（每小时第 0 分钟执行）
   - 获取 `anomalyco/opencode` 的 `dev` 分支最新 commit
   - 与 `.nightly-state` 文件中记录的上次构建 commit 对比
   - 计算新增 commit 数量

2. **触发条件**
   - 累计新 commit 数量 ≥ 5 时触发构建
   - 或手动触发时指定 `force_build=true`
   - 首次构建（无 `.nightly-state` 文件）时直接触发

3. **构建流程**
   - 编译 Go CLI 工具（三平台）
   - 克隆上游源码并应用汉化
   - 编译 OpenCode（三平台）
   - 生成包含上游更新日志的 Release Notes
   - 打包并发布到 `nightly` tag

4. **版本号策略**
   - 文件名格式: `opencode-zh-CN-nightly-{platform}.zip`
   - 使用固定的 `nightly` tag，每次构建覆盖更新
   - 下载链接始终指向最新构建
   - 标记为 `prerelease`，与正式版区分

**Release Notes 内容：**
- 构建信息（上游分支、commit SHA、新增 commit 数量、构建时间）
- 下载链接表格
- **OpenCode 官方更新日志**（自动抓取上游 git log）
- 自动更新机制说明

**手动触发：**

```bash
# 通过 GitHub CLI 触发 Nightly 构建
gh workflow run nightly.yml

# 强制构建（跳过 commit 数量检测）
gh workflow run nightly.yml -f force_build=true

# 自定义阈值（例如累计 3 个 commit 就触发）
gh workflow run nightly.yml -f min_commits=3
```

**状态文件：**

`.nightly-state` 文件记录上次构建的上游 commit SHA，用于增量检测：

```
a1b2c3d4e5f6...  # 上次构建的 commit SHA
```

每次成功构建后，Actions 会自动更新此文件并提交到仓库。

### Release (正式发布)

正式版发布流程请参阅下方"发布流程"章节。

**触发方式：**

```bash
# 方式一：通过 release.ps1 脚本（推荐）
.\release.ps1 -Version 8.4.0 -Message "新功能说明"

# 方式二：手动触发 Actions
gh workflow run release.yml -f tag_name=v8.4.0

# 方式三：推送 Tag
git tag v8.4.0
git push origin v8.4.0
```

### Nightly vs Release 对比

| 特性 | Nightly | Release |
|------|---------|---------|
| 触发频率 | 每日自动 | 手动触发 |
| 版本号 | `nightly-YYYYMMDD` | `v8.x.x` |
| 稳定性 | 可能不稳定 | 经过测试 |
| 推荐用户 | 开发者/测试者 | 普通用户 |
| Tag 类型 | 滚动覆盖 | 永久保留 |
| prerelease | ✅ 是 | ❌ 否 |

---

## 📦 发布流程

### 1. 更新版本信息

编辑 `opencode-i18n/config.json`:

```json
{
  "version": "7.1",
  "lastUpdate": "2026-01-18",
  "supportedCommit": "新的 commit hash"
}
```

### 2. 完整测试

```bash
# 完整工作流
opencodenpm full

# 验证
opencodenpm verify

# 测试运行
opencodenpm launch
```

### 3. 打包发布

```bash
# 打包所有平台
opencodenpm package -a

# 发布到 releases/ 目录
```

### 4. 提交发布

```bash
git add .
git commit -m "release(i18n): v7.1 - 更新说明"
git tag v7.1
git push && git push --tags
```

---

## 🔗 相关资源

| 链接 | 说明 |
|------|------|
| [OpenCode 官方](https://github.com/opencode-ai/opencode) | 原项目仓库 |
| [汉化项目 GitHub](https://github.com/Teznew/OpenCodeChineseTranslation) | 本项目 |
| [汉化项目 Gitee](https://gitee.com/QtCodeCreators/OpenCodeChineseTranslation) | 国内镜像 |
| [Antigravity 集成](./ANTIGRAVITY_INTEGRATION.md) | 本地 AI 网关配置 |
| [问题反馈](https://github.com/Teznew/OpenCodeChineseTranslation/issues) | 提交 Issue |

---

## 📊 汉化覆盖统计

| 模块 | 文件数 | 覆盖内容 | 状态 |
|------|--------|----------|------|
| dialogs | 20 | 所有对话框组件 | ✅ 100% |
| routes | 6 | 页面路由文本 | ✅ 100% |
| components | 6 | UI 组件文本 | ✅ 100% |
| common | 6 | 通用提示信息 | ✅ 100% |
| **总计** | **39** | **全部模块** | ✅ **100%** |

---

> **最后更新**: 2026-01-18
> **维护者**: CodeCreator
> **汉化版本**: v7.0
