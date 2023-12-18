package service

import (
	"math/rand"

	"github.com/BoyYangZai/go-server-lib/pkg/database"
	"github.com/BoyYangZai/go-server-lib/pkg/jwt"
)

// 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func getRandomAvatarURL(username string) string {
	// API URL
	apiURL := "https://api.multiavatar.com/"
	return apiURL + username
}

func getAuthUser() User {
	db := database.Db
	db.Where("id = ?", jwt.CurrentAuthUserId).First(&AuthUser)
	return AuthUser
}
