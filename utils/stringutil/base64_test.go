package stringutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBase64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "base64 decode",
			args: args{
				s: "aGVsbG8gd29ybGQ=",
			},
			want:    []byte("hello world"),
			wantErr: assert.NoError,
		},
		{
			name: "base64 decode with invalid input",
			args: args{
				s: "aGVsbG8gd29ybGQ",
			},
			want:    []byte("hello wor"),
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64(tt.args.s)
			if !tt.wantErr(t, err, fmt.Sprintf("DecodeBase64(%v)", tt.args.s)) {
				return
			}
			assert.Equalf(t, tt.want, got, "DecodeBase64(%v)", tt.args.s)
		})
	}
}
