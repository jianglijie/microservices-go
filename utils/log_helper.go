package utils

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.WarnLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
	//Log.AddHook(filename.NewHook())

	//rootDir := common.GetCurrentPath()
	//configDir := path.Join(rootDir, "log")
	//isExist, err := common.IsPathExists(configDir)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(isExist)
	//if !isExist {
	//	err := os.Mkdir(configDir, 0755)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	ConfigLocalFilesystemLogger("log", "server.log", time.Hour*24*7, time.Hour*24)
}

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M%S",
		//rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
	Log.AddHook(lfHook)
}

func LogError(content string, fields map[string]interface{}) {
	Log.WithFields(fields).Error(content)
}

func LogWarn(content string, fields map[string]interface{}) {
	Log.WithFields(fields).Warn(content)
}
