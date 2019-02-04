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

func readEntries(absBase string) []os.FileInfo {
	// read input dir
	dirfiles, err := ioutil.ReadDir(absBase)
	if err != nil {
		die("Error: Unable to read the dir.")
	}

	return dirfiles
}

func createSubDir(absSubDirPath string) {
	if !org.IsDir(absSubDirPath) {
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

func main() {
	or := org.Init(readArg())
	items := readEntries(or.Source())
	createRequiredDir(&or, org.DirList())
	for _, f := range items {
		fileName := f.Name()
		if f.IsDir() && fileName == or.BaseDirName() {
			continue
		}

		absSource := or.LocateInSource(fileName)
		finalPath := or.FinalPath(fileName, org.SubDirName(absSource))
		err := os.Rename(absSource, finalPath)
		// log error
		if err != nil {
			fmt.Println("Error: Unable to move " + fileName + ".")
		} else {
			fmt.Println("Note: Moved " + fileName + ".")
		}
	}

	removeEmptySubDir(&or, org.DirList())
}
