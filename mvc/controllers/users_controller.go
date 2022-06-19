package controllers

import (
	"encoding/json"
	"github.com/rmortale/golang-microservices/mvc/services"
	"github.com/rmortale/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		jsonValue, _ := json.Marshal(&utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		})

		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
