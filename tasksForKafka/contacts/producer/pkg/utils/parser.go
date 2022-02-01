package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context, x interface{}) {
	if body, err := ioutil.ReadAll(c.Request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetContactID(c *gin.Context) string {
	var id string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "id" {
			id = queryValue
		}
	}
	return id
}
