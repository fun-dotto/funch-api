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

	if len(dbMenuItems) == 0 {
		return []domain.MenuItem{}, nil
	}

	menuItemIDs := make([]uint, len(dbMenuItems))
	for i, dbMenuItem := range dbMenuItems {
		menuItemIDs[i] = dbMenuItem.ID
	}

	var dbPrices []database.MenuItemPrice
	if err := r.db.WithContext(ctx).Where("menu_item_id IN ?", menuItemIDs).Find(&dbPrices).Error; err != nil {
		return nil, err
	}

	pricesByMenuItemID := make(map[uint][]database.MenuItemPrice, len(dbMenuItems))
	for _, dbPrice := range dbPrices {
		pricesByMenuItemID[dbPrice.MenuItemID] = append(pricesByMenuItemID[dbPrice.MenuItemID], dbPrice)
	}

	domainMenuItems := make([]domain.MenuItem, len(dbMenuItems))
	for i, dbMenuItem := range dbMenuItems {
		domainMenuItems[i] = dbMenuItem.ToDomain(pricesByMenuItemID[dbMenuItem.ID])
	}

	return domainMenuItems, nil
}
