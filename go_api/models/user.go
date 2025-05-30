package models

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Books    []Book `json:"books" gorm:"foreignKey:UserID"`
}

const (
	RoleAdmin  = "admin"
	RoleUser   = "user"
	RoleSeller = "seller"
)
