package fabric

import (
	"bytes"
	"encoding/json"
	"github.com/gosimple/slug"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util/packer"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/fabric/model"
	"path"
	"text/template"
)

type ExternalCcConfig struct {
	config.BaseConfig
}

func NewExternalCcConfig(conf *config.Config) *ExternalCcConfig {
	return &ExternalCcConfig{config.BaseConfig{Config: conf}}
}

func (c *ExternalCcConfig) Create(name string) (string, error) {
	ccConfig := c.Config.CC

	tmpl, err := template.New("address").Parse(ccConfig.ServerAddressTemplate)
	if err != nil {
		return "", err
	}

	slugName := slug.Make(name)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, slugName)
	if err != nil {
		return "", err
	}

	connection := model.NewConnection()
	connection.SetAddress(buf.String(), ccConfig.Port)
	connectionBytes, err := json.Marshal(connection)
	if err != nil {
		return "", err
	}

	arch := packer.NewArch()
	arch.Append(packer.ArchiveArgs{
		Name: "connection.json",
		Data: connectionBytes,
	})

	codeTar, err := arch.Archive()
	if err != nil {
		return "", err
	}

	codeTarGz, err := packer.Comporess(codeTar)
	if err != nil {
		return "", err
	}

	metadata := model.NewMetadata()
	metadata.Label = slugName
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return "", err
	}

	arch = packer.NewArch()
	arch.Append(packer.ArchiveArgs{
		Name: "metadata.json",
		Data: metadataBytes,
	})
	arch.Append(packer.ArchiveArgs{
		Name: "code.tar.gz",
		Data: codeTarGz,
	})

	outFile := path.Join(ccConfig.PackageLocation, slugName+".tgz")

	ccPack := packer.NewCcPack(arch)
	err = ccPack.Write(outFile)

	return outFile, err
}
