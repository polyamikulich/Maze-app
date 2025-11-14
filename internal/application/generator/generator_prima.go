package generator

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type PrimaGenerator struct {
	*BaseGenerator
}

func NewPrimaGenerator() *PrimaGenerator {
	return &PrimaGenerator{
		BaseGenerator: NewBaseGenerator(AlgorithmPrim),
	}
}

func (g *PrimaGenerator) Generate(width, height int) (*domain.Maze, error) {
	maze, err := g.BaseGenerator.Generate(width, height)
	if err != nil {
		return nil, err
	}
	return maze, nil
}

func (g *BaseGenerator) primaGenerate(maze *domain.Maze, start domain.Point) {
	walls := make(map[domain.Point]bool)

	maze.Grid[start.Y][start.X] = domain.CellEmpty

	g.addWalls(maze, start, walls)

	for len(walls) > 0 {
		currentWall := g.getRandomWall(walls)

		delete(walls, currentWall)

		if maze.Grid[currentWall.Y][currentWall.X] != domain.CellWall {
			continue
		}

		connectedWalls := g.getConnectedWalls(walls, currentWall, maze)

		if (connectedWalls[0].X == -1 && connectedWalls[0].Y == -1) || (connectedWalls[1].X == -1 && connectedWalls[1].Y == -1) {
			continue
		}

		maze.Grid[currentWall.Y][currentWall.X] = domain.CellEmpty
		maze.Grid[connectedWalls[1].Y][connectedWalls[1].X] = domain.CellEmpty

		g.addWalls(maze, connectedWalls[1], walls)
	}
}

func (g *BaseGenerator) addWalls(maze *domain.Maze, point domain.Point, walls map[domain.Point]bool) {
	x, y := point.X, point.Y

	neighborsDir := []struct{ dx, dy int }{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for _, dir := range neighborsDir {
		nextX, nextY := x+dir.dx, y+dir.dy

		if g.isInBounds(nextX, nextY, maze) && maze.Grid[nextY][nextX] == domain.CellWall {
			walls[domain.Point{X: nextX, Y: nextY}] = true
		}
	}
}

func (g *BaseGenerator) getRandomWall(walls map[domain.Point]bool) domain.Point {
	wallKeys := make([]domain.Point, 0, len(walls))
	for wallKey := range walls {
		wallKeys = append(wallKeys, wallKey)
	}

	randomIndex := g.rand.Intn(len(wallKeys))
	currentWall := wallKeys[randomIndex]

	return currentWall
}

func (g *BaseGenerator) getConnectedWalls(walls map[domain.Point]bool, currentWall domain.Point, maze *domain.Maze) [2]domain.Point {
	var connectedWalls [2]domain.Point
	existingPath := domain.Point{X: -1, Y: -1}
	newWall := domain.Point{X: -1, Y: -1}

	neighborsDir := []struct{ dx, dy int }{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for i := 0; i < 4; i += 2 {
		prevX, prevY := currentWall.X+neighborsDir[i].dx, currentWall.Y+neighborsDir[i].dy
		nextX, nextY := currentWall.X+neighborsDir[i+1].dx, currentWall.Y+neighborsDir[i+1].dy
		if g.isInBounds(prevX, prevY, maze) && g.isInBounds(nextX, nextY, maze) {
			if maze.Grid[prevY][prevX] == domain.CellWall && maze.Grid[nextY][nextX] == domain.CellEmpty {
				existingPath = domain.Point{X: nextX, Y: nextY}
				newWall = domain.Point{X: prevX, Y: prevY}
				break
			}

			if maze.Grid[prevY][prevX] == domain.CellEmpty && maze.Grid[nextY][nextX] == domain.CellWall {
				existingPath = domain.Point{X: prevX, Y: prevY}
				newWall = domain.Point{X: nextX, Y: nextY}
				break
			}
		}
	}

	connectedWalls[0] = existingPath
	connectedWalls[1] = newWall

	return connectedWalls
}

func (g *BaseGenerator) isInBounds(x, y int, maze *domain.Maze) bool {
	return x >= 1 && x < maze.Width+1 && y >= 1 && y < maze.Height+1
}
