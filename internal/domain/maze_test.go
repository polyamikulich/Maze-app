package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaze_NewMaze(t *testing.T) {
	maze := NewMaze(5, 7)

	assert.NotNil(t, maze, "Лабиринт не создан")
	assert.Equal(t, 5, maze.Width, "Неверная ширина лабиринта")
	assert.Equal(t, 7, maze.Height, "Неверная высота лабиринта")
	assert.NotNil(t, maze.Grid, "Нет массива ячеек")
	assert.Equal(t, 9, len(maze.Grid), "Неверная высота сетки")

	for i := 0; i < maze.Height; i++ {
		assert.Equal(t, 7, len(maze.Grid[i]), "Неверная ширина строки в сетке")
	}

	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
			assert.Equal(t, CellWall, maze.Grid[i][j], "Неверный тип ячейки")
		}
	}
}

func TestMaze_ApplyPath(t *testing.T) {
	maze := NewMaze(5, 7)

	path := &Path{
		Points: []Point{
			{X: 1, Y: 1}, // Start
			{X: 2, Y: 1}, // Path
			{X: 3, Y: 1}, // Path
			{X: 3, Y: 2}, // Path
			{X: 3, Y: 3}, // End
		},
	}

	maze.ApplyPath(path)

	assert.Equal(t, CellStart, maze.Grid[1][1], "Неверный тип ячейки")
	assert.Equal(t, CellEnd, maze.Grid[3][3], "Неверный тип ячейки")
	assert.Equal(t, CellPath, maze.Grid[1][2], "Неверный тип ячейки")
	assert.Equal(t, CellPath, maze.Grid[1][3], "Неверный тип ячейки")
	assert.Equal(t, CellPath, maze.Grid[2][3], "Неверный тип ячейки")

	assert.Equal(t, CellWall, maze.Grid[5][5], "Неиспользованные ячейки должны оставаться стенами")
	assert.Equal(t, CellWall, maze.Grid[2][2], "Неиспользованные ячейки должны оставаться стенами")
}
