package main_test

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"os/exec"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("it greets me", func(t *testing.T) {
		output, err := runPreCommit(t)
		assert.NoError(t, err)
		assert.Equal(t, "Hello world!", output)
	})
}

func runPreCommit(tb testing.TB) (string, error) {
	b, err := exec.CommandContext(context.Background(), "go", "run", "./...").CombinedOutput()
	output := strings.TrimSpace(string(b))
	tb.Log("CLI output:", output)
	return output, err
}
