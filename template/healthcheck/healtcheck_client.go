package healthcheck

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Healthchecker interface {
	CheckHealth() (*HealthcheckResponse, error)
}

type HealthcheckApiClient struct {
	baseUrl string
}

type HealthcheckResponse struct {
	Status string `json:"status"`
}

func (api *HealthcheckApiClient) CheckHealth() (*HealthcheckResponse, error) {
	httpResp, err := http.Get(api.baseUrl + "/api/healthcheck")
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(httpResp.Body)

	healthcheck := HealthcheckResponse{}
	if err := json.Unmarshal(respBody, &healthcheck); err != nil {
		return nil, err
	}
	return &healthcheck, nil
}
