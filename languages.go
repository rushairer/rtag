package main

import "strings"

// Language represents supported languages
type Language string

// Supported languages
const (
	LangEN Language = "en"
	LangZH Language = "zh"
	LangFR Language = "fr"
	LangRU Language = "ru"
)

// LanguageInfo contains language metadata
type LanguageInfo struct {
	Code       Language
	Name       string
	NativeName string
	LocaleKeys []string // Locale identifiers for auto-detection
}

// GetSupportedLanguages returns all supported languages with their metadata
func GetSupportedLanguages() map[Language]LanguageInfo {
	return map[Language]LanguageInfo{
		LangEN: {
			Code:       LangEN,
			Name:       "English",
			NativeName: "English",
			LocaleKeys: []string{"en", "english"},
		},
		LangZH: {
			Code:       LangZH,
			Name:       "Chinese",
			NativeName: "中文",
			LocaleKeys: []string{"zh", "chinese", "cn", "tw", "hk", "mo"},
		},
		LangFR: {
			Code:       LangFR,
			Name:       "French",
			NativeName: "Français",
			LocaleKeys: []string{"fr", "french", "francais"},
		},
		LangRU: {
			Code:       LangRU,
			Name:       "Russian",
			NativeName: "Русский",
			LocaleKeys: []string{"ru", "russian", "русский"},
		},
	}
}

// IsValidLanguage checks if a language code is supported
func IsValidLanguage(lang string) bool {
	supportedLangs := GetSupportedLanguages()
	_, exists := supportedLangs[Language(lang)]
	return exists
}

// GetLanguageFromString converts string to Language type if valid
func GetLanguageFromString(lang string) (Language, bool) {
	if IsValidLanguage(lang) {
		return Language(lang), true
	}
	return LangEN, false
}

// DetectLanguageFromLocale detects language from locale string
func DetectLanguageFromLocale(locale string) Language {
	if locale == "" {
		return LangEN
	}

	locale = strings.ToLower(locale)
	supportedLangs := GetSupportedLanguages()

	// Check each supported language's locale keys
	for langCode, langInfo := range supportedLangs {
		for _, key := range langInfo.LocaleKeys {
			if strings.Contains(locale, key) {
				return langCode
			}
		}
	}

	return LangEN // Default to English
}
