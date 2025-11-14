package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application/solver"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure/ui"
)

type SolveOptions struct {
	algorithm string
	file      string
	output    string
	startStr  string
	endStr    string
	unicode   int
}

func NewSolveCommand() *cobra.Command {
	opts := &SolveOptions{}

	cmd := &cobra.Command{
		Use:   "solve",
		Short: "Solve a maze with specified algorithm and points",
		Long:  "Solve a maze with specified algorithm and points",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.algorithm != "astar" && opts.algorithm != "dijkstra" {
				return fmt.Errorf("Неподдерживаемый алгоритм: %s. Поддерживаемые: astar, dijkstra", opts.algorithm)
			}

			if opts.file == "" {
				return fmt.Errorf("Не указан файл с лабиринтом")
			}

			if opts.startStr == "" || opts.endStr == "" {
				return fmt.Errorf("Не указаны точки старта и/или конца пути")
			}

			start, err := parsePoint(opts.startStr)
			if err != nil {
				return fmt.Errorf("Некорректный формат точки старта %s: %w", opts.startStr, err)
			}

			if start.X <= 0 || start.Y <= 0 {
				return fmt.Errorf("Некорректные координаты точки старта: %s", opts.startStr)
			}

			end, err := parsePoint(opts.endStr)
			if err != nil {
				return fmt.Errorf("Некорректный формат точки финиша %s: %w", opts.endStr, err)
			}

			if end.X <= 0 || end.Y <= 0 {
				return fmt.Errorf("Некорректные координаты точки финиша: %s", opts.endStr)
			}

			if opts.unicode != 1 && opts.unicode != 0 {
				return fmt.Errorf("Недопустимое значение флага unicode: %d. Допустимые: 0, 1", opts.unicode)
			}

			var sol application.Solver
			switch opts.algorithm {
			case "dijkstra":
				sol = solver.NewDijkstraSolver()
			case "astar":
				sol = solver.NewAStarSolver()
			default:
				return fmt.Errorf("Неизвестный алгоритм %s", opts.algorithm)
			}

			maze, err := ui.FromFileToMaze(opts.file)
			if err != nil {
				return fmt.Errorf("Ошибка чтения лабиринта из файла: %w", err)
			}

			if maze.Width < start.X || maze.Height < start.Y {
				return fmt.Errorf("Некорректные координаты точки старта: %s", opts.startStr)
			}

			if maze.Grid[start.Y][start.X] == domain.CellWall {
				return fmt.Errorf("Указанная точка старта %s не является пустой клеткой", opts.startStr)
			}

			if maze.Width < end.X || maze.Height < end.Y {
				return fmt.Errorf("Некорректные координаты точки финиша: %s", opts.endStr)
			}

			if maze.Grid[end.Y][end.X] == domain.CellWall {
				return fmt.Errorf("Указанная точка финиша %s не является пустой клеткой", opts.endStr)
			}

			path, err := sol.Solve(maze, start, end)
			if err != nil {
				return fmt.Errorf("Путь от %v до %v не найден", start, end)
			}

			if opts.output == "" {
				maze.ApplyPath(path)
				if opts.unicode == 0 {
					ui.PrintMaze(maze)
				} else {
					ui.PrintMazeUnicode(maze, true)
				}
			} else {
				maze.ApplyPath(path)
				if opts.unicode == 0 {
					err = ui.SaveMazeToFile(opts.output, maze)
					if err != nil {
						return fmt.Errorf("Ошибка при сохранении лабиринта: %w", err)
					}
				} else {
					err = ui.SaveMazeToFileUnicode(opts.output, maze, true)
					if err != nil {
						return fmt.Errorf("Ошибка при сохранении лабиринта: %w", err)
					}
				}
			}

			return nil
		},
		SilenceUsage: true,
	}

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printSolveHelp(cmd)
	})

	cmd.Flags().StringVarP(&opts.algorithm, "algorithm", "a", "dijkstra", "Алгоритм решения лабиринта: astar, dijkstra")
	cmd.Flags().StringVarP(&opts.file, "file", "f", "", "Файл с лабиринтом")
	cmd.Flags().StringVarP(&opts.startStr, "start", "s", "", "Координаты начальной точки (x,y)")
	cmd.Flags().StringVarP(&opts.endStr, "end", "e", "", "Координаты конечной точки (x,y)")
	cmd.Flags().StringVarP(&opts.output, "output", "o", "", "Файл для сохранения решения (если не указан, вывод в консоль)")
	cmd.Flags().IntVarP(&opts.unicode, "unicode", "u", 0, "Использовать unicode символы для вывода лабиринта")

	return cmd
}

func parsePoint(pointStr string) (domain.Point, error) {
	if pointStr == "" {
		return domain.Point{}, fmt.Errorf("строка координат пуста")
	}

	parts := strings.Split(pointStr, ",")
	if len(parts) != 2 {
		return domain.Point{}, fmt.Errorf("ожидается формат x,y, получено: %s", pointStr)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return domain.Point{}, fmt.Errorf("ошибка при парсинге координаты x: %w", err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return domain.Point{}, fmt.Errorf("ошибка при парсинге координаты y: %w", err)
	}

	return domain.Point{X: x, Y: y}, nil
}

func printSolveHelp(cmd *cobra.Command) {
	fmt.Print("Usage: maze-app solve [OPTIONS]\r\n")
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
		flagLine += fmt.Sprintf("     \t %s", flag.Usage)

		// Добавляем значение по умолчанию, если оно есть
		if flag.DefValue != "" && flag.DefValue != "false" {
			flagLine += fmt.Sprintf(" (default \"%s\")", flag.DefValue)
		}

		fmt.Print(flagLine + "\r\n")
	})
}
