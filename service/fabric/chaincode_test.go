package fabric

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"io/ioutil"
	"testing"
)

type TestConfig struct {
}

func (t *TestConfig) ReadFile() ([]byte, error) {
	return ioutil.ReadFile("../../build/package/config/loyalty.yaml")
}

func TestCreateExternalConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "External Config Packer Suite")
}

var _ = Describe("External Config", func() {
	It("should be able to create a .TGZ file with a given name", func() {
		name := "test_cfg1"
		reader := new(TestConfig)
		conf := config.NewConfig(reader)
		externalConfig := NewExternalCcConfig(conf)

		path, err := externalConfig.Create(name)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(path).Should(BeEquivalentTo("/tmp/" + name + ".tgz"))
	})
})
