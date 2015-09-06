# llogger
Leveled logger wrapper for native log package

## Why?

The main idea is to add Levels to log operations and ability to show logs from exact level or above.

Leveled logger (`llogger`) library allow you to add necessary logging everywhere in your code with related level. There are several levels: `FATAL`, `ERROR`, `WARN`, `INFO`, `DEBUG`.


## Usage

```
import "gopkg.in/jurka/llogger.v1.0"

log := llogging.NewStdLogger(llogging.INFO, "")

x := 1
log.Info("This is x =", x)

log.Debug("Cheking for x != 10")
if x != 10{
	log.Error("X has wrong value!")
}

```