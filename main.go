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

	var perm string
	var file string
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
		os.Exit(1)
	}

	perm = fmt.Sprintf("%#o", fi.Mode().Perm())
	file, err = filepath.Abs(fi.Name())
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return filepath.Join("Regular file: "+green(file), "\nPerm: "+green(perm))

	case mode.IsDir():
		return filepath.Join("Directory: "+red(file), "\nPerm: "+red(perm))

	case mode&fs.ModeSymlink != 0:
		return filepath.Join("Symbolic link: "+blue(file), "\nPerm: "+blue(perm))
	}

	return file
}

func main() {

	if len(os.Args) > 1 {
		file := wif(os.Args)
		fmt.Println(file)
	} else {
		fmt.Println("No arguments were specified!")
		os.Exit(1)
	}

}
