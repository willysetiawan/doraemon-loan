package middleware

// func ValidateAccessToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := strings.Split(c.Request.Header.Get("Authorization"), " ")
// 		resAccessToken, err := db.CheckAccessToken(token[1])
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"response_code":    strconv.Itoa(http.StatusUnauthorized),
// 				"response_message": http.StatusText(http.StatusUnauthorized),
// 			})
// 			return
// 		}

// 		if resAccessToken.LoginAccessToken == "" || resAccessToken.LoginExpiredAt.Unix() < time.Now().Local().Unix() {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"response_code":    strconv.Itoa(http.StatusUnauthorized),
// 				"response_message": constant.MESSAGE_INVALID_TOKEN,
// 			})
// 			return
// 		}

// 	}
// }
