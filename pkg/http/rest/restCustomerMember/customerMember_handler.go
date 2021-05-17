package restCustomerMember

import (
	"github.com/chandantechpanion/adminpanel-service/pkg/customerMember"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type customerMemberCtrl struct {
	log logger.LogInfoFormat
	svc customerMember.Service
}

func NewCustomerMemberCtrl(log logger.LogInfoFormat, svc customerMember.Service) *customerMemberCtrl {
	return &customerMemberCtrl{log, svc}
}

// CustomerMember godoc
// @Summary Save Customer Info
// @Description Here Save customer info.
// @Tags CustomerMember
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param custM body customerMember.CustomerMemberViewModel true "Customer Details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerMember/ [post]
func (c *customerMemberCtrl) Save(ctx *gin.Context) {

	custM := customerMember.CustomerMemberViewModel{}

	if err := ctx.ShouldBindJSON(&custM); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cMember := customerMember.CustomerMember{}
	cMember.CustomerGroupEntityRelationID = custM.CustomerGroupEntityRelationID
	cMember.UserInfoUserID = custM.UserInfoUserID
	cMember.CreateDate = time.Now()

	err := c.svc.Save(&cMember)
	if err == nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//Inovice Saving after OCR OR PO Selection
// CustomerMember
// @Summary Get Customer
// @Description Here Get customer by id
// @Tags CustomerMember
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path string true "id"
// @Success 200  {array} customerMember.CustomerMember
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerMember/{id} [get]
func (c *customerMemberCtrl) Get(ctx *gin.Context) {

	id := ctx.Param("id")

	if len(id) == 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}
	i, _ := strconv.ParseInt(id, 10, 64)
	customer, err := c.svc.Get(i)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"customer": customer})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}

// CustomerMember godoc
// @Summary delete particular customer
// @Description Here we can remove particular customer
// @Tags CustomerMember
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path string true "Customer Member ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerMember/{id} [delete]
func (ug *customerMemberCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}
	i, _ := strconv.ParseInt(id, 10, 64)
	err := ug.svc.Delete(i)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	} else {
		ctx.Status(http.StatusOK)
	}
}
