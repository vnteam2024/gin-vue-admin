package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

//@author: [granty1](https://github.com/granty1)
//@function: CreateSysOperationRecord
//@description: Create record
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysOperationRecordByIds
//@description: Delete records in batches
//@param: ids request.IdsReq
//@return: err error

func (operationRecordService *OperationRecordService) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: Delete operation record
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error

func (operationRecordService *OperationRecordService) DeleteSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Delete(&sysOperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: Get a single operation record based on id
//@param: id uint
//@return: sysOperationRecord system.SysOperationRecord, err error

func (operationRecordService *OperationRecordService) GetSysOperationRecord(id uint) (sysOperationRecord system.SysOperationRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysOperationRecordInfoList
//@description: Get the operation record list in pages
//@param: info systemReq.SysOperationRecordSearch
//@return: list interface{}, total int64, err error

func (operationRecordService *OperationRecordService) GetSysOperationRecordInfoList(info systemReq.SysOperationRecordSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
//Create db
	db := global.GVA_DB.Model(&system.SysOperationRecord{})
	var sysOperationRecords []system.SysOperationRecord
// If there is a conditional search, the search statement will be automatically created below
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return sysOperationRecords, total, err
}
