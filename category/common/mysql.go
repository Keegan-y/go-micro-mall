package common


import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	Database string `json:"database"`
	Port string `json:"port"`
}

// 获取 mysql 的配置
func GetMysqlFromConsul(config config.Config,path ...string) *MysqlConfig{
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
