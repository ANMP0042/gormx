/**
 * @Author: YMBoom
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/12/14 9:49
 */
package configx

type Config struct {
	User                      string
	Passwd                    string
	Host                      string
	Port                      int
	Dbname                    string
	Dsn                       string
	DisableDatetimePrecision  bool
	DontSupportRenameIndex    bool
	DontSupportRenameColumn   bool
	SkipInitializeWithVersion bool
}
