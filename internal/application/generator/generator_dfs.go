package generator

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type DFSGenerator struct {
	*BaseGenerator
}

func NewDFSGenerator() *DFSGenerator {
	return &DFSGenerator{
		BaseGenerator: NewBaseGenerator(AlgorithmDFS),
	}
}

func (g *DFSGenerator) Generate(width, height int) (*domain.Maze, error) {
	maze, err := g.BaseGenerator.Generate(width, height)
	if err != nil {
		return nil, err
	}
	return maze, nil
}

// dfsGenerate реализует рекурсивный DFS алгоритм для генерации лабиринта
func (g *BaseGenerator) dfsGenerate(maze *domain.Maze, start domain.Point) {
	stack := make([]domain.Point, 0)

	maze.Grid[start.Y][start.X] = domain.CellEmpty
	stack = append(stack, start)

	directions := []struct{ dx, dy int }{
		{0, -2},
		{2, 0},
		{0, 2},
		{-2, 0},
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := current.X, current.Y

		g.shuffleDirections(directions)

		for _, dir := range directions {
			nextX, nextY := x+dir.dx, y+dir.dy

			if nextX >= 1 && nextX < maze.Width+1 && nextY >= 1 && nextY < maze.Height+1 && maze.Grid[nextY][nextX] == domain.CellWall {
				maze.Grid[y+dir.dy/2][x+dir.dx/2] = domain.CellEmpty
				maze.Grid[nextY][nextX] = domain.CellEmpty
				stack = append(stack, domain.Point{X: nextX, Y: nextY})
			}
		}
	}
}

// shuffleDirections случайным образом перемешивает массив направлений
func (g *BaseGenerator) shuffleDirections(directions []struct{ dx, dy int }) {
	for i := len(directions) - 1; i > 0; i-- {
		j := g.rand.Intn(i + 1)
		directions[i], directions[j] = directions[j], directions[i]
	}
}
