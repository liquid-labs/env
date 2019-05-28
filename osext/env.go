package env

import (
  "log"
  "os"
)

TODO: change this project to 'env'; the 'sqldb' stuff that was here will get factored out to catsql.

func MustGet(k string) string {
  v := os.Getenv(k)
  if v == "" {
    log.Panicf("%s environment variable not set.", k)
  }
  return v
}

func Get(k string) string { return os.Getenv(k) }

const environmentKey = `NODE_ENV`

func MustGetType() {
  return MustGet(environmentKey)
}

func GetType() {
  return Get(environmentKey)
}

func IsDev() bool {
  return GetEnv() == `dev`
}

func IsTest() bool {
  return GetEnv() == `test`
}

func IsProduction() bool {
  return GetEnv() == `produciton`
}

func IsStandardType() bool {
  return IsDev() || IsTest() || IsProduction()
}
