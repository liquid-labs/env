package sqldb

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
  "github.com/Liquid-Labs/go-api/osext"
)

var DB *sql.DB

type dbSetupFunc func(db *sql.DB)
var setupFuncs = make([]dbSetupFunc, 0, 8)
func RegisterSetup(newSetupFuncs ...dbSetupFunc) {
  setupFuncs = append(setupFuncs, newSetupFuncs...)
}

func InitDb() {
  var (
    connectionName = osext.MustGetenv("CLOUDSQL_CONNECTION_NAME")
    connectionProt = osext.MustGetenv("CLOUDSQL_CONNECTION_PROT")
    user           = osext.MustGetenv("CLOUDSQL_USER")
    password       = osext.MustGetenv("CLOUDSQL_PASSWORD") // NOTE: password may NOT be empty
    dbName         = osext.MustGetenv("CLOUDSQL_DB")
  )
  var dsn string = fmt.Sprintf("%s:%s@%s(%s)/%s", user, password, connectionProt, connectionName, dbName)

  var err error
  DB, err = sql.Open("mysql", dsn)
  if err != nil {
    log.Panicf("Could not open db: %v", err)
  }

  for _, setupFunc := range setupFuncs {
    setupFunc(DB)
  }
}
