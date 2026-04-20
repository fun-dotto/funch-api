package repository

import (
	"context"
	"time"

	"github.com/fun-dotto/funch-api/internal/database"
	"github.com/fun-dotto/funch-api/internal/domain"
	"gorm.io/gorm"
)

type menuItemRepository struct {
	db *gorm.DB
}

func NewMenuItemRepository(db *gorm.DB) *menuItemRepository {
	return &menuItemRepository{db: db}
}

func (r *menuItemRepository) GetMenuItemsByDate(ctx context.Context, date time.Time) ([]domain.MenuItem, error) {
	var dbMenuItems []database.MenuItem
	if err := r.db.WithContext(ctx).Where("date = ?", date).Find(&dbMenuItems).Error; err != nil {
		return nil, err
	}

	domainMenuItems := make([]domain.MenuItem, len(dbMenuItems))
	for i, dbMenuItem := range dbMenuItems {
		var dbPrices []database.MenuItemPrice
		if err := r.db.WithContext(ctx).Where("menu_item_id = ?", dbMenuItem.ID).Find(&dbPrices).Error; err != nil {
			return nil, err
		}
		domainMenuItems[i] = dbMenuItem.ToDomain(dbPrices)
	}

	return domainMenuItems, nil
}
