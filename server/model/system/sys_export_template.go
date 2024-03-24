// Automatically generate template SysExportTemplate
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Export template structure SysExportTemplate
type SysExportTemplate struct {
	global.GVA_MODEL
	DBName string `json:"dbName" form:"dbName" gorm:"column:db_name;comment:database name;"` //Database name
	Name string `json:"name" form:"name" gorm:"column:name;comment:template name;"` //Template name
	TableName string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:table name;"` //Table name
	TemplateID string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:template identification;"` //Template identification
	TemplateInfo string `json:"templateInfo" form:"templateInfo" gorm:"column:template_info;type:text;"` //Template information
	Limit int `json:"limit" form:"limit" gorm:"column:limit;comment:Export limit"`
	Order string `json:"order" form:"order" gorm:"column:order;comment:sort"`
	Conditions []Condition `json:"conditions" form:"conditions" gorm:"foreignKey:TemplateID;references:TemplateID;comment:condition"`
}

type Condition struct {
	global.GVA_MODEL
	TemplateID string `json:"templateID" form:"templateID" gorm:"column:template_id;comment:template ID"`
	From string `json:"from" form:"from" gorm:"column:from;comment:key taken from the condition"`
	Column string `json:"column" form:"column" gorm:"column:column;comment:field as query condition"`
	Operator string `json:"operator" form:"operator" gorm:"column:operator;comment:operator"`
}

func (Condition) TableName() string {
	return "sys_export_template_condition"
}
