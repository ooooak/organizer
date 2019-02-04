package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"../org"
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

func createSubDir(absSubDirPath string) {
	if !org.IsDir(absSubDirPath) {
		// sub dir dont exit? create subdir it
		err := org.CreateDir(absSubDirPath)
		if err != nil {
			fmt.Println("Error: Unable to create Sub Dir " + absSubDirPath + ".")
		}
	}
}

func removeEmptySubDir(wrDir *org.Organizer, subdirs []string) {
	for _, dir := range subdirs {
		absPath := (wrDir.AbsSubDir(dir))
		if org.IsEmptyDir(absPath) {
			os.Remove(absPath)
		}
	}
}

func createRequiredDir(wrDir *org.Organizer, subdir []string) {
	if !org.IsDir(wrDir.AbsBase()) {
		err := org.CreateDir(wrDir.AbsBase())
		if err != nil {
			die("Error: Unable to create required DIR.")
		}
	}

	for _, dir := range subdir {
		createSubDir(wrDir.AbsSubDir(dir))
	}
}

// TODO:
func main() {
	organizer := org.Init(readArg())
	files, dirs := readEntries(organizer.Source())

	createRequiredDir(&organizer, org.SubDirList())

	// handle files
	for _, fileName := range files {
		err := organizer.MoveFile(fileName)
		if err != nil {
			fmt.Println("Error: Unable to move " + fileName + ".")
		} else {
			fmt.Println("Note: Moved " + fileName + ".")
		}
	}

	// handle dir
	for _, dir := range dirs {
		if dir == organizer.BaseDirName() {
			continue
		}

		err := organizer.MoveDir(dir)
		if err != nil {
			fmt.Println("Error: Unable to move " + dir + ".")
		} else {
			fmt.Println("Note: Move " + dir + ".")
		}
	}

	removeEmptySubDir(&organizer, org.SubDirList())
}
