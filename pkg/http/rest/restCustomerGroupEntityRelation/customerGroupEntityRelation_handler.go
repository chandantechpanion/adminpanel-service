package restCustomerGroupEntityRelation

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroupEntityRelation"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type customerGroupEntityRelationCtrl struct {
	log logger.LogInfoFormat
	svc customerGroupEntityRelation.Service
}

func NewCustomerGroupEntityRelationCtrl(log logger.LogInfoFormat, svc customerGroupEntityRelation.Service) *customerGroupEntityRelationCtrl {
	return &customerGroupEntityRelationCtrl{log, svc}
}

// CustomerGroupEntityRelation godoc
// @Summary Save Customer Group EntityRelation Information
// @Description Here Save customer GroupEntityRelation info.
// @Tags CustomerGroupEntityRelation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param custM body customerGroupEntityRelation.CustomerGroupEntityRelationViewModel true "Customer GroupEntityRelation Details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerGroupEntityRelation/ [post]
func (cger *customerGroupEntityRelationCtrl) Save(ctx *gin.Context) {

	custGER := customerGroupEntityRelation.CustomerGroupEntityRelationViewModel{}

	if err := ctx.ShouldBindJSON(&custGER); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cGroupER := customerGroupEntityRelation.CustomerGroupEntityRelation{}
	cGroupER.CustomerEntityID = custGER.CustomerEntityID
	cGroupER.CustomerGroupID = custGER.CustomerGroupID
	cGroupER.CustomerMasterID = custGER.CustomerMasterID
	// cGroupER.ID = custGER.
	cGroupER.CreateDate = time.Now()

	err := cger.svc.Save(&cGroupER)
	if err == nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//Inovice Saving after OCR OR PO Selection
// CustomerGroupEntityRelation
// @Summary Get CustomerGroupEntityRelation
// @Description Here Get customer by id
// @Tags CustomerGroupEntityRelation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path string true "id"
// @Success 200  {array} customerGroupEntityRelation.CustomerGroupEntityRelation
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerGroupEntityRelation/{id} [get]
func (c *customerGroupEntityRelationCtrl) Get(ctx *gin.Context) {

	id := ctx.Param("id")

	if len(id) == 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}
	i, _ := strconv.ParseInt(id, 10, 64)
	customerGER, err := c.svc.Get(i)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"customerGroupEntityRelation": customerGER})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}

// CustomerGroupEntityRelation
// @Summary Get customerGroupEntityRelation
// @Description Here Get customer by id
// @Tags CustomerGroupEntityRelation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Success 200  {array} customerGroupEntityRelation.CustomerGroupEntityRelation
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerGroupEntityRelation/ [get]
func (c *customerGroupEntityRelationCtrl) GetAll(ctx *gin.Context) {

	customerGroupEntityRelationList, err := c.svc.GetAll()
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"customerGroupEntityRelationList": customerGroupEntityRelationList})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}

// CustomerGroupEntityRelation godoc
// @Summary delete particular customer
// @Description Here we can remove particular customer
// @Tags CustomerGroupEntityRelation
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path string true "Customer GroupEntityRelation ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerGroupEntityRelation/{id} [delete]
func (ug *customerGroupEntityRelationCtrl) Delete(ctx *gin.Context) {
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
