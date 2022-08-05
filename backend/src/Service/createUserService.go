package Service

import (
	model "../Model"
)

func CreateUser(body map[string]string) bool {
	dbResult := model.CreateUser(body)
	return dbResult
}
