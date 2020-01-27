package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tsauzeau/lbc/cmd/lbc/db"
	"github.com/tsauzeau/lbc/cmd/lbc/forms"
	"github.com/tsauzeau/lbc/pkg"
)

// FizzbuzzController ...
type FizzbuzzController struct{}

// Stat return the fizzbuzz stat
func (ctrl FizzbuzzController) Stat(c *gin.Context) {
	client := db.GetClient()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Redis not available"})
		c.Abort()
		return
	}
	zRevRange, err := client.ZRevRangeWithScores("set", 0, 0).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Redis error", "error": err.Error()})
		c.Abort()
		return
	} else if len(zRevRange) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Empty Stat"})
		c.Abort()
		return
	} else {
		first := zRevRange[0]
		c.JSON(http.StatusOK, first)
	}
}

// Get fizzbuz value (from FizzbuzzForm)
func (ctrl FizzbuzzController) Get(c *gin.Context) {
	var fizzbuzzForm forms.FizzbuzzForm

	if err := c.ShouldBind(&fizzbuzzForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed request", "error": err.Error()})
		c.Abort()
		return
	}

	res, err := pkg.Fizzbuzz(&fizzbuzzForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed request", "error": err.Error()})
		c.Abort()
		return
	}

	req, _ := json.Marshal(fizzbuzzForm)
	client := db.GetClient()
	if client != nil {
		_, err = client.ZIncrBy("set", 1, string(req)).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Redis error", "error": err.Error()})
			c.Abort()
			return
		}
	}

	c.String(http.StatusOK, strings.Join(res[:], ","))
}
