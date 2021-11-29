package boot

import (
	_ "farcai-go/library/dynamodb"
	_ "farcai-go/library/snowflake"
	"farcai-go/library/utils"
	_ "farcai-go/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	// view 自定义函数
	g.View().BindFunc("isODD", utils.IsODD)
	g.View().BindFunc("getNextName", utils.GetNextName)
}
