package model

import "gorm.io/gorm"

type ScienceMagazine struct {
	gorm.Model
	ID    int    `json:"id"`
	Title string `json:"title"`
	Issn  string `json:"issn"`
}
