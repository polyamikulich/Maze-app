package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRootCommand создает корневую команду CLI
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "maze-app",
		Short: "Maze generator and solver CLI application.",
		Long:  "Maze generator and solver CLI application.",
		Run: func(cmd *cobra.Command, args []string) {
			// Если команда запущена без аргументов, показываем справку
			printCustomHelp(cmd)
			//os.Exit(0)
		},
		SilenceUsage: true,
	}

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printCustomHelp(cmd)
	})

	// Добавляем флаг версии
	rootCmd.Flags().BoolP("version", "V", false, "Print version information and exit.")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if version, _ := cmd.Flags().GetBool("version"); version {
			// Вместо os.Exit просто возвращаем специальную ошибку или выводим и завершаем
			cmd.Println("maze-app v1.0.0")
			return nil // Просто завершаем выполнение
		}

		// Если не version, показываем help
		printCustomHelp(cmd)
		return nil
	}

	rootCmd.AddCommand(NewGenerateCommand())
	rootCmd.AddCommand(NewSolveCommand())

	return rootCmd
}

func printCustomHelp(cmd *cobra.Command) {
	fmt.Printf("Usage: %s [-hV] [COMMAND]\r\n", cmd.Use)
	fmt.Printf("%s\r\n", cmd.Long)
	fmt.Print("  -h, --help      Show this help message and exit.\r\n")
	fmt.Print("  -V, --version   Print version information and exit.\r\n")

	if len(cmd.Commands()) > 0 {
		fmt.Print("Commands:\r\n")
		for _, subCmd := range cmd.Commands() {
			// Пропускаем стандартные команды help и completion
			if subCmd.Name() == "help" || subCmd.Name() == "completion" {
				continue
			}

			if subCmd.Name() == "generate" {
				fmt.Print("  generate   Generate a maze with specified algorithm and dimensions.\r\n")
			}

			if subCmd.Name() == "solve" {
				fmt.Print("  solve      Solve a maze with specified algorithm and points.\r\n")
			}
		}
	}
}
