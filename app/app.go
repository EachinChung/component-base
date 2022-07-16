package app

//goland:noinspection SpellCheckingInspection
import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/eachinchung/errors"
	"github.com/eachinchung/log"

	"github.com/eachinchung/component-base/cli/flag"
	"github.com/eachinchung/component-base/cli/globalflag"
	"github.com/eachinchung/component-base/version"
	vflag "github.com/eachinchung/component-base/version/flag"
)

// RunFunc 定义了应用程序的启动回调函数。
type RunFunc func(basename string) error

// Application cli 应用程序的主要结构。
// 建议使用 app.NewApplication() 函数创建应用程序。
type Application struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool

	args cobra.PositionalArgs
	cmd  *cobra.Command
}

// Option 定义用于初始化应用程序结构的可选参数。
type Option func(*Application)

// WithOptions 打开应用程序的函数从命令行读取或从配置文件读取参数。
func WithOptions(opt CliOptions) Option {
	return func(a *Application) {
		a.options = opt
	}
}

// WithRunFunc 用于设置应用程序启动回调函数选项。
func WithRunFunc(run RunFunc) Option {
	return func(a *Application) {
		a.runFunc = run
	}
}

// WithDescription 用于设置应用程序的描述。
func WithDescription(desc string) Option {
	return func(a *Application) {
		a.description = desc
	}
}

// WithSilence 将应用程序设置为静音模式，
// 在该模式下，程序启动信息、配置信息和版本信息不会打印在控制台中。
func WithSilence() Option {
	return func(a *Application) {
		a.silence = true
	}
}

// WithNoConfig set the application does not provide config flag.
func WithNoConfig() Option {
	return func(a *Application) {
		a.noConfig = true
	}
}

// WithNoVersion set the application does not provide version flag.
func WithNoVersion() Option {
	return func(a *Application) {
		a.noVersion = true
	}
}

// WithValidArgs set the validation function to valid non-flag arguments.
func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *Application) {
		a.args = args
	}
}

// WithDefaultValidArgs set default validation function to valid non-flag arguments.
func WithDefaultValidArgs() Option {
	return func(a *Application) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

// NewApplication 基于给定的应用程序名称、二进制名称和其他选项创建一个新的应用程序实例。
func NewApplication(name string, basename string, opts ...Option) *Application {
	a := &Application{
		name:     name,
		basename: basename,
	}

	for _, o := range opts {
		o(a)
	}

	a.buildCommand()

	return a
}

// Run 用于启动应用程序。
func (a *Application) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// Command 返回应用程序内的 cobra 实例。
func (a *Application) Command() *cobra.Command {
	return a.cmd
}

func (a *Application) buildCommand() {
	cmd := cobra.Command{
		Use:   formatBaseName(a.basename),
		Short: a.name,
		Long:  a.description,
		// 命令错误时 stop printing usage
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	flag.InitFlags(cmd.Flags())

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	var namedFlagSets flag.NamedFlagSets
	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	if !a.noConfig {
		addConfigFlag(a.basename, namedFlagSets.FlagSet("global"))
	}
	if !a.noVersion {
		vflag.AddFlags(namedFlagSets.FlagSet("global"))
	}
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())
	// 将新的全局 FlagSet 添加到 cmd FlagSet
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet("global"))

	addCmdTemplate(&cmd, namedFlagSets)
	a.cmd = &cmd
}

func (a *Application) runCommand(cmd *cobra.Command, args []string) error {
	printWorkingDir()
	flag.PrintFlags(cmd.Flags())
	if !a.noVersion {
		// 显示应用程序版本信息
		vflag.PrintAndExitIfRequested()
	}

	if !a.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}

	if !a.silence {
		log.Infof("%v Starting %s ...", progressMessage, a.name)
		if !a.noVersion {
			log.Infof("%v Version: `%s`", progressMessage, version.Get().ToJSON())
		}
		if !a.noConfig {
			log.Infof("%v Config file used: `%s`", progressMessage, viper.ConfigFileUsed())
		}
	}
	if a.options != nil {
		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}
	// 运行应用程序
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}

func (a *Application) applyOptionRules() error {
	if completableOptions, ok := a.options.(CompletableOptions); ok {
		if err := completableOptions.Complete(); err != nil {
			return err
		}
	}

	if errs := a.options.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs...)
	}

	if printableOptions, ok := a.options.(PrintableOptions); ok && !a.silence {
		log.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	return nil
}
