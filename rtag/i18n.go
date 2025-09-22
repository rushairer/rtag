package rtag

import (
	"os"
	"strings"
)

// Language represents supported languages
type Language string

const (
	LangEN Language = "en"
	LangZH Language = "zh"
)

// Messages contains all translatable strings
type Messages struct {
	// Command descriptions
	RootShort string
	RootLong  string

	InitShort string
	InitLong  string

	AddShort string
	AddLong  string

	PushShort string
	PushLong  string

	ListShort string
	ListLong  string

	RmShort string
	RmLong  string

	// Flag descriptions
	PushAllFlag string

	// User messages
	ErrorReadingRtagFile    string
	RtagFileEmptyOrNotExist string
	CurrentTags             string
	AddTagFailed            string
	AddTagSuccess           string
	ReadTagsFailed          string
	NoTagsFound             string
	TagNotExistInFile       string
	SpecifyTagOrUseAll      string
	AllTags                 string
	RemoveTagFailed         string
	RemoveTagSuccess        string
	EnterTag                string
	Exit                    string
	TagCannotBeEmpty        string
	ContinueAdding          string
	ReadInputFailed         string
	StartPushingTags        string
	CreateGitTag            string
	CreateTagFailed         string
	PushingTagsToRemote     string
	PushTagsFailed          string
	PushTagsSuccess         string
	TagAlreadyExists        string
	TagNotExist             string
}

var currentLang Language = LangEN
var messages Messages

// English messages
var enMessages = Messages{
	RootShort: "A CLI tool for managing release tags",
	RootLong:  "rtag is a command line tool for managing release tags with git integration.",

	InitShort: "Initialize .rtag file",
	InitLong:  "Initialize .rtag file for storing published tags.",

	AddShort: "Add a new tag",
	AddLong:  "Add a new tag to the .rtag file. If no tag is provided, interactive mode will be used.",

	PushShort: "Push tags to remote repository",
	PushLong:  "Push tags to remote repository. Use --all flag to push all tags.",

	ListShort: "List all tags",
	ListLong:  "List all tags from the .rtag file.",

	RmShort: "Remove a tag",
	RmLong:  "Remove a tag from the .rtag file.",

	PushAllFlag: "Push all tags",

	ErrorReadingRtagFile:    "Error reading .rtag file: %v",
	RtagFileEmptyOrNotExist: ".rtag file is empty or does not exist",
	CurrentTags:             "Current tags:",
	AddTagFailed:            "Failed to add tag: %v",
	AddTagSuccess:           "Successfully added tag: %s",
	ReadTagsFailed:          "Failed to read tags: %v",
	NoTagsFound:             "No tags found",
	TagNotExistInFile:       "Tag '%s' does not exist in .rtag file",
	SpecifyTagOrUseAll:      "Please specify a tag to push or use --all flag",
	AllTags:                 "All tags:",
	RemoveTagFailed:         "Failed to remove tag: %v",
	RemoveTagSuccess:        "Successfully removed tag: %s",
	EnterTag:                "Enter a tag (or 'exit' to quit): ",
	Exit:                    "Exit",
	TagCannotBeEmpty:        "Tag cannot be empty, please try again",
	ContinueAdding:          "Continue adding? (y/n): ",
	ReadInputFailed:         "Failed to read input: %v",
	StartPushingTags:        "Starting to push tags (timestamp: %s)...",
	CreateGitTag:            "Creating git tag: %s",
	CreateTagFailed:         "Failed to create tag %s: %v",
	PushingTagsToRemote:     "Pushing tags to remote repository...",
	PushTagsFailed:          "Failed to push tags: %v",
	PushTagsSuccess:         "Successfully pushed all tags",
	TagAlreadyExists:        "tag '%s' already exists",
	TagNotExist:             "tag '%s' does not exist",
}

// Chinese messages
var zhMessages = Messages{
	RootShort: "发布标签管理工具",
	RootLong:  "rtag 是一个用于管理发布标签并集成 git 的命令行工具。",

	InitShort: "初始化 .rtag 文件",
	InitLong:  ".rtag 文件用于存储已发布的 tag。",

	AddShort: "添加新标签",
	AddLong:  "向 .rtag 文件添加新标签。如果未提供标签，将使用交互模式。",

	PushShort: "推送标签到远程仓库",
	PushLong:  "推送标签到远程仓库。使用 --all 标志推送所有标签。",

	ListShort: "列出所有标签",
	ListLong:  "列出 .rtag 文件中的所有标签。",

	RmShort: "删除标签",
	RmLong:  "从 .rtag 文件中删除标签。",

	PushAllFlag: "推送所有标签",

	ErrorReadingRtagFile:    "错误读取 .rtag 文件: %v",
	RtagFileEmptyOrNotExist: ".rtag 文件为空或不存在",
	CurrentTags:             "当前 tags:",
	AddTagFailed:            "添加 tag 失败: %v",
	AddTagSuccess:           "成功添加 tag: %s",
	ReadTagsFailed:          "读取 tags 失败: %v",
	NoTagsFound:             "没有找到任何 tags",
	TagNotExistInFile:       "Tag '%s' 不存在于 .rtag 文件中",
	SpecifyTagOrUseAll:      "请指定要推送的 tag 或使用 --all 标志",
	AllTags:                 "所有 tags:",
	RemoveTagFailed:         "删除 tag 失败: %v",
	RemoveTagSuccess:        "成功删除 tag: %s",
	EnterTag:                "请输入一个 tag (或输入 'exit' 退出): ",
	Exit:                    "退出",
	TagCannotBeEmpty:        "Tag 不能为空，请重新输入",
	ContinueAdding:          "是否继续添加? (y/n): ",
	ReadInputFailed:         "读取输入失败: %v",
	StartPushingTags:        "开始推送 tags (时间戳: %s)...",
	CreateGitTag:            "创建 git tag: %s",
	CreateTagFailed:         "创建 tag %s 失败: %v",
	PushingTagsToRemote:     "推送 tags 到远程仓库...",
	PushTagsFailed:          "推送 tags 失败: %v",
	PushTagsSuccess:         "成功推送所有 tags",
	TagAlreadyExists:        "tag '%s' 已存在",
	TagNotExist:             "tag '%s' 不存在",
}

// InitI18n initializes the internationalization
func InitI18n() {
	// Detect language from environment
	lang := detectLanguage()
	SetLanguage(lang)
}

// detectLanguage detects the user's preferred language
func detectLanguage() Language {
	// Check RTAG_LANG environment variable first
	if envLang := os.Getenv("RTAG_LANG"); envLang != "" {
		switch strings.ToLower(envLang) {
		case "zh", "zh_cn", "chinese":
			return LangZH
		case "en", "en_us", "english":
			return LangEN
		}
	}

	// Check LANG environment variable
	if envLang := os.Getenv("LANG"); envLang != "" {
		if strings.Contains(strings.ToLower(envLang), "zh") {
			return LangZH
		}
	}

	// Check LC_ALL environment variable
	if envLang := os.Getenv("LC_ALL"); envLang != "" {
		if strings.Contains(strings.ToLower(envLang), "zh") {
			return LangZH
		}
	}

	// Default to English
	return LangEN
}

// SetLanguage sets the current language
func SetLanguage(lang Language) {
	currentLang = lang
	switch lang {
	case LangZH:
		messages = zhMessages
	default:
		messages = enMessages
	}
}

// GetLanguage returns the current language
func GetLanguage() Language {
	return currentLang
}

// T returns the translated message (alias for GetMessage)
func T() Messages {
	return messages
}
