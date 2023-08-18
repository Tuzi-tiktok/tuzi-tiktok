package cfg

const (
	registryK         = "registry"
	logK              = "log"
	DefaultConfigType = "yaml"
	DefaultConfigPath = "."
)

var (
	ConfigEnv         = "Default"
	DefaultConfigName = "registry"
	// CandidateConfigPath 临时路径便于开发 配置文件搜索路径
	CandidateConfigPath string
)
