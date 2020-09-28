package  app_logger

import (
	"github.com/sirupsen/logrus"
	"os"
)
// singleton
var Logger *logrus.Logger

func NewAppLog(urlUdp string, layoutUdpLog string, logLevel string) (*logrus.Logger, error)  {
	log       := logrus.New()
	//hook, err := NewUdpHook(urlUdp, layoutUdpLog)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	//#DEBUG #ERROR #INFO #WARN
	level:=logrus.DebugLevel
	switch logLevel {
		case "DEBUG":
			level = logrus.DebugLevel
			break
		case "ERROR":
			level = logrus.ErrorLevel
			break
		case "INFO":
			level = logrus.InfoLevel
			break
		case "WARN":
			level = logrus.WarnLevel
			break
		default:
			level = logrus.DebugLevel
			break
	}
	log.SetLevel(level)

	//if err == nil {
	//	log.Hooks.Add(hook)
	//}
	log.Infof("RegisterAppLog done")
	Logger = log
	return log, nil
}

