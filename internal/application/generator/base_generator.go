package generator

import (
	"math/rand"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type AlgorithmType int

type BaseGenerator struct {
	rand      *rand.Rand
	algorithm AlgorithmType
}

const (
	AlgorithmDFS AlgorithmType = iota
	AlgorithmPrim
)

func NewBaseGenerator(algorithm AlgorithmType) *BaseGenerator {
	return &BaseGenerator{
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		algorithm: algorithm,
	}
}

func (g *BaseGenerator) Generate(width, height int) (*domain.Maze, error) {
	maze := domain.NewMaze(width, height)
	startX := 1
	startY := 1
	startPoint := domain.Point{X: startX, Y: startY}

	switch g.algorithm {
	case AlgorithmDFS:
		g.dfsGenerate(maze, startPoint)
	case AlgorithmPrim:
		g.primaGenerate(maze, startPoint)
	}

	return maze, nil
}

func (g *BaseGenerator) GenerateEnhanced(width, height int) (*domain.Maze, error) {
	maze := domain.NewMaze(width, height)

	startX := 1
	startY := 1

	startPoint := domain.Point{X: startX, Y: startY}

	switch g.algorithm {
	case AlgorithmDFS:
		g.dfsGenerate(maze, startPoint)
	case AlgorithmPrim:
		g.primaGenerate(maze, startPoint)
	}

	g.addSpecialSurfaces(maze)

	return maze, nil
}
