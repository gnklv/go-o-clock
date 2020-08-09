package ui

import (
	"fmt"
	"os"

	"../font"
	"../types"

	"github.com/nsf/termbox-go"
)

func ToText(str string) types.Text {
	symbols := make(types.Text, 0)
	for _, r := range str {
		if s, ok := font.DefaultFont[r]; ok {
			symbols = append(symbols, s)
		}
	}
	return symbols
}

func Echo(s types.Symbol, startX, startY int) {
	x, y := startX, startY
	for _, line := range s {
		for _, r := range line {
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}
		x = startX
		y++
	}
}

func Clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
}

func Flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func Stderr(s string, a ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, s, a...)
	if err != nil {
		panic(err)
	}
}
