package handler

import (
	openapi_types "github.com/oapi-codegen/runtime/types"

	api "github.com/fun-dotto/funch-api/generated"
	"github.com/fun-dotto/funch-api/internal/domain"
)

func toApiMenuItem(menuItem domain.MenuItem) api.MenuItem {
	apiPrices := make([]api.Price, len(menuItem.Prices))
	for i, price := range menuItem.Prices {
		apiPrices[i] = api.Price{
			Size:  api.Size(price.Size),
			Price: price.Price,
		}
	}

	return api.MenuItem{
		Id:       menuItem.ID,
		Date:     openapi_types.Date{Time: menuItem.Date},
		Name:     menuItem.Name,
		ImageUrl: menuItem.ImageURL,
		Category: api.Category(menuItem.Category),
		Prices:   apiPrices,
	}
}
