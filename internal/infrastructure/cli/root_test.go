package cli

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест на вывод справки функции generate
func TestRootCommand_HelpFlag(t *testing.T) {
	cmd := NewRootCommand()

	// Проверяем, что команда создана
	require.NotNil(t, cmd)

	// Перенаправление вывода в pipe
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	//defer w.Close()
	os.Stdout = w

	// Выполнение: запускаем команду с флагом --help
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()

	// Восстанавливаем stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	// Проверки
	assert.NoError(t, err) // Execute должен завершиться успешно при --help

	// Проверим, что в выводе есть все необходимые флаги
	assert.Contains(t, string(out), "Maze generator and solver CLI application.")
	assert.Contains(t, string(out), "-h, --help")
	assert.Contains(t, string(out), "-V, --version")
	assert.Contains(t, string(out), "generate   Generate a maze with specified algorithm and dimensions.")
	assert.Contains(t, string(out), "solve      Solve a maze with specified algorithm and points.")
}

func TestRootCommand_VersionOutput(t *testing.T) {
	cmd := NewRootCommand()

	require.NotNil(t, cmd)

	// Перехватываем вывод команды
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)

	// Устанавливаем аргументы для вызова версии
	cmd.SetArgs([]string{"--version"})

	// Выполняем команду
	err := cmd.Execute()

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "maze-app v1.0.0")
}

func TestRootCommand_InvalidFlag(t *testing.T) {
	cmd := NewRootCommand()

	// Устанавливаем неверный флаг
	cmd.SetArgs([]string{"--invalid-flag"})

	// Выполняем команду и проверяем ошибку
	err := cmd.Execute()

	// Проверяем, что возникла ошибка
	assert.Error(t, err)
}
