package handler

import (
	"github.com/fun-dotto/funch-api/internal/service"
)

type Handler struct {
	menuItemService *service.MenuItemService
}

func NewHandler(menuItemService *service.MenuItemService) *Handler {
	return &Handler{menuItemService: menuItemService}
}

