package api

import (

	"github.com/chunganhbk/gin-go/internal/app/services"
	"github.com/chunganhbk/gin-go/internal/app/schema"
	"github.com/chunganhbk/gin-go/pkg/app"

	"github.com/chunganhbk/gin-go/pkg/logger"
	"github.com/gin-gonic/gin"

)



// Login
type Auth struct {
	AuthService services.IAuthService
	UserService services.IUserService
}
func NewAuth(authService services.IAuthService, userService services.IUserService) *Auth {
	return &Auth{AuthService:authService, UserService:userService}
}

// Login
func (a *Auth) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.LoginParam
	if err := app.ParseJSON(c, &item); err != nil {
		app.ResError(c, err)
		return
	}


	user, err := a.AuthService.Verify(ctx, item.UserName, item.Password)
	if err != nil {
		app.ResError(c, err)
		return
	}

	userID := user.ID

	app.SetUserID(c, userID)

	ctx = logger.NewUserIDContext(ctx, userID)
	tokenInfo, err := a.AuthService.GenerateToken(userID)
	if err != nil {
		app.ResError(c, err)
		return
	}

	logger.StartSpan(ctx, logger.SetSpanTitle("User login"), logger.SetSpanFuncName("Login")).Infof("Login system")
	app.ResSuccess(c, tokenInfo)
}

// RefreshToken
func (a *Auth) RefreshToken(c *gin.Context) {
	tokenInfo, err := a.AuthService.GenerateToken(app.GetUserID(c))
	if err != nil {
		app.ResError(c, err)
		return
	}
	app.ResSuccess(c, tokenInfo)
}
func (a *Auth) Register(c *gin.Context){
	ctx := c.Request.Context()
	var item schema.User
	if err := app.ParseJSON(c, &item); err != nil {
		app.ResError(c, err)
		return
	} else if item.Password == "" {
		app.ResError(c, app.New400Response(app.ERROR_PASSWORD_REQUIRED, nil))
		return
	}

	result, err := a.UserService.Create(ctx, item)
	if err != nil {
		app.ResError(c, err)
		return
	}
	app.ResSuccess(c, result)
}
