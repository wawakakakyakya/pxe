package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	process string
	logger  *zerolog.Logger
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func NewLogger(process string) *Logger {
	// ev := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().
	// 	Timestamp().
	// 	Str("server", server).
	// 	Logger()

	// logger := zerolog.Nop()
	// logpath := cache.Path("elephantsql", fmt.Sprintf("access.%s.log", time.Now().Local().Format("20060102"))) // logpath = ~/.cache/elephantsql/access.YYYYMMDD.log
	var logger zerolog.Logger
	file, err := os.OpenFile("pxe.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("log file open error")
		fmt.Println(err.Error())
	}
	// if err != nil {
	// 	logger = zerolog.New(os.Stdout)
	// } else {
	// 	logger = zerolog.New(io.MultiWriter(
	// 		file,
	// 		zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
	// 	))
	// }
	writer := io.MultiWriter(file, zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false})
	logger = zerolog.New(writer).Level(zerolog.InfoLevel).With().
		Timestamp().
		Str("process", process).
		Logger()
	// logger = logger.Level(zerolog.DebugLevel).With().Timestamp().Logger()
	// if err != nil {
	// 	logger.Error().Interface("error", errs.Wrap(err, errs.WithContext("logpath", logpath))).Str("logpath", logpath).Msg("error in opening logfile")
	// }
	// return logger

	return &Logger{process: process, logger: &logger}
}
