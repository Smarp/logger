package logger

import "go.uber.org/zap/zapcore"

/*
DEFAULT 	(0) The log entry has no assigned severity level.
DEBUG 	(100) Debug or trace information.
INFO 	(200) Routine information, such as ongoing status or performance.
NOTICE 	(300) Normal but significant events, such as start up, shut down, or a configuration change.
WARNING 	(400) Warning events might cause problems.
ERROR 	(500) Error events are likely to cause problems.
CRITICAL 	(600) Critical events cause more severe problems or outages.
ALERT 	(700) A person must take an action immediately.
EMERGENCY 	(800) One or more systems are unusable.
*/

type GoogleSeverity int

func (l GoogleSeverity) String() string {
	switch l {
	case 0:
		return "DEFAULT"
	case 100:
		return "DEBUG"
	case 200:
		return "INFO"
	case 300:
		return "NOTICE"
	case 400:
		return "WARNING"
	case 500:
		return "ERROR"
	case 600:
		return "CRITICAL"
	case 700:
		return "ALERT"
	case 800:
		return "EMERGENCY"
	}
	return ""
}

func LevelToGoogleSeverity(level zapcore.Level) GoogleSeverity {
	switch level {
	case zapcore.DebugLevel:
		return GoogleSeverityDebug
	case zapcore.InfoLevel:
		return GoogleSeverityInfo
	case zapcore.WarnLevel:
		return GoogleSeverityWarning
	case zapcore.ErrorLevel:
		return GoogleSeverityError
	case zapcore.DPanicLevel:
		return GoogleSeverityCritical
	case zapcore.PanicLevel:
		return GoogleSeverityCritical
	case zapcore.FatalLevel:
		return GoogleSeverityEmergency
	}
	return GoogleSeverityDefault
}

const (
	GoogleSeverityDefault   GoogleSeverity = 0
	GoogleSeverityDebug     GoogleSeverity = 100
	GoogleSeverityInfo      GoogleSeverity = 200
	GoogleSeverityNotice    GoogleSeverity = 300
	GoogleSeverityWarning   GoogleSeverity = 400
	GoogleSeverityError     GoogleSeverity = 500
	GoogleSeverityCritical  GoogleSeverity = 600
	GoogleSeverityAlert     GoogleSeverity = 700
	GoogleSeverityEmergency GoogleSeverity = 800
)

func GoogleLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(LevelToGoogleSeverity(l).String())
}
