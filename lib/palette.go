package lib

import (
	"image/color"
)

const (
	INCREMENT   = 5
	GRADIENT    = 40
	BLACK       = 0
	WHITE       = 1
	RED_START   = WHITE + 1
	RED_END     = RED_START + GRADIENT
	GREEN_START = RED_END + 1
	GREEN_END   = GREEN_START + GRADIENT
	BLUE_START  = GREEN_END + 1
	BLUE_END    = BLUE_START + GRADIENT
)

func position(base int, index int) uint8 {
	result := 0xff - (uint8(base-index) * uint8(INCREMENT))
	//fmt.Println("Table Entry is", result)
	return result
}

func WhitePalette() []color.Color {
	palette := make([]color.Color, 5)

	palette[0] = color.RGBA{0xff, 0xff, 0xff, 0xff}
	palette[1] = color.RGBA{0xfe, 0xff, 0xff, 0xff}
	palette[2] = color.RGBA{0xfe, 0xff, 0xff, 0xff}
	palette[3] = color.RGBA{0xfe, 0xff, 0xff, 0xff}
	palette[4] = color.RGBA{0xfe, 0xff, 0xff, 0xff}

	return palette
}

func GradientPalette() []color.Color {
	size := BLUE_END + 1
	palette := make([]color.Color, size)

	palette[BLACK] = color.RGBA{0x00, 0x00, 0x00, 0xff}
	palette[WHITE] = color.RGBA{0xff, 0xff, 0xff, 0xff}

	for i := RED_START; i <= RED_END; i++ {
		palette[i] = color.RGBA{position(RED_END, i), 0x00, 0x00, 0xff}
	}

	for i := GREEN_START; i <= GREEN_END; i++ {
		palette[i] = color.RGBA{0x00, position(GREEN_END, i), 0x00, 0xff}
	}

	for i := BLUE_START; i <= BLUE_END; i++ {
		palette[i] = color.RGBA{0x00, 0x00, position(BLUE_END, i), 0xff}
	}

	return palette
}

func SimplePalette() []color.Color {

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, // Black
		color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
		color.RGBA{0x00, 0x00, 0xff, 0xff}, // Blue
		color.RGBA{0x00, 0xff, 0x00, 0xff}, // Green
		color.RGBA{0x00, 0xff, 0xff, 0xff}, // Cyan
		color.RGBA{0xff, 0x00, 0x00, 0xff}, // Red
		color.RGBA{0xff, 0x00, 0xff, 0xff}, // Magenta
		color.RGBA{0xff, 0xff, 0x00, 0xff}, // Yellow
	}

	return palette
}
