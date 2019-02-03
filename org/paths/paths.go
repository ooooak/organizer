package paths

import "os"

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
