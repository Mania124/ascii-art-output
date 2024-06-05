package functions

import (
	"fmt"
	"os"
	"strings"
)

func Art(s string, mYmap map[int][]string) string {
	var str string
	input := strings.Split(s, "\\n")
	for j, word := range input {
		if word == "" && j != len(input)-1 {
			str += fmt.Sprintln()
		}
		for i := 1; i <= 8; i++ {
			for _, cha := range word {
				if ok := (cha >= ' ' && cha <= rune(126)); ok {
					str += mYmap[int(cha)][i]
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
