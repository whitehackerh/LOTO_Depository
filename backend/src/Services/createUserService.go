package Services

import (
	model "loto_depository/src/Models"
)

func CreateUser(body map[string]string) bool {
	dbResult := model.CreateUser(body)
	return dbResult
}
