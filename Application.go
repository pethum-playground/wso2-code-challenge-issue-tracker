package main

type Application struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (Application) TableName() string {
	return "Application"
}
