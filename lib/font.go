package lib

import (
	"flag"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "../github.com/golang/freetype/testdata/luxisr.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	//size     = flag.Float64("size", 12, "font size in points")
	spacing = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb    = flag.Bool("whiteonblack", false, "white text on a black background")
)

func AllWhiteImage() *image.Paletted {

	rect := image.Rect(0, 0, 100, 100)
	base := image.NewPaletted(rect, WhitePalette())

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			base.Set(x, y, base.Palette[1])
		}
	}

	return base
}

func AddFlexibleLabel(img *image.Paletted, x, y int, label string, size float64) {
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		panic(err)
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	context := freetype.NewContext()
	context.SetDPI(*dpi)
	context.SetFont(f)
	context.SetFontSize(size)
	//context.SetDst(rgba)
	//context.SetSrc(fg)

	context.SetSrc(AllWhiteImage())
	context.SetClip(img.Bounds())
	context.SetDst(img)

	//size := 12.0 // font size in pixels
	pt := freetype.Pt(x, y+int(context.PointToFixed(size)>>6))

	if _, err := context.DrawString(label, pt); err != nil {
		panic("DrawString Failed")
	}
}

func AddFixedLabel(img *image.Paletted, x, y int, label string) {

	col := color.RGBA{0x10, 0x10, 0xff, 0xff} // BLUE

	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	face := basicfont.Face7x13
	draw := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}

	draw.DrawString(label)
}
