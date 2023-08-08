package cfg

const (
	registryK         = "registry"
	logK              = "log"
	databaseK         = "database"
	DefaultConfigName = "registry"
	DefaultConfigType = "yaml"
	DefaultConfigPath = "."
)

var (
	// CandidateConfigPath 临时路径便于开发 配置文件搜索路径
	CandidateConfigPath = []string{
		"E:/Projektes/go/tuzi-tiktok/shared/config/",
	}
)
