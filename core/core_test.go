package core

import (
	"github.com/eachinchung/errors"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	type args struct {
		c      *gin.Context
		err    error
		detail interface{}
	}
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name string
		args args
	}{
		{
			name: "write response",
			args: args{
				c:      c,
				err:    nil,
				detail: nil,
			},
		},
		{
			name: "write response with error",
			args: args{
				c:      c,
				err:    errors.New("test error"),
				detail: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteResponse(tt.args.c, tt.args.err, tt.args.detail)
		})
	}
}
