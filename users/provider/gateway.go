package provider

import (
	"users-microservice/gateway"

	"github.com/samber/do"
)

func ProvideGateway(i *do.Injector) {
	do.Provide(i, gateway.NewEmailGw)
}
