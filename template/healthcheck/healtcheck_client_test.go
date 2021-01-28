package healthcheck

import (
	"github.com/tink-ab/digimon-recrutiment-template/template"
	"strings"
	"testing"
)

func Test_checkHealth(t *testing.T) {
	client := HealthcheckApiClient{
		baseUrl: template.ApiBaseUlr,
	}

	resp, err := client.CheckHealth()
	if err != nil {
		t.Fatal(err)
	}
	if status := resp.Status; !strings.EqualFold(status, "OK") {
		t.Errorf("wrong status: %s", status)
	}

}
