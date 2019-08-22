package lib

import (
	"fmt"
	"image"
	"image/gif"
	"math"
	"os"
	"time"
)

const (
	BORDER = 30
)

type State struct {
	Status *Status

	Original Algorithm
	Advanced Algorithm

	// Current Size
	BottomLeft Location
	TopRight   Location
}

func (state *State) Init(factor float64, status *Status) {
	state.Status = status

	state.Original = *NewAlgorithm()

	state.Advanced = *NewAlgorithm()
	state.Advanced.Factor = factor

	// Don't reset these
	//state.BottomLeft = *NewLocation(0, 0)
	//state.TopRight = *NewLocation(0, 0)
}

func Max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func Min(left int, right int) int {
	if left < right {
		return left
	}
	return right
}

func (state *State) MergeBounds() {
	state.TopRight.X = Max(state.TopRight.X, Max(state.Original.TopRight.X, state.Advanced.TopRight.X))
	state.TopRight.Y = Max(state.TopRight.Y, Max(state.Original.TopRight.Y, state.Advanced.TopRight.Y))

	state.BottomLeft.X = Min(state.BottomLeft.X, Min(state.Original.BottomLeft.X, state.Advanced.BottomLeft.X))
	state.BottomLeft.Y = Min(state.BottomLeft.Y, Min(state.Original.BottomLeft.Y, state.Advanced.BottomLeft.Y))
}

func (state *State) Bounds() (Location, Location) {
	return state.TopRight, state.BottomLeft
}

func (state *State) ExtendBounds(topRight Location, bottomLeft Location) {
	state.TopRight.X = Max(state.TopRight.X, topRight.X)
	state.TopRight.Y = Max(state.TopRight.Y, topRight.Y)

	state.BottomLeft.X = Min(state.BottomLeft.X, bottomLeft.X)
	state.BottomLeft.Y = Min(state.BottomLeft.Y, bottomLeft.Y)
}

func (state *State) UpdateBounds(next Location) {
	if next.X < state.BottomLeft.X {
		state.BottomLeft = *NewLocation(next.X, state.BottomLeft.Y)
	}

	if next.Y < state.BottomLeft.Y {
		state.BottomLeft = *NewLocation(state.BottomLeft.X, next.Y)
	}

	if next.X > state.TopRight.X {
		state.TopRight = *NewLocation(next.X, state.TopRight.Y)
	}

	if next.Y > state.TopRight.Y {
		state.TopRight = *NewLocation(state.TopRight.X, next.Y)
	}
}

type Result int

const (
	DONE Result = iota
	TELEPORT
	CHOOSE
)

func Point2Index(next float64, count int) int {
	// Should truncate it properly?
	//return int(next * float64(count))
	base, _ := math.Modf(next * float64(count))
	return int(base)
}

func (state *State) Next(nextPoint float64, last Result) Result {
	state.NextOriginal(nextPoint)
	return state.NextAdvanced(nextPoint, last)
}

func (state *State) NextOriginal(nextPoint float64) {
	index := Point2Index(nextPoint, 4)

	next := state.Original.Last.Next(index)
	state.Original.Set(next, false)
}

func (state *State) NextAdvanced(nextPoint float64, last Result) Result {
	if last == TELEPORT {
		count := len(state.Advanced.Order)
		//index := Point2Index(random.Float64(), count)
		index := Point2Index(nextPoint, count)
		key := state.Advanced.Order[index]

		state.Advanced.Set(ParseKey(key), false)
		state.Advanced.Teleport += 1
		return DONE
	}

	if last == CHOOSE {
		index := Point2Index(nextPoint, 4)
		next := state.Advanced.Last.Next(index)
		state.Advanced.Set(next, false)
		return DONE
	}

	available := state.Advanced.Available(state.Advanced.Last)

	count := len(available)

	if count > 0 {
		index := Point2Index(nextPoint, count)
		next := state.Advanced.Last.Next(available[index])
		state.Advanced.Set(next, true)
		return DONE

	} else {
		if nextPoint > state.Advanced.Factor {
			/*
				index := Point2Index(nextPoint, 4)
				next := state.Advanced.Last.Next(index)
				state.Advanced.Set(next, false)
				return DONE
			*/
			return CHOOSE
		} else {
			return TELEPORT
		}
	}
}

func (state *State) NewPlot(file *os.File) *Plot {

	// Update everything
	state.MergeBounds()
	state.Status.SetBox("Boundry", state.BottomLeft, state.TopRight)

	// Calculate the offsets
	var xoffset, yoffset int

	if state.BottomLeft.X < 0 {
		xoffset = state.BottomLeft.X * -1
	}

	if state.BottomLeft.Y < 0 {
		yoffset = state.BottomLeft.Y * -1
	}

	rect := image.Rect(0, 0, state.TopRight.X+xoffset+2*BORDER, state.TopRight.Y+yoffset+2*BORDER)

	palette := GradientPalette()
	pixels := image.NewPaletted(rect, palette)

	return &Plot{Pixels: pixels, Xoffset: xoffset, Yoffset: yoffset}
}

