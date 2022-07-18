package verification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordPower(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_password_power",
			args: args{
				pwd: "123456",
			},
			want: false,
		},
		{
			name: "test_password_power",
			args: args{
				pwd: "123456789Qwerty123.",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PasswordPower(tt.args.pwd), "PasswordPower(%v)", tt.args.pwd)
		})
	}
}

func TestPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_phone",
			args: args{
				phone: "13800138000",
			},
			want: true,
		},
		{
			name: "test_phone_invalid",
			args: args{
				phone: "138001380000",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Phone(tt.args.phone), "Phone(%v)", tt.args.phone)
		})
	}
}

func TestEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_email",
			args: args{
				email: "test@a.com",
			},
			want: true,
		},
		{
			name: "test_email_invalid",
			args: args{
				email: "test@a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Email(tt.args.email), "Email(%v)", tt.args.email)
		})
	}
}
