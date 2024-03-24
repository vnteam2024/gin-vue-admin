package middleware

import (
	"bytes"
	"io"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	utils2 "github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func ErrorToEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string
		claims, _ := utils2.GetClaims(c)
		if claims.Username != "" {
			username = claims.Username
		} else {
			id, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			user, err := userService.FindUserById(id)
			if err != nil {
				username = "Unknown"
			}
			username = user.Username
		}
		body, _ := io.ReadAll(c.Request.Body)
//Rewrite it back to the request body, ioutil.ReadAll will clear the data in c.Request.Body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
		}
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		status := c.Writer.Status()
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
str := "The request received is" + record.Body + "\n" + "The request method is" + record.Method + "\n" + "The error message is as follows" + record.ErrorMessage + "\n" + "Time consuming" + latency.String() + "\n"
		if status != 200 {
subject := username + "" + record.Ip + "Called" + record.Path + "Error reported"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.GVA_LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
