package filesystem

import (
	"fmt"
	"os"

	"http3-server-poc/cmd/api/config"
)

type ImageStore struct{}

func NewImageStore() *ImageStore {
	return &ImageStore{}
}

func (i *ImageStore) StoreImage(imageName string, imgBytes []byte) (string, error) {
	filePath := fmt.Sprintf(
		"%s/%s.%s",
		config.Cfg.ImageConfig.ImageDirectory,
		imageName,
		config.Cfg.ImageConfig.ImageExtension,
	)
	out, err := os.Create(filePath)
	defer out.Close()
	if err != nil {
		panic(err)
	}

	// write into a file
	if _, err := out.Write(imgBytes); err != nil {
		// TODO add logging
		panic(err)
	}

	return filePath, nil
}
