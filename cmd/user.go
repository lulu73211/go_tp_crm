package cmd

import (
	"fmt"
	"go_tp_crm/models"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Gestion des utilisateurs",
	Run: func(cmd *cobra.Command, args []string) {
		u := models.User{Name: "Lucas"}
		fmt.Println("Nouvel utilisateur créé :", u.Name)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
