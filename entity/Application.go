package entity

type Application struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name,omitempty" gorm:"type:varchar(50);not null;unique"`
	Value string `json:"value,omitempty" gorm:"type:longtext;nullable"`
}

func (Application) TableName() string {
	return "Application"
}
