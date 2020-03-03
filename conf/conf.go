package conf

type (
	DbConf struct {
		Host   string   `json:"Host"`
		Port   string   `json:"Port"`
		User   string   `json:"User"`
		Pwd    string   `json:"Password"`
		DbName string   `json:"Database"`
		Tables []string `json:"Tables"`
	}
	Config struct {
		Db DbConf `json:"Db"`
		// 包名
		ModelPackage string `json:"ModelPackage"`
		// 是否覆盖已存在model
		ModelReplace bool `json:"ModelReplace"`
		// model保存路径
		ModelPath string `json:"ModelPath"`
	}
)

var Conf Config
