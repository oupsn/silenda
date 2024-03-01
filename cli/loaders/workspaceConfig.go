package loaders

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/wuhoops/silenda/cli/utils"
)

type WorkspaceConfig struct {
	WorkSpaceId string `json:"WorkSpaceId" mapstructure:"WorkSpaceId"`
}

var Wc *WorkspaceConfig

func WorkspaceConfigInit() {
	viper.SetConfigFile(utils.CONFIG_FILE_NAME)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&Wc); err != nil {
		panic(fmt.Errorf("[CONFIG] fatal loading configuration: %w, maybe due to invalid configuration format", err))
	}
}
