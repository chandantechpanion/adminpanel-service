package restCustomerGroup

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroup"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type customerGroupCtrl struct {
	log logger.LogInfoFormat
	svc customerGroup.Service
}

func NewCustomerGroupCtrl(log logger.LogInfoFormat, svc customerGroup.Service) *customerGroupCtrl {
	return &customerGroupCtrl{log, svc}
}

// CustomerGroup godoc
// @Summary Save Customer Info
// @Description Here Save customer info.
// @Tags CustomerGroup
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param cgroup body customerGroup.CustomerGroupViewModel true "Customer Details"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 409
// @Failure 500
// @Router /customerGroup/ [post]
func (c *customerGroupCtrl) Save(ctx *gin.Context) {

	custGroup := customerGroup.CustomerGroupViewModel{}

	if err := ctx.ShouldBindJSON(&custGroup); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cGroup := customerGroup.CustomerGroup{}
	cGroup.Name = custGroup.Name
	cGroup.CreatedDate = time.Now()

	err := c.svc.Save(&cGroup)
	if err == nil {
		ctx.Status(http.StatusOK)
	} else if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// CustomerGroup
// @Summary Get Customer
// @Description Here Get customer by id
// @Tags CustomerGroup
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id header int64 true "id for Auth"
// @Param Page query int true "Page Number" mininum(1) maxinum(100000)
// @Param RecordsPerPage query int true "Records Per Page" mininum(1) maxinum(500)
// @Success 200  {array} customerGroup.CustomerGroup
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerGroup/ [get]
func (c *customerGroupCtrl) Get(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.GetHeader("userId"), 10, 64)
	Page, _ := strconv.Atoi(ctx.Query("Page"))
	RecordsPerPage, _ := strconv.Atoi(ctx.Query("RecordsPerPage"))

	if Page < 1 {
		Page = 1
	}
	if RecordsPerPage < 1 {
		RecordsPerPage = 5
	}

	if userId < 0 {
		ctx.Status(http.StatusBadRequest)
		return
	}
	// i, _ := strconv.ParseInt(id, 10, 64)
	// invoiceCaptureList, pagination, err := i.svc.Get(userId, Page, RecordsPerPage)
	customerGroupList, pagination, err := c.svc.Get(userId, Page, RecordsPerPage)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	} else if len(customerGroupList) == 0 {
		ctx.Status(http.StatusNoContent)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"customerGroupList": customerGroupList, "pagination": pagination})
	}

}

// CustomerGroup godoc
// @Summary delete particular customer
// @Description Here we can remve customer
// @Tags CustomerGroup
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer {token}"
// @Param id path string true "Customer Group ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /customerGroup/{id} [delete]
func (ug *customerGroupCtrl) Delete(ctx *gin.Context) {
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
