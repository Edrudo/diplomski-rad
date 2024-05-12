package services

type DataStore interface {
	StorePart(name string, part []byte) error
}
