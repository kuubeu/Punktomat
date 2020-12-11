package model

import "gorm.io/gorm"

// ScienceMagazine is a bulleted scientific journal structure
type ScienceMagazine struct {
	gorm.Model
	ID                     int    `json:"id"`
	Title                  string `json:"title"`
	Issn                   string `json:"issn"`
	Eissn                  string `json:"e-issn"`
	SecondTitle            string `json:"secondTitle"`
	SecondIssn             string `json:"secondIssn"`
	SecondEissn            string `json:"secondE-issn"`
	Points                 int    `json:"points"`
	Archeology             bool   `json:"archeology"`
	Philosophy             bool   `json:"philosophy"`
	History                bool   `json:"history"`
	Linguistics            bool   `json:"linguistics"`
	Literary               bool   `json:"literary"`
	CultureAndReligion     bool   `json:"cultureAndReligion"`
	ArtsSciences           bool   `json:"artsSciences"`
	Architecture           bool   `json:"architecture"`
	AutomaticsElectronics  bool   `json:"automaticsElectronics"`
	TechnicalInformatics   bool   `json:"technicalInformatics"`
	Biomedical             bool   `json:"biomedical"`
	ChemicalEngineering    bool   `json:"chemicalEngineering"`
	CivilAndTransportation bool   `json:"civilAndTransportation"`
	Material               bool   `json:"material"`
	Mechanical             bool   `json:"mechanical "`
	EnvironmentalAndMining bool   `json:"environmentalAndMining"`
	Pharmaceutical         bool   `json:"pharmaceutical"`
	Medical                bool   `json:"medical"`
	PhysicalCulture        bool   `json:"physicalCulture"`
	Health                 bool   `json:"health"`
	Forest                 bool   `json:"forest"`
	Agriculture            bool   `json:"agriculture"`
	FoodAndNutrition       bool   `json:"foodAndNutrition"`
	Veterinary             bool   `json:"veterinary"`
	Zootechnics            bool   `json:"zootechnics"`
	Economics              bool   `json:"economics"`
	SocioEconomic          bool   `json:"socioEconomic"`
	Security               bool   `json:"security"`
	SocialCommunication    bool   `json:"socialCommunication"`
	Political              bool   `json:"political"`
	Management             bool   `json:"management"`
	Legal                  bool   `json:"legal"`
	Sociological           bool   `json:"sociological"`
	Education              bool   `json:"education"`
	CanonicLaw             bool   `json:"canonicLaw"`
	Psychology             bool   `json:"psychology"`
	Astronomy              bool   `json:"astronomy"`
	Informatics            bool   `json:"informatics"`
	Maths                  bool   `json:"maths"`
	Biological             bool   `json:"biological"`
	Chemical               bool   `json:"chemical"`
	Physical               bool   `json:"physical"`
	EarthAndEnvironment    bool   `json:"earthAndEnvironment"`
	Theological            bool   `json:"theological"`
}
