package ui

import (
	"image/color"

	"github.com/crislainesc/space_game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Menu struct {
	isMobile    bool
	readyToPlay bool
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func NewMenu(isMobile bool) *Menu {
	return &Menu{
		readyToPlay: false,
		isMobile:    isMobile,
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	text.Draw(
		screen,
		"GO: Space War",
		assets.ScoreFont,
		200,
		300,
		color.White,
	)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(315, 150)
	screen.DrawImage(assets.GopherPlayer, op)
	msg := "Press ENTER to start"
	if m.isMobile {
		msg = "Touch the screen to start"
	}
	text.Draw(screen, msg, assets.FontUi, 100, 400, color.White)
}

func (m *Menu) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) || len(ebiten.AppendTouchIDs(nil)) > 0 {
		m.readyToPlay = true
	}
}

func (m *Menu) IsReady() bool {
	return m.readyToPlay
}

func (m *Menu) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
