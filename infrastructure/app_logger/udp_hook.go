package  app_logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

type UdpHook struct {
	Conn         net.Conn
	UrlUdpLog    string
	LayoutUdpLog string
}


func NewUdpHook( urlUdp string, layoutUdpLog string) (*UdpHook, error) {
	conn, err := net.Dial("udp", urlUdp)
	if err != nil {
		fmt.Printf("Cannot connect to Logger Server: %s\n error: %v", urlUdp, err)
	}
	return &UdpHook{Conn:conn,LayoutUdpLog:layoutUdpLog,UrlUdpLog:urlUdp}, err
}
func (hook *UdpHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
    stringLog :=fmt.Sprintf("%s %s [[%s", hook.LayoutUdpLog, entry.Level.String(), line)
	switch entry.Level {
	case logrus.PanicLevel:
		_, err := fmt.Fprint(hook.Conn, stringLog)
		return  err
	case logrus.FatalLevel:
		_, err := fmt.Fprint(hook.Conn, stringLog)
		return  err
	case logrus.ErrorLevel:
		_, err := fmt.Fprint(hook.Conn, stringLog)
		return  err
	case logrus.WarnLevel:
		_, err := fmt.Fprint(hook.Conn, stringLog)
		return  err
	case logrus.InfoLevel:
		_, err := fmt.Fprint(hook.Conn, stringLog)
		return  err
	case logrus.DebugLevel, logrus.TraceLevel:
		_, err := fmt.Fprint(hook.Conn, stringLog)
		return  err
	default:
		return nil
	}
}

func (hook *UdpHook) Levels() []logrus.Level {
	return logrus.AllLevels
}