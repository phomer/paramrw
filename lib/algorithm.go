package lib

// Should be split for polymorphic reasons
type Algorithm struct {
	Visited     map[string]int
	Order       []string
	Last        Location
	MaxOverlaps int

	Factor   float64
	Discard  int
	Teleport int

	BottomLeft Location
	TopRight   Location
}

func NewAlgorithm() *Algorithm {
	return &Algorithm{
		Visited: make(map[string]int, 0),
		//Factor:  1.0,
		//Factor: 1.0 / 10.0,
		Factor: 1.0 / 1000.0,
		//Factor: 1.0 / 1000000.0,
	}
}

func (algorithm Algorithm) MapColor(location Location, base int) uint8 {
	count := float64(algorithm.Count(location))
	ratio := float64(GRADIENT*INCREMENT) / float64(algorithm.MaxOverlaps)
	step := count * ratio

	if step > GRADIENT {
		step = GRADIENT
	} else if step < 0.0 {
		step = 0.0
	}

	result := uint8(base + int(step))

	return result
}

func (algorithm Algorithm) Available(location Location) []int {
	base := []int{}
	for i := 0; i < 4; i++ {
		nearby := location.Next(i)
		if algorithm.IsEmpty(nearby) {
			base = append(base, i)
		}
	}
	return base
}

func (algorithm Algorithm) Count(location Location) int {
	key := location.GenerateKey()
	var value int
	var found bool
	if value, found = algorithm.Visited[key]; !found {
		return 0
	}
	return value
}

func (algorithm Algorithm) IsEmpty(location Location) bool {
	key := location.GenerateKey()
	if _, found := algorithm.Visited[key]; !found {
		return true
	}
	return false
}

func (algorithm Algorithm) AllFull(location Location) bool {
	for i := 0; i < 4; i++ {
		nearby := location.Next(i)
		if algorithm.IsEmpty(nearby) {
			return false
		}
	}
	return true
}

func (algorithm *Algorithm) Set(location Location, updateOrder bool) {
	key := location.GenerateKey()
	if value, found := algorithm.Visited[key]; !found {
		algorithm.Visited[key] = 0
		if updateOrder {
			algorithm.Order = append(algorithm.Order, key)
		}

	} else {
		algorithm.Visited[key] = value + 1
		if value+1 > algorithm.MaxOverlaps {
			algorithm.MaxOverlaps = value + 1
		}
	}
	algorithm.Last = location
	algorithm.UpdateBounds(location)
}

func (algorithm *Algorithm) UpdateBounds(next Location) {
	if next.X < algorithm.BottomLeft.X {
		algorithm.BottomLeft = *NewLocation(next.X, algorithm.BottomLeft.Y)
	}

	if next.Y < algorithm.BottomLeft.Y {
		algorithm.BottomLeft = *NewLocation(algorithm.BottomLeft.X, next.Y)
	}

	if next.X > algorithm.TopRight.X {
		algorithm.TopRight = *NewLocation(next.X, algorithm.TopRight.Y)
	}

	if next.Y > algorithm.TopRight.Y {
		algorithm.TopRight = *NewLocation(algorithm.TopRight.X, next.Y)
	}
}
