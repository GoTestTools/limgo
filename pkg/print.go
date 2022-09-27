package pkg

import (
	"fmt"
	"io"
)

type Color string

// ASCII escape codes for colors, see: https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
const (
	colorReset Color = "\033[0m"

	ColorRed    Color = "\033[31m"
	ColorGreen  Color = "\033[32m"
	ColorYellow Color = "\033[33m"
	ColorBlue   Color = "\033[34m"
	ColorPurple Color = "\033[35m"
	ColorCyan   Color = "\033[36m"
	ColorWhite  Color = "\033[37m"
)

func ColorizedFPrintln(w io.Writer, txt string, color Color) {
	fmt.Fprintf(w, "%s%s%s\n", color, txt, colorReset)
}

func ColorizedPrintln(txt string, color Color) {
	fmt.Printf("%s%s%s\n", color, txt, colorReset)
}
