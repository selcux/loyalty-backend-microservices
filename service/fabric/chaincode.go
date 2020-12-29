package fabric

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util/packer"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/fabric/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"
)

type InstallCommand string
type InstallArgs string

type ExternalCcConfig struct {
	config.BaseConfig
	cmd  InstallCommand
	args InstallArgs
}

func NewExternalCcConfig(conf *config.Config) *ExternalCcConfig {
	return &ExternalCcConfig{
		BaseConfig: config.BaseConfig{Config: conf},
		cmd:        ". %s; setGlobals %d; peer lifecycle chaincode install %s",
	}
}

func (c *ExternalCcConfig) Create(name string) (string, error) {
	fabricNetwork := c.Config.FabricNetwork

	tmpl, err := template.New("address").Parse(fabricNetwork.ServerAddressTemplate)
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
	connection.SetAddress(buf.String(), fabricNetwork.Port)
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

	outFile := path.Join(fabricNetwork.PackageLocation, slugName+".tgz")

	ccPack := packer.NewCcPack(arch)
	err = ccPack.Write(outFile)

	return outFile, err
}

func (c *ExternalCcConfig) Install(ccPackPath string) error {
	orgs, err := c.getOrganizations()
	if err != nil {
		return err
	}

	outputs, err := c.installCCPack(orgs, ccPackPath)
	if err != nil {
		return err
	}

	log.Println(outputs)

	return nil
}

func (c *ExternalCcConfig) getOrganizations() (int, error) {
	fabricNetwork := c.Config.FabricNetwork
	networkConfigFile := os.Getenv(fabricNetwork.NetworkConfigFile)
	source, err := ioutil.ReadFile(networkConfigFile)
	if err != nil {
		return 0, err
	}

	var configTx model.ConfigTx
	err = yaml.Unmarshal(source, &configTx)
	if err != nil {
		return 0, err
	}

	orgCount := 0
	for _, org := range configTx.Organizations {
		if 0 < len(org.AnchorPeers) {
			orgCount++
		}
	}

	return orgCount, nil
}

func (c *ExternalCcConfig) installCCPack(orgCount int, ccPackPath string) ([]string, error) {
	if ccPackPath == "" {
		return nil, errors.New("invalid package file")
	}

	fabricNetwork := c.Config.FabricNetwork
	outputs := make([]string, 0)
	envVarSh := os.Getenv(fabricNetwork.EnvScript)
	log.Println("envVarSh", envVarSh)

	for i := 1; i <= orgCount; i++ {
		completeCmd := fmt.Sprintf(string(c.cmd), envVarSh, i, ccPackPath)
		log.Println(completeCmd)
		out, err := exec.Command("/bin/sh", "-c", completeCmd).Output()
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, string(out))
	}

	return outputs, nil
}
