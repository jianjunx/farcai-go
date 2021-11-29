package utils

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []interface{}, i int) interface{} {
	return strs[i+1]
}
