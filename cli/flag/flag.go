package flag

import (
	stdflag "flag"
	"strings"

	"github.com/spf13/pflag"

	"github.com/eachinchung/log"
)

// WordSepNormalizeFunc 更改所有包含“_”分隔符的标志。
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// InitFlags 规范化、解析，然后记录命令行标志。
func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(stdflag.CommandLine)
}

// PrintFlags 将标志记录在 FlagSet 中。
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Debugf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
