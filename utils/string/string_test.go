package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCaseToUnderscore(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_camel_case_to_underscore",
			args: args{
				str: "TestCamelCaseToUnderscore",
			},
			want: "test_camel_case_to_underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CamelCaseToUnderscore(tt.args.str), "CamelCaseToUnderscore(%v)", tt.args.str)
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		base    []string
		exclude []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "test_diff",
			args: args{
				base:    []string{"a", "b", "c"},
				exclude: []string{"a", "c"},
			},
			wantResult: []string{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.wantResult, Diff(tt.args.base, tt.args.exclude), "Diff(%v, %v)", tt.args.base, tt.args.exclude)
		})
	}
}

func TestFindString(t *testing.T) {
	type args struct {
		array []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_find_string",
			args: args{
				array: []string{"a", "b", "c"},
				str:   "b",
			},
			want: 1,
		},
		{
			name: "test_find_string_not_found",
			args: args{
				array: []string{"a", "b", "c"},
				str:   "d",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FindString(tt.args.array, tt.args.str), "FindString(%v, %v)", tt.args.array, tt.args.str)
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_reverse",
			args: args{
				s: "abc",
			},
			want: "cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Reverse(tt.args.s), "Reverse(%v)", tt.args.s)
		})
	}
}

func TestStringIn(t *testing.T) {
	type args struct {
		str   string
		array []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_string_in",
			args: args{
				str:   "a",
				array: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "test_string_not_in",
			args: args{
				str:   "d",
				array: []string{"a", "b", "c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringIn(tt.args.str, tt.args.array), "StringIn(%v, %v)", tt.args.str, tt.args.array)
		})
	}
}

func TestUnderscoreToCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_underscore_to_camel_case",
			args: args{
				str: "test_underscore_to_camel_case",
			},
			want: "TestUnderscoreToCamelCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UnderscoreToCamelCase(tt.args.str), "UnderscoreToCamelCase(%v)", tt.args.str)
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "test_unique",
			args: args{
				ss: []string{"a", "b", "c", "a"},
			},
			wantResult: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.wantResult, Unique(tt.args.ss), "Unique(%v)", tt.args.ss)
		})
	}
}
