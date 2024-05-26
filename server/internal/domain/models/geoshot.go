package models

type Geoshot struct {
	Image           []byte      `json:"jpg"`
	Timestamp       string      `json:"ts"`
	Age             int         `json:"age"`
	DeviceId        int         `json:"id"`
	Datetime        string      `json:"dt"`
	Coordinates     [][]float64 `json:"geo"`
	VehicleStanding int         `json:"vel"`
}
