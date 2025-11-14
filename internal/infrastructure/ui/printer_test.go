package ui

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestPrinter_MazeWithoutPath(t *testing.T) {
	content := `#######
# #   #
# # # #
#   # #
#######
`

	maze := domain.NewMaze(5, 3)
	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty
	maze.Grid[1][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty
	maze.Grid[2][5] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellEmpty

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMaze(maze)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Equal(t, content, string(out))
}

func TestPrinter_MazeWithPath(t *testing.T) {
	content := `#######
#O# # #
#.#X# #
#...  #
#######
`

	maze := domain.NewMaze(5, 3)
	maze.Grid[1][1] = domain.CellStart
	maze.Grid[2][1] = domain.CellPath
	maze.Grid[3][1] = domain.CellPath
	maze.Grid[3][2] = domain.CellPath
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEnd
	maze.Grid[3][3] = domain.CellPath
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty
	maze.Grid[2][5] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellEmpty

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMaze(maze)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Equal(t, content, string(out))
}

func TestPrinter_MazeWithoutPath_Unicode(t *testing.T) {
	//	content := `#######
	//#     #
	//# # ###
	//# #   #
	//#######
	//`

	expectedContent := `┌─────┐
│ ╷ ╶─┤
└─┴───┘
`

	maze := domain.NewMaze(5, 3)
	maze.Grid[1][1] = domain.CellEmpty
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[1][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[2][3] = domain.CellEmpty

	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellEmpty

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMazeUnicode(maze, false)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Equal(t, expectedContent, string(out))
}

func TestPrinter_MazeWithPath_Unicode(t *testing.T) {
	//	content := `#######
	//#O  # #
	//#.###X#
	//#.....#
	//#######
	//`

	expectedContent := `┌───┬─┐
│O  │ │
│.╶─┘X│
│.....│
└─────┘
`

	maze := domain.NewMaze(5, 3)
	maze.Grid[1][1] = domain.CellStart
	maze.Grid[2][1] = domain.CellPath
	maze.Grid[3][1] = domain.CellPath

	maze.Grid[3][2] = domain.CellPath
	maze.Grid[3][3] = domain.CellPath
	maze.Grid[3][4] = domain.CellPath
	maze.Grid[3][5] = domain.CellPath

	maze.Grid[2][5] = domain.CellEnd

	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMazeUnicode(maze, true)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Equal(t, expectedContent, string(out))
}

func TestPrinter_MazeEnhWithoutPath_Unicode(t *testing.T) {
	//content := `#########
	//#$  # # #
	//# ### # #
	//#    &  #
	//# #######
	//# ~=~ = #
	//#########`

	expectedContent := `┌───┬─┬─┐
│○  │ │ │
│ ╶─┘ ╵ │
│    ▓  │
│ ╶─────┤
│ ░□░ □ │
└───────┘
`

	maze := domain.NewMaze(7, 5)
	maze.Grid[1][1] = domain.CellCoin
	maze.Grid[1][2] = domain.CellEmpty
	maze.Grid[1][3] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty
	maze.Grid[1][7] = domain.CellEmpty

	maze.Grid[2][1] = domain.CellEmpty
	maze.Grid[2][5] = domain.CellEmpty
	maze.Grid[2][7] = domain.CellEmpty

	maze.Grid[3][1] = domain.CellEmpty
	maze.Grid[3][2] = domain.CellEmpty
	maze.Grid[3][3] = domain.CellEmpty
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[3][5] = domain.CellSwamp
	maze.Grid[3][6] = domain.CellEmpty
	maze.Grid[3][7] = domain.CellEmpty

	maze.Grid[4][1] = domain.CellEmpty

	maze.Grid[5][1] = domain.CellEmpty
	maze.Grid[5][2] = domain.CellSand
	maze.Grid[5][3] = domain.CellSmooth
	maze.Grid[5][4] = domain.CellSand
	maze.Grid[5][5] = domain.CellEmpty
	maze.Grid[5][6] = domain.CellSmooth
	maze.Grid[5][7] = domain.CellEmpty

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMazeUnicode(maze, true)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Equal(t, expectedContent, string(out))
}

func TestPrinter_MazeEnhWithPath(t *testing.T) {
	content := `#######
#O#$# #
#.#X#~#
#... =#
#######
`

	maze := domain.NewMaze(5, 3)
	maze.Grid[1][1] = domain.CellStart
	maze.Grid[2][1] = domain.CellPath
	maze.Grid[3][1] = domain.CellPath
	maze.Grid[3][2] = domain.CellPath
	maze.Grid[1][3] = domain.CellCoin
	maze.Grid[2][3] = domain.CellEnd
	maze.Grid[3][3] = domain.CellPath
	maze.Grid[3][4] = domain.CellEmpty
	maze.Grid[1][5] = domain.CellEmpty
	maze.Grid[2][5] = domain.CellSand
	maze.Grid[3][5] = domain.CellSmooth

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMaze(maze)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.Equal(t, content, string(out))
}
