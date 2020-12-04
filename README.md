# go-mpnj

### migrate
* rel migrate -adapter="github.com/go-rel/rel/adapter/sqlite3" -driver="github.com/mattn/go-sqlite3" -dsn="mpnj.db?_foreign_keys=1&_loc=Local"

### rollback
* rel rollback -adapter="github.com/go-rel/rel/adapter/sqlite3" -driver="github.com/mattn/go-sqlite3" -dsn="mpnj.db?_foreign_keys=1&_loc=Local"