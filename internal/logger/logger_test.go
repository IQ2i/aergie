package logger

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	color.ForceOpenColor()
	is := assert.New(t)

	buf := new(bytes.Buffer)
	color.SetOutput(buf)
	Error(fmt.Errorf("error test"))
	color.ResetOutput()

	expected := "\n\x1b[97;41m              \x1b[0m\n\x1b[97;41m  error test  \x1b[0m\n\x1b[97;41m              \x1b[0m\n\n"
	is.Equal(expected, buf.String())
}

func TestStep(t *testing.T) {
	color.ForceOpenColor()
	is := assert.New(t)

	buf := new(bytes.Buffer)
	color.SetOutput(buf)
	Step("step 1")
	color.ResetOutput()

	expected := "\n\x1b[37;44m              \x1b[0m\n\x1b[37;44m  [*] step 1  \x1b[0m\n\x1b[37;44m              \x1b[0m\n\n"
	is.Equal(expected, buf.String())
}
