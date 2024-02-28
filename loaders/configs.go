package loaders

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/wuhoops/silenda/utils")

func Configs() {
	viper.SetConfigFile(utils.CONFIG_FILE_NAME)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}
