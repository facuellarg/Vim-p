package config

import "github.com/spf13/viper"

var configPath = "./config"
var configFileName = "config"
var configExtension = "toml"

//SetExtension set the extension of file, by default toml
func SetExtension(extension string) {
	configExtension = extension
}

//SetFileName set the name to config file by default config
func SetFileName(fileName string) {
	configFileName = fileName
}

//SetConfigPath set the directory where will be the configuration file
//by default the current directory
func SetConfigPath(path string) {
	configPath = path
}

//ReadConf read the configuration in the current directory
func ReadConf() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configExtension)
	return viper.ReadInConfig()

}
