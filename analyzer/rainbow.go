package analyzer

import (
	"fmt"
	"math"
	"strconv"
)

func hsvToRGB(h, s, v float64) (r, g, b int) {
	if h >= 360 {
		h = 0
	}
	h /= 60
	i := math.Floor(h)
	ff := h - i
	p := v * (1 - s)
	q := v * (1 - s*ff)
	t := v * (1 - s*(1-ff))

	var rF, gF, bF float64
	switch int(i) {
	case 0:
		rF, gF, bF = v, t, p
	case 1:
		rF, gF, bF = q, v, p
	case 2:
		rF, gF, bF = p, v, t
	case 3:
		rF, gF, bF = p, q, v
	case 4:
		rF, gF, bF = t, p, v
	case 5:
		rF, gF, bF = v, p, q
	default:
		rF, gF, bF = v, v, v
	}
	return int(rF * 255), int(gF * 255), int(bF * 255)
}

func rainbowText(i, total int, label, text string) string {
	hue := float64(i) / float64(total) * 360.0
	r, g, b := hsvToRGB(hue, 1.0, 1.0)
	ansi := "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
	reset := "\033[0m"
	return fmt.Sprintf("%s%s %s%s", ansi, label, text, reset)
}
