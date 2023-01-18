package model

type ProcessedResponse interface {
	GetItems() []ProcessedItem
}

type ProcessedItem interface {
	GetName() string
}

func CastToProcessedItem[T ProcessedItem](items []T) []ProcessedItem {
	result := []ProcessedItem{}
	for _, item := range items {
		result = append(result, item)
	}
	return result
}
