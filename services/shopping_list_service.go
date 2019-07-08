package services

import (
    "marauders-list/domain"
)

var shopping_list = domain.ShoppingList{Items: []domain.Item{}, CrossedoffItems: []domain.Item{}}

func NewShoppingList() {
    shopping_list = domain.ShoppingList{Items: []domain.Item{}, CrossedoffItems: []domain.Item{}}
}

func Add(item domain.Item) {
	shopping_list.Items = append(shopping_list.Items, item)
}

func Items() []domain.Item {
	return append([]domain.Item(nil), shopping_list.Items...)
}

func CrossedoffItems() []domain.Item {
	return append([]domain.Item(nil), shopping_list.CrossedoffItems...)
}

func ItemsContains(item domain.Item) bool {
    return sliceContains(shopping_list.Items, item)
}

func CrossedoffItemsContains(item domain.Item) bool {
    return sliceContains(shopping_list.CrossedoffItems, item)
}

func sliceContains(slice []domain.Item, item domain.Item) bool {
    for _, b := range slice {
        if b == item {
            return true
        }
    }
    return false
}

func CrossoffItem(item domain.Item) {
    var itemIndex = -1
    for index, i := range shopping_list.Items {
        if i == item {
            itemIndex = index
            break 
        }
    }

    if itemIndex != -1 {
        shopping_list.Items = append(shopping_list.Items[:itemIndex], shopping_list.Items[itemIndex+1:]...)
        shopping_list.CrossedoffItems = append(shopping_list.CrossedoffItems, item)
    }
}

func UncrossoffItem(item domain.Item) {
    var itemIndex = -1
    for index, i := range shopping_list.CrossedoffItems {
        if i == item {
            itemIndex = index
            break 
        }
    }
    
    if itemIndex != -1 {
        shopping_list.CrossedoffItems = append(shopping_list.CrossedoffItems[:itemIndex], shopping_list.CrossedoffItems[itemIndex+1:]...)
        shopping_list.Items = append(shopping_list.Items, item)
    }
}