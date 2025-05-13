package zdutil

import (
	"fmt"
	"strconv"
	"strings"
)

// ColorHexToRGB converts a color represented as a hex string to
// three individual RGB values.
//
// The returned values represent the red, green, and blue components
// of the color, each ranging from 0 to 255.
//
// An empty or invalid hex string will return three zero values.
func ColorHexToRGB(colorHex string) (red, green, blue int) {
	colorHex = strings.TrimPrefix(colorHex, "#")
	color64, err := strconv.ParseInt(colorHex, 16, 32)
	if err != nil {
		return
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF
}

// ColorRGBToHex converts three individual RGB values to a color represented as a hex string.
//
// The returned string is a valid hex color code, with a leading "#".
func ColorRGBToHex(red, green, blue int) string {
	r := strconv.FormatInt(int64(red), 16)
	g := strconv.FormatInt(int64(green), 16)
	b := strconv.FormatInt(int64(blue), 16)

	if len(r) == 1 {
		r = "0" + r
	}
	if len(g) == 1 {
		g = "0" + g
	}
	if len(b) == 1 {
		b = "0" + b
	}

	return "#" + r + g + b
}

// ConvertYmlToJson converts a YAML-compatible map to a JSON-compatible map.
//
// The input map may contain keys that are not strings, and the values
// may be maps that require recursive conversion. The returned map
// will have string keys, and only contain values that are compatible
// with JSON.
func ConvertYmlToJson(m map[interface{}]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for k, v := range m {
		switch v2 := v.(type) {
		case map[interface{}]interface{}:
			res[fmt.Sprint(k)] = ConvertYmlToJson(v2)
		default:
			res[fmt.Sprint(k)] = v
		}
	}
	return res
}
