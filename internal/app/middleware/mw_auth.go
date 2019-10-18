package middleware

import (
	"github.com/LyricTian/gin-admin/internal/app/config"
	"github.com/LyricTian/gin-admin/internal/app/errors"
	"github.com/LyricTian/gin-admin/internal/app/ginplus"
	"github.com/LyricTian/gin-admin/pkg/auth"
	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			id, err := a.ParseUserID(t)
			if err != nil {
				if err == auth.ErrInvalidToken {
					ginplus.ResError(c, errors.ErrNoPerm)
					return
				}
				ginplus.ResError(c, errors.WithStack(err))
				return
			}
			userID = id
		}

		if userID != "" {
			c.Set(ginplus.UserIDKey, userID)
		}

		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		if userID == "" {
			if config.Global().IsDebugMode() {
				c.Set(ginplus.UserIDKey, config.Global().Root.UserName)
				c.Next()
				return
			}
			ginplus.ResError(c, errors.ErrNoPerm)
		}
	}
}