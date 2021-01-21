package config

import (
	"github.com/bxcodec/faker/v3"
	"github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"testing"
)

type TestConfig struct {
}

func (t *TestConfig) ReadFile() ([]byte, error) {
	return ioutil.ReadFile("../../build/package/config/loyalty.yaml")
}

func TestConfigStructure(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	config := new(Config)
	err := faker.FakeData(config)

	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	configData, err := yaml.Marshal(config)

	g.Expect(err).ShouldNot(gomega.HaveOccurred())

	log.Printf("\n%+v", string(configData))
}

func TestConfiguration_Load(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	reader := new(TestConfig)
	config := NewConfig(reader)

	g.Expect(config).ShouldNot(gomega.BeNil())
}

func TestConfig_MongoProps(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	reader := new(TestConfig)
	config := NewConfig(reader)

	g.Expect(config).ShouldNot(gomega.BeNil())

	mongoProps := config.MongoProps()

	g.Expect(mongoProps).ShouldNot(gomega.BeNil())
	g.Expect(mongoProps.DbName).Should(gomega.BeEquivalentTo("loyalty-dlt"))
	g.Expect(mongoProps.ConnectionString).Should(gomega.BeEquivalentTo("mongodb://mongo-local-loyalty:27017"))
	g.Expect(mongoProps.Collections).ShouldNot(gomega.BeEmpty())
	g.Expect(mongoProps.Collections).Should(gomega.HaveLen(5))
	g.Expect(mongoProps.Collections).Should(gomega.BeEquivalentTo(map[string]string{
		"product":  "products",
		"consumer": "consumers",
		"item":     "items",
		"company":  "companies",
		"merchant": "merchants",
	}))
}
