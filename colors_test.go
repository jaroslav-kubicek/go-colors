package colors_test

import (
	"os"
	"testing"

	"github.com/jaroslav-kubicek/go-colors"
	"github.com/stretchr/testify/assert"
)

func TestCanAcceptMultiple_Success(t *testing.T) {
	redAndBold := colors.New(colors.Codes["red"], colors.Codes["bold"])

	assert.Equal(t, "\x1b[31m\x1b[1mred and bold text\x1b[39m\x1b[22m", redAndBold.Format("red and bold text"))
}

func TestReturnsTextWithDisabledColors_Success(t *testing.T) {
	os.Setenv("NO_COLOR", "1")
	defer os.Unsetenv("NO_COLOR")

	assert.Equal(t, "red and bold text", colors.New(colors.Codes["red"]).Format("red and bold text"))
}
