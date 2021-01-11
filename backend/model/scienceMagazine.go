package model

import pq "github.com/lib/pq"

// ScienceMagazine is a bulleted scientific journal structure
type ScienceMagazine struct {
	ID        			   uint   		 	`gorm:"primaryKey"`
	Title                  string 		 	`json:"title"`
	Issn                   string 		 	`json:"issn"`
	Eissn                  string 		 	`json:"e-issn"`
	SecondTitle            string 		 	`json:"secondTitle"`
	SecondIssn             string 		 	`json:"secondIssn"`
	SecondEissn            string 		 	`json:"secondE-issn"`
	Points                 int           	`json:"points"`
	Categories		   	   pq.StringArray  	`gorm:"type:text[]"`
}
