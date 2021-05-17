package rest

import (
	"net/http"

	"github.com/chandantechpanion/adminpanel-service/pkg/adminUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"github.com/chandantechpanion/adminpanel-service/pkg/login"

	"github.com/gin-gonic/gin"
)

//LoginCtrl for login
type LoginCtrl struct {
	log logger.LogInfoFormat
	svc login.Service
}

//NewLoginCtrl from buyer user info table
func NewLoginCtrl(log logger.LogInfoFormat, svc login.Service) *LoginCtrl {
	return &LoginCtrl{log, svc}
}

//AccessToken ...
// Login godoc
// @Summary Get Access Token
// @Description Pass valid credentials and Get the access token
// @Tags Login
// @Accept  json
// @Produce  json
// @Param login body adminUserInfo.Login true "Login"
// @Success 200 {string} string "token"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /access_token/ [post]
func (l LoginCtrl) AccessToken(ctx *gin.Context) {
	lg := adminUserInfo.Login{}
	if err := ctx.ShouldBindJSON(&lg); err != nil {
		l.log.Errorf("signin request bind failed with reason : %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, errToken := l.svc.SignIn(lg.LoginUserName, lg.Password)
	if errToken != nil {
		ctx.JSON(http.StatusUnauthorized, errToken)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": access_token})
}
