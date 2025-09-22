package rtag

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	// Initialize i18n
	InitI18n()

	// Initialize root command with translated strings
	rootCmd = &cobra.Command{
		Use:   "rtag",
		Short: T().RootShort,
		Long:  T().RootLong,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing root command: %v", err)
		os.Exit(1)
	}
}
