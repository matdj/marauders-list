package services

import (
    "marauders-list/domain"
)

var shopping_list = domain.ShoppingList{Items: []domain.Item{}, CrossedOffItems: []domain.Item{}}

func NewShoppingList() {
    shopping_list = domain.ShoppingList{Items: []domain.Item{}, CrossedOffItems: []domain.Item{}}
}

func Add(item domain.Item) {
	shopping_list.Items = append(shopping_list.Items, item)
}

func Items() []domain.Item {
	return shopping_list.Items
}

func ItemsContains(item domain.Item) bool {
    return sliceContains(shopping_list.Items, item)
}

func CrossedoffItemsContains(item domain.Item) bool {
    return sliceContains(shopping_list.CrossedOffItems, item)
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
        shopping_list.CrossedOffItems = append(shopping_list.CrossedOffItems, item)
    }
}

func UncrossoffItem(item domain.Item) {
    var itemIndex = -1
    for index, i := range shopping_list.CrossedOffItems {
        if i == item {
            itemIndex = index
            break 
        }
    }
    
    if itemIndex != -1 {
        shopping_list.CrossedOffItems = append(shopping_list.CrossedOffItems[:itemIndex], shopping_list.CrossedOffItems[itemIndex+1:]...)
        shopping_list.Items = append(shopping_list.Items, item)
    }
}