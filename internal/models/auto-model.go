package models

type AutoStruct struct {
	MaxSpeed int    `json:"max_speed"`
	Distance int    `json:"distance"`
	Handler  string `json:"handler"`
	Stock    string `json:"stock"`
}
