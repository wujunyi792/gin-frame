package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/ginFrame/logger"
	"github.com/wujunyi792/ginFrame/model/v1/common"
	"github.com/wujunyi792/ginFrame/utils"
	"net/http"
)

func JwtVerify(c *gin.Context) {
	var res common.CommonResponseApi
	if gin.Mode() == gin.DebugMode {
		c.Next()
		return
	}
	cookie, err := c.Cookie("token")
	if err == nil {
		_, err = utils.ParseToken(cookie)
		if err == nil {
			c.Next()
			return
		} else {
			res.Code = 14005
			res.Message = fmt.Sprintf("%v", err)
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}
	}
	logger.Warning.Println(err)
	res.Code = 14005
	res.Message = "cookie not set"
	c.JSON(http.StatusForbidden, res)
	c.Abort()
	return
}
