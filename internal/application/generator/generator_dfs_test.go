package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestDFSGenerator_NewDFSGenerator(t *testing.T) {
	gen := NewDFSGenerator()
	assert.NotNil(t, gen, "NewDFSGenerator() должен возвращать непустой генератор")
}

// TestDFSGenerator_Generate тестирует генерацию лабиринта с помощью DFS
func TestDFSGenerator_Generate(t *testing.T) {
	gen := NewDFSGenerator()

	// Тестируем генерацию лабиринта 10x10
	maze, err := gen.Generate(10, 10)
	assert.NoError(t, err)
	assert.NotNil(t, maze)
	assert.Equal(t, 10, maze.Width)
	assert.Equal(t, 10, maze.Height)
	assert.NotNil(t, maze.Grid)
	assert.Equal(t, 12, len(maze.Grid))

	// Проверяем, что каждая строка имеет правильную длину
	for _, row := range maze.Grid {
		assert.Equal(t, 12, len(row))
	}

	// Проверяем, что границы лабиринта являются стенами
	for i := 0; i < 10; i++ {
		assert.Equal(t, domain.CellWall, maze.Grid[0][i])
		assert.Equal(t, domain.CellWall, maze.Grid[11][i])
		assert.Equal(t, domain.CellWall, maze.Grid[i][0])
		assert.Equal(t, domain.CellWall, maze.Grid[i][11])
	}
}

// TestDFSGenerator_ShuffleDirections тестирует метод shuffleDirections
func TestDFSGenerator_ShuffleDirections(t *testing.T) {
	gen := NewDFSGenerator()

	// Создаем фиксированный генератор случайных чисел для воспроизводимости теста
	gen.rand.Seed(12345)

	// Создаем направления
	directions := []struct{ dx, dy int }{
		{0, -2},
		{2, 0},
		{0, 2},
		{-2, 0},
	}

	// Копируем оригинальный массив
	original := make([]struct{ dx, dy int }, len(directions))
	copy(original, directions)

	// Перемешиваем
	gen.shuffleDirections(directions)

	// Проверяем, что массив изменился
	changed := false
	for i := range directions {
		if directions[i] != original[i] {
			changed = true
			break
		}
	}

	assert.True(t, changed, "Направления должны быть перемешаны")

	// Проверяем, что все элементы остались
	for _, orig := range original {
		found := false
		for _, dir := range directions {
			if dir == orig {
				found = true
				break
			}
		}
		assert.True(t, found, "Все оригинальные элементы должны присутствовать после перемешивания")
	}
}
