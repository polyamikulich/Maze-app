package ui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestSaveMazeToFile_Correct(t *testing.T) {
	maze := domain.NewMaze(3, 3)
	maze.Grid[1][1] = domain.CellStart
	maze.Grid[1][2] = domain.CellPath
	maze.Grid[1][3] = domain.CellPath
	maze.Grid[2][3] = domain.CellPath
	maze.Grid[3][3] = domain.CellEnd

	tempFile := "test_maze.txt"
	defer os.Remove(tempFile)

	err := SaveMazeToFile(tempFile, maze)
	require.NoError(t, err)

	_, err = os.Stat(tempFile)
	require.NoError(t, err)

	expectedContent := `#####
#O..#
###.#
###X#
#####
`
	content, err := os.ReadFile(tempFile)
	require.NoError(t, err)
	require.Equal(t, expectedContent, string(content))
}

func TestSaveMazeToFile_InvalidPath(t *testing.T) {
	maze := domain.NewMaze(2, 2)

	err := SaveMazeToFile("/invalid/path/test.txt", maze)
	assert.Error(t, err)
}

func TestCellTypeToChar(t *testing.T) {
	tests := []struct {
		name     string
		cellType domain.CellType
		expected rune
	}{
		{
			name:     "CellWall",
			cellType: domain.CellWall,
			expected: '#',
		},
		{
			name:     "CellEmpty",
			cellType: domain.CellEmpty,
			expected: ' ',
		},
		{
			name:     "CellPath",
			cellType: domain.CellPath,
			expected: '.',
		},
		{
			name:     "CellStart",
			cellType: domain.CellStart,
			expected: 'O',
		},
		{
			name:     "CellEnd",
			cellType: domain.CellEnd,
			expected: 'X',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cellTypeToChar(tt.cellType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSaveMazeToFile_WithoutPath_Unicode(t *testing.T) {
	maze := domain.NewMaze(3, 5)
	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellEmpty

	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty

	maze.Grid[4][1] = domain.CellEmpty
	maze.Grid[4][3] = domain.CellEmpty

	maze.Grid[5][1] = domain.CellEmpty
	maze.Grid[5][3] = domain.CellEmpty

	tempFile := "test_maze.txt"
	defer os.Remove(tempFile)

	err := SaveMazeToFileUnicode(tempFile, maze, false)
	require.NoError(t, err)

	_, err = os.Stat(tempFile)
	require.NoError(t, err)

	//content := `#####
	//#   #
	//# ###
	//#   #
	//# # #
	//# # #
	//#####
	//`

	expectedContent := `┌───┐
│ ╶─┤
│ ╷ │
└─┴─┘
`
	content, err := os.ReadFile(tempFile)
	require.NoError(t, err)
	require.Equal(t, expectedContent, string(content))
}

func TestSaveMazeToFile_WithPath_Unicode(t *testing.T) {
	maze := domain.NewMaze(3, 5)
	maze.Grid[1][1] = domain.CellStart
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellPath

	maze.Grid[3][1] = domain.CellPath
	maze.Grid[3][2] = domain.CellPath
	maze.Grid[3][3] = domain.CellPath

	maze.Grid[4][1] = domain.CellEmpty
	maze.Grid[4][3] = domain.CellPath

	maze.Grid[5][1] = domain.CellEmpty
	maze.Grid[5][3] = domain.CellEnd

	tempFile := "test_maze.txt"
	defer os.Remove(tempFile)

	err := SaveMazeToFileUnicode(tempFile, maze, true)
	require.NoError(t, err)

	_, err = os.Stat(tempFile)
	require.NoError(t, err)

	// content := `#####
	//#O  #
	//#.###
	//#...#
	//# #.#
	//# #X#
	//#####`

	expectedContent := `┌───┐
│O  │
│.╶─┤
│...│
│ ╷.│
│ │X│
└─┴─┘
`
	content, err := os.ReadFile(tempFile)
	require.NoError(t, err)
	require.Equal(t, expectedContent, string(content))
}

func TestSaveEnhMazeToFile_WithPath_Unicode(t *testing.T) {
	maze := domain.NewMaze(3, 5)
	maze.Grid[1][1] = domain.CellStart
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellCoin

	maze.Grid[2][1] = domain.CellPath

	maze.Grid[3][1] = domain.CellPath
	maze.Grid[3][2] = domain.CellPath
	maze.Grid[3][3] = domain.CellPath

	maze.Grid[4][1] = domain.CellSmooth
	maze.Grid[4][3] = domain.CellPath

	maze.Grid[5][1] = domain.CellEmpty
	maze.Grid[5][3] = domain.CellEnd

	tempFile := "test_maze.txt"
	defer os.Remove(tempFile)

	err := SaveMazeToFileUnicode(tempFile, maze, true)
	require.NoError(t, err)

	_, err = os.Stat(tempFile)
	require.NoError(t, err)

	// content := `#####
	//#O $#
	//#.###
	//#...#
	//#=#.#
	//# #X#
	//#####`

	expectedContent := `┌───┐
│O ○│
│.╶─┤
│...│
│□╷.│
│ │X│
└─┴─┘
`

	content, err := os.ReadFile(tempFile)
	require.NoError(t, err)
	require.Equal(t, expectedContent, string(content))
}
