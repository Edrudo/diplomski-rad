package models

type Part struct {
	DataHash   string
	PartNumber int
	TotalParts int
	PartData   []byte
}
