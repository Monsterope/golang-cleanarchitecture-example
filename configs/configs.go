package configs

import (
	"strings"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read envoriment")
	}
}

func GetEnv(name string) string {
	nameVarEnv := strings.ToUpper(name)
	nameVarEnv = strings.ReplaceAll(nameVarEnv, ".", "_")
	return viper.GetString(nameVarEnv)
}
