package wrapper

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	AccessToken string
	TokenType   string
	ExpiresIn   string
	UserID      int64
}
