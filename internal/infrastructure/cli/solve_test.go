package cli

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveCommand_HelpFlag(t *testing.T) {
	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"--help", "solve"})
	err := cmd.Execute()

	// Восстанавливаем stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.NoError(t, err)

	assert.Contains(t, string(out), "Solve a maze with specified algorithm and points")
	assert.Contains(t, string(out), "astar")
	assert.Contains(t, string(out), "dijkstra")
	assert.Contains(t, string(out), "-a, --algorithm")
	assert.Contains(t, string(out), "-e, --end")
	assert.Contains(t, string(out), "-s, --start")
	assert.Contains(t, string(out), "-f, --file")
	assert.Contains(t, string(out), "-o, --output")
	assert.Contains(t, string(out), "-u, --unicode")
}

func TestSolveCommand_InvalidAlgorithm(t *testing.T) {
	cmd := NewSolveCommand()

	require.NotNil(t, cmd)

	// Выполнение: запускаем команду с флагом --help
	cmd.SetArgs([]string{"-a", "invalid_algo"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Неподдерживаемый алгоритм: invalid_algo")
}

func TestSolveCommand_InvalidStartFinish(t *testing.T) {
	// Сначала создаем временный файл с лабиринтом
	tempFile := "test_maze.txt"
	defer os.Remove(tempFile) // Удаляем после теста

	// Генерируем лабиринт и сохраняем в файл
	generateCmd := NewGenerateCommand()
	generateCmd.SetArgs([]string{"-W", "14", "-H", "16", "-a", "dfs", "-o", tempFile})
	err := generateCmd.Execute()
	require.NoError(t, err, "Failed to generate maze for test")

	// Проверяем что файл создался
	_, err = os.Stat(tempFile)
	require.NoError(t, err, "Maze file was not created")

	tests1 := []struct {
		name     string
		start    string
		end      string
		expected string
	}{
		{
			name:     "invalid start point",
			start:    "invalid",
			end:      "1,1",
			expected: "Некорректный формат точки старта invalid: ожидается формат x,y",
		},
		{
			name:     "invalid end point",
			start:    "1,1",
			end:      "invalid",
			expected: "Некорректный формат точки финиша invalid: ожидается формат x,y",
		},
		{
			name:     "invalid start point only x",
			start:    "11",
			end:      "1,1",
			expected: "Некорректный формат точки старта 11: ожидается формат x,y",
		},
		{
			name:     "invalid end point only x",
			start:    "1,1",
			end:      "11",
			expected: "Некорректный формат точки финиша 11: ожидается формат x,y",
		},
		{
			name:     "invalid start point y",
			start:    "1,-8",
			end:      "10,5",
			expected: "Некорректные координаты точки старта: 1,-8",
		},
		{
			name:     "invalid end point y",
			start:    "10,5",
			end:      "1,-8",
			expected: "Некорректные координаты точки финиша: 1,-8",
		},
		{
			name:     "invalid start point x",
			start:    "0,6",
			end:      "1,5",
			expected: "Некорректные координаты точки старта: 0,6",
		},
		{
			name:     "invalid end point x",
			start:    "1,6",
			end:      "-7,5",
			expected: "Некорректные координаты точки финиша: -7,5",
		},
	}

	tests2 := []struct {
		name     string
		file     string
		start    string
		end      string
		expected string
	}{
		{
			name:     "invalid start point",
			start:    "16,1",
			end:      "1,1",
			file:     tempFile,
			expected: "Некорректные координаты точки старта: 16,1",
		},
		{
			name:     "invalid end point",
			start:    "1,1",
			end:      "1,19",
			file:     tempFile,
			expected: "Некорректные координаты точки финиша: 1,19",
		},
		{
			name:     "start point not empty",
			start:    "14,16",
			end:      "1,1",
			file:     tempFile,
			expected: "Указанная точка старта 14,16 не является пустой клеткой",
		},
		{
			name:     "end point not empty",
			start:    "1,1",
			end:      "14,16",
			file:     tempFile,
			expected: "Указанная точка финиша 14,16 не является пустой клеткой",
		},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewSolveCommand()
			require.NotNil(t, cmd)

			cmd.SetArgs([]string{"-f", tempFile, "-s", tt.start, "-e", tt.end})
			err := cmd.Execute()

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewSolveCommand()
			require.NotNil(t, cmd)

			cmd.SetArgs([]string{"-f", tempFile, "-s", tt.start, "-e", tt.end})
			err := cmd.Execute()

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}

func TestSolveCommand_WithoutFile(t *testing.T) {
	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-s", "1,1", "-e", "10,10"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Не указан файл с лабиринтом")
}

func TestSolveCommand_InvalidUnicode(t *testing.T) {
	cmd := NewSolveCommand()

	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-f", "text.txt", "-s", "1,1", "-e", "10,10", "--unicode", "34"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Недопустимое значение флага unicode: 34")
}

func TestSolveCommand_WithoutStartEndPoint(t *testing.T) {
	tests := []struct {
		name  string
		start string
		end   string
		file  string
	}{
		{
			name:  "without start point",
			start: "",
			end:   "12,12",
			file:  "test.txt", // Не будем создавать его даже для теста, потому что программа проверит наличие точек старта и финиша
			// раньше, чем попытается открыть файл
		},
		{
			name:  "without end point",
			start: "12,12",
			end:   "",
			file:  "test.txt",
		},
		{
			name:  "without start end point",
			start: "",
			end:   "",
			file:  "test.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewSolveCommand()
			require.NotNil(t, cmd)

			cmd.SetArgs([]string{"-s", tt.start, "-e", tt.end, "-f", tt.file})
			err := cmd.Execute()

			assert.Error(t, err)
			assert.Contains(t, err.Error(), "Не указаны точки старта и/или конца пути")
		})
	}
}

