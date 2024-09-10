package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
)

func CreateCategory(categ models.Category) (categID uint, err error) {
	var category models.Category

	category, _ = repository.GetCategoryByName(categ.CategoryName)
	if category.ID != 0 {
		return category.ID, errs.ErrCategoryNameUniquenessFailed
	}

	if categID, err = repository.CreateCategory(categ); err != nil {
		return 0, err
	}

	return categID, nil
}

func UpdateCategory(categoryID uint, categ models.Category) (categID uint, err error) {
	if categID, err = repository.UpdateCategory(categoryID, categ); err != nil {
		return 0, err
	}

	return categID, nil
}
