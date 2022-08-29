package Services

import (
	model "../Models"
)

func CreateUser(body map[string]string) bool {
	dbResult := model.CreateUser(body)
	return dbResult
}
