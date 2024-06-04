package functions

import (
	"fmt"
	"os"
)

//Checks the properties of banner file with focus on size and non-exist error
func CheckStatus(s string) {
	_, err := os.Stat(s)
	if os.IsNotExist(err) {
		var choice string
		fmt.Println("Banner file missing! Would you wish to download and continue? yes/no")
		fmt.Scanf("%s", &choice)
		switch choice {
		case "yes", "y", "YES", "Yes", "Y":
			err := Download(s)
			if err != nil {
				fmt.Printf("error donloading file: %s", err)
			}
		case "no", "No", "NO", "n", "N":
			fmt.Println("Can't download without user authorization. Bye.")
			os.Exit(1)
		default:
			fmt.Println("Wrong USER input. Try Again later.")
			os.Exit(1)
		}

	}
	switch s {
	case "standard.txt":
		if info, _ := os.Stat("standard.txt"); info.Size() != 6623 {
			fmt.Println("File content changed. Please confirm!")
			os.Exit(1)
		}
		// if info.IsDir() {
		// 	fmt.Println("This is a directory. Please confirm!")
		// 	os.Exit(1)
		// }
	case "shadow.txt":
		if info, _ := os.Stat("shadow.txt"); info.Size() != 7463 {
			fmt.Println("File content changed. Please confirm!")
			os.Exit(1)
		}
		// if stat.IsDir() {
		// 	fmt.Println("This is a directory. Please confirm!")
		// 	os.Exit(1)
		// }
	case "thinkertoy.txt":
		if info, _ := os.Stat("thinkertoy.txt"); info.Size() != 5558 {
			fmt.Println("File content changed. Please confirm!")
			os.Exit(1)
		}
		// if stat.IsDir() {
		// 	fmt.Println("This is a directory. Please confirm!")
		// 	os.Exit(1)
		// }
	}

}
