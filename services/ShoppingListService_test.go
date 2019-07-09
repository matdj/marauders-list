package services

import (
	"marauders-list/domain"
	"testing"
)

func TestAddItemToShoppingList(t *testing.T) {
	service := new(ShoppingListService)
	weetbix := domain.Item{Name: "Weetbix"}
	service.Add(weetbix)

	if !service.ItemsContains(weetbix) {
		t.Errorf("Shopping list must contain 'Weetbix'")
	}
}

func TestCopyItems(t *testing.T) {
	service := new(ShoppingListService)

	itemsThen := service.Items()

	weetbix := domain.Item{Name: "Weetbix"}
	service.Add(weetbix)

	if len(itemsThen) != 0 {
		t.Errorf("Items must be a copy, to prevent external mutation")
	}
}

func TestCopyCrossedoffItems(t *testing.T) {
	service := new(ShoppingListService)

	crossedoffIemsThen := service.CrossedoffItems()

	weetbix := domain.Item{Name: "Weetbix"}
	service.Add(weetbix)
	service.CrossoffItem(weetbix)

	if len(crossedoffIemsThen) != 0 {
		t.Errorf("CrossedoffItems must be a copy, to prevent external mutation")
	}
}

func TestCrossoffItem(t *testing.T) {
	service := new(ShoppingListService)

	bananas := domain.Item{Name: "bananas"}
	service.Add(bananas)

	service.CrossoffItem(bananas)

	if service.ItemsContains(bananas) {
		t.Errorf("Items must not contain bananas")
	}

	if !service.CrossedoffItemsContains(bananas) {
		t.Errorf("Crossedoff items must contain bananas")
	}
}

func TestUncrossoffItem(t *testing.T) {
	service := new(ShoppingListService)

	pineapples := domain.Item{Name: "pineapples"}
	service.Add(pineapples)

	service.CrossoffItem(pineapples)
	service.UncrossoffItem(pineapples)

	if service.CrossedoffItemsContains(pineapples) {
		t.Errorf("Crossedoff items must NOT contain pineapples")
	}

	if !service.ItemsContains(pineapples) {
		t.Errorf("Items must contain pineapples")
	}
}
