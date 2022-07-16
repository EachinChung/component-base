package flag

import (
	"fmt"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestAddFlags(t *testing.T) {
	type args struct {
		fs *pflag.FlagSet
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "add flags",
			args: args{
				fs: pflag.NewFlagSet("test", pflag.ContinueOnError),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddFlags(tt.args.fs)
		})
	}
}

func TestPrintAndExitIfRequested(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "print version and exit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAndExitIfRequested()
		})
	}
}

func TestVersion(t *testing.T) {
	type args struct {
		name      string
		shorthand string
		value     versionValue
		usage     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "version",
			args: args{
				name:      "test",
				shorthand: "t",
				value:     VersionFalse,
				usage:     "打印版本信息并退出。",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version(tt.args.name, tt.args.shorthand, tt.args.value, tt.args.usage)
		})
	}
}

func TestVersionVar(t *testing.T) {
	type args struct {
		p         *versionValue
		name      string
		shorthand string
		value     versionValue
		usage     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "versionVar",
			args: args{
				p:         new(versionValue),
				name:      "test2",
				shorthand: "2",
				value:     VersionFalse,
				usage:     "打印版本信息并退出。",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VersionVar(tt.args.p, tt.args.name, tt.args.shorthand, tt.args.value, tt.args.usage)
		})
	}
}

func Test_versionValue_Get(t *testing.T) {
	tests := []struct {
		name string
		v    versionValue
	}{
		{
			name: "get",
			v:    VersionFalse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Get()
		})
	}
}

func Test_versionValue_IsBoolFlag(t *testing.T) {
	tests := []struct {
		name string
		v    versionValue
		want bool
	}{
		{
			name: "isBoolFlag",
			v:    VersionFalse,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v.IsBoolFlag(), "IsBoolFlag()")
		})
	}
}

func Test_versionValue_Set(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		v       versionValue
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "set",
			v:    VersionFalse,
			args: args{
				s: "true",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.v.Set(tt.args.s), fmt.Sprintf("Set(%v)", tt.args.s))
		})
	}
}

func Test_versionValue_String(t *testing.T) {
	tests := []struct {
		name string
		v    versionValue
		want string
	}{
		{
			name: "string",
			v:    VersionFalse,
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v.String(), "String()")
		})
	}
}

func Test_versionValue_Type(t *testing.T) {
	tests := []struct {
		name string
		v    versionValue
		want string
	}{
		{
			name: "type",
			v:    VersionFalse,
			want: "version",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v.Type(), "Type()")
		})
	}
}
