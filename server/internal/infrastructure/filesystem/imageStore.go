package filesystem

import (
	"fmt"
	"os"
)

type DataStore struct{}

func NewDataStore() *DataStore {
	return &DataStore{}
}

func (i *DataStore) StorePart(dataHash string, data []byte) error {
	out, err := os.Create(
		fmt.Sprintf(
			// TODO change this path to be configurable
			"./images/%s",
			dataHash,
		),
	)
	defer out.Close()
	if err != nil {
		panic(err)
	}

	// write into a file
	if _, err := out.Write(data); err != nil {
		// TODO add logging
		panic(err)
	}

	return nil
}
