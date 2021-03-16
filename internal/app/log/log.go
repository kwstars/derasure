package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

var ProviderSet = wire.NewSet(New, NewLogOptions)

// Options is log configuration struct
type LogOptions struct {
	Level      string `mapstructure:"level"`
	StdoutFile struct {
		Open       bool   `mapstructure:"open"`
		Path       string `mapstructure:"path"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxAge     int    `mapstructure:"max_age"`
	} `mapstructure:"stdout_file"`
	StdoutConsole bool `mapstructure:"stdout_console"`
}

func NewLogOptions(v *viper.Viper) (o *LogOptions, err error) {
	o = &LogOptions{}
	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}

	return o, err
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// New for init zap log library
func New(o *LogOptions) (logger *zap.Logger, cf func(), err error) {
	level := zap.NewAtomicLevel()
	if err = level.UnmarshalText([]byte(o.Level)); err != nil {
		return
	},

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.StdoutFile.Path,
		MaxSize:    o.StdoutFile.MaxSize, // megabytes
		MaxBackups: o.StdoutFile.MaxBackups,
		MaxAge:     o.StdoutFile.MaxAge, // days
	})

	cw := zapcore.Lock(os.Stdout)

	cores := make([]zapcore.Core, 0, 2)

	// 输出到文件
	if o.StdoutFile.Open {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = SyslogTimeEncoder
		//encoderConfig.FunctionKey = "function"
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fw, level))
	}

	// 输出到Console
	if o.StdoutConsole {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), cw, level))
	}

	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	zap.ReplaceGlobals(logger)

	cf = func() {
		if err := logger.Sync(); err != nil {
			log.Println(err)
		}
	}

	return
}
