package options

import (
	"github.com/spf13/pflag"
)

// RedisOptions 为 redis 定义选项
type RedisOptions struct {
	Host     string `json:"host"                     mapstructure:"host"`
	Port     int    `json:"port"                     mapstructure:"port"`
	Password string `json:"password"                 mapstructure:"password"`
	Database int    `json:"database"                 mapstructure:"database"`
}

// NewRedisOptions 创建一个带有默认参数的 RedisOptions 对象。
func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123456",
		Database: 0,
	}
}

// Validate 验证选项字段。
func (o *RedisOptions) Validate() []error {
	var errs []error

	return errs
}

// AddFlags 将 redis 的各个字段追加到传入的 pflag.FlagSet 变量中。
func (o *RedisOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "redis.host", o.Host, "Redis服务器的主机名。")
	fs.IntVar(&o.Port, "redis.port", o.Port, "Redis服务器监听的端口。")
	fs.StringVar(&o.Password, "redis.password", o.Password, "Redis 身份验证密码。")
	fs.IntVar(
		&o.Database,
		"redis.database",
		o.Database,
		"默认情况下，数据库为0。redis集群不支持设置数据库。因此，如果您具有-redis.enable-cluster = true，则应省略此值或显式设置为0。",
	)
}
