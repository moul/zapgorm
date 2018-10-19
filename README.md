# zapgorm

[![GuardRails badge](https://badges.production.guardrails.io/moul/zapgorm.svg)](https://www.guardrails.io)

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
