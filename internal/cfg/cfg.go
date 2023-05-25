package cfg

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	GOOSEDRIVER       = ""
	GOOSEDBSTRING     = ""
	GOOSEMIGRATIONDIR = ""
	GOOSENOCOLOR      = ""
)

var (
	DefaultMigrationDir = "."
)

func init() {
	_, err := os.Stat("conf/migrations.yaml")
	if err != nil {
		log.Printf(" 读取conf/migrations.yaml配置文件不存在")
		return
	}
	v := viper.New()
	v.SetConfigFile("conf/migrations.yaml")
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		log.Printf(" 读取conf/migrations.yaml配置文件错误")
		return
	}
	GOOSEDRIVER = v.GetString("migrations.driver")
	GOOSEDBSTRING = v.GetString("migrations.url")
	GOOSEMIGRATIONDIR = v.GetString("migrations.dir")
	GOOSENOCOLOR = v.GetString("migrations.no_color")

}

type EnvVar struct {
	Name  string
	Value string
}

func List() []EnvVar {
	return []EnvVar{
		{Name: "GOOSE_DRIVER", Value: GOOSEDRIVER},
		{Name: "GOOSE_DBSTRING", Value: GOOSEDBSTRING},
		{Name: "GOOSE_MIGRATION_DIR", Value: GOOSEMIGRATIONDIR},
		{Name: "NO_COLOR", Value: GOOSENOCOLOR},
	}
}
