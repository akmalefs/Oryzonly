package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Massage string `json:"massage"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseJSON(Massage string, Code int, Status string, Data interface{}) Response {
	Meta := Meta{
		Massage: Massage,
		Code:    Code,
		Status:  Status,
	}

	responseJSON := Response{
		Meta: Meta,
		Data: Data,
	}

	return responseJSON
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
