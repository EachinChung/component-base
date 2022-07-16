package version

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		want Info
	}{
		{
			name: "Get",
			want: Info{
				GitVersion:   "v0.0.0-master+$Format:%h$",
				GitCommit:    "$Format:%H$",
				GitTreeState: "",
				BuildDate:    "1970-01-01T00:00:00Z",
				GoVersion:    runtime.Version(),
				Compiler:     runtime.Compiler,
				Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValuesf(t, tt.want, Get(), "Get()")
		})
	}
}

func TestInfo_String(t *testing.T) {
	type fields struct {
		GitVersion   string
		GitCommit    string
		GitTreeState string
		BuildDate    string
		GoVersion    string
		Compiler     string
		Platform     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String",
			want: "  gitVersion: \n   gitCommit: \ngitTreeState: \n   buildDate: \n   goVersion: \n    compiler: \n    platform: ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := Info{
				GitVersion:   tt.fields.GitVersion,
				GitCommit:    tt.fields.GitCommit,
				GitTreeState: tt.fields.GitTreeState,
				BuildDate:    tt.fields.BuildDate,
				GoVersion:    tt.fields.GoVersion,
				Compiler:     tt.fields.Compiler,
				Platform:     tt.fields.Platform,
			}
			assert.Equalf(t, tt.want, info.String(), "String()")
		})
	}
}

func TestInfo_Text(t *testing.T) {
	type fields struct {
		GitVersion   string
		GitCommit    string
		GitTreeState string
		BuildDate    string
		GoVersion    string
		Compiler     string
		Platform     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Text",
			want:    []byte("  gitVersion: \n   gitCommit: \ngitTreeState: \n   buildDate: \n   goVersion: \n    compiler: \n    platform: "),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := Info{
				GitVersion:   tt.fields.GitVersion,
				GitCommit:    tt.fields.GitCommit,
				GitTreeState: tt.fields.GitTreeState,
				BuildDate:    tt.fields.BuildDate,
				GoVersion:    tt.fields.GoVersion,
				Compiler:     tt.fields.Compiler,
				Platform:     tt.fields.Platform,
			}
			got, err := info.Text()
			if !tt.wantErr(t, err, fmt.Sprintf("Text()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "Text()")
		})
	}
}

func TestInfo_ToJSON(t *testing.T) {
	type fields struct {
		GitVersion   string
		GitCommit    string
		GitTreeState string
		BuildDate    string
		GoVersion    string
		Compiler     string
		Platform     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ToJSON",
			want: "{\"gitVersion\":\"\",\"gitCommit\":\"\",\"gitTreeState\":\"\",\"buildDate\":\"\",\"goVersion\":\"\",\"compiler\":\"\",\"platform\":\"\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := Info{
				GitVersion:   tt.fields.GitVersion,
				GitCommit:    tt.fields.GitCommit,
				GitTreeState: tt.fields.GitTreeState,
				BuildDate:    tt.fields.BuildDate,
				GoVersion:    tt.fields.GoVersion,
				Compiler:     tt.fields.Compiler,
				Platform:     tt.fields.Platform,
			}
			assert.Equalf(t, tt.want, info.ToJSON(), "ToJSON()")
		})
	}
}
