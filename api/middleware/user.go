package middleware

import (
	"backend_ukmik/config"
	"backend_ukmik/domain"
	"backend_ukmik/utils"
	"fmt"
	"net/http"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUserRole(userDomain domain.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status_code": 401, "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfigEnv(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status_code": 401, "message": err.Error()})
			return
		}

		id, _ := strconv.Atoi(fmt.Sprint(sub))
		result, err := userDomain.FindUserById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status_code": 401, "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Set("currentUserId", result.ID)
		ctx.Next()
	}
}
