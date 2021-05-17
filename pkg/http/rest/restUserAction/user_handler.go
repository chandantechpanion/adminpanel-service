package restUserAction

import (
	"net/http"

	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"github.com/chandantechpanion/adminpanel-service/pkg/user"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type userCtrl struct {
	log logger.LogInfoFormat
	svc user.Service
}

func NewUserCtrl(log logger.LogInfoFormat, svc user.Service) *userCtrl {
	return &userCtrl{log, svc}
}

func (u *userCtrl) GetAll(ctx *gin.Context) {
	users, err := u.svc.GetAll()
	if len(users) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// User godoc
// @Summary Get Particular User
// @Description Here we will get particular user details
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} user.User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /users/{id} [get]
func (u *userCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := u.svc.GetByID(id)
	if user == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// User godoc
// @Summary Create User
// @Description Please create the user for providing permission
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body user.CreateUser true "Create User Details"
// @Success 201
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /users/ [post]
func (u *userCtrl) Store(ctx *gin.Context) {
	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := u.svc.Store(&user)
	if err != nil {
		ctx.Status(http.StatusCreated)
	} else {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}

func (u *userCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.UserId = id
	u.svc.Update(&user)
	ctx.Status(http.StatusOK)
}

func (u *userCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	u.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}
