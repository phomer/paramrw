package main

import (
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/phomer/paramrw/lib"
	"github.com/phomer/paramrw/lib/log"
)

var seed = time.Now().UnixNano()
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	factors := []float64{
		1.0,
		0.7, 0.5, 0.3, 0.1,
		0.07, 0.05, 0.03, 0.01,
		0.007, 0.005, 0.003, 0.001,
		0.0007, 0.0005, 0.0003, 0.0001,
		0.00007, 0.00005, 0.00003, 0.00001,
		0.000007, 0.000005, 0.000003, 0.000001,
	}

	state := make([]lib.State, len(factors))

	//iterations := 5000000 // TODO: Get from argv
	iterations := 1000000 // TODO: Get from argv

	images := make([]*image.Paletted, 0)

	state[0].Init(0.0, &lib.Status{})
	file := state[0].CreateGif()

	for index, factor := range factors {

		random = rand.New(rand.NewSource(seed))
		fmt.Println("Parameterized Random Walk:", iterations)

		state[index].Init(factor, &lib.Status{})
		state[index].Status.SetInteger("Iterations", iterations)

		var last lib.Result = lib.DONE
		for i := 0; i < iterations; i++ {
			value := random.Float64()
			last = state[index].Next(value, last)
		}
		state[index].MergeBounds()
	}

	var top, bottom lib.Location

	first := true
	for index, _ := range factors {
		if !first {
			state[index].ExtendBounds(top, bottom)
		} else {
			first = false
		}
		top, bottom = state[index].Bounds()
	}

	for index, _ := range factors {
		state[index].TopRight = top
		state[index].BottomLeft = bottom
		plot := state[index].UpdateGif(file)
		log.DEBUG("Bounds", state[index].BottomLeft, state[index].TopRight)
		images = append(images, plot)
	}

	state[0].CloseGif(file, images)
}
