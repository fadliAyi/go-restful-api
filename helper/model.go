package helper

import (
	"go-restful-api/model/domain"
	"go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   int(category.Id),
		Name: category.Name,
	}
}
