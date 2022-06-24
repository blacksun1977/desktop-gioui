package main

import (
	"fmt"
	"gioui.org/widget/material"
	"github.com/blacksun1977/desktop-gioui/ui"
	"github.com/blacksun1977/desktop-gioui/ui/fonts"
	"image"
	"image/color"
)

func main() {
	w := ui.NewWindow()
	w.SetTitle("test window")
	w.SetSize(800, 600)
	//w.SetWindowMode(app.Fullscreen)
	//w.SetWindowMode(app.Maximized) // not work on Mac
	//w.SetWindowMode(app.Minimized) // not work on Mac
	w.OnResize(func(w ui.Window, from, to *image.Point) {
		fmt.Println("OnResize", from, to)
	})
	mt := material.NewTheme(fonts.GetFonts())
	mt.Bg = color.NRGBA{A: 0xFF, R: 0x00, G: 0xFF, B: 0x00}
	mt.Fg = color.NRGBA{A: 0xFF, R: 0x00, G: 0xFF, B: 0xFF}
	w.SetTheme(mt)
	w.Main()
}
