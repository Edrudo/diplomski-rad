package controller

type Part struct {
	DataHash   string `json:"dataHash"`
	PartNumber int    `json:"partNumber"`
	TotalParts int    `json:"totalParts"`
	PartData   []byte `json:"partData"`
}
