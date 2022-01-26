package shoppinglist

import "foodHelper/data"

type Item struct {
	volume     data.Volume
	ingredient data.Ingredient
}

type ShoppingList struct {
	//id -> item
	items map[string]Item[]
}
