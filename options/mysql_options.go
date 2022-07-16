package options

import (
	"time"

	"github.com/eachinchung/component-base/db/logger"

	"gorm.io/gorm"

	"github.com/spf13/pflag"

	"github.com/eachinchung/component-base/db"
)

// MySQLOptions 为mysql数据库定义选项
type MySQLOptions struct {
	Host                  string        `json:"host,omitempty"                     mapstructure:"host"`
	Username              string        `json:"username,omitempty"                 mapstructure:"username"`
	Password              string        `json:"-"                                  mapstructure:"password"`
	Database              string        `json:"database"                           mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `json:"log-level"                          mapstructure:"log-level"`
	LogEnableColor        bool          `json:"log-enable-color"                   mapstructure:"log-enable-color"`
}

// NewMySQLOptions 创建一个带有默认参数的 MySQLOptions 对象。
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:                  "127.0.0.1:3306",
		Username:              "",
		Password:              "",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Minute,
		LogLevel:              1,
	}
}

// Validate 验证选项字段。
func (o *MySQLOptions) Validate() []error {
	return []error{}
}

// AddFlags 将 mysql 的各个字段追加到传入的 pflag.FlagSet 变量中。
func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host, "MySQL服务主机地址。如果留空，以下相关的mysql选项将被忽略")

	fs.StringVar(&o.Username, "mysql.username", o.Username, "访问mysql服务的用户名")

	fs.StringVar(&o.Password, "mysql.password", o.Password, "用于访问mysql的密码，应与密码配对使用")

	fs.StringVar(&o.Database, "mysql.database", o.Database, "服务器要使用的数据库名称")

	fs.IntVar(
		&o.MaxIdleConnections,
		"mysql.max-idle-connections",
		o.MaxOpenConnections,
		"允许连接到mysql的最大空闲连接",
	)

	fs.IntVar(
		&o.MaxOpenConnections,
		"mysql.max-open-connections",
		o.MaxOpenConnections,
		"允许连接到mysql的最大开放连接",
	)

	fs.DurationVar(
		&o.MaxConnectionLifeTime,
		"mysql.max-connection-life-time",
		o.MaxConnectionLifeTime,
		"允许连接到mysql的最大连接寿命",
	)

	fs.IntVar(&o.LogLevel, "mysql.log-level", o.LogLevel, "指定gorm日志级别")

	fs.BoolVar(&o.LogEnableColor,
		"mysql.log-enable-color",
		o.LogEnableColor,
		"是否开启颜色输出，true:是，false:否",
	)
}

// NewClient 使用给定的配置创建 MySQL DB。
func (o *MySQLOptions) NewClient() (*gorm.DB, error) {
	opts := &db.Options{
		Host:                  o.Host,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
		Logger:                logger.New(o.LogLevel, o.LogEnableColor),
	}

	return db.New(opts)
}
