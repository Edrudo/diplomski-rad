package models

type ImagePart struct {
	ImageName  string
	PartNumber int
	TotalParts int
	PartData   []byte
}
