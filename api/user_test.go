package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"good case", `{"username": "kyr", "password": "123321"}`, "注册成功"},
	}

	r := gin.Default()
	r.POST("/user/register", Register)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/user/register",            // 请求URL
				strings.NewReader(tt.param), // 请求参数
			)
			// mock一个响应记录器
			w := httptest.NewRecorder()
			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)
			// 解析并检验响应内容是否复合预期
			var resp map[string]string
			_ = json.Unmarshal([]byte(w.Body.String()), &resp)
			//assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"good case", `{"username": "kyr", "password": "123321"}`, "登录成功"},
	}

	r := gin.Default()
	r.POST("/user/login", Login)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// httptest 这个包是为了mock一个HTTP请求
			// 参数分别是请求方法，请求URL，请求Body
			// Body只能使用io.Reader
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/user/login",               // 请求URL
				strings.NewReader(tt.param), // 请求参数
			)
			// mock一个响应记录器
			w := httptest.NewRecorder()
			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			// 校验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)
			// 解析并检验响应内容是否复合预期
			var resp map[string]string
			_ = json.Unmarshal([]byte(w.Body.String()), &resp)
			//assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
