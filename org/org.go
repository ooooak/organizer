package org

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"./paths"
)

func isDir(absPath string) bool {
	fi, err := os.Stat(absPath)
	return (err == nil && fi.IsDir())
}

func createDir(absDir string) error {
	return os.Mkdir(absDir, os.ModePerm)
}

func proccessDir(org paths.Organizer, dirs []string) error {
	// create Folder sub dir
	if !isDir(org.AbsSubDir(fileTypeDirectory)) {
		// sub dir dont exit? create subdir it
		err := createDir(org.AbsSubDir(fileTypeDirectory))
		if err != nil {
			return err
		}
	}

	// handle dir
	for _, fileName := range dirs {
		err := os.Rename(org.AbsSource(fileName),
			org.FinalPath(fileName, fileTypeDirectory))
		if err != nil {
			return err
		}
	}

	return nil
}

func proccessFiles(org paths.Organizer, files []string) error {
	for _, fileName := range files {
		absSource := org.AbsSource(fileName)
		fileType := guessFileType(getExt(absSource))
		subDir := org.AbsSubDir(fileType)
		finalPath := org.FinalPath(fileName, fileType)

		if !isDir(subDir) {
			// sub dir dont exit? create subdir it
			err := createDir(subDir)
			if err != nil {
				return err
			}
		}

		err := os.Rename(absSource, finalPath)
		if err != nil {
			return err
		}

		if fileType == fileTypeHTML {
			// move data dir
			htmlDataDir := strings.Split(fileName, ".")[0] + "_files"
			absHTMLDataPath := org.FinalPath(htmlDataDir, fileType)
			if isDir(absHTMLDataPath) {
				// move data file
				err = os.Rename(org.AbsSource(htmlDataDir), absHTMLDataPath)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Run file organizer
func Run(basePath string) {
	org := paths.Init(basePath)

	// read input dir
	dirfiles, err := ioutil.ReadDir(org.Base())
	if err != nil {
		log.Fatal(err)
	}

	// create new base
	if !isDir(org.NewBase()) {
		err = createDir(org.NewBase())
		if err != nil {
			log.Fatal(err)
		}
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

	// handle files
	err = proccessFiles(org, files)
	if err != nil {
		log.Println(err)
	}

	// handle dirs
	err = proccessDir(org, dirs)
	if err != nil {
		log.Println(err)
	}
}
