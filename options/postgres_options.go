package options

import (
	"time"

	"gorm.io/gorm"

	"github.com/spf13/pflag"

	"github.com/eachinchung/component-base/db/logger"
	"github.com/eachinchung/component-base/db/postgres"
)

// PostgresOptions 为Postgres数据库定义选项
type PostgresOptions struct {
	Host                  string        `json:"host,omitempty"                     mapstructure:"host"`
	Port                  int           `json:"port"                               mapstructure:"port"`
	Username              string        `json:"username,omitempty"                 mapstructure:"username"`
	Password              string        `json:"-"                                  mapstructure:"password"`
	Database              string        `json:"database"                           mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `json:"log-level"                          mapstructure:"log-level"`
	LogEnableColor        bool          `json:"log-enable-color"                   mapstructure:"log-enable-color"`
}

// NewPostgresOptions 创建一个带有默认参数的 PostgresOptions 对象。
func NewPostgresOptions() *PostgresOptions {
	return &PostgresOptions{
		Host:                  "127.0.0.1",
		Port:                  5432,
		Username:              "postgres",
		Password:              "123456",
		Database:              "postgres",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Minute,
		LogLevel:              1,
	}
}

// Validate 验证选项字段。
func (o *PostgresOptions) Validate() []error {
	return []error{}
}

// AddFlags 将 Postgres 的各个字段追加到传入的 pflag.FlagSet 变量中。
func (o *PostgresOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "Postgres.host", o.Host, "Postgres服务主机地址。如果留空，以下相关的Postgres选项将被忽略")

	fs.IntVar(&o.Port, "Postgres.port", o.Port, "Postgres服务端口")

	fs.StringVar(&o.Username, "Postgres.username", o.Username, "访问Postgres服务的用户名")

	fs.StringVar(&o.Password, "Postgres.password", o.Password, "用于访问Postgres的密码，应与密码配对使用")

	fs.StringVar(&o.Database, "Postgres.database", o.Database, "服务器要使用的数据库名称")

	fs.IntVar(
		&o.MaxIdleConnections,
		"Postgres.max-idle-connections",
		o.MaxOpenConnections,
		"允许连接到Postgres的最大空闲连接",
	)

	fs.IntVar(
		&o.MaxOpenConnections,
		"Postgres.max-open-connections",
		o.MaxOpenConnections,
		"允许连接到Postgres的最大开放连接",
	)

	fs.DurationVar(
		&o.MaxConnectionLifeTime,
		"Postgres.max-connection-life-time",
		o.MaxConnectionLifeTime,
		"允许连接到Postgres的最大连接寿命",
	)

	fs.IntVar(&o.LogLevel, "Postgres.log-level", o.LogLevel, "指定gorm日志级别")

	fs.BoolVar(&o.LogEnableColor,
		"Postgres.log-enable-color",
		o.LogEnableColor,
		"是否开启颜色输出，true:是，false:否",
	)
}

// NewClient 使用给定的配置创建 Postgres DB。
func (o *PostgresOptions) NewClient() (*gorm.DB, error) {
	opts := &postgres.Options{
		Host:                  o.Host,
		Port:                  o.Port,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
		Logger:                logger.New(o.LogLevel, o.LogEnableColor),
	}

	return postgres.New(opts)
}
