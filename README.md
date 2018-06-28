# logrus-levelpad-formatter

Logrus (https://github.com/sirupsen/logrus) formatter with messages like

```
[INFO]: 2006-01-02T15:04:05Z07:00 - Log message
```

Configure

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