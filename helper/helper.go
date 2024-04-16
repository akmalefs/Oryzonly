package helper

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
