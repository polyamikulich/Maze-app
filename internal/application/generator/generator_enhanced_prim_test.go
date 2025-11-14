package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestPrimnhancedSGenerator_NewGenerator(t *testing.T) {
	gen := NewPrimEnhancedGenerator()
	assert.NotNil(t, gen, "NewDFSGenerator() должен возвращать непустой генератор")
}

func TestPrimEnhancedGenerator_Generate(t *testing.T) {
	gen := NewPrimEnhancedGenerator()

	// Тестируем генерацию лабиринта 10x10
	maze, err := gen.Generate(10, 10)
	assert.NoError(t, err)
	assert.NotNil(t, maze)
	assert.Equal(t, 10, maze.Width)
	assert.Equal(t, 10, maze.Height)
	assert.NotNil(t, maze.Grid)
	assert.Equal(t, 12, len(maze.Grid))

	// Проверяем, что каждая строка имеет правильную длину
	for _, row := range maze.Grid {
		assert.Equal(t, 12, len(row))
	}

	// Проверяем, что границы лабиринта являются стенами
	for i := 0; i < 10; i++ {
		assert.Equal(t, domain.CellWall, maze.Grid[0][i])
		assert.Equal(t, domain.CellWall, maze.Grid[11][i])
		assert.Equal(t, domain.CellWall, maze.Grid[i][0])
		assert.Equal(t, domain.CellWall, maze.Grid[i][11])
	}

	// Проверяем, что внутри лабиринта есть различные типы поверхностей
	hasSwamp := false
	hasSand := false
	hasCoin := false
	hasSmooth := false
	hasEmpty := false

	for y := 1; y <= maze.Height; y++ {
		for x := 1; x <= maze.Width; x++ {
			switch maze.Grid[y][x] {
			case domain.CellEmpty:
				hasEmpty = true
			case domain.CellSwamp:
				hasSwamp = true
			case domain.CellSand:
				hasSand = true
			case domain.CellCoin:
				hasCoin = true
			case domain.CellSmooth:
				hasSmooth = true
			}
		}
	}

	// Проверяем, что в лабиринте есть хотя бы некоторые типы поверхностей
	assert.True(t, hasEmpty, "Лабиринт должен содержать пустые ячейки")
	assert.True(t, hasSwamp, "Лабиринт должен содержать ячейки с болотом")
	assert.True(t, hasSand, "Лабиринт должен содержать ячейки с песком")
	assert.True(t, hasCoin, "Лабиринт должен содержать ячейки с монетками")
	assert.True(t, hasSmooth, "Лабиринт должен содержать ячейки с гладкими участками")
}
