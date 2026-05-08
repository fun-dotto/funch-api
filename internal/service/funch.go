package service

import (
	"context"
	"time"

	"github.com/fun-dotto/funch-api/internal/domain"
)

type MenuItemRepository interface {
	GetMenuItemsByDate(ctx context.Context, date time.Time) ([]domain.MenuItem, error)
}

type MenuItemService struct {
	menuItemRepository MenuItemRepository
}

func NewMenuItemService(menuItemRepository MenuItemRepository) *MenuItemService {
	return &MenuItemService{menuItemRepository: menuItemRepository}
}

func (s *MenuItemService) GetMenuItemsByDate(ctx context.Context, date time.Time) ([]domain.MenuItem, error) {
	return s.menuItemRepository.GetMenuItemsByDate(ctx, date)
}
