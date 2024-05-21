package handler

import (
	"github.com/Haibarapink/OpenDictionary/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService model.UserService
}

type Config struct {
	R           *gin.Engine
	UserService model.UserService
}

func NewHandler(c *Config) {
	h := &Handler{
		UserService: c.UserService,
	}
	g := c.R.Group("/api/account")
	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}
