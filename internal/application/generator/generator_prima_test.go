package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestPrimaGenerator_NewPrimaGenerator(t *testing.T) {
	gen := NewPrimaGenerator()
	assert.NotNil(t, gen, "NewPrimaGenerator() должен возвращать непустой генератор")
}

// TestDFSGenerator_Generate тестирует генерацию лабиринта с помощью Prima
func TestPrimaGenerator_Generate(t *testing.T) {
	gen := NewDFSGenerator()

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
}

func TestPrimaGenerator_AddWalls(t *testing.T) {
	gen := NewPrimaGenerator()
	maze := domain.NewMaze(4, 4)

	// Создаем пустую карту стен
	walls := make(map[domain.Point]bool)

	point1 := domain.Point{X: 1, Y: 1}
	gen.addWalls(maze, point1, walls)

	assert.Equal(t, 2, len(walls))

	expectedWalls := []domain.Point{
		{X: 1, Y: 0}, // Верхняя стена
		{X: 1, Y: 2}, // Нижняя стена
		{X: 0, Y: 1}, // Левая стена
		{X: 2, Y: 1}, // Правая стена
	}

	// Проверяем, что все ожидаемые стены добавлены
	for _, expectedWall := range expectedWalls {
		// Так как лабиринт окружен стенами, некоторые координаты могут быть недоступны
		// Проверяем только те, которые находятся в границах
		if gen.isInBounds(expectedWall.X, expectedWall.Y, maze) {
			assert.True(t, walls[expectedWall], "Стена %v должна быть добавлена", expectedWall)
		}
	}

	// Добавляем стены вокруг точки (3,3)
	point2 := domain.Point{X: 3, Y: 3}
	gen.addWalls(maze, point2, walls)

	assert.Equal(t, 6, len(walls))

	// Проверяем, что стены были добавлены
	// Вокруг точки (3,3) должны быть стены (3,2), (3,4), (2,3), (4,3)
	// Но так как лабиринт окружен стенами, некоторые из них могут быть недоступны

	expectedWalls2 := []domain.Point{
		{X: 3, Y: 2}, // Верхняя стена
		{X: 3, Y: 4}, // Нижняя стена
		{X: 2, Y: 3}, // Левая стена
		{X: 4, Y: 3}, // Правая стена
	}

	// Проверяем, что все ожидаемые стены добавлены
	for _, expectedWall := range expectedWalls2 {
		// Так как лабиринт окружен стенами, некоторые координаты могут быть недоступны
		// Проверяем только те, которые находятся в границах
		if gen.isInBounds(expectedWall.X, expectedWall.Y, maze) {
			assert.True(t, walls[expectedWall], "Стена %v должна быть добавлена", expectedWall)
		}
	}

	point3 := domain.Point{X: 2, Y: 4}
	gen.addWalls(maze, point3, walls)

	assert.Equal(t, 7, len(walls))

	// Проверяем, что стены были добавлены
	// Вокруг точки (3,3) должны быть стены (3,2), (3,4), (2,3), (4,3)
	// Но так как лабиринт окружен стенами, некоторые из них могут быть недоступны

	expectedWalls3 := []domain.Point{
		{X: 2, Y: 3}, // Верхняя стена
		{X: 2, Y: 5}, // Нижняя стена
		{X: 1, Y: 4}, // Левая стена
		{X: 3, Y: 4}, // Правая стена
	}

	// Проверяем, что все ожидаемые стены добавлены
	for _, expectedWall := range expectedWalls3 {
		// Так как лабиринт окружен стенами, некоторые координаты могут быть недоступны
		// Проверяем только те, которые находятся в границах
		if gen.isInBounds(expectedWall.X, expectedWall.Y, maze) {
			assert.True(t, walls[expectedWall], "Стена %v должна быть добавлена", expectedWall)
		}
	}

}

