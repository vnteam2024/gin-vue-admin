package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	email_response "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailApi struct{}

// EmailTest
// @Tags      System
// @Summary Send test email
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Sent successfully"}"
// @Router    /email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	err := service.ServiceGroupApp.EmailTest()
	if err != nil {
global.GVA_LOG.Error("Sending failed!", zap.Error(err))
response.FailWithMessage("Sending failed", c)
		return
	}
response.OkWithMessage("Sent successfully", c)
}

// SendEmail
// @Tags      System
// @Summary Send email
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param data body email_response.Email true "Required parameters for sending emails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Sent successfully"}"
// @Router    /email/sendEmail [post]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body)
	if err != nil {
global.GVA_LOG.Error("Sending failed!", zap.Error(err))
response.FailWithMessage("Sending failed", c)
		return
	}
response.OkWithMessage("Sent successfully", c)
}
