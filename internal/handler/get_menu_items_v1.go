package handler

import (
	"context"

	api "github.com/fun-dotto/funch-api/generated"
)

func (h *Handler) MenuItemsV1List(ctx context.Context, request api.MenuItemsV1ListRequestObject) (api.MenuItemsV1ListResponseObject, error) {
	date := request.Params.Date.Time

	menuItems, err := h.menuItemService.GetMenuItemsByDate(ctx, date)
	if err != nil {
		return nil, err
	}

	apiMenuItems := make([]api.MenuItem, len(menuItems))
	for i, menuItem := range menuItems {
		apiMenuItems[i] = toApiMenuItem(menuItem)
	}

	return api.MenuItemsV1List200JSONResponse{
		MenuItems: apiMenuItems,
	}, nil
}
