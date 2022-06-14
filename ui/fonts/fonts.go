package fonts

import (
	_ "embed"
	"fmt"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
)

//go:embed simkai.ttf
var ttfSimKai []byte

/*
    phtf, err := ioutil.ReadFile("assets/phtl.ttf")
    phtft, err := opentype.Parse(phtf)
    if err != nil {
        panic(err)
    }
    fonts := []text.FontFace{
        text.FontFace{Face: phtft},
    }
    th := material.NewTheme(fonts)

    /**
    yhfs, err := ioutil.ReadFile("assets/msyhl.ttc")
    fonts, err :=opentype.ParseCollection(yhfs)
    if err != nil {
        panic(err)
    }
    fmt.Printf("num of font:%d\n", fonts.NumFonts());

    sft, err := fonts.Font(0)
    fcarr := []text.FontFace{
        text.FontFace{Face: sft},
    }
    th := material.NewTheme(fcarr)//assets.FontCollection()
//gofont.Collection()
*/

func GetFonts() []text.FontFace {
	var fonts []text.FontFace
	// 原始英文字体
	fonts = append(fonts, gofont.Collection()...)
	// 自定义中文字体
	register := func(fnt text.Font, ttf []byte) {
		face, err := opentype.Parse(ttf)
		if err != nil {
			panic(fmt.Errorf("failed to parse font: %v", err))
		}
		fnt.Typeface = "Go"
		fonts = append(fonts, text.FontFace{Font: fnt, Face: face})
	}
	register(text.Font{}, ttfSimKai)
	return fonts
}
