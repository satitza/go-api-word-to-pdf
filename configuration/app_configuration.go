package configuration

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfiguration struct {
	Port          int `mapstructure:"port"`
	ConvertConfig struct {
		WordTempPath string `mapstructure:"word-temp-path"`
		PdfTempPath  string `mapstructure:"pdf-temp-path"`
	} `mapstructure:"convert-config"`
	LibreofficeConfig struct {
		WindowsPath string `mapstructure:"windows-path"`
		LinuxPath   string `mapstructure:"linux-path"`
	} `mapstructure:"libreoffice-config"`
}

var instance *AppConfiguration

func GetConfig() (*AppConfiguration, error) {

	if instance == nil {

		viper.SetConfigName("config")        // name of the config file (without extension)
		viper.SetConfigType("yaml")          // the type of config file (yaml, json, etc.)
		viper.AddConfigPath("configuration") // path to look for the config file (current directory)

		// Read the config file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		// Unmarshal the config into the struct
		if err := viper.Unmarshal(&instance); err != nil {
			log.Fatalf("Unable to decode into struct: %v", err)
		}
	}

	return instance, nil

}
