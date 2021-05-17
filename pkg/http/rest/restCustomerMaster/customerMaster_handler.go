package restCustomerMaster

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerMaster"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type customerMasterCtrl struct {
	log logger.LogInfoFormat
	svc customerMaster.Service
}

func NewCustomerMasterCtrl(log logger.LogInfoFormat, svc customerMaster.Service) *customerMasterCtrl {
	return &customerMasterCtrl{log, svc}
}

// CustomerMaster godoc
// @Summary Save Customer Info
// @Description Here Save customer info.
// @Tags CustomerMaster
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param cMaster body customerMaster.CustomerMasterViewModel true "Customer Details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /customerMaster/ [post]
func (c *customerMasterCtrl) Save(ctx *gin.Context) {

	custMaster := customerMaster.CustomerMasterViewModel{}

	if err := ctx.ShouldBindJSON(&custMaster); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cMaster := customerMaster.CustomerMaster{}
	cMaster.Name = custMaster.Name
	cMaster.CreateDate = time.Now()

	err := c.svc.Save(&cMaster)
	if err == nil {
		ctx.Status(http.StatusCreated)
	} else if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// CustomerMaster
// @Summary Get Customer
// @Description Here Get customer by id
// @Tags CustomerMaster
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Success 200  {array} customerMaster.CustomerMaster
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerMaster/ [get]
func (c *customerMasterCtrl) Get(ctx *gin.Context) {

	customerList, err := c.svc.Get()
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"customerList": customerList})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}

// CustomerMaster godoc
// @Summary delete particular customer
// @Description Here we can remve customer
// @Tags CustomerMaster
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path string true "Customer Master ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerMaster/{id} [delete]
func (ug *customerMasterCtrl) Delete(ctx *gin.Context) {
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
