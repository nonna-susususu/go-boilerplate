package domain

type Todo struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
