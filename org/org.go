package org

import (
	"io"
	"os"
	"strings"

	"./filetype"
)

// Organizer has all the info about directory
type Organizer struct {
	base string
	name string
	sp   string
}

// AbsSource return exact path of working dir
func (org *Organizer) AbsSource(fileName string) string {
	return org.base + org.sp + fileName
}

// Base returns new base dir
func (org *Organizer) Base() string {
	return org.base
}

// NewBaseName !
func (org *Organizer) NewBaseName() string {
	return org.name
}

// NewBase returns new base dir
func (org *Organizer) NewBase() string {
	return org.base + org.sp + org.name
}

// AbsSubDir retrun abs path of guess new sub dir
func (org *Organizer) AbsSubDir(subDirType string) string {
	return org.NewBase() + org.sp + subDirType
}

// FinalPath of the file that will be created
func (org *Organizer) FinalPath(fileName, subDirType string) string {
	return org.AbsSubDir(subDirType) + org.sp + fileName
}

// Init base Organizer
func Init(basePath string) Organizer {
	return Organizer{
		base: basePath,
		name: "__organized",
		sp:   string(os.PathSeparator),
	}
}

// IsEmptyDir directory
func IsEmptyDir(absDir string) bool {
	f, err := os.Open(absDir)
	if err != nil {
		return false
	}

	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}

	return false
}

// IsDir !
func IsDir(absPath string) bool {
	fi, err := os.Stat(absPath)
	return (err == nil && fi.IsDir())
}

// CreateDir !
func CreateDir(absDir string) error {
	return os.Mkdir(absDir, os.ModePerm)
}

// MoveDir to its sub dir
func MoveDir(org *Organizer, dir string) error {
	source := org.AbsSource(dir)
	subDir := filetype.FileTypeDirectory
	if IsEmptyDir(source) {
		subDir = filetype.FileTypeEmptyDir
	}

	return os.Rename(source, org.FinalPath(dir, subDir))
}

// MoveFile in
func MoveFile(org *Organizer, fileName string) error {
	absSource := org.AbsSource(fileName)
	fileType := filetype.Get(filetype.Ext(absSource))

	finalPath := org.FinalPath(fileName, fileType)

	err := os.Rename(absSource, finalPath)
	if err != nil {
		return err
	}

	if fileType == filetype.FileTypeHTML {
		// move data dir
		htmlDataDir := strings.Split(fileName, ".")[0] + "_files"
		absHTMLDataPath := org.FinalPath(htmlDataDir, fileType)
		if IsDir(absHTMLDataPath) {
			// move data file
			err = os.Rename(org.AbsSource(htmlDataDir), absHTMLDataPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SubDirList !
func SubDirList() []string {
	return []string{
		filetype.FileTypeGraphics,
		filetype.FileTypeDataFiles,
		filetype.FileTypeExecutable,
		filetype.FileTypeImages,
		filetype.FileTypeArchive,
		filetype.FileTypeDocs,
		filetype.FileTypeBooks,
		filetype.FileTypeAudio,
		filetype.FileTypeVideos,
		filetype.FileTypeScripts,
		filetype.FileTypeHTML,
		filetype.FileTypeDirectory,
		filetype.FileTypeTorrent,
		filetype.FileTypeText,
		filetype.FileTypeShortCut,
		filetype.FileTypeEmptyDir,
		filetype.FileTypeUnknown,
	}
}
