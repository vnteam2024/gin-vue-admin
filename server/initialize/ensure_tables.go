package initialize

import (
	"context"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
)

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// auto run
func init() {
	system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (ensureTables) InitializerName() string {
	return "ensure_tables_created"
}
func (e *ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *ensureTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *ensureTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		sysModel.SysApi{},
		sysModel.SysUser{},
		sysModel.SysBaseMenu{},
		sysModel.SysAuthority{},
		sysModel.JwtBlacklist{},
		sysModel.SysDictionary{},
		sysModel.SysAutoCodeHistory{},
		sysModel.SysOperationRecord{},
		sysModel.SysDictionaryDetail{},
		sysModel.SysBaseMenuParameter{},
		sysModel.SysBaseMenuBtn{},
		sysModel.SysAuthorityBtn{},
		sysModel.SysAutoCode{},
		sysModel.SysExportTemplate{},
		sysModel.Condition{},

		adapter.CasbinRule{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
// The view authority_menu will be created as a table, causing a conflict error (newer versions of gorm don't seem to)
// Since AutoMigrate() basically does not need to consider errors, it is explicitly ignored
	}
	return ctx, nil
}

func (e *ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		sysModel.SysApi{},
		sysModel.SysUser{},
		sysModel.SysBaseMenu{},
		sysModel.SysAuthority{},
		sysModel.JwtBlacklist{},
		sysModel.SysDictionary{},
		sysModel.SysAutoCodeHistory{},
		sysModel.SysOperationRecord{},
		sysModel.SysDictionaryDetail{},
		sysModel.SysBaseMenuParameter{},
		sysModel.SysBaseMenuBtn{},
		sysModel.SysAuthorityBtn{},
		sysModel.SysAutoCode{},
		sysModel.SysExportTemplate{},
		sysModel.Condition{},

		adapter.CasbinRule{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
