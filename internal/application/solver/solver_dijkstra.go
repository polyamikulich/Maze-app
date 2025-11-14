package solver

import (
	"container/heap"
	"errors"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

// Item представляет элемент очереди с приоритетом
type ItemDijkstra struct {
	point    domain.Point
	distance int
	index    int
}

// PriorityQueue реализует очередь с приоритетом на основе heap.Interface
type PriorityQueueDijkstra []*ItemDijkstra

func (pq PriorityQueueDijkstra) Len() int {
	return len(pq)
}

func (pq PriorityQueueDijkstra) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueueDijkstra) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueDijkstra) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ItemDijkstra)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueDijkstra) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

// DijkstraSolver реализует алгоритм Дейкстры для поиска пути в лабиринте
type DijkstraSolver struct{}

// NewDijkstraSolver создает новый экземпляр DijkstraSolver
func NewDijkstraSolver() *DijkstraSolver {
	return &DijkstraSolver{}
}

func (d *DijkstraSolver) Solve(maze *domain.Maze, start, end domain.Point) (*domain.Path, error) {
	dist := make([][]int, maze.Height+2)
	prev := make([][]domain.Point, maze.Height+2)
	visited := make([][]bool, maze.Height+2)

	for i := range dist {
		dist[i] = make([]int, maze.Width+2)
		prev[i] = make([]domain.Point, maze.Width+2)
		visited[i] = make([]bool, maze.Width+2)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[start.Y][start.X] = 0

	pq := make(PriorityQueueDijkstra, 0)

	d.dijkstraSolve(maze, start, end, dist, prev, visited, pq)

	if dist[end.Y][end.X] == math.MaxInt32 {
		return nil, errors.New("no path found")
	}

	path := d.reconstructPath(prev, start, end)

	return &domain.Path{Points: path}, nil
}

func (d *DijkstraSolver) dijkstraSolve(maze *domain.Maze, start, end domain.Point, dist [][]int, prev [][]domain.Point, visited [][]bool, pq PriorityQueueDijkstra) {
	heap.Init(&pq)
	heap.Push(&pq, &ItemDijkstra{
		point:    start,
		distance: 0,
	})

	for pq.Len() > 0 {
		current_item := heap.Pop(&pq).(*ItemDijkstra)
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

		for _, dir := range neighborsDir {
			nextX, nextY := current.X+dir.dx, current.Y+dir.dy
			if !d.isInBounds(nextX, nextY, maze) || visited[nextY][nextX] {
				continue
			}

			weight := domain.GetCellWeight(maze.Grid[nextY][nextX])
			newDistance := dist[current.Y][current.X] + weight
			if newDistance < dist[nextY][nextX] {
				dist[nextY][nextX] = newDistance
				prev[nextY][nextX] = current

				heap.Push(&pq, &ItemDijkstra{
					point:    domain.Point{X: nextX, Y: nextY},
					distance: newDistance,
				})
			}
		}
	}
}

func (d DijkstraSolver) reconstructPath(prev [][]domain.Point, start, end domain.Point) []domain.Point {
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

func (d DijkstraSolver) isInBounds(x, y int, maze *domain.Maze) bool {
	if !(x >= 1 && x < maze.Width+1 && y >= 1 && y < maze.Height+1) {
		return false
	}

	return maze.Grid[y][x] != domain.CellWall
}
