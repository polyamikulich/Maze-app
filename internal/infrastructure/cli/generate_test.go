package cli

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест на вывод справки функции generate
func TestGenerateCommand_HelpFlag(t *testing.T) {
	cmd := NewGenerateCommand()

	// Проверяем, что команда создана
	require.NotNil(t, cmd)

	// Перенаправление вывода в pipe
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	//defer w.Close()
	os.Stdout = w

	// Выполнение: запускаем команду с флагом --help
	cmd.SetArgs([]string{"--help", "generate"})
	err := cmd.Execute()

	// Восстанавливаем stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	// Проверки
	assert.NoError(t, err) // Execute должен завершиться успешно при --help

	// Проверим, что в выводе есть все необходимые флаги
	assert.Contains(t, string(out), "Generate a maze with specified algorithm and dimensions")
	assert.Contains(t, string(out), "-a, --algorithm")
	assert.Contains(t, string(out), "dfs")
	assert.Contains(t, string(out), "prim")
	assert.Contains(t, string(out), "-W, --width")
	assert.Contains(t, string(out), "-H, --height")
	assert.Contains(t, string(out), "-o, --output")
	assert.Contains(t, string(out), "--unicode")
}

func TestGenerateCommand_InvalidAlgorithm(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	// Выполнение: запускаем команду с флагом --help
	cmd.SetArgs([]string{"-a", "invalid_algo"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Неподдерживаемый алгоритм: invalid_algo")
}

func TestGenerateCommand_InvalidSize(t *testing.T) {
	tests := []struct {
		name     string
		width    string
		height   string
		expected string
	}{
		{
			name:     "both negative",
			width:    "-9",
			height:   "-1",
			expected: "Ширина и высота должны быть положительными числами, получено: -9 * -1",
		},
		{
			name:     "zero width",
			width:    "0",
			height:   "5",
			expected: "Ширина и высота должны быть положительными числами, получено: 0 * 5",
		},
		{
			name:     "zero height",
			width:    "5",
			height:   "0",
			expected: "Ширина и высота должны быть положительными числами, получено: 5 * 0",
		},
		{
			name:     "both zero",
			width:    "0",
			height:   "0",
			expected: "Ширина и высота должны быть положительными числами, получено: 0 * 0",
		},
		{
			name:     "negative width",
			width:    "-5",
			height:   "10",
			expected: "Ширина и высота должны быть положительными числами, получено: -5 * 10",
		},
		{
			name:     "negative height",
			width:    "5",
			height:   "-10",
			expected: "Ширина и высота должны быть положительными числами, получено: 5 * -10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewGenerateCommand()
			require.NotNil(t, cmd)

			cmd.SetArgs([]string{"-W", tt.width, "-H", tt.height})
			err := cmd.Execute()

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expected)
		})
	}
}

func TestGenerateCommand_DFSStdout(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-W", "5", "-H", "5", "-a", "dfs"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.NoError(t, err)
	assert.Contains(t, string(out), "#")
}

func TestGenerateCommand_PrimaStdout(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-W", "5", "-H", "5", "-a", "prim"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.NoError(t, err)
	assert.Contains(t, string(out), "#")
}

func TestGenerateCommand_DFSEnhStdout(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-W", "5", "-H", "5", "-a", "dfs-enhanced"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.NoError(t, err)
	assert.Contains(t, string(out), "#")
}

func TestGenerateCommand_DFSEnhStdout_Unicode(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-W", "5", "-H", "5", "-a", "dfs-enhanced", "-u", "1"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.NoError(t, err)
	assert.Contains(t, string(out), "┌")
}

func TestGenerateCommand_PrimaEnhStdout(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs([]string{"-W", "5", "-H", "5", "-a", "prim-enhanced"})
	err := cmd.Execute()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	assert.NoError(t, err)
	assert.Contains(t, string(out), "#")
}

func TestGenerateCommand_ToFileCorrect(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	tempfile := "test_maze.txt"
	defer os.Remove(tempfile)

	cmd.SetArgs([]string{"-o", tempfile})
	err := cmd.Execute()

	assert.NoError(t, err)

	_, statErr := os.Stat(tempfile)
	assert.NoError(t, statErr)

	content, readErr := os.ReadFile(tempfile)
	assert.NoError(t, readErr)
	assert.Contains(t, string(content), "#")
}

func TestGenerateCommand_ToFileCorrectUnicode(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	tempfile := "test_maze.txt"
	defer os.Remove(tempfile)

	cmd.SetArgs([]string{"-o", tempfile, "-u", "1"})
	err := cmd.Execute()

	assert.NoError(t, err)

	_, statErr := os.Stat(tempfile)
	assert.NoError(t, statErr)

	content, readErr := os.ReadFile(tempfile)
	assert.NoError(t, readErr)
	assert.Contains(t, string(content), "┌")
}

func TestGenerateCommand_ToFileCorrectUnicode_Enh(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	tempfile := "test_maze.txt"
	defer os.Remove(tempfile)

	cmd.SetArgs([]string{"-o", tempfile, "-u", "1", "-a", "prim-enhanced"})
	err := cmd.Execute()

	assert.NoError(t, err)

	_, statErr := os.Stat(tempfile)
	assert.NoError(t, statErr)

	content, readErr := os.ReadFile(tempfile)
	assert.NoError(t, readErr)
	assert.Contains(t, string(content), "┌")
}

func TestGenerateCommand_InvalidUnicode(t *testing.T) {
	cmd := NewGenerateCommand()

	require.NotNil(t, cmd)

	cmd.SetArgs([]string{"--unicode", "34"})
	err := cmd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Недопустимое значение флага unicode: 34")
}
