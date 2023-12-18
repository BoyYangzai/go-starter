package service

import (
	"fmt"
	"reflect"
	"time"

	"github.com/BoyYangZai/go-server-lib/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	CreatedTime     time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime     time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_time"`
	Password        string    `gorm:"type:varchar(80);not null" json:"password"`
	Username        string    `gorm:"type:varchar(80);not null" json:"username"`
	Email           string    `gorm:"type:varchar(255);not null" json:"email"`
	EmailVerify     int       `gorm:"type:tinyint;not null;default:0" json:"email_verify"`
	EmailVerifyCode string    `gorm:"type:varchar(80);not null" json:"email_verify_code"`
	AvatarURL       string    `gorm:"type:varchar(255);not null" json:"avatar_url"`
	Extra           string    `gorm:"type:varchar(255);not null" json:"extra"`
}

func UpdateVerifyCode(email string, code string) {
	db := database.Db
	user := User{}

	result := db.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		// 表示未找到匹配的记录
		newUser := &User{Email: email, EmailVerifyCode: code}
		db.Create(&newUser)
	} else if result.Error != nil {
		// 发生其他错误
		fmt.Println("Query error:", result.Error)
	} else {
		// 找到了匹配的记录

		userEmail := email
		newEmailVerifyCode := code

		// 使用 Update 更新已有用户的 EmailVerifyCode

		UpdOneKeyWhereAnoKey("email", userEmail, "email_verify_code", newEmailVerifyCode)

		fmt.Printf("Email '%s' found in the users table\n", user.Email)
	}
}

func UpdOneKeyWhereAnoKey(whereKey string, whereKeyValue any, changeKey string, changeKeyValue any) {
	db := database.Db
	result := db.Model(&User{}).Where(whereKey+" = ?", whereKeyValue).Update(changeKey, changeKeyValue)
	if result.Error != nil {
		fmt.Println("Error updating key:", result.Error)
		return
	}
}

func MatchEmailAndKey(email string, keyValue string, matchKey string) (bool, User) {
	db := database.Db
	user := User{}
	result := db.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		// 表示未找到匹配的记录
		fmt.Printf("Email '%s' not found in the users table\n", email)
		return false, user
	} else if result.Error != nil {
		// 发生其他错误
		fmt.Println("Query error:", result.Error)
		return false, user
	}
	fieldValue, found := getField(&user, matchKey)
	if found {
		fmt.Printf("%s 的值是：%v\n", matchKey, fieldValue)
	} else {
		fmt.Printf("未找到字段：%s\n", matchKey)
	}
	return fieldValue == keyValue, user
}

func InitUser(email string, password string) {
	db := database.Db
	user := User{}
	db.Where("email = ?", email).First(&user)
	username := GenerateRandomString(10)
	db.Model(&user).Updates(User{Username: username, CreatedTime: time.Now(), Password: password, EmailVerify: 1, AvatarURL: getRandomAvatarURL(username)})
	db.Save(&user)
}

func getField(obj interface{}, fieldName string) (interface{}, bool) {

	val := reflect.ValueOf(obj)

	// 如果传递的是指针，获取其指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 确保 obj 是结构体
	if val.Kind() != reflect.Struct {
		return nil, false
	}

	// 获取字段
	field := val.FieldByName(fieldName)

	if !field.IsValid() {
		return nil, false
	}

	return field.Interface(), true
}
