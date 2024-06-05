package functions

import (
	"fmt"
	"os"
	"strings"
)

func Art(s string, mYmap map[int][]string) string {
	var str string
	s = strings.ReplaceAll(s, "\n", "\\n")
	input := strings.Split(s, "\\n")
	fmt.Println(input)
	for j, word := range input {
		if word == "" && j != len(input)-1 {
			str += fmt.Sprintln()
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, cha := range word {
				if ok := (cha >= ' ' && cha <= rune(126)) || (cha == '\n'); ok {
					// if cha == '\n' {
					// 	str += "\n"
					// } else {
					str += mYmap[int(cha)][i]
					//}
				} else if !ok {
					fmt.Println("Unprintable character within input")
					os.Exit(1)
				}
			}
			str += fmt.Sprintln()
		}
	}

	return str
}
