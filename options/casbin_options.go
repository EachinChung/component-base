package options

import (
	"github.com/spf13/pflag"
)

// CasbinOptions jwt 配置选项
type CasbinOptions struct {
	Model string `json:"model" mapstructure:"realm"`
}

// NewCasbinOptions 创建一个带有默认参数的 CasbinOptions 对象。
func NewCasbinOptions() *CasbinOptions {
	return &CasbinOptions{
		Model: "configs/model.conf",
	}
}

// Validate 验证选项字段。
func (s *CasbinOptions) Validate() []error {
	return []error{}
}

// AddFlags 将 casbin 的各个字段追加到传入的 pflag.FlagSet 变量中。
func (s *CasbinOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.Model, "casbin.model", s.Model, "casbin 权限模型文件地址")
}
