package handler

type responseList struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}
