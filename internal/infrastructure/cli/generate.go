package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/generator"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/ui"
)

type GenerateOptions struct {
	algorithm string
	width     int
	height    int
	output    string
	unicode   int
}

func NewGenerateCommand() *cobra.Command {
	opts := &GenerateOptions{}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a maze with specified algorithm and dimensions",
		Long:  "Generate a maze with specified algorithm and dimensions",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.algorithm != "dfs" && opts.algorithm != "prim" && opts.algorithm != "prim-enhanced" && opts.algorithm != "dfs-enhanced" {
				return fmt.Errorf("Неподдерживаемый алгоритм: %s. Поддерживаемые: dfs, prim, dfs-enhanced, prim-enhanced", opts.algorithm)
			}

			if opts.width <= 0 || opts.height <= 0 {
				return fmt.Errorf("Ширина и высота должны быть положительными числами, получено: %d * %d", opts.width, opts.height)
			}

			var gen application.Generator
			// enh маркирует что используется расширенная версия алгоритма (с доп. поверхностями)
			// это влияет на отрисовку лабиринта
			var enh int
			switch opts.algorithm {
			case "dfs":
				gen = generator.NewDFSGenerator()
				enh = 0
			case "prim":
				gen = generator.NewPrimaGenerator()
				enh = 0
			case "dfs-enhanced":
				gen = generator.NewDFSEnhancedGenerator()
				enh = 1
			case "prim-enhanced":
				gen = generator.NewPrimEnhancedGenerator()
				enh = 1
			}

			if opts.unicode != 1 && opts.unicode != 0 {
				return fmt.Errorf("Недопустимое значение флага unicode: %d. Допустимые: 0, 1", opts.unicode)
			}

			if opts.unicode == 1 {
				opts.width = opts.width*2 - 1
				opts.height = opts.height*2 - 1
			}

			maze, err := gen.Generate(opts.width, opts.height)
			if err != nil {
				return fmt.Errorf("Ошибка при генерации лабиринта: %w", err)
			}

			if opts.output == "" {
				if opts.unicode == 0 {
					ui.PrintMaze(maze)
					// ui.PrintMazeUnicode(maze, false)
				} else {
					// в зависимости от enh рисуем лабиринт либо с true, либо с false
					if enh == 1 {
						ui.PrintMazeUnicode(maze, true)
					} else {
						ui.PrintMazeUnicode(maze, false)
					}
					// ui.PrintMaze(maze)
				}
			} else {
				if opts.unicode == 0 {
					err = ui.SaveMazeToFile(opts.output, maze)
					if err != nil {
						return fmt.Errorf("Ошибка при сохранении лабиринта: %w", err)
					}
				} else {
					if enh == 1 {
						err = ui.SaveMazeToFileUnicode(opts.output, maze, true)
						if err != nil {
							return fmt.Errorf("Ошибка при сохранении лабиринта: %w", err)
						}
					} else {
						err = ui.SaveMazeToFileUnicode(opts.output, maze, false)
						if err != nil {
							return fmt.Errorf("Ошибка при сохранении лабиринта: %w", err)
						}
					}
				}
			}
			return nil
		},
		SilenceUsage: true,
	}

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printGenerateHelp(cmd)
	})

	cmd.Flags().StringVarP(&opts.algorithm, "algorithm", "a", "dfs", "Алгоритм генерации лабиринта: dfs, prim, dfs-enhanced, prim-enhanced")
	cmd.Flags().IntVarP(&opts.width, "width", "W", 10, "Ширина лабиринта")
	cmd.Flags().IntVarP(&opts.height, "height", "H", 10, "Высота лабиринта")
	cmd.Flags().StringVarP(&opts.output, "output", "o", "", "Файл для сохранения лабиринта (если не указан, вывод в консоль)")
	cmd.Flags().IntVarP(&opts.unicode, "unicode", "u", 0, "Использовать unicode символы для вывода лабиринта")

	return cmd
}

func printGenerateHelp(cmd *cobra.Command) {
	fmt.Print("Usage: maze-app generate [OPTIONS]\r\n")
	fmt.Printf("%s", cmd.Long)
	fmt.Print("\r\n")
	fmt.Print("Options:\r\n")

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		// Форматируем вывод флага
		var flagLine string
		if flag.Shorthand != "" {
			flagLine = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
		} else {
			flagLine = fmt.Sprintf("      --%s", flag.Name)
		}

		// Добавляем описание
		flagLine += fmt.Sprintf("    \t %s", flag.Usage)

		// Добавляем значение по умолчанию, если оно есть
		if flag.DefValue != "" && flag.DefValue != "false" {
			flagLine += fmt.Sprintf(" (default \"%s\")", flag.DefValue)
		}

		fmt.Print(flagLine + "\r\n")
	})
}
