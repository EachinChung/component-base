package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test hash password",
			args: args{
				password: "123456",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			t.Log(got)
			if !tt.wantErr(t, err, fmt.Sprintf("HashPassword(%v)", tt.args.password)) {
				return
			}
			tt.wantErr(t, ComparePasswordHash(got, "123456"), fmt.Sprintf("ComparePasswordHash(%v)", got))
		})
	}
}
