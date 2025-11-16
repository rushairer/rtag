# rtag - Release Tag 管理工具

[![Go Version](https://img.shields.io/github/go-mod/go-version/rushairer/rtag)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/rushairer/rtag)](https://github.com/rushairer/rtag/releases)
[![License](https://img.shields.io/github/license/rushairer/rtag)](https://github.com/rushairer/rtag/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rushairer/rtag)](https://goreportcard.com/report/github.com/rushairer/rtag)
[![GitHub Issues](https://img.shields.io/github/issues/rushairer/rtag)](https://github.com/rushairer/rtag/issues)
[![GitHub Stars](https://img.shields.io/github/stars/rushairer/rtag)](https://github.com/rushairer/rtag/stargazers)

中文文档 | [English](README.md)

rtag 是一个基于 Cobra 的 Go 命令行工具，用于管理项目的发布标签。

## 功能特性

- 管理本地 `.rtag` 文件中的标签
- 交互式添加标签
- 批量或单个推送 Git 标签
- 列出所有标签
- 删除标签
- **🌍 国际化支持**：支持英文、中文、法文、俄文多语言界面

## 安装

### 使用 go install
```bash
go install github.com/rushairer/rtag@latest
```

确保 `$HOME/go/bin` 在你的 `PATH` 环境变量中。

## 使用方法

### 基本命令

#### 1. 初始运行
```bash
rtag init
```
- 如果 `.rtag` 文件不存在或为空，会提示用户交互式添加标签
- 如果文件存在，会显示当前所有标签

#### 2. 添加标签
```bash
# 交互式添加标签
rtag add

# 直接添加指定标签
rtag add api
rtag add cron
rtag add debug
```

#### 3. 列出所有标签
```bash
rtag list
```

#### 4. 推送标签
```bash
# 推送所有标签
rtag push --all

# 推送指定标签
rtag push api
```

#### 5. 删除标签
```bash
rtag rm api
```

## .rtag 文件格式

`.rtag` 文件每行包含一个标签名，例如：
```
api
cron
debug
```

## Git 标签格式

推送时会创建格式为 `release-YYYYMMDDHHMM-{tag}` 的 Git 标签，例如：
- `release-202409221900-api`
- `release-202409221900-cron`
- `release-202409221900-debug`

## 示例工作流

1. 初始化项目标签：
```bash
rtag init
# 交互式添加 api, cron, debug
```

2. 添加新标签：
```bash
rtag add web
```

3. 查看所有标签：
```bash
rtag list
```

4. 推送所有标签：
```bash
rtag push --all
```

5. 推送单个标签：
```bash
rtag push api
```

6. 删除不需要的标签：
```bash
rtag rm web
```

7. 语言设置：
```bash
# 查看当前语言设置
rtag lang

# 设置为英文
rtag lang en

# 设置为中文
rtag lang zh
```

## 国际化支持

rtag 支持多语言界面（英文、中文、法文、俄文），提供多种语言设置方式：

### 🌍 智能语言检测优先级
1. **`RTAG_LANG` 环境变量**（`en` 或 `zh`）- **最高优先级**
2. **保存的用户偏好设置**（`~/.config/rtag/config`）
3. **系统语言自动检测**：
   - 中文系统（`zh_CN`, `zh_TW`, `zh_HK` 等）→ 自动使用中文界面
   - 其他语言系统（`en`, `fr`, `de`, `ja`, `ko` 等）→ 自动使用英文界面
4. **默认使用英文**（后备选项）

### 持久化语言设置
```bash
# 查看当前语言设置
rtag lang

# 永久设置为中文（保存到配置文件）
rtag lang zh

# 永久设置为英文（保存到配置文件）
rtag lang en
```

### 临时语言设置
```bash
# 临时使用中文（仅当次命令有效）
RTAG_LANG=zh rtag --help

# 临时使用英文（仅当次命令有效）
RTAG_LANG=en rtag --help

# 永久环境变量设置（添加到 ~/.zshrc 或 ~/.bashrc）
export RTAG_LANG=zh
```

### 系统语言示例
```bash
# 中文系统自动使用中文界面
LANG=zh_CN.UTF-8 rtag --help  # 显示中文

# 英文系统自动使用英文界面  
LANG=en_US.UTF-8 rtag --help  # 显示英文

# 其他语言系统默认使用英文界面
LANG=ja_JP.UTF-8 rtag --help  # 显示英文
```

### 配置文件位置
- Linux/macOS: `~/.config/rtag/config`
- Windows: `%USERPROFILE%\.rtag\config`

## 注意事项

- 确保在 Git 仓库中运行此工具
- 推送标签前请确保有推送权限
- 标签名不能重复
- 删除标签只会从 `.rtag` 文件中删除，不会删除已推送的 Git 标签

### 🌍 国际化特性
- **智能检测**：首次使用时根据系统语言自动选择合适的界面语言
- **持久化保存**：语言偏好自动保存，设置一次永久生效
- **环境变量覆盖**：`RTAG_LANG` 具有最高优先级，可临时覆盖任何设置
- **跨平台支持**：支持 Linux、macOS、Windows 系统语言检测

## 🌐 贡献新语言支持

欢迎为 rtag 添加新语言支持！目前已支持：
- 🇺🇸 English (en)
- 🇨🇳 中文 (zh)
- 🇫🇷 Français (fr)
- 🇷🇺 Русский (ru)

### 添加新语言只需 3 个步骤：

#### 步骤 1：在 `languages.go` 中添加语言定义

```go
// 在常量定义中添加新语言
const (
    LangEN Language = "en"
    LangZH Language = "zh"
    LangFR Language = "fr"
    LangRU Language = "ru"
    LangJA Language = "ja"  // 新增日语
)

// 在 GetSupportedLanguages() 函数中添加语言信息
LangJA: {
    Code:       LangJA,
    Name:       "Japanese",
    NativeName: "日本語",
    LocaleKeys: []string{"ja", "japanese", "jp"},
},
```

#### 步骤 2：在 `messages.go` 中添加翻译函数

```go
// 创建新的翻译函数
func getJapaneseMessages() Messages {
    return Messages{
        RootShort: "リリースタグ管理ツール",
        RootLong:  "rtag は git 統合でリリースタグを管理するコマンドラインツールです。",
        // ... 添加所有消息的日语翻译
    }
}
```

#### 步骤 3：更新 `GetAllMessages()` 函数

```go
func GetAllMessages() map[Language]Messages {
    return map[Language]Messages{
        LangEN: getEnglishMessages(),
        LangZH: getChineseMessages(),
        LangFR: getFrenchMessages(),
        LangRU: getRussianMessages(),
        LangJA: getJapaneseMessages(),  // 添加新语言映射
    }
}
```

### 🎯 翻译指南

1. **保持一致性**：确保术语翻译在整个应用中保持一致
2. **本地化适配**：考虑目标语言的文化和使用习惯
3. **格式化字符串**：注意保留 `%s`、`%v` 等格式化占位符
4. **测试验证**：添加翻译后请测试所有命令确保正常工作

### 📝 提交贡献

1. Fork 本项目
2. 创建新分支：`git checkout -b add-language-xx`
3. 按照上述 3 个步骤添加新语言支持
4. 测试新语言：`RTAG_LANG=xx rtag --help`
5. 提交 Pull Request

### 🔍 需要翻译的消息类型

- **命令描述**：各个子命令的简短和详细描述
- **标志说明**：命令行参数的说明文字
- **用户消息**：错误提示、成功消息、交互提示等
- **语言命令**：语言设置相关的提示信息

每种语言大约需要翻译 **40+ 条消息**，完整的消息列表请参考 `messages.go` 中的 `Messages` 结构体。

## 许可证

本项目采用 MIT 许可证 - 详情请查看 [LICENSE](LICENSE) 文件。

## 贡献

欢迎贡献！请随时提交 Pull Request。
