package types

import "time"

type AddExtraCurricularReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
}

type UpdateExtraCurricularReq struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
}
