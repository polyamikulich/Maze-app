package generator

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"

type PrimEnhancedGenerator struct {
	*BaseGenerator
}

func NewPrimEnhancedGenerator() *PrimEnhancedGenerator {
	return &PrimEnhancedGenerator{
		BaseGenerator: NewBaseGenerator(AlgorithmPrim),
	}
}

func (g *PrimEnhancedGenerator) Generate(width, height int) (*domain.Maze, error) {
	maze, err := g.BaseGenerator.GenerateEnhanced(width, height)
	if err != nil {
		return nil, err
	}

	return maze, nil
}
