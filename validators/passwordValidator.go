package validators

import (
	"BizMart/validators/utils"
)

func Password(password string) bool {
	specialSymbols := utils.GetValidateSymbols()
	alphabetSymbols := utils.GetAlphabetSymbols()

	for _, char := range password {
		// Проверка на специальный символ
		for _, sym := range specialSymbols {
			if string(char) == sym {
				return false
			}
		}

		// Проверка на принадлежность к алфавитным символам (например, A-Z, a-z)
		for _, sym := range alphabetSymbols {
			if string(char) == sym {
				return false
			}
		}
	}

	return true
}
