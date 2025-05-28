package model

import "github.com/gosimple/slug"

func Enrich(data *AwesomeData) {
	for i := range *data {
		enrichCategory(&(*data)[i])
	}
}

func enrichCategory(cat *Category) {
	// slug
	if cat.Slug == "" {
		cat.Slug = slug.Make(cat.Title)
	}

	for i := range cat.Subcategories {
		enrichCategory(&cat.Subcategories[i])
	}
}
