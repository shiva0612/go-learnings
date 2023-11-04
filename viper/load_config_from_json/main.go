package main

import (
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

type Person struct {
	Name string `mapstructure:"name"`
}

func initLogger() {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	l, _ := config.Build()
	log = l.Sugar()
}
func main() {

	initLogger()

	p := new(Person)

	LoadConfig("./person.json", p)

	log.Infof("Person = %+v", p)
	log.Infof("done...")
}

func LoadConfig(path string, cls interface{}) error {
	viper.AddConfigPath(filepath.Dir(path))     //dir path
	viper.SetConfigName(filepath.Base(path))    //filename
	viper.SetConfigType(filepath.Ext(path)[1:]) //.json

	err := viper.ReadInConfig()
	if err != nil {
		log.Infof("error while reading config file: ", path, err.Error())
		return err
	}

	err = viper.Unmarshal(cls)
	if err != nil {
		log.Infof("error while unmarshalling: ", err.Error())
		return err
	}

	return nil

}
