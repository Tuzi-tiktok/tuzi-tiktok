package cfg

const (
	registryK         = "registry"
	logK              = "log"
	databaseK         = "database"
	DefaultConfigType = "yaml"
	DefaultConfigPath = "."
	AppENVIRONMENT    = "TUZI_ENV"
)

var (
	ConfigEnv         = "Default"
	DefaultConfigName = "registry"
	// CandidateConfigPath 临时路径便于开发 配置文件搜索路径
	CandidateConfigPath = []string{
		"E:/Projektes/go/tuzi-tiktok/shared/config/",
		`C:\Users\Admin\GolandProjects\tuzi-tiktok\shared\config`,
	}
)
