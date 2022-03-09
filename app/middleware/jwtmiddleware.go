package middleware

import (
	"gosagaapi/app/global"
	"github.com/gin-gonic/gin"
)

type JwtMiddleware struct {
}

var exception = map[string]byte{
	"/home":  0,
	"/login": 0,
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{}
}

func (m *JwtMiddleware) Handle() gin.HandlerFunc {
	return global.Jwt.Middleware(exception)
}
