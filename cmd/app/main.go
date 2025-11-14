// Package main содержит точку входа для приложения
package main

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/cli"
)

func main() {
	rootCmd := cli.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("%w", err)
		os.Exit(0)
	}
}