func DrawLine(plot *Plot, start *Location, end *Location, color uint8) {
	for x := Min(start.X, end.X); x <= Max(start.X, end.X); x++ {
		for y := Min(start.Y, end.Y); y <= Max(start.Y, end.Y); y++ {
			for xo := range []int{1, 0, -1} {
				for yo := range []int{1, 0, -1} {
					plot.Set(x+xo+BORDER, y+yo+BORDER, color)
				}
			}
		}
	}
}

func DrawBox(plot *Plot, start *Location, end *Location, color uint8) {
	DrawLine(plot, NewLocation(start.X, start.Y), NewLocation(start.X, end.Y), color)
	DrawLine(plot, NewLocation(start.X, end.Y), NewLocation(end.X, end.Y), color)
	DrawLine(plot, NewLocation(end.X, end.Y), NewLocation(end.X, start.Y), color)
	DrawLine(plot, NewLocation(end.X, start.Y), NewLocation(start.X, start.Y), color)
}

func Shift(point *Location, xoffset int, yoffset int) *Location {
	return NewLocation(point.X+xoffset, point.Y+yoffset)
}

func (state *State) AddAxis(plot *Plot) {

	left := NewLocation(state.Original.BottomLeft.X, state.Original.BottomLeft.Y)
	right := NewLocation(state.Original.TopRight.X, state.Original.TopRight.Y)
	DrawBox(plot, left, right, BLUE_END-(GRADIENT/2))

	left = NewLocation(state.Advanced.BottomLeft.X, state.Advanced.BottomLeft.Y)
	right = NewLocation(state.Advanced.TopRight.X, state.Advanced.TopRight.Y)
	DrawBox(plot, left, right, GREEN_END-(GRADIENT/2))

	set := []int{-1, 0, 1}
	for y := range set {
		start := NewLocation(state.BottomLeft.X, y)
		end := NewLocation(state.TopRight.X, y)
		DrawLine(plot, start, end, WHITE)
	}

	for x := range set {
		start := NewLocation(x, state.BottomLeft.Y)
		end := NewLocation(x, state.TopRight.Y)
		DrawLine(plot, start, end, WHITE)
	}
}

func (state *State) CreateGif() (file *os.File) {
	return state.OpenFile()
}

func (state *State) UpdateGif(file *os.File) *image.Paletted {

	plot := state.NewPlot(file)

	state.AddAxis(plot)

	var ocount, ncount int

	// Plot the original
	for x := state.BottomLeft.X; x <= state.TopRight.X; x++ {
		for y := state.BottomLeft.Y; y <= state.TopRight.Y; y++ {
			location := *NewLocation(x, y)
			count := state.Original.Count(location)
			if count > 0 {
				plot.Set(x, y, state.Original.MapColor(location, BLUE_START))
				ocount++
			}
		}
	}
	state.Status.Set("Original", "BLUE")
	state.Status.SetInteger("- Pixels", ocount)
	state.Status.SetInteger("- MaxOverlaps", state.Original.MaxOverlaps)

	// Plot the new
	for x := state.BottomLeft.X; x <= state.TopRight.X; x++ {
		for y := state.BottomLeft.Y; y <= state.TopRight.Y; y++ {
			location := *NewLocation(x, y)
			count := state.Advanced.Count(location)
			if count > 0 {
				plot.Set(x, y, state.Advanced.MapColor(location, GREEN_START))
				ncount++
			}
		}
	}
	state.Status.Set("Advanced", "GREEN")
	state.Status.SetInteger("-- Pixels", ncount)
	state.Status.SetInteger("-- MaxOverlaps", state.Advanced.MaxOverlaps)
	state.Status.SetInteger("-- Discards", state.Advanced.Discard)
	state.Status.SetInteger("-- Teleports", state.Advanced.Teleport)
	state.Status.SetFloat64("-- Factor", state.Advanced.Factor)

	state.Status.AddDescription(plot.Pixels)

	return plot.Pixels
}

func (state *State) CloseGif(file *os.File, images []*image.Paletted) {

	delay := make([]int, len(images))
	for i := 0; i < len(images); i++ {
		delay[i] = 25
	}

	anim := gif.GIF{Delay: delay, Image: images}

	err := gif.EncodeAll(file, &anim)
	if err != nil {
		panic(err)
	}

	file.Close()
}

func (state State) OpenFile() *os.File {
	now := time.Now()
	date := fmt.Sprintf("%04d%02d%02d", now.Year(), now.Month(), now.Day())
	time := fmt.Sprintf("%02d%02d%02d", now.Hour(), now.Minute(), now.Second())

	state.Status.Set("Date", date)
	state.Status.Set("Time", time)

	fileName := "examples/randomWalk-" + date + "-" + time + ".gif"

	state.Status.Set("Filename", fileName)

	fmt.Println("Gif: ", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return file
}
