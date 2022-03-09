package logic

import (
	"gosagaapi/app/agreement"
	"gosagaapi/app/entity"
	"gosagaapi/app/global"
	"fmt"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

type AuthLogic struct {
}

var (
	authLogic     *AuthLogic
	authLogicOnce sync.Once
)

func NewAuthLogic() *AuthLogic {
	if authLogic == nil {
		authLogicOnce.Do(func() {
			authLogic = &AuthLogic{}
		})
	}

	return authLogic
}

func (l *AuthLogic) Auth(ctx *gin.Context, params agreement.AuthRequest) (result agreement.AuthResponse, err error) {
	expiresAt := time.Now().Add(time.Second * config.GetDuration("Jwt.Ttl"))
	user := entity.User{
		Id:    1,
		Name:  "ZhangSan",
		Email: params.Email,
		Age:   25,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token, err := global.Jwt.GenerateToken(user)
	if err != nil {
		logger.ErrorfX(ctx, "generate token failed: user=%v err=%v", user, err)
		return result, err
	}

	result = agreement.AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt.Unix(),
	}

	return result, nil
}

func (l *AuthLogic) Refresh(ctx *gin.Context) (response agreement.RefreshResponse, err error) {
	expiresAt := time.Now().Add(time.Second * config.GetDuration("Jwt.Ttl"))
	user := entity.User{
		Id:    1,
		Name:  "ZhangSan",
		Email: "xxx",
		Age:   25,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token, err := global.Jwt.GenerateToken(user)
	if err != nil {
		logger.ErrorfX(ctx, "generate token failed: user=%v err=%v", user, err)
		return response, err
	}

	response = agreement.RefreshResponse{
		Token:     token,
		ExpiresAt: expiresAt.Unix(),
	}

	return response, nil
}

func (l *AuthLogic) Verify(claims jwt.Claims) bool {
	// 填充验证逻辑
	fmt.Println(claims)

	return true
}

func (l *AuthLogic) GetEmptyMyClaims() jwt.Claims {
	return &entity.User{}
}
