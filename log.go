package nekolog

import (
	"fmt"
	"github.com/avellea/nekolog/internal/utils"
	"os"
	"sync"
)

type Logger struct {
	mu sync.Mutex
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) log(level utils.Level, format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := fmt.Sprintf(format, args...)
	color := utils.LevelColor(level)
	fmt.Fprintf(os.Stdout, "[%s] %s[%s]%s %s\n",
		utils.GetTime(),
		color,
		level,
		utils.Reset,
		message)
}

// TODO: Consolidate these into just their formatted versions? Unsure.

func (l *Logger) Trace(msg string) { l.log(utils.TRACE, msg) }
func (l *Logger) Tracef(format string, args ...any) { l.log(utils.TRACE, fmt.Sprintf(format, args...)) }

func (l *Logger) Debug(msg string) { l.log(utils.DEBUG, msg) }
func (l *Logger) Debugf(format string, args ...any) { l.log(utils.DEBUG, fmt.Sprintf(format, args...)) }

func (l *Logger) Info(msg string) { l.log(utils.INFO, msg) }
func (l *Logger) Infof(format string, args ...any) { l.log(utils.INFO, fmt.Sprintf(format, args...)) }

func (l *Logger) Warn(msg string) { l.log(utils.WARN, msg) }
func (l *Logger) Warnf(format string, args ...any) { l.log(utils.WARN, fmt.Sprintf(format, args...)) }

func (l *Logger) Error(msg string) { l.log(utils.ERROR, msg) }
func (l *Logger) Errorf(format string, args ...any) { l.log(utils.ERROR, fmt.Sprintf(format, args...)) }

func (l *Logger) Fatal(msg string) { l.log(utils.FATAL, msg) }
func (l *Logger) Fatalf(format string, args ...any) { l.log(utils.FATAL, fmt.Sprintf(format, args...)) }
