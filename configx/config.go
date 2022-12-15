/**
 * @Author: YMBoom
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/12/14 9:49
 */
package configx

import "github.com/ANMP0042/gormx/common/logger"

type Config struct {
	User                      string
	Passwd                    string
	Host                      string
	Port                      int
	Dbname                    string
	Dsn                       string
	LogLevel                  logger.LogLevel
	DisableDatetimePrecision  bool
	DontSupportRenameIndex    bool
	DontSupportRenameColumn   bool
	SkipInitializeWithVersion bool
	SingularTable             bool
}
