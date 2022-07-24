package iputil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalIP(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_get_local_ip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GetLocalIP())
		})
	}
}

func TestGetRemoteIP(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_get_remote_ip",
			args: args{
				req: httptest.NewRequest("GET", "/", nil),
			},
			want: "192.0.2.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetRemoteIP(tt.args.req), "GetRemoteIP(%v)", tt.args.req)
		})
	}
}
