package org

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"./paths"
)

// List of sub dir
const (
	FileTypeGraphics   string = "Graphics"
	FileTypeDataFiles  string = "Datafiles"
	FileTypeExecutable string = "Executable"
	FileTypeImages     string = "Images"
	FileTypeArchive    string = "Archive"
	FileTypeDocs       string = "Docs"
	FileTypeBooks      string = "Books"
	FileTypeAudio      string = "Audio"
	FileTypeVideos     string = "Videos"
	FileTypeScripts    string = "Scripts"
	FileTypeHTML       string = "Html"
	FileTypeDirectory  string = "Folders"
	FileTypeTorrent    string = "Torrent"
	FileTypeText       string = "Text"
	FileTypeShortCut   string = "ShortCut"
	FileTypeEmptyDir   string = "Empty"
	FileTypeUnknown    string = "Unknown"
)

// TODO: find away to define subdir
// SubDirList !
func SubDirList() []string {
	return []string{
		FileTypeGraphics, FileTypeDataFiles,
		FileTypeExecutable, FileTypeImages,
		FileTypeArchive, FileTypeDocs,
		FileTypeBooks, FileTypeAudio,
		FileTypeVideos, FileTypeScripts,
		FileTypeHTML, FileTypeDirectory,
		FileTypeTorrent, FileTypeText,
		FileTypeShortCut, FileTypeEmptyDir,
		FileTypeUnknown,
	}
}

// GetExt file extension
func GetExt(absSource string) string {
	return filepath.Ext(filepath.Base(absSource))
}

// GuessFileType get file type
func GuessFileType(ext string) string {
	switch strings.ToLower(ext) {
	case ".psd", ".eps", ".ai", ".flinto", ".sketch":
		return FileTypeGraphics

	case ".sql", ".xml", ".json":
		return FileTypeDataFiles

	case ".exe", ".msi":
		return FileTypeExecutable

	case ".png", ".jpg", ".svg", ".gif", ".jpeg":
		return FileTypeImages

	case ".zip", ".rar", ".7z", ".gz":
		return FileTypeArchive

	case ".text", ".txt":
		return FileTypeText

	case ".docx", ".doc", ".xlsx", ".md", ".pub", ".pt":
		return FileTypeDocs

	case ".epub", ".pdf", ".djvu", ".chm":
		return FileTypeBooks

	case ".mp3", ".m3u":
		return FileTypeAudio

	case ".mp4", ".flv", ".3gp", ".mpg", ".wmv", ".mov":
		return FileTypeVideos

	case ".php", ".c", ".js", ".cpp", ".fs", ".hs", ".ml",
		".rs", ".go", ".d", ".java", ".h", ".py", ".rb", ".lua",
		".r", ".rkt", ".clj", ".cljs", ".coffee", ".ts":
		return FileTypeScripts

	case ".torrent":
		return FileTypeTorrent

	case ".htm", ".html":
		return FileTypeHTML

	case ".lnk":
		return FileTypeShortCut

	default:
		return FileTypeUnknown
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
func MoveDir(org *paths.Organizer, dir string) error {
	source := org.AbsSource(dir)
	subDir := FileTypeDirectory
	if IsEmptyDir(source) {
		subDir = FileTypeEmptyDir
	}

	return os.Rename(source, org.FinalPath(dir, subDir))
}

// MoveFile in
func MoveFile(org *paths.Organizer, fileName string) error {
	absSource := org.AbsSource(fileName)
	fileType := GuessFileType(GetExt(absSource))

	finalPath := org.FinalPath(fileName, fileType)

	err := os.Rename(absSource, finalPath)
	if err != nil {
		return err
	}

	if fileType == FileTypeHTML {
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
