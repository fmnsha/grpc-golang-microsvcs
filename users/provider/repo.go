package provider

import (
	"users-microservice/pkg/repo"

	"github.com/samber/do"
)

func ProvideRepo(i *do.Injector) {
	do.Provide(i, repo.NewUserRepo)
	do.Provide(i, repo.NewRoleRepo)
	do.Provide(i, repo.NewCapabilityRepo)
}
