package filetype

import (
	"path/filepath"
	"strings"
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

// Ext return file ext from abs path
func Ext(absSource string) string {
	return filepath.Ext(filepath.Base(absSource))
}

// Get File type
func Get(ext string) string {
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
