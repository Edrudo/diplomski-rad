package services

type JsonStore interface {
	StoreJson(jsonName string, json []byte) (string, error)
}
