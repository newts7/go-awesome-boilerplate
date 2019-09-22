package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"goBoilerPlate/config"
	"io/ioutil"
	"os"
	"time"
)
var LogH = logrus.New()

type DefaultFieldHook struct {
	GetValue func() string
}
type fieldKey string
type FieldMap map[fieldKey]string

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	e.Data["service-name"] = "myapp"
	env := os.Getenv("environment")
	switch  env {
	case "development":
		e.Data["env"] = "development"
	case "test":
		e.Data["env"] = "test"
	case "production":
		e.Data["env"] = "production"
	}
	return nil
}
func Init(){
	logLevels := map[string] logrus.Level {
		"PANIC": logrus.PanicLevel,
		"FATAL": logrus.FatalLevel,
		"ERROR": logrus.ErrorLevel,
		"INFO": logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"TRACE": logrus.TraceLevel,
	}
	env := os.Getenv("environment")
	loggingConfiguration := config.GetConfig().Logging

	LogH.SetReportCaller(true)
	LogH.SetLevel(logLevels[loggingConfiguration.Level])
	LogH.AddHook(&DefaultFieldHook{})
	if loggingConfiguration.StdoutLoggingEnable == false{
		LogH.Out = ioutil.Discard
	}
	infoLogPath  := loggingConfiguration.AppName+"-"+env+"-info"
	infoLogWriter, _ := rotatelogs.New(
		infoLogPath+"-%Y-%m-%d"+ ".log",
		rotatelogs.WithLinkName(infoLogPath+".log"),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	errorLogPath := loggingConfiguration.AppName+"-"+env+"-error"
	errorLogWriter, _ := rotatelogs.New(
		errorLogPath+"-%Y-%m-%d"+".log",
		rotatelogs.WithLinkName(errorLogPath+".log"),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	switch  env {
	case  "development":
		LogH.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel: infoLogWriter,
				logrus.ErrorLevel: errorLogWriter,
			},
			&logrus.JSONFormatter{
				FieldMap:logrus.FieldMap{
					logrus.FieldKeyTime: "timestamp",
				},
			},
		))
	case "test":
		LogH.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel: infoLogWriter,
				logrus.ErrorLevel: errorLogWriter,
			},
			&logrus.JSONFormatter{
				FieldMap:logrus.FieldMap{
					logrus.FieldKeyTime: "timestamp",
				},
			},
		))
	case "production":
		LogH.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel: infoLogWriter,
				logrus.ErrorLevel: errorLogWriter,
			},
			&logrus.JSONFormatter{
				FieldMap:logrus.FieldMap{
					logrus.FieldKeyTime: "timestamp",
				},
			},
		))
	}
}