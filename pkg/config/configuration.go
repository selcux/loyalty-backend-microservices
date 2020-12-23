package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

const configFilePath = "/etc/loyalty.yaml"

type PropReader interface {
	MongoProps() Mongo
}
type FileReader interface {
	ReadFile() ([]byte, error)
}

type Mongo struct {
	Database    `yaml:",inline"`
	Collections map[string]string `yaml:"collections"`
}

type Tables struct {
	Company string `yaml:"company,omitempty"`
	Product string `yaml:"product,omitempty"`
}

type DbProperties struct {
	Mongo Mongo `yaml:"mongo,omitempty"`
}

type Database struct {
	ConnectionString string `yaml:"connection_string"`
	DbName           string `yaml:"db_name"`
}

type Service struct {
	Host     string `yaml:"host"`
	ApiPort  int    `yaml:"api_port"`
	GrpcPort int    `yaml:"grpc_port"`
}

type FabNetwork struct {
	ServerAddressTemplate string `yaml:"server_address_template"`
	Port                  int    `yaml:"port"`
	PackageLocation       string `yaml:"package_location"`
	PeerAddressDefinition string `yaml:"peer_address_definition"`
	NetworkConfigFile     string `yaml:"network-config"`
	EnvScript             string `yaml:"env-script"`
}

type Config struct {
	DbProperties  DbProperties       `yaml:"db_properties"`
	Services      map[string]Service `yaml:"services"`
	FabricNetwork FabNetwork         `yaml:"fabric_network"`
}

func (c *Config) MongoProps() Mongo {
	return c.DbProperties.Mongo
}

type YamlConfig struct {
}

func (y YamlConfig) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(configFilePath)
}

func NewYamlConfig() FileReader {
	return new(YamlConfig)
}

func initConfig(reader FileReader) *Config {
	configData := new(Config)
	data, err := reader.ReadFile()
	if err != nil {
		log.Fatalf("Unable to read %s\n %v", configFilePath, err)
	}

	err = yaml.Unmarshal(data, configData)
	if err != nil {
		log.Fatalf("Unable to unmarshal config %v", err)
	}

	return configData
}

func NewConfig(reader FileReader) *Config {
	return initConfig(reader)
}

type BaseConfig struct {
	Config *Config
}
