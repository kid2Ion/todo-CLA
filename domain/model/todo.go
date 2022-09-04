package model

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	LimitDate string `json:"limitDate"`
	Status    bool   `json:"status"`
}
