package lib

import "fmt"

type Location struct {
	X int
	Y int
}

func NewLocation(x, y int) *Location {
	return &Location{
		X: x,
		Y: y,
	}
}

func ParseKey(key string) Location {
	var x, y int
	_, err := fmt.Sscanf(key, "%d-%d", &x, &y)
	if err != nil {
		panic(err)
	}

	return *NewLocation(x, y)
}

func (location Location) GenerateKey() string {
	return fmt.Sprintf("%d-%d", location.X, location.Y)
}

func (location Location) Next(direction int) Location {
	switch direction {
	case LEFT:
		return *NewLocation(location.X-1, location.Y)
	case RIGHT:
		return *NewLocation(location.X+1, location.Y)
	case DOWN:
		return *NewLocation(location.X, location.Y-1)
	case UP:
		return *NewLocation(location.X, location.Y+1)
	}
	return location
}

func (location Location) Print() {
	fmt.Println("Location: ", location.X, ",", location.Y)
}
