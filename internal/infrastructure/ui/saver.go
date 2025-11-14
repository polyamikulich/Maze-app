package ui

import (
	"bufio"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func SaveMazeToFile(filename string, maze *domain.Maze) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for y := 0; y < maze.Height+2; y++ {
		for x := 0; x < maze.Width+2; x++ {
			cell := maze.Grid[y][x]
			char := cellTypeToChar(cell)

			_, err := writer.WriteRune(char)
			if err != nil {
				return err
			}
		}

		_, err := writer.WriteRune('\n')
		if err != nil {
			return err
		}
	}

	return nil
}

func cellTypeToChar(cellType domain.CellType) rune {
	switch cellType {
	case domain.CellWall:
		return '#'
	case domain.CellEmpty:
		return ' '
	case domain.CellStart:
		return 'O'
	case domain.CellEnd:
		return 'X'
	case domain.CellPath:
		return '.'
	case domain.CellCoin:
		return '$'
	case domain.CellSand:
		return '~'
	case domain.CellSwamp:
		return '&'
	case domain.CellSmooth:
		return '='
	default:
		return '?' // Неизвестный тип ячейки
	}
}

func SaveMazeToFileUnicode(filename string, maze *domain.Maze, isPath bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Если лабиринт выводится без пути, то цикл по y делаем с шагом 2, так красивее
	// Если лабиринт выводится с путем, то цикл по y делаем с шагом 1, потому что иначе не вывести
	var step int
	if isPath {
		step = 1
	} else {
		step = 2
	}

	chars := []rune{' ', '╴', '╷', '┐', '╶', '─', '┌', '┬', '╵', '┘', '│', '┤', '└', '┴', '├', '┼', '.', 'O', 'X', '░', '▓', '○', '□'}

	for y := 0; y < maze.Height+2; y += step {
		for x := 0; x < maze.Width+2; x++ {
			var char rune

			if maze.Grid[y][x] == domain.CellEmpty {
				charIndex := 0
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellPath {
				charIndex := 16
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellStart {
				charIndex := 17
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellEnd {
				charIndex := 18
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellSand {
				charIndex := 19
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellSwamp {
				charIndex := 20
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellCoin {
				charIndex := 21
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}
			if maze.Grid[y][x] == domain.CellSmooth {
				charIndex := 22
				char = chars[charIndex]
				_, err := writer.WriteRune(char)
				if err != nil {
					return err
				}
				continue
			}

			var u, l, d, r int

			// Проверим, есть ли стена сверху
			if y > 0 && maze.Grid[y-1][x] == domain.CellWall {
				u = 1
			}

			// Проверим, есть ли стена справа
			if x < maze.Width+1 && maze.Grid[y][x+1] == domain.CellWall {
				r = 1
			}

			// Проверим, есть ли стена снизу
			if y < maze.Height+1 && maze.Grid[y+1][x] == domain.CellWall {
				d = 1
			}

			// Проверим, есть ли стена слева
			if x > 0 && maze.Grid[y][x-1] == domain.CellWall {
				l = 1
			}

			// Символы в массиве расположены таким образом, что данная формула напрямую определяет символ,
			// который соответствует расположению стен вокруг клетки.
			charIndex := 8*u + 4*r + 2*d + l
			if charIndex < len(chars) {
				char = chars[charIndex]
			} else {
				char = '?'
			}

			_, err := writer.WriteRune(char)
			if err != nil {
				return err
			}
		}

		_, err := writer.WriteRune('\n')
		if err != nil {
			return err
		}
	}

	return nil
}
