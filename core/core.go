package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eachinchung/errors"
	"github.com/eachinchung/log"
)

type response struct {
	httpStatus int
	err        error
}

type Option func(*response)

func WithHttpStatus(httpStatus int) Option {
	return func(r *response) {
		r.httpStatus = httpStatus
	}
}

func WithError(err error) Option {
	return func(r *response) {
		r.err = err
	}
}

// ErrResponse 定义发生错误时的返回消息。
type ErrResponse struct {
	// ErrCode 定义业务错误代码。
	ErrCode int `json:"err_code"`

	// Message 包含此消息的详细信息。
	// 此消息适合暴露于外部
	Message string `json:"message"`

	// Detail 返回可能对解决此错误有用的详细信息。
	Detail interface{} `json:"detail,omitempty"`
}

// WriteResponse 将错误或响应数据写入http响应主体。
// 它使用 errors.ParseCoder 将任何错误解析为 errors.Coder
// errors.Coder 包含错误代码、用户安全错误消息和 http 状态代码。
func WriteResponse(c *gin.Context, detail interface{}, opts ...Option) {
	r := &response{
		httpStatus: http.StatusOK,
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.err != nil {
		coder := errors.ParseCoder(r.err)
		if coder.Code() == 1 {
			log.Errorf("检测到未知错误, err: %+v", r.err)
		}
		c.JSON(coder.HTTPStatus(), ErrResponse{
			ErrCode: coder.Code(),
			Message: coder.String(),
			Detail:  detail,
		})
		return
	}

	c.JSON(r.httpStatus, detail)
}
