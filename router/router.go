package router

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	web(r)
}
