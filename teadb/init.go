package teadb

import (
	"github.com/iwind/TeaGo"
	"github.com/iwind/TeaGo/Tea"
)

func init() {
	// 在测试环境下直接建立数据库，在二进制环境下需要等服务启动的时候才启动
	if Tea.IsTesting() {
		SetupDB()
	} else {
		TeaGo.BeforeStart(func(server *TeaGo.Server) {
			SetupDB()
		})
	}
}
