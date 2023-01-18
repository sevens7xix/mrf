package model

type ProcesedReponse interface {
	GetItems() []ProcessedItem
}
type ProcessedItem interface {
	GetName() string
}
