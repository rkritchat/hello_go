package model

type Student struct{
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int8   `json:"age"`
}

type AbsResponse struct {
	StatusCode string  `json:"statusCode"`
	StatusDesc string  `json:"statusDesc"`
}