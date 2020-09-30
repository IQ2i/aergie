package logger

import (
	"strings"

	"github.com/gookit/color"
)

// Error print an error message in console and quit application
func Error(err error) {
	color.Println("")
	color.Println("<error>  " + strings.Repeat(" ", len(err.Error())) + "  </>")
	color.Println("<error>  " + err.Error() + "  </>")
	color.Println("<error>  " + strings.Repeat(" ", len(err.Error())) + "  </>")
	color.Println("")
}

// Step print an message in console with the current step
func Step(step string) {
	color.Println("")
	color.Println("<fg=white;bg=blue;>      " + strings.Repeat(" ", len(step)) + "  </>")
	color.Println("<fg=white;bg=blue;>  [*] " + step + "  </>")
	color.Println("<fg=white;bg=blue;>      " + strings.Repeat(" ", len(step)) + "  </>")
	color.Println("")
}
