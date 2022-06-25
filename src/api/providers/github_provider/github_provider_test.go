package github_provider

import (
	"errors"
	"github.com/rmortale/golang-microservices/src/api/clients/restclient"
	"github.com/rmortale/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}
func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc")
	assert.EqualValues(t, "token abc", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.FlushMockups()
	errorMessage := "invalid restclient response"
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New(errorMessage),
	})

	response, err := CreateRepo("invalid", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, errorMessage, err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockups()

	invalidCloser, _ := os.Open("-asf3")
	errorMessage := "invalid response body"
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})

	response, err := CreateRepo("invalid", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, errorMessage, err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()

	errorMessage := "invalid json response body"
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := CreateRepo("invalid", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, errorMessage, err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockups()

	errorMessage := "Requires authentication"
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})

	response, err := CreateRepo("invalid", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, errorMessage, err.Message)
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockups()

	errorMessage := "error unmarshalling github create repo response"
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})

	response, err := CreateRepo("invalid", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, errorMessage, err.Message)
}

func TestCreateRepoSuccess(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1}`)),
		},
	})

	response, err := CreateRepo("valid", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 1, response.Id)
}
