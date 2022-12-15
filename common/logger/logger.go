/**
 * @Author: YMBoom
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2022/12/15 14:23
 */
package logger

type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)
