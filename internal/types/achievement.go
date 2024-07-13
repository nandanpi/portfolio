package types

import "time"

type AddAchievementReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
}

type UpdateAchievementReq struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
}
