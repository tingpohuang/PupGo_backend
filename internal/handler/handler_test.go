package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestRootHandler(t *testing.T) {
	w := httptest.NewRecorder() // initialize httptest Recorder
	c, _ := gin.CreateTestContext(w) // initialize context
	rootHandler := RootHandler()
	rootHandler(c)
	assert.Equal(t, http.StatusOK, w.Code) // check whether the http response is OK

	var payload gin.H
	err := json.Unmarshal(w.Body.Bytes(), &payload) // Unmarshal context body
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hello world", payload["data"])  // check whether the payload is correct

}
