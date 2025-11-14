package generator

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type DFSEnhancedGenerator struct {
	*BaseGenerator
}

func NewDFSEnhancedGenerator() *DFSEnhancedGenerator {
	return &DFSEnhancedGenerator{
		BaseGenerator: NewBaseGenerator(AlgorithmDFS),
	}
}

func (g *DFSEnhancedGenerator) Generate(width, height int) (*domain.Maze, error) {
	maze, err := g.BaseGenerator.GenerateEnhanced(width, height)
	if err != nil {
		return nil, err
	}

	return maze, nil
}
