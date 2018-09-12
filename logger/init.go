package logger

import (
	"go.uber.org/zap"
)

// default logger use zap sugar
var (
	defaultLogger *zap.SugaredLogger = nil
)

// init
func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defaultLogger = l.Sugar()
}

// Set default sugar is param l
//
// @param l *zap.SugaredLogger
//
func SetSugarLogger(l *zap.SugaredLogger) {
	defaultLogger = l
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//   sugaredLogger.With(
//     "hello", "world",
//     "failure", errors.New("oh no"),
//     Stack(),
//     "count", 42,
//     "user", User{Name: "alice"},
//  )
// is the equivalent of
//   unsugared.With(
//     String("hello", "world"),
//     String("failure", "oh no"),
//     Stack(),
//     Int("count", 42),
//     Object("user", User{Name: "alice"}),
//   )
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func With(args ...interface{}) *zap.SugaredLogger {
	return defaultLogger.With(args...)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Print uses fmt.Sprint to construct and log a message.
func Print(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	defaultLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// wrap args split by space
func wrap(args []interface{}) (a []interface{}) {
	for n, arg := range args {
		if len(args) > 0 && n < len(args)-1 {
			a = append(a, arg, " ")
		} else {
			a = append(a, arg)
		}
	}
	return
}

// Debugln uses fmt.Sprint to construct and log a message.
func Debugln(args ...interface{}) {
	defaultLogger.Debug(wrap(args)...)
}

// Infoln uses fmt.Sprint to construct and log a message.
func Infoln(args ...interface{}) {
	defaultLogger.Info(wrap(args)...)
}

// Println uses fmt.Sprint to construct and log a message.
func Println(args ...interface{}) {
	defaultLogger.Info(wrap(args)...)
}

// Warnln uses fmt.Sprint to construct and log a message.
func Warnln(args ...interface{}) {
	defaultLogger.Warn(wrap(args)...)
}

// Errorln uses fmt.Sprint to construct and log a message.
func Errorln(args ...interface{}) {
	defaultLogger.Error(wrap(args)...)
}

// DPanicln uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicln(args ...interface{}) {
	defaultLogger.DPanic(wrap(args)...)
}

// Panicln uses fmt.Sprint to construct and log a message, then panics.
func Panicln(args ...interface{}) {
	defaultLogger.Panic(wrap(args)...)
}

// Fatalln uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatalln(args ...interface{}) {
	defaultLogger.Fatal(wrap(args)...)
}
