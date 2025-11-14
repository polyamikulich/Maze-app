package generator

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"

func (g *BaseGenerator) addSpecialSurfaces(maze *domain.Maze) {
	wallsCells := g.findWallCells(maze)
	g.destroyWalls(maze, wallsCells, 0.08)

	emptyCells := g.findEmptyCells(maze)

	if len(emptyCells) < 8 {
		return
	}

	g.placeSurfaces(maze, emptyCells, domain.CellSand, 0.09)
	g.placeSurfaces(maze, emptyCells, domain.CellSmooth, 0.09)
	g.placeSurfaces(maze, emptyCells, domain.CellSwamp, 0.07)
	g.placeSurfaces(maze, emptyCells, domain.CellCoin, 0.07)
}

// findEmptyCells находит все пустые клетки в лабиринте
func (g *BaseGenerator) findEmptyCells(maze *domain.Maze) []domain.Point {
	emptyCells := make([]domain.Point, 0)

	for i := 1; i < maze.Height+1; i++ {
		for j := 1; j < maze.Width+1; j++ {
			if maze.Grid[i][j] == domain.CellEmpty {
				emptyCells = append(emptyCells, domain.Point{X: j, Y: i})
			}
		}
	}

	return emptyCells
}

// placeSurfaces размещает поверхности определенного типа
func (g *BaseGenerator) placeSurfaces(maze *domain.Maze, emptyCells []domain.Point, surfaceType domain.CellType, probability float64) {
	count := int(float64(len(emptyCells)) * probability)

	shuffled := make([]domain.Point, len(emptyCells))
	copy(shuffled, emptyCells)
	g.shufflePoints(shuffled)

	// Размещаем поверхности
	placed := 0
	for i := 0; i < len(shuffled) && placed < count; i++ {
		point := shuffled[i]

		// Проверяем что клетка всё ещё пустая (не была заменена другим типом)
		if maze.Grid[point.Y][point.X] == domain.CellEmpty {
			maze.Grid[point.Y][point.X] = surfaceType
			placed++
		}
	}
}

func (g *BaseGenerator) destroyWalls(maze *domain.Maze, wallsCells []domain.Point, probability float64) {
	count := int(float64(len(wallsCells)) * probability)

	shuffled := make([]domain.Point, len(wallsCells))
	copy(shuffled, wallsCells)
	g.shufflePoints(shuffled)

	// Размещаем поверхности
	placed := 0
	for i := 0; i < len(shuffled) && placed < count; i++ {
		point := shuffled[i]

		// Проверяем что клетка всё ещё пустая (не была заменена другим типом)
		if maze.Grid[point.Y][point.X] == domain.CellWall {
			maze.Grid[point.Y][point.X] = domain.CellEmpty
			placed++
		}
	}
}

// shufflePoints перемешивает массив точек
func (g *BaseGenerator) shufflePoints(points []domain.Point) {
	for i := len(points) - 1; i > 0; i-- {
		j := g.rand.Intn(i + 1)
		points[i], points[j] = points[j], points[i]
	}
}

// findEmptyCells находит все клетки стены в лабиринте
func (g *BaseGenerator) findWallCells(maze *domain.Maze) []domain.Point {
	wallCells := make([]domain.Point, 0)

	for i := 1; i < maze.Height+1; i++ {
		for j := 1; j < maze.Width+1; j++ {
			if maze.Grid[i][j] == domain.CellWall {
				wallCells = append(wallCells, domain.Point{X: j, Y: i})
			}
		}
	}

	return wallCells
}
