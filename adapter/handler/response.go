package handler

import (
	"errors"
	"golangTest/pkg/errs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

var errorStatus = map[string]int{
	errs.ErrTaskNotFound.Error():    404,
	errs.ErrInvalidInput.Error():    400,
	errs.ErrDatabase.Error():        500,
	errs.ErrStatusUnchanged.Error(): 400,
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type errorResponse struct {
	Success bool     `json:"success" example:"false"`
	Message []string `json:"message" example:"Error message"`
}

func ResponseCreateSuccess(c *gin.Context, message string, data any) {
	res := response{
		Success: true,
		Message: message,
		Data:    data,
	}
	c.JSON(201, res)
}

func ResponseSuccess(c *gin.Context, data any) {
	res := response{
		Success: true,
		Message: "Success",
		Data:    data,
	}
	c.JSON(200, res)
}

func ResponseError(message []string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}

func handleError(c *gin.Context, err error) {
	code, ok := errorStatus[err.Error()]
	if !ok {
		code = 500
	}
	errMsg := parseError(err)
	errResponse := ResponseError(errMsg)
	c.JSON(code, errResponse)
}
