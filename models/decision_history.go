package models

type DecisionHistory struct {
	ID        string     `json:"id"`
	Status    string     `json:"status"`
	Wave      int        `json:"wave"`
	Reviewer  string     `json:"reviewer"`
	Timestamp int64      `json:"timestamp"`
	History   []Decision `json:"history"`
}
