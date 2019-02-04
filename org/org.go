package org

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// List of sub dir
const (
	DirGraphics   string = "Graphics"
	DirDataFiles  string = "Datafiles"
	DirExecutable string = "Executable"
	DirImages     string = "Images"
	DirArchive    string = "Archive"
	DirDocs       string = "Docs"
	DirBooks      string = "Books"
	DirAudio      string = "Audio"
	DirVideos     string = "Videos"
	DirScripts    string = "Scripts"
	DirHTML       string = "Html"
	DirTorrent    string = "Torrent"
	DirText       string = "Text"
	DirShortCut   string = "ShortCut"
	DirUnknown    string = "Unknown"

	// Used to handle Folder
	DirFolder   string = "Folders"
	DirEmptyDir string = "Empty Folders"
)

// DirList returns all dir that can be created
func DirList() []string {
	return []string{
		DirGraphics,
		DirDataFiles,
		DirExecutable,
		DirImages,
		DirArchive,
		DirDocs,
		DirBooks,
		DirAudio,
		DirVideos,
		DirScripts,
		DirHTML,
		DirTorrent,
		DirText,
		DirShortCut,
		DirUnknown,
		DirEmptyDir,
		DirFolder,
	}
}

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

// AbsSubDir return abs path of guess new sub dir
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
	subDir := DirFolder
	if IsEmptyDir(source) {
		subDir = DirEmptyDir
	}

	return os.Rename(source, org.FinalPath(dir, subDir))
}

// Ext return file ext from abs path
func Ext(absSource string) string {
	return filepath.Ext(filepath.Base(absSource))
}

// SubDirNameByExt dir type by ext
func SubDirNameByExt(ext string) string {
	switch strings.ToLower(ext) {
	case ".psd", ".eps", ".ai", ".flinto", ".sketch":
		return DirGraphics

	case ".sql", ".xml", ".json":
		return DirDataFiles

	case ".exe", ".msi":
		return DirExecutable

	case ".png", ".jpg", ".svg", ".gif", ".jpeg":
		return DirImages

	case ".zip", ".rar", ".7z", ".gz":
		return DirArchive

	case ".text", ".txt":
		return DirText

	case ".docx", ".doc", ".xlsx", ".md", ".pub", ".pt":
		return DirDocs

	case ".epub", ".pdf", ".djvu", ".chm":
		return DirBooks

	case ".mp3", ".m3u":
		return DirAudio

	case ".mp4", ".flv", ".3gp", ".mpg", ".wmv", ".mov":
		return DirVideos

	case ".php", ".c", ".js", ".cpp", ".fs", ".hs", ".ml",
		".rs", ".go", ".d", ".java", ".h", ".py", ".rb", ".lua",
		".r", ".rkt", ".clj", ".cljs", ".coffee", ".ts":
		return DirScripts

	case ".torrent":
		return DirTorrent

	case ".htm", ".html":
		return DirHTML

	case ".lnk":
		return DirShortCut

	default:
		return DirUnknown
	}
}

// SubDirName !
func SubDirName(absSource string) string {
	if IsDir(absSource) {
		if IsEmptyDir(absSource) {
			return DirEmptyDir
		}
		return DirFolder
	} else {
		return SubDirNameByExt(Ext(absSource))
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
