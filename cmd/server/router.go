package server

import (
	"github.com/chandantechpanion/adminpanel-service/pkg/buyerUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerEntity"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroup"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroupEntityRelation"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerMaster"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerMember"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest"
	restbuyeruserinfo "github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restBuyerUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restCustomerEntity"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restCustomerGroup"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restCustomerGroupEntityRelation"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restCustomerMaster"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restCustomerMember"
	"github.com/chandantechpanion/adminpanel-service/pkg/http/rest/restUserAction"
	"github.com/chandantechpanion/adminpanel-service/pkg/login"
	"github.com/chandantechpanion/adminpanel-service/pkg/middleware"
	"github.com/chandantechpanion/adminpanel-service/pkg/user"

	"fmt"

	"github.com/gin-gonic/gin"
)

func (ds *dserver) MapRoutes() {

	// Group : v1
	apiV1 := ds.router.Group("api/v1")
	ds.loginRoutes(apiV1)
	ds.userRoutes(apiV1)
	ds.buyerUserRoutes(apiV1)
	ds.customerGroupRoutes(apiV1)
	ds.customerMemberRoutes(apiV1)
	ds.customerGroupEntityRelationRoutes(apiV1)
	ds.customerMasterRoutes(apiV1)
	ds.customerEntityRoutes(apiV1)
}

func (ds *dserver) customerGroupRoutes(api *gin.RouterGroup) {
	var Svc customerGroup.Service
	err := ds.cont.Invoke(func(l customerGroup.Service) {
		Svc = l
	})
	if err != nil {
		fmt.Println("Error In Creating DI function for customerGroup")
		fmt.Println(err)
	} else {
		customerGroupRoutes := api.Group("/customerGroup")
		{
			customerGroupRoutes.Use(middleware.AuthorizeJWT())
			f := restCustomerGroup.NewCustomerGroupCtrl(ds.logger, Svc)
			customerGroupRoutes.POST("/", f.Save)
			customerGroupRoutes.GET("/", f.Get)
			customerGroupRoutes.DELETE("/:id", f.Delete)
		}
	}
}

func (ds *dserver) customerMemberRoutes(api *gin.RouterGroup) {
	var Svc customerMember.Service
	err := ds.cont.Invoke(func(l customerMember.Service) {
		Svc = l
	})
	if err != nil {
		fmt.Println("Error In Creating DI function for customerMember")
		fmt.Println(err)
	} else {
		customerMemberRoutes := api.Group("/customerMember")
		{
			customerMemberRoutes.Use(middleware.AuthorizeJWT())
			f := restCustomerMember.NewCustomerMemberCtrl(ds.logger, Svc)
			customerMemberRoutes.POST("/", f.Save)
			customerMemberRoutes.GET("/:id", f.Get)
			customerMemberRoutes.DELETE("/:id", f.Delete)
		}
	}
}

func (ds *dserver) customerGroupEntityRelationRoutes(api *gin.RouterGroup) {
	var Svc customerGroupEntityRelation.Service
	err := ds.cont.Invoke(func(l customerGroupEntityRelation.Service) {
		Svc = l
	})
	if err != nil {
		fmt.Println("Error In Creating DI function for customerGroupEntityRelation")
		fmt.Println(err)
	} else {
		customerGroupEntityRelationRoutes := api.Group("/customerGroupEntityRelation")
		{
			customerGroupEntityRelationRoutes.Use(middleware.AuthorizeJWT())
			f := restCustomerGroupEntityRelation.NewCustomerGroupEntityRelationCtrl(ds.logger, Svc)
			customerGroupEntityRelationRoutes.POST("/", f.Save)
			customerGroupEntityRelationRoutes.GET("/:id", f.Get)
			customerGroupEntityRelationRoutes.GET("/", f.GetAll)
			customerGroupEntityRelationRoutes.DELETE("/:id", f.Delete)
		}
	}
}

func (ds *dserver) customerMasterRoutes(api *gin.RouterGroup) {
	var Svc customerMaster.Service
	err := ds.cont.Invoke(func(l customerMaster.Service) {
		Svc = l
	})
	if err != nil {
		fmt.Println("Error In Creating DI function for customerMaster")
		fmt.Println(err)
	} else {
		customerMasterRoutes := api.Group("/customerMaster")
		{
			customerMasterRoutes.Use(middleware.AuthorizeJWT())
			f := restCustomerMaster.NewCustomerMasterCtrl(ds.logger, Svc)
			customerMasterRoutes.POST("/", f.Save)
			customerMasterRoutes.GET("/", f.Get)
			customerMasterRoutes.DELETE("/:id", f.Delete)
		}
	}
}

func (ds *dserver) customerEntityRoutes(api *gin.RouterGroup) {
	var Svc customerEntity.Service
	err := ds.cont.Invoke(func(l customerEntity.Service) {
		Svc = l
	})
	if err != nil {
		fmt.Println("Error In Creating DI function for customerEntity")
		fmt.Println(err)
	} else {
		customerEntityRoutes := api.Group("/customerEntity")
		{
			customerEntityRoutes.Use(middleware.AuthorizeJWT())
			f := restCustomerEntity.NewCustomerEntityCtrl(ds.logger, Svc)
			customerEntityRoutes.POST("/", f.Save)
			// customerEntityRoutes.GET("/", f.GetAll)
			customerEntityRoutes.GET("/:customer_master_id", f.Get)
		}
	}
}

func (ds *dserver) loginRoutes(api *gin.RouterGroup) {
	var loginSvc login.Service
	err := ds.cont.Invoke(func(l login.Service) {
		loginSvc = l
	})
	if err != nil {
		fmt.Println("Error In Creating DI function for Login")
		fmt.Println(err)
	} else {
		loginRoutes := api.Group("/access_token")
		{
			f := rest.NewLoginCtrl(ds.logger, loginSvc)
			loginRoutes.POST("/", f.AccessToken)
		}
	}
}

func (ds *dserver) userRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	{
		var userSvc user.Service
		err := ds.cont.Invoke(func(u user.Service) {
			userSvc = u
		})

		if err != nil {
			fmt.Println("Error In Creating DI function for User")
			fmt.Println(err)
		} else {
			usr := restUserAction.NewUserCtrl(ds.logger, userSvc)

			//userRoutes.GET("/", usr.GetAll)
			userRoutes.POST("/", usr.Store)
			userRoutes.GET("/:id", usr.GetByID)
			userRoutes.PUT("/:id", usr.Update)
			userRoutes.DELETE("/:id", usr.Delete)
		}
	}
}

func (ds *dserver) buyerUserRoutes(api *gin.RouterGroup) {
	buyerUserRoutes := api.Group("/buyerUsers")
	{
		var userSvc buyerUserInfo.Service
		err := ds.cont.Invoke(func(u buyerUserInfo.Service) {
			userSvc = u
		})

		if err != nil {
			fmt.Println("Error In Creating DI function for User")
			fmt.Println(err)
		} else {
			usr := restbuyeruserinfo.NewBuyerUserInfoCtrl(ds.logger, userSvc)

			buyerUserRoutes.GET("/:id", usr.GetBuyerUserInfo)
			buyerUserRoutes.POST("/", usr.Store)
		}
	}
}
