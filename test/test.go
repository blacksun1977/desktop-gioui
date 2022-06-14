package main

import (
	"fmt"
	"github.com/blacksun1977/desktop-gioui/ui"
	"image"
)

func main() {
	w := ui.NewWindow()
	w.SetTitle("test window")
	w.SetSize(800, 600)
	w.OnResize(func(w ui.Window, from, to *image.Point) {
		fmt.Println("OnResize", from, to)
	})
	w.Main()
}
