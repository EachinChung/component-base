package idutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenSecretID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test_gen_secret_id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GenSecretID())
		})
	}
}

func TestGenSecretKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test_gen_secret_key"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GenSecretKey())
		})
	}
}

func TestGenUint64ID(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{name: "test_gen_uint64_id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GenUint64ID())
		})
	}
}

func Test_randString(t *testing.T) {
	type args struct {
		letters string
		n       int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_rand_string",
			args: args{
				letters: Alphabet36,
				n:       10,
			},
		},
		{
			name: "test_rand_string_with_alphabet62",
			args: args{
				letters: Alphabet62,
				n:       32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randString(tt.args.letters, tt.args.n)
			t.Log(got)
			assert.Lenf(t, got, tt.args.n, "randString(%v, %v)", tt.args.letters, tt.args.n)
		})
	}
}

func TestGetInstanceID(t *testing.T) {
	type args struct {
		uid    uint64
		prefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_get_instance_id",
			args: args{
				uid:    1,
				prefix: "test",
			},
			want: "test-3935mq",
		},
		{
			name: "test_get_instance_id_int64",
			args: args{
				uid:    9223372036854775807,
				prefix: "test",
			},
			want: "test-dk7kpn8j1ddrnlo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetInstanceID(tt.args.uid, tt.args.prefix), "GetInstanceID(%v, %v)", tt.args.uid, tt.args.prefix)
		})
	}
}
