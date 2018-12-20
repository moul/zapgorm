# zapgorm
zap logger for gorm

## Usage

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
