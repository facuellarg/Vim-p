package config

import (
	"freddy.facuellarg.com/domain/connection"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	configPath      = "./config"
	configFileName  = "config"
	configExtension = "toml"
	//defaultValuesForConnection
	databaseFields = connection.DataBaseConnection{
		User: "root",
		Pass: "root1234",
		Host: "localhost",
		Port: 3306,
		DB:   "vim",
	}
)

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

//LoadConf read the configuration in the current directory
func LoadConf() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configExtension)
	return viper.ReadInConfig()
}

//GetDatabaseConf return the fields for database connection
func GetDatabaseConf() (connection.DataBaseConnection, error) {
	if err := mapstructure.Decode(
		viper.GetStringMap("database"),
		&databaseFields,
	); err != nil {
		return connection.DataBaseConnection{}, err
	}
	return databaseFields, nil
}
