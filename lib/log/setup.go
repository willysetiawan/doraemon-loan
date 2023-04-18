package log

import (
	"encoding/json"
	"fmt"
	"os"
	"process-loan/lib/constant"
	"process-loan/lib/env"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	year  int
	month time.Month
	day   int

	log   = logrus.New()
	logDb = logrus.New()
)

func init() {
	year, month, day = time.Now().Local().Date()
	logPrettyPrintFile, _ := strconv.ParseBool(env.String("Logging.logFile.PrettyPrint", "false"))
	logPrettyPrint, _ := strconv.ParseBool(env.String("Logging.logPrint.PrettyPrint", "false"))
	logPrettyPrintDB, _ := strconv.ParseBool(env.String("Logging.logErrorInsertDB.PrettyPrint", "false"))

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: env.String("Logging.logFile.TimeFormat", "01-02-2006 15:04:05"),
		PrettyPrint:     logPrettyPrintFile,
	})

	log.SetOutput(os.Stderr)
	logDb.SetOutput(os.Stderr)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: env.String("Logging.logPrint.TimeFormat", "01-02-2006 15:04:05"),
		PrettyPrint:     logPrettyPrint,
	})

	logDb.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: env.String("Logging.logErrorInsertDB.TimeFormat", "01-02-2006 15:04:05"),
		PrettyPrint:     logPrettyPrintDB,
	})
}

func LogFmtTemp(v ...string) {
	fmt.Printf("\n >>>>>>>>>  %4s  <<<<<<<<< \n", v)
}

func LogPrintErrorInsertDB(data interface{}, info string) {
	ActiveLogFile, _ := strconv.ParseBool(env.String("Logging.logFile.Active", "false"))
	if !ActiveLogFile {
		return
	}
	formatDate := fmt.Sprintf("%s_%d_%s_%d.log", env.String("Logging.logErrorInsertDB.FileName", "DB_log"), day, month, year)
	file, err := os.OpenFile(formatDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logDb.Out = file
	} else {
		logrus.Info("Failed create file log print db")
	}
	dataByte, _ := json.Marshal(data)
	item := logrus.Fields{
		"data": string(dataByte),
	}
	logDb.WithFields(item).Warning(info)
}

func LogPrintFile(trace_id string, pkgName string, actName string, level int, data interface{}) {
	formatDate := fmt.Sprintf("%s_%d_%s_%d.log", env.String("Logging.logFile.FileName", "log"), day, month, year)
	file, err := os.OpenFile(formatDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		logrus.Info("Failed create file log print file")
	}
	item := logrus.Fields{
		"service_name":    env.String("MainSetup.ServiceName", ""),
		"service_type":    env.String("MainSetup.ServiceType", ""),
		"version_release": env.String("Version.ReleaseVersion", ""),
		"version_type":    env.String("Version.VersionType", ""),
		"data":            data,
		"action":          actName,
		"trace_id":        trace_id,
	}
	if level == constant.LEVEL_LOG_INFO {
		log.WithFields(item).Infof(pkgName)
	}
	if level == constant.LEVEL_LOG_WARNING {
		log.WithFields(item).Warning(pkgName)
	}
	if level == constant.LEVEL_LOG_ERROR {
		log.WithFields(item).Error(pkgName)
	}
	if level == constant.LEVEL_LOG_FATAL {
		log.WithFields(item).Fatal(pkgName)
	}
}

func LogPrint(trace_id, pkgName, actName string, level int, data interface{}) {
	fmt.Print("\n")
	itemPrint := logrus.Fields{
		"service_name":    env.String("MainSetup.ServiceName", ""),
		"service_type":    env.String("MainSetup.ServiceType", ""),
		"version_release": env.String("Version.ReleaseVersion", ""),
		"version_type":    env.String("Version.VersionType", ""),
		"data":            data,
		"action":          actName,
		"trace_id":        trace_id,
	}

	if level == constant.LEVEL_LOG_INFO {
		logrus.WithFields(itemPrint).Info(pkgName)
	}
	if level == constant.LEVEL_LOG_WARNING {
		logrus.WithFields(itemPrint).Warning(pkgName)
	}
	if level == constant.LEVEL_LOG_ERROR {
		logrus.WithFields(itemPrint).Error(pkgName)
	}
	if level == constant.LEVEL_LOG_FATAL {
		logrus.WithFields(itemPrint).Fatal(pkgName)
	}
	fmt.Print("\n")
}

func LogData(trace_id, pkgName, actName string, level int, data interface{}) {
	ActiveLogFile, _ := strconv.ParseBool(env.String("Logging.logFile.Active", "false"))
	if ActiveLogFile {
		LogPrintFile(trace_id, pkgName, actName, level, data)
	}
	ActiveLogPrint, _ := strconv.ParseBool(env.String("Logging.logPrint.Active", "true"))
	if ActiveLogPrint {
		LogPrint(trace_id, pkgName, actName, level, data)
	}
}

func LogSendRequest(trace_id, pkgName string, url string, header interface{}, method string, req, res interface{}) {
	item := logrus.Fields{
		"host":            url,
		"host_method":     method,
		"host_header":     header,
		"host_request":    req,
		"host_response":   res,
		"service_name":    env.String("MainSetup.ServiceName", ""),
		"service_type":    env.String("MainSetup.ServiceType", ""),
		"version_release": env.String("Version.ReleaseVersion", ""),
		"version_type":    env.String("Version.VersionType", ""),
		"trace_id":        trace_id,
	}
	ActiveLogFile, _ := strconv.ParseBool(env.String("Logging.logFile.Active", "false"))
	if ActiveLogFile {
		formatDate := fmt.Sprintf("%s_%d_%s_%d.log", env.String("Logging.logFile.FileName", "log"), day, month, year)
		file, err := os.OpenFile(formatDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			logrus.Info("Failed create file log print file")
		}
		log.WithFields(item).Info(pkgName)
	}
	ActiveLogPrint, _ := strconv.ParseBool(env.String("Logging.logPrint.Active", "true"))
	if ActiveLogPrint {
		logrus.WithFields(item).Info(pkgName)
	}
}
