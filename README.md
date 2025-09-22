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

rtag 支持中文和英文双语界面：

### 语言检测优先级
1. `RTAG_LANG` 环境变量（`en` 或 `zh`）
2. `LANG` 环境变量（包含 `zh` 则使用中文）
3. `LC_ALL` 环境变量（包含 `zh` 则使用中文）
4. 默认使用英文

### 使用示例
```bash
# 临时使用中文
RTAG_LANG=zh rtag --help

# 临时使用英文
RTAG_LANG=en rtag --help

# 永久设置（添加到 ~/.zshrc 或 ~/.bashrc）
export RTAG_LANG=zh
```

## 注意事项

- 确保在 Git 仓库中运行此工具
- 推送标签前请确保有推送权限
- 标签名不能重复
- 删除标签只会从 `.rtag` 文件中删除，不会删除已推送的 Git 标签
- 语言设置通过环境变量控制，每次命令执行时生效