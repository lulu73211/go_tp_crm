package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "CRM CLI avec Cobra, Viper et Gorm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bienvenue dans ton CRM CLI 🚀")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
