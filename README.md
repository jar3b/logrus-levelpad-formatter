# logrus-levelpad-formatter

Logrus (https://github.com/sirupsen/logrus) formatter allows to specify log level with padding and produces messages 
like:

```
[INFO    ][2019-02-23 03:11:36.414] Service starting
```

## parameters

- `TimestampFormat` - datetime format
- `LogFormat` - whole log entry format, you can specify extra params such as `%lvl%` (level: info, warn, error...) and 
`%time%` (timestamp previously formatted using `TimestampFormat` directive) 
- `LevelPad` - padding for `%lvl%` parameter, when `0` then no padding will be applied.

## example configuration

```
import (
	log "github.com/sirupsen/logrus"
	"github.com/jar3b/logrus-levelpad-formatter"
	"os"
)

func InitLog() {
	log.SetFormatter(&levelpad.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		LogFormat:       "[%lvl%][%time%] %msg%\n",
		LevelPad:        8,
	})

	log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)
}
```