package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
)

func CreateStore(store models.Store) error {
	storeCheck, err := repository.GetStoreByName(store.Name)
	if storeCheck.ID != 0 {
		return errs.ErrStoreNameUniquenessFailed
	}

	if err != nil {
		return err
	}

	err = repository.CreateStore(&store)
	if err != nil {
		return err
	}

	return nil
}
