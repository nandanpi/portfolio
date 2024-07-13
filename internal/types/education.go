package types

import "time"

type AddEducationReq struct {
	School    string    `json:"school"`
	Degree    string    `json:"degree"`
	Aggregate string    `json:"aggregate"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type UpdateEducationReq struct {
	ID        uint      `json:"id"`
	School    string    `json:"school"`
	Degree    string    `json:"degree"`
	Aggregate string    `json:"aggregate"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
