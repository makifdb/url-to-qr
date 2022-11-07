package color

import (
	"image/color"
	"strconv"
)

func ParseColor(s string) (color.Color, error) {
	if s == "" {
		return color.Transparent, nil
	}

	// remove the # from the string
	s = s[1:]

	// convert to rgb
	r, err := strconv.ParseUint(s[0:2], 16, 8)
	if err != nil {
		return nil, err
	}

	g, err := strconv.ParseUint(s[2:4], 16, 8)
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseUint(s[4:6], 16, 8)
	if err != nil {
		return nil, err
	}

	a, err := strconv.ParseUint(s[6:8], 16, 8)
	if err != nil {
		return nil, err
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}, nil
}
