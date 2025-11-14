package ui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestReader_EmptyFile(t *testing.T) {
	filename := "temp.txt"
	content := ""

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	_, err := FromFileToMaze(filename)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Файл пустой")
}

func TestReader_InvalidFile(t *testing.T) {
	filename := "./invalid/temp.txt"

	_, err := FromFileToMaze(filename)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Не удалось открыть файл")
}

func TestReader_IncorrectFile_WithoutSymb(t *testing.T) {
	filename := "incorrect.txt"
	content := `#####
# # #
# # #
#   #
####`

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	_, err := FromFileToMaze(filename)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Некорректный формат лабиринта: строка 4 имеет длину 4, ожидалось 5")
}

func TestReader_IncorrectFile_Badsymb(t *testing.T) {
	filename := "incorrect.txt"
	content := `#####
# # #
# - #
#   #
#####`

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	_, err := FromFileToMaze(filename)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Неизвестный символ '-' в позиции (2, 2)")
}

func TestReader_CorrectFile(t *testing.T) {
	filename := "correct.txt"
	content := `######
#   ##
# # ##
# # ##
######`

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	maze, err := FromFileToMaze(filename)
	assert.Nil(t, err)

	assert.Equal(t, 4, maze.Width)
	assert.Equal(t, 3, maze.Height)

	assert.Equal(t, 5, len(maze.Grid))
	for _, row := range maze.Grid {
		assert.Equal(t, 6, len(row))
	}

	// Проверяем содержимое лабиринта
	// Первая строка должна быть стенами
	assert.Equal(t, domain.CellWall, maze.Grid[0][0])
	assert.Equal(t, domain.CellWall, maze.Grid[0][1])
	assert.Equal(t, domain.CellWall, maze.Grid[0][2])
	assert.Equal(t, domain.CellWall, maze.Grid[0][3])
	assert.Equal(t, domain.CellWall, maze.Grid[0][4])

	// Вторая строка: стена, пусто, пусто, пусто, стена, стена
	assert.Equal(t, domain.CellWall, maze.Grid[1][0])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][1])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][2])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][3])
	assert.Equal(t, domain.CellWall, maze.Grid[1][4])
	assert.Equal(t, domain.CellWall, maze.Grid[1][5])

	// Третья строка: стена, пусто, стена, пусто, стена, стена
	assert.Equal(t, domain.CellWall, maze.Grid[2][0])
	assert.Equal(t, domain.CellEmpty, maze.Grid[2][1])
	assert.Equal(t, domain.CellWall, maze.Grid[2][2])
	assert.Equal(t, domain.CellEmpty, maze.Grid[2][3])
	assert.Equal(t, domain.CellWall, maze.Grid[2][4])
	assert.Equal(t, domain.CellWall, maze.Grid[2][5])

	// Четвертая строка: стена, пусто, стена, пусто, стена, стена
	assert.Equal(t, domain.CellWall, maze.Grid[3][0])
	assert.Equal(t, domain.CellEmpty, maze.Grid[3][1])
	assert.Equal(t, domain.CellWall, maze.Grid[3][2])
	assert.Equal(t, domain.CellEmpty, maze.Grid[3][3])
	assert.Equal(t, domain.CellWall, maze.Grid[3][4])
	assert.Equal(t, domain.CellWall, maze.Grid[3][5])

	// Пятая строка должна быть стенами
	assert.Equal(t, domain.CellWall, maze.Grid[4][0])
	assert.Equal(t, domain.CellWall, maze.Grid[4][1])
	assert.Equal(t, domain.CellWall, maze.Grid[4][2])
	assert.Equal(t, domain.CellWall, maze.Grid[4][3])
	assert.Equal(t, domain.CellWall, maze.Grid[4][4])
	assert.Equal(t, domain.CellWall, maze.Grid[4][5])
}

func TestReaderEnhMaze_CorrectFile(t *testing.T) {
	filename := "correct.txt"
	content := `#######
#    &#
# #$###
#=# ~ #
#######`

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	maze, err := FromFileToMaze(filename)
	assert.Nil(t, err)

	assert.Equal(t, 5, maze.Width)
	assert.Equal(t, 3, maze.Height)

	assert.Equal(t, 5, len(maze.Grid))
	for _, row := range maze.Grid {
		assert.Equal(t, 7, len(row))
	}

	// Проверяем содержимое лабиринта
	// Первая строка должна быть стенами
	assert.Equal(t, domain.CellWall, maze.Grid[0][0])
	assert.Equal(t, domain.CellWall, maze.Grid[0][1])
	assert.Equal(t, domain.CellWall, maze.Grid[0][2])
	assert.Equal(t, domain.CellWall, maze.Grid[0][3])
	assert.Equal(t, domain.CellWall, maze.Grid[0][4])
	assert.Equal(t, domain.CellWall, maze.Grid[0][5])
	assert.Equal(t, domain.CellWall, maze.Grid[0][6])

	// Вторая строка: стена, пусто, пусто, пусто, стена, стена
	assert.Equal(t, domain.CellWall, maze.Grid[1][0])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][1])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][2])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][3])
	assert.Equal(t, domain.CellEmpty, maze.Grid[1][4])
	assert.Equal(t, domain.CellSwamp, maze.Grid[1][5])
	assert.Equal(t, domain.CellWall, maze.Grid[1][6])

	// Третья строка: стена, пусто, стена, пусто, стена, стена
	assert.Equal(t, domain.CellWall, maze.Grid[2][0])
	assert.Equal(t, domain.CellEmpty, maze.Grid[2][1])
	assert.Equal(t, domain.CellWall, maze.Grid[2][2])
	assert.Equal(t, domain.CellCoin, maze.Grid[2][3])
	assert.Equal(t, domain.CellWall, maze.Grid[2][4])
	assert.Equal(t, domain.CellWall, maze.Grid[2][5])
	assert.Equal(t, domain.CellWall, maze.Grid[2][6])

	// Четвертая строка: стена, пусто, стена, пусто, стена, стена
	assert.Equal(t, domain.CellWall, maze.Grid[3][0])
	assert.Equal(t, domain.CellSmooth, maze.Grid[3][1])
	assert.Equal(t, domain.CellWall, maze.Grid[3][2])
	assert.Equal(t, domain.CellEmpty, maze.Grid[3][3])
	assert.Equal(t, domain.CellSand, maze.Grid[3][4])
	assert.Equal(t, domain.CellEmpty, maze.Grid[3][5])
	assert.Equal(t, domain.CellWall, maze.Grid[3][6])

	// Пятая строка должна быть стенами
	assert.Equal(t, domain.CellWall, maze.Grid[4][0])
	assert.Equal(t, domain.CellWall, maze.Grid[4][1])
	assert.Equal(t, domain.CellWall, maze.Grid[4][2])
	assert.Equal(t, domain.CellWall, maze.Grid[4][3])
	assert.Equal(t, domain.CellWall, maze.Grid[4][4])
	assert.Equal(t, domain.CellWall, maze.Grid[4][5])
	assert.Equal(t, domain.CellWall, maze.Grid[4][6])
}

func TestReader_UnicodeFile(t *testing.T) {
	filename := "correct.txt"
	content := `┌─────┐
│ ╷ ╶─┤
└─┴───┘`
	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	_, err := FromFileToMaze(filename)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Файл содержит лабиринт в unicode формате. Для корректного чтения файла лабиринт должен быть в ASCII формате")
}
