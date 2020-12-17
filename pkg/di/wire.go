//+build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util/packer"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/fabric"
)

func InitializeConfig() *config.Config {
	wire.Build(config.NewConfig, config.NewYamlConfig)
	return &config.Config{}
}

func InitializeCcPack() packer.FilePacker {
	wire.Build(packer.NewCcPack, packer.NewArch)
	return &packer.CcPack{}
}

func InitializeExternalCcConfig() *fabric.ExternalCcConfig {
	wire.Build(fabric.NewExternalCcConfig, InitializeConfig)
	return &fabric.ExternalCcConfig{}
}
