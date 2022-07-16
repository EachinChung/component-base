package app

import "github.com/eachinchung/component-base/cli/flag"

// CliOptions abstracts 从命令行读取参数的配置选项。
type CliOptions interface {
	Flags() (fss flag.NamedFlagSets)
	Validate() []error
}

// CompletableOptions abstracts 完成的选项。
type CompletableOptions interface {
	Complete() error
}

// PrintableOptions abstracts 打印的选项。
type PrintableOptions interface {
	String() string
}
