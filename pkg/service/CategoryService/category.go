package CategoryService

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/repository/categoryRepository"
)

func CreateCategory(categ models.Category) (categID uint, err error) {
	var category models.Category

	category, _ = categoryRepository.GetCategoryByName(categ.CategoryName)
	if category.ID != 0 {
		return category.ID, errs.ErrCategoryNameUniquenessFailed
	}

	if categID, err = categoryRepository.CreateCategory(categ); err != nil {
		return 0, err
	}

	return categID, nil
}

func UpdateCategory(categoryID uint, categ models.Category) (categID uint, err error) {
	var category models.Category

	category, _ = categoryRepository.GetCategoryByName(categ.CategoryName)
	if category.ID != 0 {
		return category.ID, errs.ErrCategoryNameUniquenessFailed
	}

	if categID, err = categoryRepository.UpdateCategory(categoryID, categ); err != nil {
		return 0, err
	}

	return categID, nil
}
