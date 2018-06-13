package main

import "strings"

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

func unpadShortName(string) string {

	return ""
}
