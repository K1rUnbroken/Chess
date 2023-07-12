package api

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name  string
		param string
	}{
		{"good case", `{"username": "kyr", "password": "123321"}`},
	}

	r := gin.Default()
	r.POST("/user/connect", Connect)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader
			req := httptest.NewRequest(
				"GET",                        // 请求方法
				"/user/connect?username=kyr", // 请求URL
				nil,                          // 请求参数
			)
			// mock一个响应记录器
			w := httptest.NewRecorder()
			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
		})
	}
}
