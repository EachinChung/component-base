package flag

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/eachinchung/component-base/version"
)

type versionValue int

const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)

const (
	versionFlagName  = "version"
	versionShorthand = "v"
)

const strRawVersion string = "raw"

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() any {
	return v
}

func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", *v == VersionTrue)
}

// Type pflag.Value 接口要求的pflag 类型。
func (v *versionValue) Type() string {
	return "version"
}

// VersionVar 定义了一个具有指定名称和用法字符串的 pflag。
func VersionVar(p *versionValue, name, shorthand string, value versionValue, usage string) {
	*p = value
	pflag.VarP(p, name, shorthand, usage)
	// "--version" 将被视为 "--version=true"
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version 包装了 VersionVar 函数。
func Version(name, shorthand string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, shorthand, value, usage)
	return p
}

var versionFlag = Version(versionFlagName, versionShorthand, VersionFalse, "打印版本信息并退出。")

// AddFlags 在任意 FlagSet 上注册此包的标志，以便它们指向与全局标志相同的值。
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested 将检查 -version 是否被传递，如果是，则打印版本并退出。
func PrintAndExitIfRequested() {
	switch *versionFlag {
	case VersionRaw:
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	case VersionTrue:
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
