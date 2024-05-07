package filesystem

import (
	"fmt"
	"os"
)

type ImageStore struct{}

func NewImageStore() *ImageStore {
	return &ImageStore{}
}

func (i *ImageStore) StoreImage(imageName string, imgBytes []byte) error {
	out, err := os.Create(
		fmt.Sprintf(
			// TODO change this path to be configurable
			"./images/%s",
			imageName,
		),
	)
	defer out.Close()
	if err != nil {
		panic(err)
	}

	// write into a file
	if _, err := out.Write(imgBytes); err != nil {
		// TODO add logging
		panic(err)
	}

	return nil
}
