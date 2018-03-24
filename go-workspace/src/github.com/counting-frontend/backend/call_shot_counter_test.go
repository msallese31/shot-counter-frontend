package backend

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/counting-frontend/data"
	"github.com/stretchr/testify/assert"
)

func TestCallShotCounter(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/doesntmatter", nil)
	if err != nil {
		t.Fatal(err)
	}

	countData := data.CountObject{}
	countData.Request = *req

	rr := httptest.NewRecorder()
	countData.Writer = rr

	CallShotCounter(countData)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Unexpected error when testing simple readiness endpoint")
}

func TestCallShotCounterBadBody(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	req, err := http.NewRequest("POST", "/doesntmatter", bytes.NewBuffer(make([]byte, 5)))
	if err != nil {
		t.Fatal(err)
	}

	countData := data.CountObject{}
	countData.Request = *req

	rr := httptest.NewRecorder()
	countData.Writer = rr

	CallShotCounter(countData)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Unexpected error when testing simple readiness endpoint")
}

func TestCallShotCounterEmptyBody(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	req, err := http.NewRequest("POST", "/doesntmatter", bytes.NewBuffer(make([]byte, 0)))
	if err != nil {
		t.Fatal(err)
	}

	countData := data.CountObject{}
	countData.Request = *req

	rr := httptest.NewRecorder()
	countData.Writer = rr

	CallShotCounter(countData)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Unexpected error when testing simple readiness endpoint")
}
