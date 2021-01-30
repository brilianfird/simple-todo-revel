package model

import "github.com/revel/revel"

type Response struct {
	StatusCode int `json:"statusCode"`
	Error []*revel.ValidationError `json:"error"`
	Data interface{} `json:"data"`
}