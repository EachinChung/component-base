package app

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/eachinchung/log"

	"github.com/eachinchung/component-base/cli/flag"
	"github.com/eachinchung/component-base/terminal"
)

var progressMessage = color.GreenString("==>")

func printWorkingDir() {
	wd, _ := os.Getwd()
	log.Infof("%v Working Dir: %s", progressMessage, wd)
}

func addCmdTemplate(cmd *cobra.Command, namedFlagSets flag.NamedFlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := terminal.Size(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		_, _ = fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		flag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)

		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		flag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
}

// formatBaseName 根据给定的名称格式化为不同操作系统下的可执行文件名。
func formatBaseName(basename string) string {
	// 使大小写不敏感，并剥离可执行后缀 (如果存在)
	//goland:noinspection GoBoolExpressions
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}
