package dir

import (
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func padShortName(str string) string {
	if len(str) > 12 {
		return ""
	}
	parts := strings.Split(str, ".")
	name := str
	ext := ""
	if len(parts) > 1 {
		name = strings.Join(parts[0:len(parts)-1], ".")
		ext = parts[len(parts)-1]
	}
	if len(name) > 8 || len(ext) > 3 {
		return ""
	}
	spacesNeeded := (8 - len(name)) + (3 - len(ext))

	spaces := ""
	for i := 0; i < spacesNeeded; i++ {
		spaces += " "
	}

	return strings.ToUpper(name + spaces + ext)
}

func unpadShortName(bytes []byte) string {
	if len(bytes) != 11 {
		return "☠️ "
	}
	name := make([]rune, 0, 8)
	ext := make([]rune, 0, 3)
	for i, b := range bytes {
		c := charmap.CodePage437.DecodeByte(b)
		if i < 8 {
			name = append(name, c)
		} else {
			ext = append(ext, c)
		}
	}
	strName := strings.TrimRight(string(name), " ")
	strExt := strings.TrimRight(string(ext), " ")
	if len(strExt) == 0 {
		return strName
	}

	return strName + "." + strExt
}
