package provider

import (
	"users-microservice/pkg/service"

	"github.com/samber/do"
)

func ProvideService(i *do.Injector) {
	do.Provide(i, service.NewAuthSvcs)
	do.Provide(i, service.NewCapabilitySvcs)
	do.Provide(i, service.NewJwtManager)
	do.Provide(i, service.NewRoleSvcs)
	do.Provide(i, service.NewUserSvcs)

}
