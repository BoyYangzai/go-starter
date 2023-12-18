package service

import (
	"math/rand"
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
