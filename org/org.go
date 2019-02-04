package org

import (
	"os"
	"strings"

	"./filetype"
)

// Organizer has all the info about directory
type Organizer struct {
	source   string
	baseName string
	sp       string
}

// LocateInSource return exact path of working dir
func (org *Organizer) LocateInSource(fileName string) string {
	return org.source + org.sp + fileName
}

// Source returns new base dir
func (org *Organizer) Source() string {
	return org.source
}

// BaseDirName !
func (org *Organizer) BaseDirName() string {
	return org.baseName
}

// AbsBase returns new base dir
func (org *Organizer) AbsBase() string {
	return org.source + org.sp + org.baseName
}

// AbsSubDir retrun abs path of guess new sub dir
func (org *Organizer) AbsSubDir(subDirType string) string {
	return org.AbsBase() + org.sp + subDirType
}

// FinalPath of the file that will be created
func (org *Organizer) FinalPath(fileName, subDirType string) string {
	return org.AbsSubDir(subDirType) + org.sp + fileName
}

// Init base Organizer
func Init(inputDir string) Organizer {
	return Organizer{
		source:   inputDir,
		baseName: "__organized",
		sp:       string(os.PathSeparator),
	}
}

// MoveDir to its sub dir
func (org *Organizer) MoveDir(dir string) error {
	source := org.LocateInSource(dir)
	subDir := filetype.FileTypeDirectory
	if IsEmptyDir(source) {
		subDir = filetype.FileTypeEmptyDir
	}

	return os.Rename(source, org.FinalPath(dir, subDir))
}

// MoveFile in
func (org *Organizer) MoveFile(fileName string) error {
	absSource := org.LocateInSource(fileName)
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
			err = os.Rename(org.LocateInSource(htmlDataDir), absHTMLDataPath)
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
