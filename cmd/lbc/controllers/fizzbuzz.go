package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsauzeau/lbc/cmd/lbc/forms"
)

//FizzbuzzController ...
type FizzbuzzController struct{}

//Fizzbuzz ...
func Fizzbuzz(fizzbuzz *forms.FizzbuzzForm) (res []string, err error) {
	if fizzbuzz.Limit <= 0 || fizzbuzz.Int1 <= 0 || fizzbuzz.Int2 <= 0 {
		return nil, errors.New("Limit, Int1 and Int2 needs to be positive values")
	}
	for i := 1; i <= fizzbuzz.Limit+1; i++ {
		if i%(fizzbuzz.Int1*fizzbuzz.Int2) == 0 {
			res = append(res, fmt.Sprintf("%s%s", fizzbuzz.String1, fizzbuzz.String2))
		} else if i%fizzbuzz.Int1 == 0 {
			res = append(res, fmt.Sprintf(fizzbuzz.String1))
		} else if i%fizzbuzz.Int2 == 0 {
			res = append(res, fmt.Sprintf(fizzbuzz.String2))
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	res = append(res, "...")
	return res, nil
}

//Get ...
func (ctrl FizzbuzzController) Get(c *gin.Context) {
	var fizzbuzzForm forms.FizzbuzzForm

	if err := c.ShouldBind(&fizzbuzzForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed request", "error": err.Error()})
		c.Abort()
		return
	}

	res, err := Fizzbuzz(&fizzbuzzForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed request", "error": err.Error()})
		c.Abort()
		return
	}

	c.String(http.StatusOK, strings.Join(res[:], ","))
}
