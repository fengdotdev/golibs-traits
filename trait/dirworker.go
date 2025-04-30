package trait


// NOT READY FOR PRODUCTION
type DirWorker interface {
	SetWorkingDir(path string) error
	GetWorkingDir() string
}
