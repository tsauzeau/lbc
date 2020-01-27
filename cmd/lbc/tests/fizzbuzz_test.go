package tests

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tsauzeau/lbc/cmd/lbc/controllers"
	"github.com/tsauzeau/lbc/cmd/lbc/forms"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	v1 := r.Group("/v1")
	{
		/*** START Fizzbuzz ***/
		fizzbuzz := new(controllers.FizzbuzzController)

		v1.GET("/fizzbuzz", fizzbuzz.Get)
	}

	return r
}

func main() {
	r := SetupRouter()
	r.Run()
}

var signinCookie string

var testEmail = "test-gin-boilerplate@test.com"
var testPassword = "123456"

var articleID int

/**
* TestFizzbuzzOK
* Test fizzbuzz in normal condition
*
* Must return response code 200
 */
func TestFizzbuzzOK(t *testing.T) {
	testRouter := SetupRouter()

	var fizzbuzForm forms.FizzbuzzForm

	fizzbuzForm.Int1 = 3
	fizzbuzForm.Int2 = 5
	fizzbuzForm.String1 = "fizz"
	fizzbuzForm.String2 = "buzz"
	fizzbuzForm.Limit = 15

	v, _ := query.Values(fizzbuzForm)
	req, err := http.NewRequest("GET", "/v1/fizzbuzz", nil)
	req.URL.RawQuery = v.Encode()

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.String(), "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...")
}

/**
* TestFizzbuzzMissing
* Test fizzbuzz with missing args
*
* Must return response code 400
 */
func TestFizzbuzzMissing(t *testing.T) {
	testRouter := SetupRouter()

	var fizzbuzForm forms.FizzbuzzForm

	fizzbuzForm.Int1 = 3
	fizzbuzForm.Int2 = 5
	fizzbuzForm.String1 = "fizz"
	fizzbuzForm.Limit = 15

	v, _ := query.Values(fizzbuzForm)
	req, err := http.NewRequest("GET", "/v1/fizzbuzz", nil)
	req.URL.RawQuery = v.Encode()

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 400)
}

/**
* TestFizzbuzzInvalid
* Test fizzbuzz with invalid args
*
* Must return response code 400
 */
func TestFizzbuzzInvalid(t *testing.T) {
	testRouter := SetupRouter()

	var fizzbuzForm forms.FizzbuzzForm

	fizzbuzForm.Int1 = -1
	fizzbuzForm.Int2 = 5
	fizzbuzForm.String1 = "fizz"
	fizzbuzForm.String2 = "fizz"
	fizzbuzForm.Limit = -1

	v, _ := query.Values(fizzbuzForm)
	req, err := http.NewRequest("GET", "/v1/fizzbuzz", nil)
	req.URL.RawQuery = v.Encode()

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 400)
}

/**
* TestFizzbuzzNotFound
* Test fizzbuzz with invalid args
*
* Must return response code 404
 */
func TestFizzbuzzNotFound(t *testing.T) {
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/not_found", nil)

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 404)
}
