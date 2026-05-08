package database

import (
	"time"

	"github.com/fun-dotto/funch-api/internal/domain"
)

type MenuItem struct {
	ID       string    `gorm:"primaryKey;type:uuid"`
	Date     time.Time `gorm:"not null;index;type:date"`
	Name     string    `gorm:"not null"`
	ImageURL string    `gorm:"not null"`
	Category string    `gorm:"not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type MenuItemPrice struct {
	ID         string `gorm:"primaryKey;type:uuid"`
	MenuItemID string `gorm:"not null;index;type:uuid"`
	Size       string `gorm:"not null"`
	Price      int32  `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	MenuItem MenuItem `gorm:"foreignKey:MenuItemID"`
}

func (m *MenuItem) ToDomain(prices []MenuItemPrice) domain.MenuItem {
	domainPrices := make([]domain.Price, len(prices))
	for i, p := range prices {
		domainPrices[i] = domain.Price{
			Size:  domain.Size(p.Size),
			Price: p.Price,
		}
	}

	return domain.MenuItem{
		ID:       m.ID,
		Date:     m.Date,
		Name:     m.Name,
		ImageURL: m.ImageURL,
		Category: domain.Category(m.Category),
		Prices:   domainPrices,
	}
}

func FromDomainMenuItem(menuItem domain.MenuItem) MenuItem {
	return MenuItem{
		ID:       menuItem.ID,
		Date:     menuItem.Date,
		Name:     menuItem.Name,
		ImageURL: menuItem.ImageURL,
		Category: string(menuItem.Category),
	}
}

func FromDomainMenuItemPrices(menuItemID string, prices []domain.Price) []MenuItemPrice {
	result := make([]MenuItemPrice, len(prices))
	for i, p := range prices {
		result[i] = MenuItemPrice{
			MenuItemID: menuItemID,
			Size:       string(p.Size),
			Price:      p.Price,
		}
	}
	return result
}
