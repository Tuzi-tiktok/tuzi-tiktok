package cfg

var (
	Registration registration
	LoggerConfig loggerConfig
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

// Named 内置模块命名 (待用)
type Named interface {
	AppName() string
}
type App struct{}

func AppName() string {
	return "Config"
}
