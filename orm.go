package main

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Note   string `json:"note"`
	IsDone string `json:"is_done"`
}
