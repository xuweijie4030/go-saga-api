package middleware

import (
	"github.com/carefreex-io/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var middlewares = map[string]gin.HandlerFunc{
	"cors": cors.Default(),
	"jwt":  NewJwtMiddleware().Handle(),
}

func Register(r *gin.Engine) {
	for name, middleware := range middlewares {
		if name == "jwt" && !config.GetBool("Jwt.Enable") {
			continue
		}
		if name == "cors" && !config.GetBool("Cors.Enable") {
			continue
		}
		r.Use(middleware)
	}
}
