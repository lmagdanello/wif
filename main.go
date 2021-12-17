package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/mgutz/ansi"
)

func wif(target []string) string {

	var answer string
	green := ansi.ColorFunc("green+")
	red := ansi.ColorFunc("red+")
	blue := ansi.ColorFunc("blue+")

	if len(target) > 2 {
		fmt.Println("Too many arguments!")
		os.Exit(1)
	}

	fi, err := os.Lstat(target[1])

	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		answer, err = filepath.Abs(fi.Name())
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}

		return filepath.Join("Regular file: " + green(answer))

	case mode.IsDir():
		answer, err = filepath.Abs(fi.Name())
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}

		return filepath.Join("Directory: " + red(answer))

	case mode&fs.ModeSymlink != 0:
		answer, err = filepath.Abs(fi.Name())
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}

		return filepath.Join("Symbolic link: " + blue(answer))
	}

	return answer
}

func main() {

	if len(os.Args) > 1 {
		answer := wif(os.Args)
		fmt.Println(answer)
	} else {
		fmt.Println("No arguments were specified!")
		os.Exit(1)
	}

}
