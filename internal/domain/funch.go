package domain

import "time"

type Category string

const (
	CategorySetAndSingle Category = "SetAndSingle"
	CategoryBowlAndCurry Category = "BowlAndCurry"
	CategoryNoodle       Category = "Noodle"
	CategorySide         Category = "Side"
	CategoryDessert      Category = "Dessert"
)

type Size string

const (
	SizeSmall  Size = "Small"
	SizeMedium Size = "Medium"
	SizeLarge  Size = "Large"
)

type MenuItem struct {
	ID       string
	Date     time.Time
	Name     string
	ImageURL string
	Category Category
	Prices   []Price
}

type Price struct {
	Size  Size
	Price int32
}
