package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestAStarSolver_NewAStarSolver(t *testing.T) {
	solver := NewAStarSolver()
	assert.NotNil(t, solver, "NewAStarSolver() должен возвращать непустой решатель")
}

func TestAStarSolver_Solve(t *testing.T) {
	solver := NewAStarSolver()

	// Создаем простой лабиринт для тестирования
	// #####
	// #S  #
	// ### #
	// #  E#
	// #####
	maze := domain.NewMaze(3, 3)

	// Создаем проходы
	maze.Grid[1][1] = domain.CellStart // Start
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty
	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEnd // End

	start := domain.Point{X: 1, Y: 1}
	end := domain.Point{X: 3, Y: 3}

	path, err := solver.Solve(maze, start, end)

	assert.NoError(t, err, "Решатель не должен возвращать ошибку для достижимой цели")
	assert.NotNil(t, path, "Путь не должен быть nil")
	assert.Greater(t, len(path.Points), 0, "Путь должен содержать хотя бы одну точку")
}

func TestAStarSolver_Solve_EnhMaze(t *testing.T) {
	solver := NewAStarSolver()

	//#######
	//#S $  #
	//#&#=###
	//#$#=# #
	//# #~# #
	//#~E  =#
	//#######

	maze := domain.NewMaze(5, 5)

	// Создаем проходы
	maze.Grid[1][1] = domain.CellStart // Start
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellCoin
	maze.Grid[1][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellSwamp
	maze.Grid[2][3] = domain.CellSmooth

	maze.Grid[3][1] = domain.CellCoin
	maze.Grid[3][3] = domain.CellSmooth
	maze.Grid[3][5] = domain.CellEmpty

	maze.Grid[4][1] = domain.CellEmpty
	maze.Grid[4][3] = domain.CellSand
	maze.Grid[4][5] = domain.CellEmpty

	maze.Grid[5][1] = domain.CellSand
	maze.Grid[5][2] = domain.CellEnd
	maze.Grid[5][3] = domain.CellEmpty
	maze.Grid[5][4] = domain.CellEmpty
	maze.Grid[5][5] = domain.CellSmooth // End

	start := domain.Point{X: 1, Y: 1}
	end := domain.Point{X: 2, Y: 5}

	path, err := solver.Solve(maze, start, end)

	assert.NoError(t, err, "Решатель не должен возвращать ошибку для достижимой цели")
	assert.NotNil(t, path, "Путь не должен быть nil")
	assert.Greater(t, len(path.Points), 0, "Путь должен содержать хотя бы одну точку")
}

func TestAStarSolver_SolveUnreachable(t *testing.T) {
	solver := NewAStarSolver()

	// Создаем лабиринт с недостижимой точкой
	// #####
	// #S#E#
	// # # #
	// # # #
	// #####
	maze := domain.NewMaze(3, 3)

	// Создаем проходы
	maze.Grid[1][1] = domain.CellStart // Start
	maze.Grid[1][3] = domain.CellEnd   // End (недостижима)
	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty
	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty

	start := domain.Point{X: 1, Y: 1}
	end := domain.Point{X: 3, Y: 1}

	path, err := solver.Solve(maze, start, end)

	assert.Error(t, err, "Решатель должен возвращать ошибку для недостижимой цели")
	assert.Nil(t, path, "Путь должен быть nil для недостижимой цели")
}

func TestAStarSolver_IsInBounds(t *testing.T) {
	solver := NewAStarSolver()

	// Создаем тестовый лабиринт
	maze := domain.NewMaze(3, 3)
	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[2][2] = domain.CellWall

	// Тест внутри границ и не стена
	assert.True(t, solver.isInBounds(1, 1, maze), "Точка (1,1) должна быть в границах и не стеной")

	// Тест внутри границ, но стена
	assert.False(t, solver.isInBounds(2, 2, maze), "Точка (2,2) является стеной")

	// Тест вне границ
	assert.False(t, solver.isInBounds(0, 1, maze), "Точка (0,1) должна быть вне границ")
	assert.False(t, solver.isInBounds(4, 1, maze), "Точка (4,1) должна быть вне границ")
	assert.False(t, solver.isInBounds(1, 0, maze), "Точка (1,0) должна быть вне границ")
	assert.False(t, solver.isInBounds(1, 4, maze), "Точка (1,4) должна быть вне границ")
}

func TestAStarSolver_ReconstructPath(t *testing.T) {
	solver := NewAStarSolver()

	// Создаем простой лабиринт для тестирования
	// #####
	// #S  #
	// ### #
	// #  E#
	// #####
	// maze := domain.NewMaze(3, 3)

	// Создаем тестовые данные для восстановления пути
	prev := make([][]domain.Point, 5)
	for i := range prev {
		prev[i] = make([]domain.Point, 5)
	}

	start := domain.Point{X: 1, Y: 1}
	end := domain.Point{X: 3, Y: 3}

	prev[1][2] = start
	prev[1][3] = domain.Point{X: 2, Y: 1}
	prev[2][3] = domain.Point{X: 3, Y: 1}
	prev[3][3] = domain.Point{X: 3, Y: 2}

	path := solver.reconstructPath(prev, start, end)

	assert.Equal(t, []domain.Point{start, domain.Point{X: 2, Y: 1}, domain.Point{X: 3, Y: 1}, domain.Point{X: 3, Y: 2}, end}, path)
}

func TestAStarSolver_Heuristic(t *testing.T) {
	solver := NewAStarSolver()

	// Тест 1: точки совпадают
	point1 := domain.Point{X: 1, Y: 1}
	point2 := domain.Point{X: 1, Y: 1}
	distance := solver.heuristic(point1, point2, false)
	assert.Equal(t, 0, distance, "Расстояние между совпадающими точками должно быть 0")

	// Тест 2: точки на одной горизонтали
	point1 = domain.Point{X: 1, Y: 1}
	point2 = domain.Point{X: 4, Y: 1}
	distance = solver.heuristic(point1, point2, false)
	assert.Equal(t, 30, distance, "Расстояние по горизонтали должно быть 3")

	// Тест 3: точки на одной вертикали
	point1 = domain.Point{X: 1, Y: 1}
	point2 = domain.Point{X: 1, Y: 5}
	distance = solver.heuristic(point1, point2, false)
	assert.Equal(t, 40, distance, "Расстояние по вертикали должно быть 4")

	// Тест 4: точки в разных квадрантах
	point1 = domain.Point{X: 2, Y: 3}
	point2 = domain.Point{X: 5, Y: 7}
	distance = solver.heuristic(point1, point2, false)
	assert.Equal(t, 70, distance, "Манхэттенское расстояние должно быть 7 (3+4)")
}
