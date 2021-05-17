package restbuyeruserinfo

import (
	"fmt"
	"net/http"

	"github.com/chandantechpanion/adminpanel-service/pkg/buyerUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//BuyerUserInfoCtrl ...
type BuyerUserInfoCtrl struct {
	log logger.LogInfoFormat
	svc buyerUserInfo.Service
}

//NewBuyerUserInfoCtrl ...
func NewBuyerUserInfoCtrl(log logger.LogInfoFormat, svc buyerUserInfo.Service) *BuyerUserInfoCtrl {
	return &BuyerUserInfoCtrl{log, svc}
}

//GetBuyerUserInfo ...
// BuyerUserInfo godoc
// @Summary Get Particular UserInfo
// @Description Here we will get particular user details
// @Tags BuyerUserInfo
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /buyerUserinfo/{id} [get]
func (bui BuyerUserInfoCtrl) GetBuyerUserInfo(ctx *gin.Context) {
	userID := ctx.Param("id")
	if _, err := uuid.FromString(userID); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := bui.svc.Get(userID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// if user != nil && len(user.MobileNumber) > 0 {
	// 	user.MobileNumber = user.MobileNumber
	// }

	ctx.JSON(http.StatusOK, gin.H{"userFirstName": user.FirstName, "userLastName": user.LastName, "mobileNumber": user.MobileNumber})
}

// BuyerUserInfo godoc
// @Summary Create Buyer User
// @Description Please create the buyer user for providing permission for buyer portal
// @Tags BuyerUserInfo
// @Accept  json
// @Produce  json
// @Param user body buyerUserInfo.CreateBuyerUserInfo true "Create User Details"
// @Success 201
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /buyerUsers/ [post]
func (u *BuyerUserInfoCtrl) Store(ctx *gin.Context) {
	var user buyerUserInfo.CreateBuyerUserInfo

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := u.svc.Save(&user)
	if err != nil {
		ctx.Status(http.StatusCreated)
	} else {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}