func TestPrimaGenerator_GetConnectedWalls(t *testing.T) {
	gen := NewPrimaGenerator()
	maze := domain.NewMaze(5, 5)

	// Создаем тестовую конфигурацию стен
	// Устанавливаем ячейку (3,3) как пустую (путь)
	maze.Grid[3][3] = domain.CellEmpty

	// Создаем пустую карту стен
	walls := make(map[domain.Point]bool)

	// Тестируем стену сверху от ячейки (3,3)
	currentWall := domain.Point{X: 3, Y: 2}
	connected := gen.getConnectedWalls(walls, currentWall, maze)

	// Проверяем, что функция правильно определила соединенные ячейки
	// Первая ячейка должна быть стеной (3,2), вторая - путем (3,3) или наоборот
	assert.True(t,
		(connected[0].X == 3 && connected[0].Y == 3 && connected[1].X == 3 && connected[1].Y == 1),
		"Функция должна правильно определить соединенные ячейки")

	// Тестируем стену снизу от ячейки (3,3)
	currentWall = domain.Point{X: 3, Y: 4}
	connected = gen.getConnectedWalls(walls, currentWall, maze)

	// Проверяем, что функция правильно определила соединенные ячейки
	assert.True(t,
		(connected[0].X == 3 && connected[0].Y == 3 && connected[1].X == 3 && connected[1].Y == 5),
		"Функция должна правильно определить соединенные ячейки")

	// Тестируем стену слева от ячейки (3,3)
	currentWall = domain.Point{X: 2, Y: 3}
	connected = gen.getConnectedWalls(walls, currentWall, maze)

	// Проверяем, что функция правильно определила соединенные ячейки
	assert.True(t,
		(connected[0].X == 3 && connected[0].Y == 3 && connected[1].X == 1 && connected[1].Y == 3),
		"Функция должна правильно определить соединенные ячейки")

	// Тестируем стену справа от ячейки (3,3)
	currentWall = domain.Point{X: 4, Y: 3}
	connected = gen.getConnectedWalls(walls, currentWall, maze)

	// Проверяем, что функция правильно определила соединенные ячейки
	assert.True(t,
		(connected[0].X == 3 && connected[0].Y == 3 && connected[1].X == 5 && connected[1].Y == 3),
		"Функция должна правильно определить соединенные ячейки")
}

func TestPrimaGenerator_GetRandomWall(t *testing.T) {
	gen := NewPrimaGenerator()

	// Создаем тестовый набор стен
	walls := make(map[domain.Point]bool)

	// Добавляем несколько стен
	walls[domain.Point{X: 1, Y: 1}] = true
	walls[domain.Point{X: 2, Y: 2}] = true
	walls[domain.Point{X: 3, Y: 3}] = true
	walls[domain.Point{X: 4, Y: 4}] = true

	selectedWalls := make(map[domain.Point]int)

	// Вызываем функцию несколько раз и проверяем, что возвращаются допустимые значения
	for i := 0; i < 100; i++ {
		randomWall := gen.getRandomWall(walls)

		// Проверяем, что возвращенная стена содержится в исходном наборе
		assert.True(t, walls[randomWall], "Возвращенная стена должна содержаться в исходном наборе")

		// Проверяем, что координаты находятся в допустимом диапазоне
		assert.True(t, randomWall.X >= 1 && randomWall.X <= 4, "Координата X должна быть в диапазоне от 1 до 4")
		assert.True(t, randomWall.Y >= 1 && randomWall.Y <= 4, "Координата Y должна быть в диапазоне от 1 до 4")

		selectedWalls[randomWall]++
	}

	// Проверяем, что все стены были выбраны хотя бы раз
	assert.Equal(t, len(walls), len(selectedWalls), "Все стены должны быть выбраны хотя бы раз")

	// Проверяем, что каждая стена была выбрана
	for wall := range walls {
		assert.True(t, selectedWalls[wall] > 0, "Стена %v должна быть выбрана хотя бы раз", wall)
	}
}

// TestPrimaGenerator_IsInBounds тестирует корректность проверки
// точки на принадлежность области лабиринта
func TestPrimaGenerator_IsInBounds(t *testing.T) {
	gen := NewPrimaGenerator()
	maze, _ := gen.Generate(5, 5)

	// Тест внутри границ
	assert.True(t, gen.isInBounds(1, 1, maze), "Точка (1,1) должна быть в границах")
	assert.True(t, gen.isInBounds(5, 5, maze), "Точка (5,5) должна быть в границах")

	// Тест вне границ
	assert.False(t, gen.isInBounds(0, 1, maze), "Точка (0,1) должна быть вне границ")
	assert.False(t, gen.isInBounds(6, 1, maze), "Точка (6,1) должна быть вне границ")
	assert.False(t, gen.isInBounds(1, 0, maze), "Точка (1,0) должна быть вне границ")
	assert.False(t, gen.isInBounds(1, 6, maze), "Точка (1,6) должна быть вне границ")
}
