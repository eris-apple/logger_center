package utils

import (
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Filter struct {
	Limit  int
	Offset int
	Order  string
}

func GetDefaultsFilter(filter *Filter, prefix ...string) *Filter {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Order == "" {
		if len(prefix) > 0 {
			filter.Order = prefix[0] + ".id desc"
		} else {
			filter.Order = "id desc"
		}
	}

	return &Filter{
		Limit:  filter.Limit,
		Offset: filter.Offset,
		Order:  filter.Order,
	}
}

func GetDefaultsFilterFromQuery(ctx *gin.Context, prefix ...string) *Filter {
	var limitInt int
	var offsetInt int

	limit := ctx.Query("limit")
	if limit != "" {
		li, err := strconv.Atoi(limit)
		if err != nil {
			ErrorResponseHandler(ctx, http.StatusInternalServerError, config.ErrBadRequest)
			return nil
		}
		limitInt = li
	}
	offset := ctx.Query("offset")
	if offset != "" {
		oi, err := strconv.Atoi(offset)
		if err != nil {
			ErrorResponseHandler(ctx, http.StatusInternalServerError, config.ErrBadRequest)
			return nil
		}
		offsetInt = oi
	}

	order := ctx.Query("order")
	if order == "" {
		order = "id desc"
	}
	if len(prefix) > 0 {
		order = prefix[0] + "." + order
	}

	f := &Filter{
		Limit:  limitInt,
		Offset: offsetInt,
		Order:  order,
	}

	return GetDefaultsFilter(f)

}
