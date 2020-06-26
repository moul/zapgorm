# zapgorm
zap logger for gorm 1 and gorm 2

## Usage with gorm 2

```go
import "moul.io/zapgorm/zapgorm2"

logger := zapgorm2.New(zap.L())
logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
db, err = gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{Logger: logger})
```

## Usage with gorm 1

```go
import "moul.io/zapgorm"

db, _ := gorm.Open("sqlite3", "./db.sqlite")
db.SetLogger(zapgorm.New(zap.L().Named("gorm")))
db.LogMode(true)
```

## Examples in the wild

* https://github.com/moul/depviz
* https://github.com/pathwar/pathwar
* https://akatsuki-chan/go-sample
* _add your repo here_
