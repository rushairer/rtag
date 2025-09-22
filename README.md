# rtag - Release Tag 管理工具

rtag 是一个基于 Cobra 的 Go 命令行工具，用于管理项目的发布标签。

## 功能特性

- 管理本地 `.rtag` 文件中的标签
- 交互式添加标签
- 批量或单个推送 Git 标签
- 列出所有标签
- 删除标签
- **🌍 国际化支持**：支持中文和英文双语界面

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
rtag
```
- 如果 `.rtag` 文件不存在或为空，会提示用户交互式添加标签
- 如果文件存在，会显示当前所有标签

#### 2. 添加标签
```bash
# 交互式添加标签
./rtag add

# 直接添加指定标签
./rtag add api
./rtag add cron
./rtag add debug
```

#### 3. 列出所有标签
```bash
./rtag list
```

#### 4. 推送标签
```bash
# 推送所有标签
./rtag push --all

# 推送指定标签
./rtag push api
```

#### 5. 删除标签
```bash
./rtag rm api
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

rtag 支持中文和英文双语界面，提供多种语言设置方式：

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