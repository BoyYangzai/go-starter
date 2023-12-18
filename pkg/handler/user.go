package handler

import (
	"go-app/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyCodeRquest struct {
	Email string `json:"email"`
}

func VerifyCode(c *gin.Context) {
	var requestBody VerifyCodeRquest

	// 通过 ShouldBindJSON 解析 JSON 请求体到结构体
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code := generateVerificationCode(6)
	user := "1484502768@qq.com"
	password := "jyderqttsuyyiagf"
	host := "smtp.qq.com:587"
	to := requestBody.Email
	subject := "verifycode:"
	body := `
	<html>
	<body>
	<h3>
	` + code +
		`
	</h3>
	</body>
	</html>
	`
	println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		println("send mail error!")
		println(err)
	} else {
		println("send mail success!")
	}

	service.UpdateVarifyCode(to, code)
	c.JSON(http.StatusOK, gin.H{
		"msg": "verifyCode sent",
	})
}

type RegistryRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

func Registry(c *gin.Context) {
	var requestBody RegistryRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isMatched := service.MatchEmailAndKey(requestBody.Email, requestBody.VerifyCode, "EmailVerifyCode")
	if isMatched {
		service.InitUser(requestBody.Email, requestBody.Password)
		c.JSON(http.StatusOK, gin.H{
			"msg": "registry success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "email and verifyCode not match",
		})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var requestBody LoginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isMatched := service.MatchEmailAndKey(requestBody.Email, requestBody.Password, "Password")
	if isMatched {
		c.JSON(http.StatusOK, gin.H{
			"msg": "login success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "email and password not match",
		})
	}
}