func TestSolveCommand_FileNotExists(t *testing.T) {
	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-s", "1,1", "-e", "1,1", "-f", "test.txt"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Ошибка чтения лабиринта из файла")
}

func TestSolveCommand_FileEmpty(t *testing.T) {
	filename := "empty.txt"
	content := ``

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-f", filename, "-s", "1,1", "-e", "3,3"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Ошибка чтения лабиринта из файла")
}

func TestSolveCommand_MazeNotCorrect_WithoutSymb(t *testing.T) {
	filename := "incorrect.txt"
	content := `#####
# # #
# # #
#   #
####`

	err := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, err)
	defer os.Remove(filename)

	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-f", filename, "-s", "1,1", "-e", "3,3"})
	err = cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Ошибка чтения лабиринта из файла")
}

func TestSolveCommand_MazeNotCorrect_BadSymb(t *testing.T) {
	filename := "incorrect.txt"
	content := `#####
# # #
# - #
#   #
#####`

	err := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, err)
	defer os.Remove(filename)

	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-f", filename, "-s", "1,1", "-e", "3,3"})
	err = cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Ошибка чтения лабиринта из файла")
}

func TestSolveCommand_SolveUnreachable(t *testing.T) {
	filename := "no_path.txt"
	content := `#####
# # #
# # #
# # #
#####`

	err := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, err)
	defer os.Remove(filename)

	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"-f", filename, "-s", "1,1", "-e", "3,3"})
	err = cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Путь от {1 1} до {3 3} не найден")
}

func TestSolveCommand_OutputConsoleCorrect(t *testing.T) {
	filename := "test_maze.txt"
	content := `#####
# # #
# # #
#   #
#####`
	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	stdoutOld := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-f", filename, "-s", "1,1", "-e", "3,1"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdoutOld

	assert.NoError(t, err)

	assert.Contains(t, string(out), "#####")
	assert.Contains(t, string(out), "#O#X#")
	assert.Contains(t, string(out), "#.#.#")
	assert.Contains(t, string(out), "#...#")
	assert.Contains(t, string(out), "#####")
}

func TestSolveCommand_OutputConsoleCorrect_Unicode(t *testing.T) {
	filename := "test_maze.txt"
	content := `#####
# # #
# # #
#   #
#####`

	//expectedContent := `┌─┬─┐
	//│O│X│
	//│.╵.│
	//│...│
	//└───┘`

	errWr := os.WriteFile(filename, []byte(content), 0644)
	assert.Nil(t, errWr)
	defer os.Remove(filename)

	cmd := NewSolveCommand()
	require.NotNil(t, cmd)

	stdoutOld := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-f", filename, "-s", "1,1", "-e", "3,1", "-u", "1"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdoutOld

	assert.NoError(t, err)

	assert.Contains(t, string(out), "┌─┬─┐")
	assert.Contains(t, string(out), "│O│X│")
	assert.Contains(t, string(out), "│.╵.│")
	assert.Contains(t, string(out), "│...│")
	assert.Contains(t, string(out), "└───┘")
}

func TestParsePoint_Incorrect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "строка координат пуста",
		},
		{
			name:     "one symbol",
			input:    "1",
			expected: "ожидается формат x,y, получено: 1",
		},
		{
			name:     "two symbols",
			input:    "12",
			expected: "ожидается формат x,y, получено: 12",
		},
		{
			name:     "three symbols",
			input:    "1,2,3",
			expected: "ожидается формат x,y, получено: 1,2,3",
		},
		{
			name:     "empty y",
			input:    "1,",
			expected: "ошибка при парсинге координаты y: ",
		},
		{
			name:     "empty x",
			input:    ",1",
			expected: "ошибка при парсинге координаты x: ",
		},
		{
			name:     "invalid x",
			input:    "a,1",
			expected: "ошибка при парсинге координаты x: ",
		},
		{
			name:     "invalid y",
			input:    "1,a",
			expected: "ошибка при парсинге координаты y: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parsePoint(tt.input)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}
