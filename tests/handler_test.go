package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google-service/models"
	"google-service/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string, body *string) *httptest.ResponseRecorder {
	var req *http.Request

	if body != nil {
		req, _ = http.NewRequest(method, path, bytes.NewBufferString(*body))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestJsonDecodeToTimesEquals(t *testing.T) {
	r := router.SetupRouter()

	exampleTimes := `[
		{
			"start": "2019-09-09T14:00:00.000+00:00",
			"end": "2019-09-09T14:00:00.000+00:00"
		}	
	]`

	w := performRequest(r, "POST", "/api/v1/ldc", &exampleTimes)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	val, exists := response["code"]
	mes, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, float64(http.StatusBadRequest), val)
	assert.Equal(t, models.ErrTimeEquals.Error(), mes)
}

func TestJsonDecodeToTimesWhenEndBeforeStart(t *testing.T) {
	r := router.SetupRouter()

	exampleTimes := `[
		{
			"start": "2019-09-09T16:00:00.000+00:00",
			"end": "2019-09-09T14:00:00.000+00:00"
		}	
	]`

	w := performRequest(r, "POST", "/api/v1/ldc", &exampleTimes)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	val, exists := response["code"]
	mes, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, float64(http.StatusBadRequest), val)
	assert.Equal(t, models.ErrTimeIncorrect.Error(), mes)
}

func TestJsonDecodeToTimesWhenDurationOneHour(t *testing.T) {
	r := router.SetupRouter()

	exampleTimes := `[
		{
			"start": "2019-09-09T12:00:00.000+00:00",
			"end": "2019-09-09T14:00:00.000+00:00"
		},
		{
			"start": "2019-09-09T13:00:00.000+00:00",
			"end": "2019-09-09T14:00:00.000+00:00"
		}	
	]`

	w := performRequest(r, "POST", "/api/v1/ldc", &exampleTimes)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	val, exists := response["code"]
	mes, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, float64(http.StatusBadRequest), val)
	assert.Equal(t, models.ErrTimeDuration.Error(), mes)
}

func TestJsonDecodeToTimesWithErrorPayload(t *testing.T) {
	r := router.SetupRouter()

	exampleTimes := `[
		{
			"start": "2019-09-09T14:00:00",
			"end": "2019-09-09T14:00:00"
		}	
	]`

	w := performRequest(r, "POST", "/api/v1/ldc", &exampleTimes)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	val, exists := response["code"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, float64(http.StatusBadRequest), val)
}

func TestGetEtag(t *testing.T) {
	r := router.SetupRouter()

	w := performRequest(r, "POST", "/api/v1/ldc/clinic/123", nil)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	val, exists := response["code"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, float64(http.StatusOK), val)
}

func TestGetEtagWithErrorParam(t *testing.T) {
	r := router.SetupRouter()

	w := performRequest(r, "POST", "/api/v1/ldc/clinic/", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
