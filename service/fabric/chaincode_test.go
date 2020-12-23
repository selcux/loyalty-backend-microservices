package fabric

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

type TestConfig struct {
}

func (t *TestConfig) ReadFile() ([]byte, error) {
	return ioutil.ReadFile("../../build/package/config/loyalty.yaml")
}

func TestInstallExternalChaincode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "External Chaincode Suite")
}

var _ = Describe("External Chaincode", func() {
	var externalConfig *ExternalCcConfig
	name := "test_cfg1"

	BeforeEach(func() {
		reader := new(TestConfig)
		conf := config.NewConfig(reader)
		externalConfig = NewExternalCcConfig(conf)
	})

	Describe("External Config", func() {
		It("should be able to create a .TGZ file with a given name", func() {
			path, err := externalConfig.Create(name)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(path).Should(BeEquivalentTo("/tmp/" + name + ".tgz"))
		})
	})

	Describe("External Chaincode Package", func() {
		It("should be able to get all related peers", func() {
			orgs, err := externalConfig.getOrganizations()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(orgs).To(BeNumerically(">", 0))
			Expect(orgs).To(Equal(2))
		})

		It("should be able to install cc pack", func() {
			path, err := externalConfig.Create(name)
			Expect(err).ShouldNot(HaveOccurred())

			orgs, err := externalConfig.getOrganizations()
			Expect(err).ShouldNot(HaveOccurred())

			externalConfig.cmd = "echo '%s - %d - %s'"

			outputs, err := externalConfig.installCCPack(orgs, path)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(outputs).ShouldNot(BeEmpty())

			log.Println(strings.Join(outputs, "\n"))
		})
	})
})
