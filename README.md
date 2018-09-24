# zapgorm
zap logger for gorm

## Usage

```go
import "moul.io/zapgorm"

db, _ := gorm.Open("sqlite3", "./db.sqlite")
db.SetLogger(zapgorm.New(zap.L().Named("gorm")))
db.LogMode(true)
```

## Examples

* https://github.com/moul/depviz
