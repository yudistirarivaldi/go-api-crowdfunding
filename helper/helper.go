package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` //agar nilai nya bebas berubah dan fleksibel // data bisa bernilai apa saja
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse

}