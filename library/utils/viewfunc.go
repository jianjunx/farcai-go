package utils

import "github.com/gogf/gf/os/gtime"

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []interface{}, i int) interface{} {
	return strs[i+1]
}

func DateDay(d string) string {
	return gtime.NewFromStr(d).Format("Y-m-d")
}
