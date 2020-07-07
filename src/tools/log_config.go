package tools

import (
	"github.com/cihub/seelog"
	"os"
)

func InitLog(file string) {
	logger, err := seelog.LoggerFromConfigAsFile(file)
	if err != nil {
		panic("err parsing config log file")
		os.Exit(2)
	}
	seelog.ReplaceLogger(logger)
	//defer seelog.Flush()
}
