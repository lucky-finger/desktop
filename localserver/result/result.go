package result

import (
	"encoding/json"
	"os"
)

// StdoutResult 标准输出结果
type StdoutResult struct {
	// Code 状态码
	Code ErrorCode `json:"code"`
	// Msg 错误消息
	Msg string `json:"msg"`
	// Success 是否成功
	Success bool `json:"success"`
	// Data 数据
	Data any `json:"data"`
}

func OutputSuccessData(data any) {
	OutputStdoutResult(&StdoutResult{
		Code:    Success,
		Msg:     "success",
		Success: true,
		Data:    data,
	})
}

func OutputError(code ErrorCode, msg string) {
	OutputStdoutResult(&StdoutResult{
		Code:    code,
		Msg:     msg,
		Success: false,
	})
	os.Exit(int(code))
}

func OutputStdoutResult(result *StdoutResult) {
	marshal, _ := json.Marshal(result)
	_, _ = os.Stdout.Write(marshal)
}
