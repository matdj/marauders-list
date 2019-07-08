package services

import (
	"marauders-list/domain"
	"testing"
)

func TestAddItemToShoppingList(t *testing.T) {
	NewShoppingList()

	weetbix := domain.Item{Name: "Weetbix"}
	Add(weetbix)

	if !ItemsContains(weetbix) {
		t.Errorf("Shopping list must contain 'Weetbix'")
	}
}

func TestCrossoffItem(t *testing.T) {
	NewShoppingList()

	bananas := domain.Item{Name: "bananas"}
	Add(bananas)

	CrossoffItem(bananas)

	if ItemsContains(bananas) {
		t.Errorf("Items must not contain bananas")
	}

	if !CrossedoffItemsContains(bananas) {
		t.Errorf("Crossedoff items must contain bananas")
	}
} 

func TestUncrossoffItem(t *testing.T) {
	NewShoppingList()

	pineapples := domain.Item{Name: "pineapples"}
	Add(pineapples)

	CrossoffItem(pineapples)
	UncrossoffItem(pineapples)

	if CrossedoffItemsContains(pineapples) {
		t.Errorf("Crossedoff items must NOT contain pineapples")
	}

	if !ItemsContains(pineapples) {
		t.Errorf("Items must contain pineapples")
	}
} 