package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

type Config struct {
	DbProperties DbProperties       `yaml:"db_properties"`
	Services     map[string]Service `yaml:"services"`
}

func (c *Config) MongoProps() Mongo {
	return c.DbProperties.Mongo
}

type YamlConfig struct {
}

func (y *YamlConfig) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(configFilePath)
}

func initializeConfig(reader FileReader) (*Config, error) {
	configData := new(Config)
	data, err := reader.ReadFile()
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, configData)

	return configData, err
}

func NewConfig() (*Config, error) {
	reader := new(YamlConfig)
	return initializeConfig(reader)
}
