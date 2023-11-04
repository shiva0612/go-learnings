package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func init() {
	config := zap.NewProductionConfig()

	config.Level.SetLevel(zapcore.DebugLevel)
	config.OutputPaths = []string{"out.log"}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.Encoding = "console"

	zlog, _ := config.Build(zap.AddStacktrace(zap.ErrorLevel)) //prints stacktrace only for error level and above
	log = zlog.Sugar()
}

func main() {
	log.Infof("infoln")

	name := "shiva"
	log.Infof("infof %s", name)

	log.Errorf("errof %s", name)

	log.Warnf("warning")
}
