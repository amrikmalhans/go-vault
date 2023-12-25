package govault

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault is a fast system storage explorer",
	Long:  `A fast and flexible system storage explorer to see file sizes, stats and more built with Go`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Vault")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
