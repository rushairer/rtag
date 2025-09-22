# rtag - Release Tag Management Tool

[![Go Version](https://img.shields.io/github/go-mod/go-version/rushairer/rtag)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/rushairer/rtag)](https://github.com/rushairer/rtag/releases)
[![License](https://img.shields.io/github/license/rushairer/rtag)](https://github.com/rushairer/rtag/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rushairer/rtag)](https://goreportcard.com/report/github.com/rushairer/rtag)
[![GitHub Issues](https://img.shields.io/github/issues/rushairer/rtag)](https://github.com/rushairer/rtag/issues)
[![GitHub Stars](https://img.shields.io/github/stars/rushairer/rtag)](https://github.com/rushairer/rtag/stargazers)

[中文文档](README_zh.md) | English

rtag is a Cobra-based Go command-line tool for managing project release tags.

## Features

- Manage tags in local `.rtag` file
- Interactive tag addition
- Batch or individual Git tag pushing
- List all tags
- Delete tags
- **🌍 Internationalization Support**: Multi-language interface supporting English, Chinese, French, and Russian

## Installation

### Using go install
```bash
go install github.com/rushairer/rtag@latest
```

Make sure `$HOME/go/bin` is in your `PATH` environment variable.

## Usage

### Basic Commands

#### 1. Initial Run
```bash
rtag
```
- If `.rtag` file doesn't exist or is empty, prompts user for interactive tag addition
- If file exists, displays all current tags

#### 2. Add Tags
```bash
# Interactive tag addition
./rtag add

# Add specific tags directly
./rtag add api
./rtag add cron
./rtag add debug
```

#### 3. List All Tags
```bash
./rtag list
```

#### 4. Push Tags
```bash
# Push all tags
./rtag push --all

# Push specific tag
./rtag push api
```

#### 5. Delete Tags
```bash
./rtag rm api
```

## .rtag File Format

The `.rtag` file contains one tag name per line, for example:
```
api
cron
debug
```

## Git Tag Format

When pushing, creates Git tags in format `release-YYYYMMDDHHMM-{tag}`, for example:
- `release-202409221900-api`
- `release-202409221900-cron`
- `release-202409221900-debug`

## Example Workflow

1. Initialize project tags:
```bash
rtag init
# Interactively add api, cron, debug
```

2. Add new tag:
```bash
rtag add web
```

3. View all tags:
```bash
rtag list
```

4. Push all tags:
```bash
rtag push --all
```

5. Push single tag:
```bash
rtag push api
```

6. Delete unwanted tag:
```bash
rtag rm web
```

7. Language settings:
```bash
# View current language setting
rtag lang

# Set to English
rtag lang en

# Set to Chinese
rtag lang zh
```

## Internationalization Support

rtag supports multi-language interfaces (English, Chinese, French, Russian) with multiple language setting methods:

### 🌍 Smart Language Detection Priority
1. **`RTAG_LANG` Environment Variable** (`en` or `zh`) - **Highest Priority**
2. **Saved User Preference** (`~/.config/rtag/config`)
3. **System Language Auto-detection**:
   - Chinese systems (`zh_CN`, `zh_TW`, `zh_HK`, etc.) → Automatically use Chinese interface
   - Other language systems (`en`, `fr`, `de`, `ja`, `ko`, etc.) → Automatically use English interface
4. **Default to English** (fallback option)

### Persistent Language Settings
```bash
# View current language setting
rtag lang

# Permanently set to Chinese (saved to config file)
rtag lang zh

# Permanently set to English (saved to config file)
rtag lang en
```

### Temporary Language Settings
```bash
# Temporarily use Chinese (effective for current command only)
RTAG_LANG=zh rtag --help

# Temporarily use English (effective for current command only)
RTAG_LANG=en rtag --help

# Permanent environment variable setting (add to ~/.zshrc or ~/.bashrc)
export RTAG_LANG=zh
```

### System Language Examples
```bash
# Chinese system automatically uses Chinese interface
LANG=zh_CN.UTF-8 rtag --help  # Shows Chinese

# English system automatically uses English interface  
LANG=en_US.UTF-8 rtag --help  # Shows English

# Other language systems default to English interface
LANG=ja_JP.UTF-8 rtag --help  # Shows English
```

### Config File Location
- Linux/macOS: `~/.config/rtag/config`
- Windows: `%USERPROFILE%\.rtag\config`

## Notes

- Ensure running this tool in a Git repository
- Ensure push permissions before pushing tags
- Tag names cannot be duplicated
- Deleting tags only removes from `.rtag` file, doesn't delete pushed Git tags

### 🌍 Internationalization Features
- **Smart Detection**: Automatically selects appropriate interface language based on system language on first use
- **Persistent Storage**: Language preferences are automatically saved, set once and effective permanently
- **Environment Variable Override**: `RTAG_LANG` has highest priority, can temporarily override any setting
- **Cross-platform Support**: Supports Linux, macOS, Windows system language detection

## 🌐 Contributing New Language Support

Welcome to add new language support for rtag! Currently supported:
- 🇺🇸 English (en)
- 🇨🇳 中文 (zh)
- 🇫🇷 Français (fr)
- 🇷🇺 Русский (ru)

### Adding a new language requires only 3 steps:

#### Step 1: Add language definition in `languages.go`

```go
// Add new language in constants
const (
    LangEN Language = "en"
    LangZH Language = "zh"
    LangFR Language = "fr"
    LangRU Language = "ru"
    LangJA Language = "ja"  // Add Japanese
)

// Add language info in GetSupportedLanguages() function
LangJA: {
    Code:       LangJA,
    Name:       "Japanese",
    NativeName: "日本語",
    LocaleKeys: []string{"ja", "japanese", "jp"},
},
```

#### Step 2: Add translation function in `messages.go`

```go
// Create new translation function
func getJapaneseMessages() Messages {
    return Messages{
        RootShort: "リリースタグ管理ツール",
        RootLong:  "rtag は git 統合でリリースタグを管理するコマンドラインツールです。",
        // ... Add Japanese translations for all messages
    }
}
```

#### Step 3: Update `GetAllMessages()` function

```go
func GetAllMessages() map[Language]Messages {
    return map[Language]Messages{
        LangEN: getEnglishMessages(),
        LangZH: getChineseMessages(),
        LangFR: getFrenchMessages(),
        LangRU: getRussianMessages(),
        LangJA: getJapaneseMessages(),  // Add new language mapping
    }
}
```

### 🎯 Translation Guidelines

1. **Maintain Consistency**: Ensure terminology translations are consistent throughout the application
2. **Localization Adaptation**: Consider cultural and usage habits of the target language
3. **Format Strings**: Be careful to preserve format placeholders like `%s`, `%v`
4. **Testing**: Test all commands after adding translations to ensure proper functionality

### 📝 Contributing

1. Fork this project
2. Create new branch: `git checkout -b add-language-xx`
3. Follow the 3 steps above to add new language support
4. Test new language: `RTAG_LANG=xx rtag --help`
5. Submit Pull Request

### 🔍 Message Types to Translate

- **Command Descriptions**: Short and detailed descriptions of subcommands
- **Flag Descriptions**: Command-line parameter descriptions
- **User Messages**: Error messages, success messages, interactive prompts, etc.
- **Language Commands**: Language setting related prompts

Each language requires translation of approximately **40+ messages**. For the complete message list, refer to the `Messages` struct in `messages.go`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.