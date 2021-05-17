package di

import (
	"github.com/chandantechpanion/adminpanel-service/pkg/buyerUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/config"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerEntity"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroup"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroupEntityRelation"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerMaster"
	"github.com/chandantechpanion/adminpanel-service/pkg/customerMember"
	"github.com/chandantechpanion/adminpanel-service/pkg/login"
	"github.com/chandantechpanion/adminpanel-service/pkg/user"

	"github.com/chandantechpanion/adminpanel-service/pkg/jwtService"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage"

	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormBuyerUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormCustomerEntity"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormCustomerGroup"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormCustomerGroupEntityRelation"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormCustomerMaster"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormCustomerMember"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormUser"
	"github.com/chandantechpanion/adminpanel-service/pkg/storage/orm/ormlogin"

	"go.uber.org/dig"
)

var container = dig.New()

//BuildContainer Building all Container here
func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// DB
	container.Provide(storage.NewDb)

	// logger
	container.Provide(logger.NewLogger)

	// JWT Service
	container.Provide(jwtService.NewJWTAuthService)

	//Login Service
	container.Provide(login.NewLoginService)
	container.Provide(ormlogin.NewloginRepo)

	//User Service
	container.Provide(user.NewUserService)
	container.Provide(ormUser.NewUserRepo)

	//Buyer User Service
	container.Provide(buyerUserInfo.NewBuyerUserInfoService)
	container.Provide(ormBuyerUserInfo.NewBuyerUserInfoRepo)

	//Customer Group
	container.Provide(customerGroup.NewCustomerGroupService)
	container.Provide(ormCustomerGroup.NewCustomerGroupRepo)

	//Customer Member
	container.Provide(customerMember.NewCustomerMemberService)
	container.Provide(ormCustomerMember.NewCustomerMemberRepo)

	//CustomerGroupEntityRelation
	container.Provide(customerGroupEntityRelation.NewCustomerGroupEntityRelationService)
	container.Provide(ormCustomerGroupEntityRelation.NewCustomerGroupEntityRelationRepo)

	//Customer Master
	container.Provide(customerMaster.NewCustomerMasterService)
	container.Provide(ormCustomerMaster.NewCustomerMasterRepo)

	//Customer Entity
	container.Provide(customerEntity.NewCustomerEntityService)
	container.Provide(ormCustomerEntity.NewCustomerEntityRepo)

	return container
}

//Invoke Where we can invoke all the DI
func Invoke(i interface{}) error {
	return container.Invoke(i)
}
