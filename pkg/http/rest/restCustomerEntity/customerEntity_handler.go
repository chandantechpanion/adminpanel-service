package restCustomerEntity

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerEntity"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type customerEntityCtrl struct {
	log logger.LogInfoFormat
	svc customerEntity.Service
}

func NewCustomerEntityCtrl(log logger.LogInfoFormat, svc customerEntity.Service) *customerEntityCtrl {
	return &customerEntityCtrl{log, svc}
}

// CustomerEntity godoc
// @Summary Save Customer Info
// @Description Here Save customer info.
// @Tags CustomerEntity
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param cEntity body customerEntity.CustomerEntityViewModel true "Customer Details"
// @Success 200
// @Created 201
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /customerEntity/ [post]
func (c *customerEntityCtrl) Save(ctx *gin.Context) {

	custEntity := customerEntity.CustomerEntityViewModel{}
	// i, _ := strconv.ParseInt(custEntity.ID, 10, 64)
	if err := ctx.ShouldBindJSON(&custEntity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cEntity := customerEntity.CustomerEntity{}
	cEntity.Name = custEntity.Name
	cEntity.ID = custEntity.ID
	cEntity.CustomerMasterID = custEntity.CustomerMasterID
	cEntity.CreatedDate = time.Now()

	err := c.svc.Save(&cEntity)
	if err == nil {
		ctx.Status(http.StatusCreated)
	} else if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// CustomerEntity
// @Summary Get Customer
// @Description Here Get entity by id
// @Tags CustomerEntity
// @Accept  json
// @Produce  json
// @Success 200  {array} customerEntity.CustomerEntity
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerEntity/ [get]
// func (c *customerEntityCtrl) GetAll(ctx *gin.Context) {

// 	customerEntityList, err := c.svc.GetAll()
// 	if err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{"customerEntityList": customerEntityList})
// 	} else {
// 		ctx.Status(http.StatusNoContent)
// 	}
// }

// CustomerEntity godoc
// @Summary Get particular entity
// @Description Here we can remve entity
// @Tags CustomerEntity
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param customer_master_id path string true "Customer Entity ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerEntity/{customer_master_id} [get]
func (ug *customerEntityCtrl) Get(ctx *gin.Context) {
	id := ctx.Param("customer_master_id")
	customerMasterId, _ := strconv.ParseInt(id, 10, 64)
	if customerMasterId < 1 {
		ctx.Status(http.StatusBadRequest)
		return
	}

	customerEntityList, err := ug.svc.Get(customerMasterId)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"customerEntityList": customerEntityList})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
