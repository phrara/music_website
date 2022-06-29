package tool

import "errors"

type Res map[string]any

// 成功的请求
func GetGoodResult(data ...any) Res {
	switch len(data) {
	// 常规请求（不携带返回数据）
	case 0:
		return Res{
			"msg": "ok",
			"data": nil,
		}
	// 常规请求（携带返回数据）
	case 1:
		return Res{
			"msg":  "ok",
			"data": data[0],
		}
	// 登录请求
	case 2:
		dataMap := map[string]any{
			"user":   data[0],
			"uToken": data[1],
		}
		return Res{
			"msg":  "ok",
			"data": dataMap,
		}
	default:
		panic(errors.New("too many args"))
	}
}

// 失败的请求
func GetBadResult(msg string) Res {
	return Res{
		"msg": msg,
		"data": nil,
	}
}