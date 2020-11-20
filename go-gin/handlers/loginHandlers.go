package handlers

import (
	"fmt"
	models "go-gin/models"

	service "go-gin/services"

	"github.com/gin-gonic/gin"
)

type test_struct struct {
	Test string
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInfo models.UserInfo
		var postParam models.Login
		// var t test_struct
		// fmt.Println(`chck this `, c.Param())
		// reqeustBody, _ := ioutil.ReadAll(c.Request.Body)
		// fmt.Println("%s", string(reqeustBody))
		// c.ShouldBindJSON(&postParam)
		c.Bind(&postParam)
		// num, _ := c.GetRawData()
		// buf := make([]byte, 1024)
		// decoder := json.NewDecoder(c.Request.Body)
		// err := decoder.Decode(&t)
		// if err != nil {
		// 	panic(err.Error())
		// }
		// println(`check the postparam **`, t.Test)
		userInfo = service.Login(postParam)

		c.Header("Authorization", userInfo.Token)
		c.JSON(200, userInfo)
	}
}

func SaveUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postParam models.UserDetails
		c.Bind(&postParam)
		fmt.Println(`check the postParam`, postParam)
		service.Save(postParam)
	}
}
