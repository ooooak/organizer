package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"../org"
	"../org/paths"
)

func die(text string) {
	fmt.Println(text)
	os.Exit(0)
}

func readArg() string {
	if len(os.Args) < 2 {
		die("Error: Please Specify path you want to organize.")
	}

	mainDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		die("Error: Unable to parse the path.")
	}

	return mainDir
}

func readEntries(absBase string) ([]string, []string) {
	// read input dir
	dirfiles, err := ioutil.ReadDir(absBase)
	if err != nil {
		die("Error: Unable to read the dir.")
	}

	var files []string
	var dirs []string
	for _, f := range dirfiles {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		} else {
			files = append(files, f.Name())
		}
	}
	return files, dirs
}

// create new base
func createNewBaseDir(newBase string) {
	if !org.IsDir(newBase) {
		err := org.CreateDir(newBase)
		if err != nil {
			die("Error: Unable to create DIR. That was required.")
		}
	}
}

func createSubDir(absSubDirPath string) {
	if !org.IsDir(absSubDirPath) {
		// sub dir dont exit? create subdir it
		err := org.CreateDir(absSubDirPath)
		if err != nil {
			fmt.Println("Error: Unable to create Sub Dir " + absSubDirPath + ".")
		}
	}
}

func main() {
	wrDir := paths.Init(readArg())
	createNewBaseDir(wrDir.NewBase())

	files, dirs := readEntries(wrDir.Base())

	// handle files
	for _, fileName := range files {
		createSubDir(wrDir.AbsSubDir(org.GuessFileType(org.GetExt(fileName))))
		err := org.MoveFile(&wrDir, fileName)
		if err != nil {
			fmt.Println("Error: Unable to move " + fileName + ".")
		} else {
			fmt.Println("Note: Moved " + fileName + ".")
		}
	}

	// handle dirs
	if len(dirs) > 0 {
		createSubDir(wrDir.AbsSubDir(org.FileTypeDirectory))
		// TODO: Handle it better
		createSubDir(wrDir.AbsSubDir(org.FileTypeEmptyDir))
	}

	for _, dir := range dirs {
		// TODO: Define Constant
		if dir == "__organized" {
			continue
		}

		err := org.MoveDir(&wrDir, dir)
		if err != nil {
			fmt.Println("Error: Unable to move " + dir + ".")
		} else {
			fmt.Println("Note: Move " + dir + ".")
		}
	}
}
