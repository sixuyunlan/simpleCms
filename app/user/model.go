package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"simpleCms/app/common"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Work
}

type Work struct {
	Username  string       `gorm:"type:varchar(255);uniqueIndex" json:"username"`
	Password  string       `gorm:"type:varchar(255)" json:"password"`
	Token     string       `gorm:"type:varchar(4096)" json:"token"`
	TokenTIme sql.NullTime `gorm:"" json:"token_time"`
	Name      string       `gorm:"type:varchar(255);index" json:"name"`
}

func init() {
	err := common.DB.AutoMigrate(&User{})
	if err != nil {
		return
	}

	u1 := &User{}
	u1.Username = "root2"
	u1.Password = password("root2")

	u1.Name = "sixu"
	common.DB.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(u1)
}
func password(plain string) string {
	mac := hmac.New(sha256.New, []byte(common.SecretKey))
	mac.Write([]byte(plain))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
