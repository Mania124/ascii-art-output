package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var outputFile string
	var input string
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	if len(os.Args) == 2 {
		input = os.Args[1]
		outputFile = "banner.txt"
	}
	if len(os.Args) == 4 {
		if !(os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		if !strings.HasPrefix(os.Args[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		outputFile = os.Args[1][9:]
		input = os.Args[2]
	}
	if len(os.Args) == 3 {
		if !strings.HasPrefix(os.Args[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		// if !(os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkertoy") {
		// 	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		// 	os.Exit(0)
		// }
		input = os.Args[1]
		outputFile = os.Args[1][9:]
	}
	fmt.Println(input)
	fmt.Println(outputFile)
	//banner := o.FileName()
	//fmt.Println(banner)
	//art := o.BannerArt(banner)
	//fmt.Println(art)
}
