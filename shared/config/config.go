package cfg

var (
	Registration   registration
	LoggerConfig   loggerConfig
	DatabaseConfig databaseConfig
)

type registration struct {
	Host        string
	Port        uint64
	NamespaceId string
	Group       string
	DataId      string
}

type loggerConfig struct {
	Level       string
	Development bool
	Encoding    string
}

type databaseConfig struct {
	Dsn          string // 数据来源名称
	Username     string // 用户
	Password     string // 密码
	Host         string // 数据库链接
	Port         uint64 // 数据库端口
	DataBaseName string // 数据库名
	Timeout      string // 连接超时
}

// Named 内置模块命名 (待用)
type Named interface {
	AppName() string
}
type App struct{}

func AppName() string {
	return "Config"
}
