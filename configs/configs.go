package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// //go:embed .env
// var envFile embed.FS

// func Load() {
// 	content, err := envFile.ReadFile(".env")
// 	if err != nil {
// 		panic(fmt.Errorf("cannot read embedded env: %w", err))
// 	}

// 	viper.SetConfigType("env")
// 	err = viper.ReadConfig(strings.NewReader(string(content)))
// 	if err != nil {
// 		panic(fmt.Errorf("cannot load env to viper: %w", err))
// 	}

// 	viper.AutomaticEnv()
// }

// func GetEnv(name string) string {
// 	nameVarEnv := strings.ToUpper(name)
// 	nameVarEnv = strings.ReplaceAll(nameVarEnv, ".", "_")
// 	return viper.GetString(nameVarEnv)
// }

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
