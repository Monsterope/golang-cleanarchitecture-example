package controllers

import (
	"cleanarchitecture-example/modules/responses"
	"cleanarchitecture-example/pkg/databases"
	"cleanarchitecture-example/pkg/utils"
)

type Controller struct {
	DBConfig   *databases.DatabaseConfig
	RedisStore *utils.RedisAuthStore
}

func NewController(dbconfig *databases.DatabaseConfig, redisstore *utils.RedisAuthStore) *Controller {
	return &Controller{
		DBConfig:   dbconfig,
		RedisStore: redisstore,
	}
}

type (
	ResponseSuccessLogin struct {
		Status       string                 `json:"status"`
		AccessToken  string                 `json:"access_token"`
		RefreshToken string                 `json:"refresh_token"`
		User         responses.SafeCustomer `json:"user"`
	}
	ResponseSuccessRefresh struct {
		AccessToken string `json:"access_token"`
	}
	ResponseSuccess struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}
	ResponseFailure struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)

func ResponseSuccessLoginData(status string, accessToken string, refreshToken string, user responses.SafeCustomer) ResponseSuccessLogin {
	return ResponseSuccessLogin{
		Status:       status,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}
}

func ResponseSuccessRefreshData(accessToken string) ResponseSuccessRefresh {
	return ResponseSuccessRefresh{
		AccessToken: accessToken,
	}
}

func ResponseSuccessData(status string, data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Status: status,
		Data:   data,
	}
}

func ResponseFailureData(status string, message string) ResponseFailure {
	return ResponseFailure{
		Status:  status,
		Message: message,
	}
}
