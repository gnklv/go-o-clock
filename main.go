package main

import (
	"fmt"
	"os"
	"time"

	"./ui"

	"github.com/nsf/termbox-go"
)

const tick = time.Second

var (
	timer          *time.Timer
	ticker         *time.Ticker
	queues         chan termbox.Event
	startDone      bool
	startX, startY int
)

func draw() {
	w, h := termbox.Size()
	ui.Clear()

	str := format()
	text := ui.ToText(str)

	if !startDone {
		startDone = true
		startX, startY = w/2-text.Width()/2, h/2-text.Height()/2
	}

	x, y := startX, startY
	for _, s := range text {
		ui.Echo(s, x, y)
		x += s.Width()
	}

	ui.Flush()
}

func format() string {
	now := time.Now()

	y := now.Year()
	mm := now.Month()
	d := now.Day()
	h := now.Hour()
	m := now.Minute()
	s := now.Second()

	return fmt.Sprintf("%02d.%02d.%02d %02d:%02d:%02d", y, mm, d, h, m, s)
}

func countup() {
	var exitCode int
	ticker = time.NewTicker(tick)

loop:
	for {
		select {
		case ev := <-queues:
			if ev.Type == termbox.EventKey && (ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC) {
				exitCode = 1
				break loop
			}
		case <-ticker.C:
			draw()
		}
	}

	termbox.Close()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	queues = make(chan termbox.Event)
	go func() {
		for {
			queues <- termbox.PollEvent()
		}
	}()

	draw()
	countup()
}
