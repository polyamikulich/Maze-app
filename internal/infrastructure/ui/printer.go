package ui

import (
	"fmt"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func PrintMaze(maze *domain.Maze) {
	if maze == nil || maze.Grid == nil {
		return
	}

	for y := 0; y < maze.Height+2; y++ {
		for x := 0; x < maze.Width+2; x++ {
			switch maze.Grid[y][x] {
			case domain.CellWall:
				fmt.Print("#")
			case domain.CellEmpty:
				fmt.Print(" ")
			case domain.CellStart:
				fmt.Print("O")
			case domain.CellEnd:
				fmt.Print("X")
			case domain.CellPath:
				fmt.Print(".")
			case domain.CellCoin:
				fmt.Print("$")
			case domain.CellSand:
				fmt.Print("~")
			case domain.CellSwamp:
				fmt.Print("&")
			case domain.CellSmooth:
				fmt.Print("=")
			default:
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func PrintMazeUnicode(maze *domain.Maze, isPath bool) {
	if maze == nil || maze.Grid == nil {
		return
	}

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
			if maze.Grid[y][x] == domain.CellEmpty {
				charIndex := 0
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellPath {
				charIndex := 16
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellStart {
				charIndex := 17
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellEnd {
				charIndex := 18
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellSand {
				charIndex := 19
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellSwamp {
				charIndex := 20
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellCoin {
				charIndex := 21
				fmt.Printf("%c", chars[charIndex])
				continue
			}
			if maze.Grid[y][x] == domain.CellSmooth {
				charIndex := 22
				fmt.Printf("%c", chars[charIndex])
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
				fmt.Printf("%c", chars[charIndex])
			} else {
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}
