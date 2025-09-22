package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var currentLang Language = LangEN
var messages Messages

// InitI18n initializes the internationalization
func InitI18n() {
	// Detect language from environment
	lang := detectLanguage()
	SetLanguage(lang)
}

// detectLanguage detects the user's preferred language
func detectLanguage() Language {
	// Check RTAG_LANG environment variable first (highest priority)
	if envLang := os.Getenv("RTAG_LANG"); envLang != "" {
		if lang, valid := GetLanguageFromString(strings.ToLower(envLang)); valid {
			return lang
		}
	}

	// Check saved language preference (second priority)
	if savedLang := loadSavedLanguage(); savedLang != "" {
		if lang, valid := GetLanguageFromString(savedLang); valid {
			return lang
		}
	}

	// Auto-detect system language (third priority)
	return detectSystemLanguage()
}

// detectSystemLanguage detects language from system environment
func detectSystemLanguage() Language {
	// Check multiple environment variables for locale
	envVars := []string{"LANG", "LC_ALL", "LC_MESSAGES", "LANGUAGE"}

	for _, envVar := range envVars {
		if envValue := os.Getenv(envVar); envValue != "" {
			if lang := DetectLanguageFromLocale(envValue); lang != LangEN {
				return lang
			}
		}
	}

	// Additional check for macOS system language
	if macLang := getMacOSSystemLanguage(); macLang != LangEN {
		return macLang
	}

	// Additional check for Windows system language
	if winLang := getWindowsSystemLanguage(); winLang != LangEN {
		return winLang
	}

	return LangEN // Default to English
}

// getMacOSSystemLanguage gets macOS system language
func getMacOSSystemLanguage() Language {
	cmd := exec.Command("defaults", "read", "-g", "AppleLanguages")
	output, err := cmd.Output()
	if err != nil {
		return LangEN
	}

	return DetectLanguageFromLocale(string(output))
}

// getWindowsSystemLanguage gets Windows system language
func getWindowsSystemLanguage() Language {
	// Check Windows-specific environment variables
	if userLang := os.Getenv("USERLANG"); userLang != "" {
		if lang := DetectLanguageFromLocale(userLang); lang != LangEN {
			return lang
		}
	}

	// Try to get Windows system locale using PowerShell
	cmd := exec.Command("powershell", "-Command", "Get-Culture | Select-Object -ExpandProperty Name")
	output, err := cmd.Output()
	if err != nil {
		return LangEN
	}

	return DetectLanguageFromLocale(string(output))
}

// loadSavedLanguage loads the saved language preference from config file
func loadSavedLanguage() string {
	configPath := getConfigPath()
	if configPath == "" {
		return ""
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return ""
	}

	lang := strings.TrimSpace(string(data))
	if IsValidLanguage(lang) {
		return lang
	}
	return ""
}

// saveLanguage saves the language preference to config file
func saveLanguage(lang Language) error {
	configPath := getConfigPath()
	if configPath == "" {
		return fmt.Errorf("unable to determine config path")
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	// Write language preference
	langStr := string(lang)
	if err := os.WriteFile(configPath, []byte(langStr), 0644); err != nil {
		return fmt.Errorf("failed to save language preference: %v", err)
	}

	return nil
}

// getConfigPath returns the path to the config file
func getConfigPath() string {
	// Try XDG_CONFIG_HOME first
	if configHome := os.Getenv("XDG_CONFIG_HOME"); configHome != "" {
		return filepath.Join(configHome, "rtag", "config")
	}

	// Fall back to ~/.config/rtag/config
	if homeDir := os.Getenv("HOME"); homeDir != "" {
		return filepath.Join(homeDir, ".config", "rtag", "config")
	}

	// Windows fallback
	if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
		return filepath.Join(userProfile, ".rtag", "config")
	}

	return ""
}

// SetLanguage sets the current language
func SetLanguage(lang Language) {
	currentLang = lang
	allMessages := GetAllMessages()
	if msg, exists := allMessages[lang]; exists {
		messages = msg
	} else {
		messages = allMessages[LangEN] // Default to English
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
