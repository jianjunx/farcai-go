package boot

import (
	_ "farcai-go/library/cos"
	_ "farcai-go/library/snowflake"
	"farcai-go/library/utils"
	_ "farcai-go/packed"
	"os"

	"github.com/gogf/gf/frame/g"
)

func init() {
	bindMysqlConf()
	bindViewFunc()
}

func bindViewFunc() {
	// view 自定义函数
	g.View().BindFunc("isODD", utils.IsODD)
	g.View().BindFunc("getNextName", utils.GetNextName)
}

func bindMysqlConf() {
	link := os.Getenv("MYSQL_LINK")
	if link != "" {
		g.Cfg().Set("database.link", link)
	}

}
