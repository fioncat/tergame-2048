package game

import (
	"fmt"

	"github.com/fioncat/ter2048/draw"
	"github.com/nsf/termbox-go"
)

type Config struct {
	Options []int
	Size    int
	Grow    int
}

type Game struct {
	area *area

	steps int

	cfg *Config

	over bool
}

func Start(cfg *Config) {
	g := &Game{cfg: cfg}
	if err := g.start(); err != nil {
		fmt.Printf("start game failed: %v\n", err)
	}
}

func (g *Game) init() {
	data := make([][]int, g.cfg.Size)
	for i := 0; i < g.cfg.Size; i++ {
		data[i] = make([]int, g.cfg.Size)
	}
	a := &area{
		data:    data,
		options: g.cfg.Options,
		nRows:   g.cfg.Size,
		nCols:   g.cfg.Size,
	}
	a.fill(1)
	g.area = a
	g.steps = 0
	g.over = false
}

func (g *Game) start() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	g.init()
	g.draw()
loop:
	for {
		ev := termbox.PollEvent()
		if ev.Type != termbox.EventKey {
			continue loop
		}
		switch ev.Key {
		case termbox.KeyEsc:
			break loop

		case termbox.KeyEnter:
			if g.over {
				draw.Clear()
				g.init()
				g.draw()
				continue loop
			}
			termbox.Flush()

		default:
			if g.over {
				continue loop
			}
			draw.Clear()
			g.mainloop(ev.Ch)
		}
	}
	return nil
}

const (
	keyH = rune(104)
	keyJ = rune(106)
	keyK = rune(107)
	keyL = rune(108)
)

func (g *Game) mainloop(key rune) {
	var direct int
	switch key {
	case keyK:
		direct = directUp

	case keyJ:
		direct = directDown

	case keyH:
		direct = directLeft

	case keyL:
		direct = directRight

	default:
		return
	}
	g.over = !g.step(direct)

	g.draw()

	if g.over {
		fmt.Println("Game over! Press <Enter> to restart!")
	}
}

func (g *Game) draw() {
	draw.Area(g.area.data)
	draw.Points(g.area.point, g.steps)
	draw.Help()
}

func (g *Game) step(direct int) bool {
	merged := g.area.mergeAll(direct)
	if merged {
		g.steps++
		g.area.fill(g.cfg.Grow)
	}
	return g.area.canMerge()
}
