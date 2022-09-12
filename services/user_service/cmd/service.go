package cmd

import (
	"user_service/common"

	goservice "github.com/phathdt/go-sdk"
	"github.com/phathdt/go-sdk/plugin/storage/sdkgorm"
)

var (
	serviceName = "app-service"
	version     = "1.0.0"
)

func newService() goservice.Service {
	s := goservice.New(
		goservice.WithName(serviceName),
		goservice.WithVersion(version),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.KeyMainDB)),
	)

	return s
}
