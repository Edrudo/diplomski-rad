package filesystem

import (
	"fmt"
	"os"

	"http3-server-poc/cmd/api/config"
)

type JsonStore struct{}

func NewJsonStore() *JsonStore {
	return &JsonStore{}
}

func (i *JsonStore) StoreJson(jsonName string, json []byte) (string, error) {
	filePath := fmt.Sprintf(
		"%s/%s.json",
		config.Cfg.JsonConfig.Directory,
		jsonName,
	)
	out, err := os.Create(filePath)
	defer out.Close()
	if err != nil {
		panic(err)
	}

	// write into a file
	if _, err := out.Write(json); err != nil {
		// TODO add logging
		panic(err)
	}

	return filePath, nil
}
