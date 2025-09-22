package rtag

import (
	"fmt"

	"github.com/spf13/cobra"
)

var langCmd *cobra.Command

func runLang(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		// Display current language
		currentLang := GetLanguage()
		switch currentLang {
		case LangEN:
			fmt.Println("Current language: English (en)")
		case LangZH:
			fmt.Println("当前语言: 中文 (zh)")
		}
		fmt.Println()
		fmt.Println("Available languages:")
		fmt.Println("  en - English")
		fmt.Println("  zh - 中文")
		fmt.Println()
		fmt.Println("Usage: rtag lang [en|zh]")
		fmt.Println("Environment variable: RTAG_LANG=[en|zh]")
		return
	}

	// Set language
	lang := args[0]
	switch lang {
	case "en", "english":
		SetLanguage(LangEN)
		fmt.Println("Language set to English")
	case "zh", "chinese", "中文":
		SetLanguage(LangZH)
		fmt.Println("语言已设置为中文")
	default:
		fmt.Printf("Unsupported language: %s\n", lang)
		fmt.Println("Supported languages: en, zh")
		return
	}

	fmt.Println("Note: Language change will take effect on next command execution.")
}
