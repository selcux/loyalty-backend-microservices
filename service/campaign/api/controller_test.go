package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var campaignStruct *model.Campaign
var campaignJson string

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Campaign Controller Suite")
}

var _ = BeforeSuite(func() {
	campaignStruct = &model.Campaign{
		Name:       "Campaign_1",
		PublicKey:  "D87598G7AF8G",
		PrivateKey: "D8S7GA98S7AG7",
		Timestamp:  1609011433403,
		Components: []model.Container{{
			ComponentType: "point",
			Order:         1,
			Component: model.ComponentComponent{
				Items: &map[string]int{
					"milk biscuit": 100,
					"chips":        200,
					"soda":         50,
					"cola":         75,
				},
			},
		}},
	}

	campaignData, err := json.Marshal(campaignStruct)
	Expect(err).ShouldNot(HaveOccurred())
	campaignJson = string(campaignData)
})

var _ = Describe("Campaign", func() {
	Describe("POST /", func() {
		It("should create a new campaign", func() {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(campaignJson))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			h := handler{campaignStruct}

			Expect(h.createCampaign(c)).ShouldNot(HaveOccurred())
			Expect(rec.Code).Should(Equal(http.StatusCreated))
			Expect(rec.Body.String()).Should(MatchJSON(campaignJson))
		})
	})
})
