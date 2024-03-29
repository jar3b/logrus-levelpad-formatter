package levelpad

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%lvl%]: %time% - %msg%"
	defaultTimestampFormat = time.RFC3339
)

func padRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
	LevelPad  int
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	if f.TimestampFormat == "" {
		f.TimestampFormat = defaultTimestampFormat
	}

	// replace timestamp
	output = strings.Replace(output, "%time%", entry.Time.Format(f.TimestampFormat), 1)
	// replace message text
	output = strings.Replace(output, "%msg%", entry.Message, 1)
	// replace level
	level := strings.ToUpper(entry.Level.String())
	if f.LevelPad != 0 {
		level = padRight(level, " ", f.LevelPad)
	}
	output = strings.Replace(output, "%lvl%", level, 1)

	// replace data
	for k, v := range entry.Data {
		if s, ok := v.(string); ok {
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}
