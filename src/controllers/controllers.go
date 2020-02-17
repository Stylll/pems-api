package controllers

import "github.com/gin-gonic/gin"

func Default(c *gin.Context) {
	c.JSON(200, gin.H{
		"PEMS": "ALIVE",
	})
}

func WelcomeAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"PEMS API": "WELCOME TO PEMS API",
	})
}
