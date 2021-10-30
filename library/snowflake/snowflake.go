package snowflake

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func init() {
	var t time.Time
	t, err := time.Parse("2006-01-02", "2021-01-01")
	if err != nil {
		fmt.Println("Snowflake 初始化错误：", err.Error())
	}
	sf.Epoch = t.UnixNano() / 1000000
	node, err = sf.NewNode(1)
	if err != nil {
		fmt.Println("Snowflake 初始化错误：", err.Error())
	}
}

// 生成通用ID
func GenerateId() int64 {
	return node.Generate().Int64()
}
