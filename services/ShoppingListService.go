package services

import (
	"marauders-list/domain"
)

type ShoppingListService struct {
	shoppingList domain.ShoppingList
}

func (service *ShoppingListService) Add(item domain.Item) {
	service.shoppingList.Items = append(service.shoppingList.Items, item)
}

func (service *ShoppingListService) Items() []domain.Item {
	b := make([]domain.Item, len(service.shoppingList.Items))
	copy(b, service.shoppingList.Items)
	return b
}

func (service *ShoppingListService) CrossedoffItems() []domain.Item {
    b := make([]domain.Item, len(service.shoppingList.CrossedoffItems))
	copy(b, service.shoppingList.Items)
	return b
}

func (service *ShoppingListService) ItemsContains(item domain.Item) bool {
	return sliceContains(service.shoppingList.Items, item)
}

func (service *ShoppingListService) CrossedoffItemsContains(item domain.Item) bool {
	return sliceContains(service.shoppingList.CrossedoffItems, item)
}

func (service *ShoppingListService) CrossoffItem(item domain.Item) {
	var itemIndex = -1
	for index, i := range service.shoppingList.Items {
		if i == item {
			itemIndex = index
			break
		}
	}

	if itemIndex != -1 {
		service.shoppingList.Items = append(service.shoppingList.Items[:itemIndex], service.shoppingList.Items[itemIndex+1:]...)
		service.shoppingList.CrossedoffItems = append(service.shoppingList.CrossedoffItems, item)
	}
}

func (service *ShoppingListService) UncrossoffItem(item domain.Item) {
	var itemIndex = -1
	for index, i := range service.shoppingList.CrossedoffItems {
		if i == item {
			itemIndex = index
			break
		}
	}

	if itemIndex != -1 {
		service.shoppingList.CrossedoffItems = append(service.shoppingList.CrossedoffItems[:itemIndex], service.shoppingList.CrossedoffItems[itemIndex+1:]...)
		service.shoppingList.Items = append(service.shoppingList.Items, item)
	}
}

func sliceContains(slice []domain.Item, item domain.Item) bool {
	for _, b := range slice {
		if b == item {
			return true
		}
	}
	return false
}
