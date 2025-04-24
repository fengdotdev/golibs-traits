package trait




type Saver interface {
	Save() error
	Load() error
	LoadFrom(path string) error
	SaveTo(path string) error
}