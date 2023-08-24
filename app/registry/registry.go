package registry

import (
	"database/sql"

	"github.com/yamato0204/go-test.git/app/controller"
	"github.com/yamato0204/go-test.git/app/infra"
	"github.com/yamato0204/go-test.git/app/usecase"
)


func NewSql() *sql.DB {
	sqlhandler := infra.NewSqlHandler()
	return sqlhandler
} 




func NewRsolver() controller.IUserController {

	sql := NewSql()
	i := infra.NewUserInfra(sql)
	u := usecase.NewUserUsecase(i)

	return controller.NewUserController(u)
}