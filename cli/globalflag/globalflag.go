package globalflag

import (
	"fmt"

	"github.com/spf13/pflag"
)

// AddGlobalFlags 显式注册库从 flag 向全局 FlagSet 注册的标志。
func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	fs.BoolP("help", "h", false, fmt.Sprintf("有关 %s 的帮助信息", name))
}
