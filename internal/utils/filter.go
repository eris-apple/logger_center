package utils

import (
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Filter struct {
	Limit  int
	Offset int
	Order  string
}

func GetDefaultsFilter(filter *Filter, prefix ...string) *Filter {
	if filter.Limit == 0 {
		filter.Limit = 100
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

func FilterArray[T any](data []T, filter *Filter) []T {
	dataValue := reflect.ValueOf(data)
	log.Print()
	if dataValue.Kind() != reflect.Slice {
		panic("array is not a slice")
	}

	if filter.Order != "" {
		order := strings.Split(filter.Order, " ")
		key := order[0]
		sortType := order[1]

		sort.Slice(data, func(i, j int) (less bool) {
			elem1 := dataValue.Index(i).FieldByName(key).Int()
			elem2 := dataValue.Index(j).FieldByName(key).Int()

			if sortType == "desc" {
				return elem1 > elem2
			} else {
				return elem1 < elem2
			}
		})
	}

	length := dataValue.Len()

	start := filter.Offset
	end := filter.Offset + filter.Limit
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	return data[start:end]
}
