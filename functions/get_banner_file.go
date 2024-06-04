package functions

import "os"

func FileName() string {
	if len(os.Args) == 2 {
		return "standard.txt"
	}
	switch os.Args[3] {
	case "standard":
		return "standard.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	case "shadow":
		return "shadow.txt"
	default:
		return "standard.txt"
	}
}
