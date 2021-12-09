package model

type Response struct {
	Error string      `json:"error"` // 错误信息
	Data  interface{} `json:"data"`  // 数据
	Code  int         `json:"code"`  // 状态码
}
