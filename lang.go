package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var langCmd *cobra.Command

func runLang(cmd *cobra.Command, args []string) {
	supportedLangs := GetSupportedLanguages()

	if len(args) == 0 {
		// Display current language
		currentLang := GetLanguage()
		langInfo := supportedLangs[currentLang]

		fmt.Printf("%s: %s (%s)\n\n", T().CurrentLanguage, langInfo.NativeName, langInfo.Code)
		fmt.Println(T().AvailableLanguages + ":")

		// Display all supported languages
		for _, info := range supportedLangs {
			fmt.Printf("  %s - %s\n", info.Code, info.NativeName)
		}

		fmt.Println()

		// Create language codes list for usage
		var langCodes []string
		for code := range supportedLangs {
			langCodes = append(langCodes, string(code))
		}

		fmt.Printf(T().LanguageUsage+"\n", strings.Join(langCodes, "|"))
		fmt.Printf(T().EnvironmentVariable+"\n", strings.Join(langCodes, "|"))
		fmt.Println(T().ConfigFile)
		return
	}

	// Set language
	langArg := strings.ToLower(args[0])
	newLang, valid := GetLanguageFromString(langArg)

	if !valid {
		var langCodes []string
		for code := range supportedLangs {
			langCodes = append(langCodes, string(code))
		}
		fmt.Printf(T().InvalidLanguage+"\n", langArg, strings.Join(langCodes, ", "))
		return
	}

	// Save language preference
	if err := saveLanguage(newLang); err != nil {
		fmt.Printf("Failed to save language preference: %v\n", err)
		return
	}

	// Set current language
	SetLanguage(newLang)

	langInfo := supportedLangs[newLang]
	fmt.Printf(T().LanguageSetTo+"\n", langInfo.NativeName)
	fmt.Println(T().LanguagePreferenceSaved)
}
