package trait

type DirWorker interface {
	SetWorkingDir(path string) error
	GetWorkingDir() string
}
