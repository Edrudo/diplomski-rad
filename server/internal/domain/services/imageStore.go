package services

type ImageStore interface {
	StoreImage(imageName string, image []byte) (string, error)
}
