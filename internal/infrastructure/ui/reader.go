package ui

import (
	"bufio"
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func FromFileToMaze(filename string) (*domain.Maze, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Не удалось открыть файл: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла: %w", err)
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("Файл пустой")
	}

	file.Seek(0, 0)

	reader := bufio.NewReader(file)
	firstChar, _, err := reader.ReadRune()
	if err != nil {
		return nil, err
	}

	if firstChar == '┌' {
		return nil, fmt.Errorf("Файл содержит лабиринт в unicode формате. Для корректного чтения файла лабиринт должен быть в ASCII формате")
	}

	file.Seek(0, 0)

	height := len(lines) - 2
	width := len(lines[0]) - 2

	for i, line := range lines {
		if len(line) != width+2 {
			return nil, fmt.Errorf("Некорректный формат лабиринта: строка %d имеет длину %d, ожидалось %d", i, len(line), width+2)
		}
	}

	maze := domain.NewMaze(width, height)

	for i := 1; i < height+1; i++ {
		for j := 1; j < width+1; j++ {
			curr_line := lines[i]
			char := curr_line[j]
			switch char {
			case '#':
				maze.Grid[i][j] = domain.CellWall
			case ' ':
				maze.Grid[i][j] = domain.CellEmpty
			case '&':
				maze.Grid[i][j] = domain.CellSwamp
			case '~':
				maze.Grid[i][j] = domain.CellSand
			case '$':
				maze.Grid[i][j] = domain.CellCoin
			case '=':
				maze.Grid[i][j] = domain.CellSmooth
			case 'O':
				return nil, fmt.Errorf("Пустой лабиринт не может содержать символ начала пути 'O'")
			case 'X':
				return nil, fmt.Errorf("Пустой лабиринт не может содержать символ конца пути 'X'")
			case '.':
				return nil, fmt.Errorf("Пустой лабиринт не может содержать символ пути '.'")
			default:
				return nil, fmt.Errorf("Неизвестный символ '%c' в позиции (%d, %d)", char, i, j)
			}
		}
	}

	return maze, nil
}
