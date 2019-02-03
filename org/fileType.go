package org

import (
	"path/filepath"
	"strings"
)

const (
	fileTypeGraphics   string = "Graphics"
	fileTypeDataFiles  string = "Datafiles"
	fileTypeExecutable string = "Executable"
	fileTypeImages     string = "Images"
	fileTypeArchive    string = "Archive"
	fileTypeDocs       string = "Docs"
	fileTypeBooks      string = "Books"
	fileTypeAudio      string = "Audio"
	fileTypeVideos     string = "Videos"
	fileTypeScripts    string = "Scripts"
	fileTypeHTML       string = "Html"
	fileTypeDirectory  string = "Folders"
	fileTypeTorrent    string = "Torrent"
	fileTypeText       string = "Text"
	fileTypeUnknown    string = "Unknown"
)

func getExt(absSource string) string {
	return filepath.Ext(filepath.Base(absSource))
}

func guessFileType(ext string) string {
	switch strings.ToLower(ext) {
	case ".psd", ".eps", ".ai", ".flinto", ".sketch":
		return fileTypeGraphics

	case ".sql", ".xml", ".json":
		return fileTypeDataFiles

	case ".exe", ".msi":
		return fileTypeExecutable

	case ".png", ".jpg", ".svg", ".gif", ".jpeg":
		return fileTypeImages

	case ".zip", ".rar", ".7z", ".gz":
		return fileTypeArchive

	case ".text", ".txt":
		return fileTypeText

	case ".docx", ".doc", ".xlsx", ".md", ".pub", ".pt":
		return fileTypeDocs

	case ".epub", ".pdf", ".djvu", ".chm":
		return fileTypeBooks

	case ".mp3", ".m3u":
		return fileTypeAudio

	case ".mp4", ".flv", ".3gp", ".mpg", ".wmv", ".mov":
		return fileTypeVideos

	case ".php", ".c", ".js", ".cpp", ".fs":
		return fileTypeScripts

	case ".torrent":
		return fileTypeTorrent

	case ".htm", ".html":
		return fileTypeHTML

	default:
		return fileTypeUnknown
	}
}
