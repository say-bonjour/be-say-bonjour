package model

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null;type:varchar(255);" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(255);" binding:"required,email"`
	Password string `json:"password" gorm:"not null;type:varchar(255);" binding:"required,gte=5"` //gte > greater than equal
	Role     Role   `json:"role" gorm:"not null;type:varchar(32);"`
}

