package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)
var lg *zap.Logger
func Init() (err error){
	writerSyncer := getLogWriter(viper.GetString("App.LogFileName"),viper.GetInt("App.MaxSize"),viper.GetInt("App.MaxBackups"),viper.GetInt("App.MaxAge"))
	encoder := getEncoder()
	var log = new(zapcore.Level)
	err = log.UnmarshalText([]byte(viper.GetString("debug")))
	if err != nil {
		return err
	}
	var core zapcore.Core
	if(viper.GetString("Server.mode") == "dev"){
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder,writerSyncer,log),
			zapcore.NewCore(consoleEncoder,zapcore.Lock(os.Stdout),zapcore.DebugLevel),
		)
	}else{
		core = zapcore.NewCore(encoder,writerSyncer,log)
	}
	lg = zap.New(core,zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string,maxsize,maxBackup,maxAge int) zapcore.WriteSyncer{
	lumberjackLogger := &lumberjack.Logger{
		Filename:filename,    //日志文件位置
		MaxSize:maxsize,      //在进行切割之前，日志文件的最大大小
		MaxBackups:maxBackup, //保留文件最大个数
		MaxAge:maxAge,        //保留文件最大天数
	}
	return zapcore.AddSync(lumberjackLogger)
}
