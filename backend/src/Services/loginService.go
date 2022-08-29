package Services

import (
	"encoding/json"
	"strconv"
	"time"

	model "../Models"
	"github.com/dgrijalva/jwt-go"
)

func Login(body map[string]string) []byte {
	response := make(map[string]string)
	user := model.GetUser(body["user_name"])
	if body["user_name"] != user[0].User_name || body["password"] != user[0].Password {
		response["error"] = "Invalid username or password."
		response["result"] = "failed"
		encodedJson, _ := json.Marshal(response)
		return encodedJson
	}
	if body["user_name"] == user[0].User_name || body["password"] == user[0].Password {
		claims := jwt.StandardClaims{
			Issuer:    strconv.Itoa(user[0].User_id),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, _ := jwtToken.SignedString([]byte("secret"))

		response["token"] = token
		response["user_id"] = strconv.Itoa(user[0].User_id)
		response["user_name"] = user[0].User_name
		response["mail_address"] = user[0].Mail_address
		response["admin_flag"] = strconv.Itoa(user[0].Admin_flag)
		response["result"] = "success"
	}
	encodedJson, _ := json.Marshal(response)

	return encodedJson
}
