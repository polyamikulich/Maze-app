package domain

// Maze представляет лабиринт как двумерный массив ячеек
type Maze struct {
	Width  int
	Height int
	Grid   [][]CellType
}

// NewMaze создает новый лабиринт
// Считаем, что размер лабиринта не учитывает стены
// Поэтому при создании лабиринта мы добавляем к лабиринту по 1 стене с каждой стороны
func NewMaze(width, height int) *Maze {
	maze := &Maze{
		Width:  width,
		Height: height,
		Grid:   make([][]CellType, height+2),
	}

	for i := range maze.Grid {
		maze.Grid[i] = make([]CellType, width+2)
		for j := range maze.Grid[i] {
			maze.Grid[i][j] = CellWall
		}
	}

	return maze
}

// ApplyPath применяет путь к лабиринту, помечая ячейки пути как CellPath
func (m *Maze) ApplyPath(path *Path) {
	n := len(path.Points)
	start := path.Points[0]
	end := path.Points[n-1]

	// Помечаем стартовую и конечную точки
	m.Grid[start.Y][start.X] = CellStart
	m.Grid[end.Y][end.X] = CellEnd

	// Помечаем ячейки пути
	for _, point := range path.Points {
		// Пропускаем стартовую и конечную точки
		if (point == start) || (point == end) {
			continue
		}
		m.Grid[point.Y][point.X] = CellPath
	}
}
