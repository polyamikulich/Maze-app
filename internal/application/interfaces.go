package application

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type Generator interface {
	Generate(width, height int) (*domain.Maze, error)
}

type Solver interface {
	Solve(maze *domain.Maze, start, end domain.Point) (*domain.Path, error)
}
