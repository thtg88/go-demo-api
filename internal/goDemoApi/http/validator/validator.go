package validator

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidateJSON validates the given Gin context and rules,
// returning the errors
func ValidateJSON(c *gin.Context, rules govalidator.MapData) url.Values {
	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		RequiredDefault: true,      // all the field to be pass the rules
		Data:            &data,
	}

	v := govalidator.New(opts)
	errors := v.ValidateJSON()

	return errors
}
