package ui

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/blacksun1977/desktop-gioui/ui/fonts"
	"image"
	"os"
)

type Window interface {
	GetTitle() (title string)
	SetTitle(title string)
	SetSize(width, height int)
	GetSize() *image.Point
	SetWindowMode(mode app.WindowMode)
	GetWindowMode() app.WindowMode
	GetTheme() *material.Theme
	SetTheme(t *material.Theme)
	Main()
	OnResize(onResize func(w Window, from, to *image.Point))
}

func NewWindow() Window {
	return &mWindow{}
}

type mWindow struct {
	gw         *app.Window
	gt         *material.Theme
	fullScreen bool
	title      string
	size       *image.Point
	mode       app.WindowMode
	onResize   func(m Window, from, to *image.Point)
}

func (m *mWindow) OnResize(onResize func(w Window, from, to *image.Point)) {
	m.onResize = onResize
}

func (m *mWindow) GetTitle() (title string) {
	if m.title == "" {
		m.title = "Desktop-GioUi"
	}
	return m.title
}

func (m *mWindow) SetTitle(title string) {
	m.title = title
	if m.gw != nil {
		m.gw.Option(app.Title(m.GetTitle()))
	}
}

func (m *mWindow) SetWindowMode(mode app.WindowMode) {
	m.mode = mode
	if m.gw != nil {
		m.gw.Option(mode.Option())
	}
}

func (m *mWindow) GetWindowMode() app.WindowMode {
	return m.mode
}

func (m *mWindow) SetSize(width, height int) {
	m.size = &image.Point{
		X: width,
		Y: height,
	}
	if m.gw != nil {
		size := m.GetSize()
		m.gw.Option(app.MinSize(unit.Dp(size.X), unit.Dp(size.Y)),
			app.Size(unit.Dp(size.X), unit.Dp(size.Y)))
		m.gw.Perform(system.ActionCenter)
	}
}

func (m *mWindow) GetSize() *image.Point {
	if m.size == nil || m.size.X <= 0 || m.size.Y <= 0 {
		m.size = &image.Point{X: 800, Y: 600}
	}
	return m.size
}

func (m *mWindow) GetTheme() *material.Theme {
	if m.gt == nil {
		m.gt = material.NewTheme(fonts.GetFonts())
	}
	return m.gt
}

func (m *mWindow) SetTheme(t *material.Theme) {
	if t != nil {
		m.gt = t
		if m.gw != nil {
			m.gw.Invalidate()
		}
	}
}

func (m *mWindow) Main() {
	go m.run()
	app.Main()
}

func (m *mWindow) run() {
	size := m.GetSize()
	m.gw = app.NewWindow(
		app.MinSize(unit.Dp(size.X), unit.Dp(size.Y)),
		app.Size(unit.Dp(size.X), unit.Dp(size.Y)),
		app.Title(m.GetTitle()),
		m.mode.Option())
	m.gw.Perform(system.ActionCenter)

	var ops op.Ops
	for {
		select {
		case e := <-m.gw.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				//fmt.Println("system.DestroyEvent", e)
				goto Done
			case system.FrameEvent:
				//fmt.Println("system.FrameEvent", e)
				gtx := layout.NewContext(&ops, e)
				//router.Layout(gtx, th)
				e.Frame(gtx.Ops)
			case app.ConfigEvent:
				m.mode = e.Config.Mode
				if m.onResize != nil {
					from := m.GetSize()
					to := e.Config.Size
					if from.X != to.X || from.Y != to.Y {
						m.onResize(m, from, &to)
					}
				}
				*m.size = e.Config.Size
				//	fmt.Println("app.ConfigEvent", e.Config.Mode, e)
				//case app.ViewEvent:
				//	fmt.Println("app.ViewEvent", e)
				//case system.StageEvent:
				//	fmt.Println("system.StageEvent", e)
				//default:
				//	if e != nil {
				//		fmt.Println(e)
				//	}
			}
		}
	}
Done:
	os.Exit(0)
}
