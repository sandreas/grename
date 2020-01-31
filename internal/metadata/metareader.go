package metadata

type MetaReader interface {
	ReadFromFile(filename string) error
}
