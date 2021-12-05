package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_getServerHealth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(getServerHealth))
	defer ts.Close()

	t.Run("hits endpoint correctly", func(t *testing.T) {
		res, err := http.Get(ts.URL)
		if err != nil {
			t.Errorf("request failed: %v", err)
		}

		resBytes, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("failed to read body: %v", err)
		}

		respJSON := &struct {
			Healthy bool `json:"healthy"`
		}{}
		err = json.Unmarshal(resBytes, respJSON)

		if err != nil {
			t.Errorf("failed to unmarshal json: %v", err)
		}
		equal := reflect.DeepEqual(*respJSON, struct {
			Healthy bool `json:"healthy"`
		}{Healthy: true})

		if !equal {
			t.Errorf("responded incorrectly, equality resulted in: %t, respJSON.Healthy: %t", equal, respJSON.Healthy)
		}
	})
}
