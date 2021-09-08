package app

import (
	"github.com/rockiecn/p2pdemo/operator"
	"github.com/rockiecn/p2pdemo/provider"
	"github.com/rockiecn/p2pdemo/user"
)

type App struct {
	Op   operator.Operator
	User user.User
	Pro  provider.Provider
}

func (app *App) Init() {
	app.Op = operator.Operator{}
	app.Op.Init()

	app.User = user.User{}
	app.User.Init()

	app.Pro = provider.Provider{}
	app.Pro.Init()

}

func (app *App) Exit() {
	app.Op.CloseDB()
	app.User.CloseDB()
	app.Pro.CloseDB()
}
