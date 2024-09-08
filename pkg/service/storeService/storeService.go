package storeService

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/repository/storeRepository"
)

func CreateStore(store models.Store) error {
	storeCheck, err := storeRepository.GetStoreByName(store.Name)
	if storeCheck.ID != 0 {
		return errs.ErrStoreNameUniquenessFailed
	}

	if err != nil {
		return err
	}

	err = storeRepository.CreateStore(&store)
	if err != nil {
		return err
	}

	return nil
}
