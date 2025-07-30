package model

type NumberRequest struct {
    Number int `json:"number" binding:"required"`
}