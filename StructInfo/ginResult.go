package StructInfo

var (
	OK  = response(200, "ok") // 通用成功
	Err = response(500, "")   // 通用错误

	FileReadErr   = response(1001, "JSON文件读取失败")
	FilePathErr   = response(1001, "获取文件路径失败")
	UnmarshalErr  = response(1003, "序列化失败")
	FileCreateErr = response(1004, "文件创建失败")
	GetSoliderErr = response(1004, "获取士兵失败")
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
	Err  error       `json:"err"`
}

// 自定义响应信息

func (res *Response) WithMsg(message string, err error) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
		Err:  err,
	}
}

// 追加响应数据

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
		Err:  nil,
	}
}
