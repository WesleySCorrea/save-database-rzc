package models

type Client struct {
	ID             int     `db:"id" json:"id"`
	Name           string  `db:"name" json:"name"`
	Surname        string  `db:"surname" json:"surname"`
	ContractedPlan string  `db:"contracted_plan" json:"contractedPlan"`
	Payment        bool    `db:"payment" json:"payment"`
	Files          []Files `gorm:"foreignKey:ClientID" json:"files"`
}
