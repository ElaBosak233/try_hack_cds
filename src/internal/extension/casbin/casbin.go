package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/elabosak233/cloudsdale/internal/app/db"
	"github.com/elabosak233/cloudsdale/internal/files"
	"go.uber.org/zap"
)

var (
	Enforcer *casbin.Enforcer
)

func InitCasbin() {
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(
		db.Db(),
		&gormadapter.CasbinRule{},
		"casbins",
	)
	cfg, err := files.F().ReadFile("configs/casbin.conf")
	md, _ := model.NewModelFromString(string(cfg))
	Enforcer, err = casbin.NewEnforcer(md, adapter)
	if err != nil {
		zap.L().Fatal("Casbin module inits failed.", zap.Error(err))
	}
	Enforcer.ClearPolicy()
	_ = Enforcer.SavePolicy()
	initDefaultPolicy()
	zap.L().Info("Casbin module inits successfully.")
}
