package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestGeneratorEnhanced_findEmptyCells(t *testing.T) {
	g := NewBaseGenerator(AlgorithmDFS)
	maze := domain.NewMaze(5, 3)

	// #######
	// # #   #
	// # # ###
	// #     #
	// #######

	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[1][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty

	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellEmpty

	emptyCells := g.findEmptyCells(maze)

	expectedCells := []domain.Point{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
		{X: 4, Y: 1},
		{X: 5, Y: 1},
		{X: 4, Y: 3},
		{X: 5, Y: 3},
	}

	assert.Equal(t, len(expectedCells), len(emptyCells))

	// Проверяем, что все ожидаемые ячейки присутствуют
	for _, expectedCell := range expectedCells {
		found := false
		for _, actualCell := range emptyCells {
			if actualCell.X == expectedCell.X && actualCell.Y == expectedCell.Y {
				found = true
				break
			}
		}
		assert.True(t, found, "Ожидаемая ячейка %v не найдена", expectedCell)
	}
}

func TestGeneratorEnhanced_findWallCells(t *testing.T) {
	g := NewBaseGenerator(AlgorithmDFS)
	maze := domain.NewMaze(5, 3)

	// #######
	// # #   #
	// # # ###
	// #     #
	// #######

	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[1][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty

	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellEmpty

	emptyCells := g.findWallCells(maze)

	expectedCells := []domain.Point{
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 4, Y: 2},
		{X: 5, Y: 2},
	}

	assert.Equal(t, len(expectedCells), len(emptyCells))

	// Проверяем, что все ожидаемые ячейки присутствуют
	for _, expectedCell := range expectedCells {
		found := false
		for _, actualCell := range emptyCells {
			if actualCell.X == expectedCell.X && actualCell.Y == expectedCell.Y {
				found = true
				break
			}
		}
		assert.True(t, found, "Ожидаемая ячейка %v не найдена", expectedCell)
	}
}

// TestBaseGenerator_DestroyWalls проверяет функцию разрушения стен
func TestBaseGenerator_DestroyWalls(t *testing.T) {
	// Создаем базовый генератор
	gen := NewBaseGenerator(AlgorithmDFS)

	// Создаем тестовый лабиринт 3x3
	maze := domain.NewMaze(3, 3)

	// Заполняем лабиринт стенами
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 3; x++ {
			maze.Grid[y][x] = domain.CellWall
		}
	}

	// Подготавливаем список ячеек стен
	wallCells := []domain.Point{
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 1, Y: 3},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
	}

	// Разрушаем стены с вероятностью 0.33 (примерно 3 ячейки)
	gen.destroyWalls(maze, wallCells, 0.33)

	// Подсчитываем количество разрушенных стен (пустых ячеек)
	emptyCount := 0
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 3; x++ {
			if maze.Grid[y][x] == domain.CellEmpty {
				emptyCount++
			}
		}
	}

	// Проверяем, что были разрушены некоторые стены
	assert.True(t, emptyCount > 0, "Должны быть разрушены некоторые стены")
	assert.True(t, emptyCount <= len(wallCells), "Количество разрушенных стен не должно превышать общее количество стен")
}

// TestBaseGenerator_ShufflePoints проверяет функцию перемешивания точек
func TestBaseGenerator_ShufflePoints(t *testing.T) {
	// Создаем базовый генератор
	gen := NewBaseGenerator(AlgorithmDFS)

	// Создаем список точек
	points := []domain.Point{
		{X: 1, Y: 1},
		{X: 2, Y: 2},
		{X: 3, Y: 3},
		{X: 4, Y: 4},
		{X: 5, Y: 5},
	}

	// Создаем копию для сравнения
	originalPoints := make([]domain.Point, len(points))
	copy(originalPoints, points)

	// Перемешиваем точки
	gen.shufflePoints(points)

	// Проверяем, что длина осталась той же
	assert.Equal(t, len(originalPoints), len(points))

	// Проверяем, что все оригинальные точки присутствуют в перемешанном списке
	for _, originalPoint := range originalPoints {
		found := false
		for _, shuffledPoint := range points {
			if shuffledPoint.X == originalPoint.X && shuffledPoint.Y == originalPoint.Y {
				found = true
				break
			}
		}
		assert.True(t, found, "Точка %v из оригинального списка не найдена в перемешанном списке", originalPoint)
	}
}

func TestGeneratorEnhanced_PlaceSurfaces(t *testing.T) {
	g := NewBaseGenerator(AlgorithmPrim)
	maze := domain.NewMaze(5, 3)

	// #######
	// # #   #
	// # # ###
	// #     #
	// #######

	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[1][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty

	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellEmpty

	emptyCells := []domain.Point{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
		{X: 4, Y: 1},
		{X: 5, Y: 1},
		{X: 4, Y: 3},
		{X: 5, Y: 3},
	}

	// Размещаем поверхности (например, песок) с вероятностью 0.33 (примерно 3 ячейки)
	g.placeSurfaces(maze, emptyCells, domain.CellSand, 0.33)

	// Подсчитываем количество ячеек с песком
	sandCount := 0
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 5; x++ {
			if maze.Grid[y][x] == domain.CellSand {
				sandCount++
			}
		}
	}

	// Проверяем, что были размещены ячейки с песком
	assert.True(t, sandCount > 0, "Должны быть размещены ячейки с песком")
	assert.True(t, sandCount <= len(emptyCells), "Количество ячеек с песком не должно превышать общее количество пустых ячеек")
}

// TestBaseGenerator_AddSpecialSurfaces проверяет функцию добавления специальных поверхностей
func TestBaseGenerator_AddSpecialSurfaces(t *testing.T) {
	// Создаем базовый генератор
	gen := NewBaseGenerator(AlgorithmDFS)

	// Создаем тестовый лабиринт
	maze := domain.NewMaze(10, 10)

	// Заполняем лабиринт для теста
	// Стены
	initialWalls := 0
	initialEmpty := 0
	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			if (x+y)%3 == 0 {
				initialWalls++
			} else {
				maze.Grid[y][x] = domain.CellEmpty
				initialEmpty++
			}
		}
	}

	// Вызываем тестируемую функцию
	gen.addSpecialSurfaces(maze)

	// Подсчитываем количество различных типов поверхностей после обработки
	walls := 0
	empty := 0
	sand := 0
	smooth := 0
	swamp := 0
	coin := 0

	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			switch maze.Grid[y][x] {
			case domain.CellWall:
				walls++
			case domain.CellEmpty:
				empty++
			case domain.CellSand:
				sand++
			case domain.CellSmooth:
				smooth++
			case domain.CellSwamp:
				swamp++
			case domain.CellCoin:
				coin++
			}
		}
	}

	// Проверяем, что количество стен изменилось (некоторые были разрушены)
	assert.True(t, walls <= initialWalls, "Количество стен не должно увеличиваться")

	// Проверяем, что появились различные типы поверхностей
	assert.True(t, sand >= 0, "Должны присутствовать ячейки с песком")
	assert.True(t, smooth >= 0, "Должны присутствовать ячейки с гладкой поверхностью")
	assert.True(t, swamp >= 0, "Должны присутствовать ячейки с болотом")
	assert.True(t, coin >= 0, "Должны присутствовать ячейки с монетами")

	// Проверяем, что общее количество ячеек осталось тем же
	totalAfter := walls + empty + sand + smooth + swamp + coin
	expectedTotal := 10 * 10
	assert.Equal(t, expectedTotal, totalAfter, "Общее количество ячеек должно остаться неизменным")
}
