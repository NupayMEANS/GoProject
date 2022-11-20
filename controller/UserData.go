package controller

import (
	"fmt"
	"goProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUsers(ctx *gin.Context) {
	var input UserDetails
	fmt.Println(input)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	u := models.UserDetails{}
	u.FirstName = input.FirstName
	u.LastName = input.LastName
	u.Mobile = input.Mobile
	u.Email = input.Email

	fmt.Println(u)

	_, err := u.SaveUser()
	if err != nil {
		fmt.Println("Not able to create User")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User Created successfully",
	})
}
