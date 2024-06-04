package main

import (
	o "ascii-art-output/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
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
	}

	banner := o.FileName()
	fmt.Println(banner)
	art := o.BannerArt(banner)
	fmt.Println(art)
}
