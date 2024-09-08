package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
)

func GetStoreByID(storeID uint) (models.Store, error) {
	var store models.Store
	store, err := repository.GetStoreByID(storeID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return store, errs.ErrStoreNotFound
		}
		
		return store, err
	}

	return store, nil
}

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
