package options

import (
	"github.com/spf13/pflag"
)

// RedisOptions 为 redis 定义选项
type RedisOptions struct {
	Host                  string   `json:"host"                     mapstructure:"host"`
	Port                  int      `json:"port"                     mapstructure:"port"`
	Addrs                 []string `json:"addrs"                    mapstructure:"addrs"`
	Username              string   `json:"username"                 mapstructure:"username"`
	Password              string   `json:"password"                 mapstructure:"password"`
	Database              int      `json:"database"                 mapstructure:"database"`
	MasterName            string   `json:"master-name"              mapstructure:"master-name"`
	MaxIdle               int      `json:"optimisation-max-idle"    mapstructure:"optimisation-max-idle"`
	MaxActive             int      `json:"optimisation-max-active"  mapstructure:"optimisation-max-active"`
	Timeout               int      `json:"timeout"                  mapstructure:"timeout"`
	EnableCluster         bool     `json:"enable-cluster"           mapstructure:"enable-cluster"`
	UseSSL                bool     `json:"use-ssl"                  mapstructure:"use-ssl"`
	SSLInsecureSkipVerify bool     `json:"ssl-insecure-skip-verify" mapstructure:"ssl-insecure-skip-verify"`
}

// NewRedisOptions 创建一个带有默认参数的 RedisOptions 对象。
func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:                  "127.0.0.1",
		Port:                  6379,
		Addrs:                 []string{},
		Username:              "",
		Password:              "",
		Database:              0,
		MasterName:            "",
		MaxIdle:               2000,
		MaxActive:             4000,
		Timeout:               0,
		EnableCluster:         false,
		UseSSL:                false,
		SSLInsecureSkipVerify: false,
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
	fs.StringSliceVar(&o.Addrs, "redis.addrs", o.Addrs, "一组redis地址 (格式: 127.0.0.1:6379)。")
	fs.StringVar(&o.Username, "redis.username", o.Username, "访问redis服务的用户名。")
	fs.StringVar(&o.Password, "redis.password", o.Password, "Redis 身份验证密码。")

	fs.IntVar(
		&o.Database,
		"redis.database",
		o.Database,
		"默认情况下，数据库为0。redis集群不支持设置数据库。因此，如果您具有-redis.enable-cluster = true，则应省略此值或显式设置为0。",
	)

	fs.StringVar(&o.MasterName, "redis.master-name", o.MasterName, "主redis实例的名称。")

	fs.IntVar(
		&o.MaxIdle,
		"redis.optimisation-max-idle",
		o.MaxIdle,
		"此设置将配置空闲时 (无流量) 池中维护的连接数量。将 --redis.optimisation-max-active 设置为较大的值，我们通常将其保留在2000左右，以进行HA部署。",
	)

	fs.IntVar(
		&o.MaxActive,
		"redis.optimisation-max-active",
		o.MaxActive,
		"为了防止Redis服务器的连接过载，我们可能会限制到Redis的活动连接总数。我们建议将其设置为4000左右。",
	)

	fs.IntVar(
		&o.Timeout,
		"redis.timeout",
		o.Timeout,
		"连接到redis服务时超时 (以秒为单位)。",
	)

	fs.BoolVar(
		&o.EnableCluster,
		"redis.enable-cluster",
		o.EnableCluster,
		"如果您使用的是Redis集群，请在此启用它以启用插槽模式。",
	)

	fs.BoolVar(&o.UseSSL, "redis.use-ssl", o.UseSSL, "如果设置，将假定与Redis的连接已加密。")

	fs.BoolVar(
		&o.SSLInsecureSkipVerify,
		"redis.ssl-insecure-skip-verify",
		o.SSLInsecureSkipVerify,
		"允许在连接到加密的Redis数据库时使用自签名证书。",
	)
}
