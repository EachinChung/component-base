package options

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/pflag"
)

// JWTOptions jwt 配置选项
type JWTOptions struct {
	Realm      string        `json:"realm"       mapstructure:"realm"`
	Key        string        `json:"key"         mapstructure:"key"`
	Timeout    time.Duration `json:"timeout"     mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}

// NewJWTOptions 创建一个带有默认参数的 JWTOptions 对象。
func NewJWTOptions() *JWTOptions {
	return &JWTOptions{
		Realm:      "JWT",
		Key:        "Key",
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
	}
}

// Validate 验证选项字段。
func (s *JWTOptions) Validate() []error {
	var errs []error

	if !govalidator.StringLength(s.Key, "6", "32") {
		errs = append(errs, fmt.Errorf("--secret-key 必须大于5且小于33"))
	}

	return errs
}

// AddFlags 将 jwt 的各个字段追加到传入的 pflag.FlagSet 变量中。
func (s *JWTOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.Realm, "jwt.realm", s.Realm, "要显示给用户的领域名称")
	fs.StringVar(&s.Key, "jwt.key", s.Key, "用于签署jwt令牌的私钥")
	fs.DurationVar(&s.Timeout, "jwt.timeout", s.Timeout, "JWT令牌超时")

	fs.DurationVar(
		&s.MaxRefresh,
		"jwt.max-refresh",
		s.MaxRefresh,
		"此字段允许客户端刷新其令牌，直到超过 MaxRefresh 时间",
	)
}
