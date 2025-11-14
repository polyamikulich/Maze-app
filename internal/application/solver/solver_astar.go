package solver

import (
	"container/heap"
	"errors"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type ItemAstar struct {
	point  domain.Point
	fScore int
	index  int
}

// Приоритетная очередь для точек
// На каждом шаге будем доставать оттуда точку с наименьшим fScore
// Реализуем при помощи "container/heap"
type PriorityQueueAstar []*ItemAstar

func (pq PriorityQueueAstar) Len() int {
	return len(pq)
}

func (pq PriorityQueueAstar) Less(i, j int) bool {
	return pq[i].fScore < pq[j].fScore
}

func (pq PriorityQueueAstar) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueAstar) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ItemAstar)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueAstar) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

type AStarSolver struct {
}

func NewAStarSolver() *AStarSolver {
	return &AStarSolver{}
}

func (a *AStarSolver) Solve(maze *domain.Maze, start, end domain.Point) (*domain.Path, error) {
	// Создаём все необходимые массивы
	gScore := make([][]int, maze.Height+2)
	prev := make([][]domain.Point, maze.Height+2)
	visited := make([][]bool, maze.Height+2)

	for i := range gScore {
		gScore[i] = make([]int, maze.Width+2)
		prev[i] = make([]domain.Point, maze.Width+2)
		visited[i] = make([]bool, maze.Width+2)
		for j := range gScore[i] {
			gScore[i][j] = math.MaxInt32
		}
	}

	gScore[start.Y][start.X] = 0

	pq := make(PriorityQueueAstar, 0)

	// Проверяем простой лабиринт или нет
	is_enhanced := a.isEnhancedMaze(maze)

	// Запускаем алгоритм
	a.aStarSolve(maze, start, end, gScore, prev, visited, pq, is_enhanced)

	if gScore[end.Y][end.X] == math.MaxInt32 {
		return nil, errors.New("no path found")
	}

	// Восстанавливаем путь
	path := a.reconstructPath(prev, start, end)

	return &domain.Path{Points: path}, nil
}

func (a *AStarSolver) aStarSolve(maze *domain.Maze, start, end domain.Point, gScore [][]int, prev [][]domain.Point, visited [][]bool, pq PriorityQueueAstar, is_enh bool) {
	// Создаем приоритетную очередь
	heap.Init(&pq)

	// Смотрим начальную эвристику от старта до конца
	initialFScore := a.heuristic(start, end, is_enh)

	// Добавляем в приоритетную очередь начальную точку
	heap.Push(&pq, &ItemAstar{
		point:  start,
		fScore: initialFScore,
	})

	for pq.Len() > 0 {
		// Извлекаем с точку с наименьшим расстоянием.
		// Сортируется автоматически
		current_item := heap.Pop(&pq).(*ItemAstar)
		current := current_item.point

		if current == end {
			break
		}

		if visited[current.Y][current.X] {
			continue
		}

		visited[current.Y][current.X] = true

		neighborsDir := []struct{ dx, dy int }{
			{0, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
		}

		// Смотрим по соседям
		for _, dir := range neighborsDir {
			nextX, nextY := current.X+dir.dx, current.Y+dir.dy
			if !a.isInBounds(nextX, nextY, maze) || visited[nextY][nextX] {
				continue
			}

			// Получаем вес следующей клетки
			weight := domain.GetCellWeight(maze.Grid[nextY][nextX])
			tentativeGScore := gScore[current.Y][current.X] + weight

			// Если улучшилось значение, то обновляем
			if tentativeGScore < gScore[nextY][nextX] {
				gScore[nextY][nextX] = tentativeGScore
				prev[nextY][nextX] = current

				nextPoint := domain.Point{X: nextX, Y: nextY}
				nextFScore := tentativeGScore + a.heuristic(nextPoint, end, is_enh)

				heap.Push(&pq, &ItemAstar{
					point:  nextPoint,
					fScore: nextFScore,
				})
			}
		}
	}

}

// В зависимости от того, улучшенный лабиринт или нет, определяем эвристическую функцию
func (a *AStarSolver) heuristic(aPoint, bPoint domain.Point, isEnh bool) int {
	manhettenDist := int(math.Abs(float64(aPoint.X-bPoint.X)) + math.Abs(float64(aPoint.Y-bPoint.Y)))
	var weight int
	if !isEnh {
		// Для обычного лабиринта вес каждой клетки = 10 (вес пустой клетки (см. domain/cell.go))
		weight = domain.GetCellWeight(domain.CellEmpty)
	} else {
		// Для enhanced лабиринта вес эвристики = 5 (вес самой "быстрой" клетки (см. domain/cell.go))
		weight = domain.GetCellWeight(domain.CellSmooth)
	}

	return manhettenDist * weight
}

// reconstructPath возвращает путь, восстанавливая его по массиву prev
func (a *AStarSolver) reconstructPath(prev [][]domain.Point, start, end domain.Point) []domain.Point {
	path := make([]domain.Point, 0)
	for current := end; current != start; current = prev[current.Y][current.X] {
		path = append(path, current)
	}

	path = append(path, start)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (a AStarSolver) isInBounds(x, y int, maze *domain.Maze) bool {
	if !(x >= 1 && x < maze.Width+1 && y >= 1 && y < maze.Height+1) {
		return false
	}

	return maze.Grid[y][x] != domain.CellWall
}

// isEnhancedMaze определяет является ли лабиринт улучшенным (от этого зависит эвристика)
func (a AStarSolver) isEnhancedMaze(maze *domain.Maze) bool {
	for y := 1; y < maze.Height+1; y++ {
		for x := 1; x < maze.Width+1; x++ {
			cellType := maze.Grid[y][x]
			// Если находим любую специальную поверхность кроме обычных CellEmpty/CellWall
			if cellType == domain.CellSand || cellType == domain.CellSwamp ||
				cellType == domain.CellCoin || cellType == domain.CellSmooth {
				return true
			}
		}
	}
	return false
}
