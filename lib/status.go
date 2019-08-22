package lib

import (
	"fmt"
	"image"
)

type Status struct {
	Parameters map[string]string
	Order      []string

	RunDate  string
	RunTime  string
	FileName string

	ImageBoundry string
	Iterations   string

	OriginalColor    string
	OriginalCount    string
	OriginalOverlaps string

	AdvancedColor    string
	AdvancedCount    string
	AdvancedOverlaps string
}

func (status *Status) SetBox(name string, bottom Location, top Location) {
	status.Set(name, fmt.Sprintf("(%d,%d),(%d,%d)", bottom.X, bottom.Y, top.X, top.Y))
}

func (status *Status) SetInteger(name string, value int) {
	status.Set(name, fmt.Sprintf("%d", value))
}

func (status *Status) SetFloat64(name string, value float64) {
	status.Set(name, fmt.Sprintf("%e", value))
}

func (status *Status) Set(name string, value string) {
	if status.Parameters == nil {
		status.Parameters = make(map[string]string, 0)
	}

	if status.Order == nil {
		status.Order = make([]string, 0)
	}

	status.Parameters[name] = value
	status.Order = append(status.Order, name)
}

func (status Status) AddDescription(plot *image.Paletted) {
	size := 40.0
	row := 43
	column := 40
	nl := int(size) + 10

	for _, key := range status.Order {
		text := fmt.Sprintf("%s: %s", key, status.Parameters[key])
		AddFlexibleLabel(plot, column, row, text, size)
		row += nl
	}

}
