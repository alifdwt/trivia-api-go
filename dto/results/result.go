package dto

type SuccessResult struct {
	Code    int    		`json:"code" example:"200"`
	Data 	interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int    		`json:"code" example:"400"`
	Message string 		`json:"message" example:"Bad Request"`
}

