package env

import (
  "log"
  "os"
  "strings"
)

func MustGet(k string) string {
  v := os.Getenv(k)
  if v == `` {
    log.Panicf("'%s' environment variable not set.", k)
  }
  return v
}

func Get(k string) string { return os.Getenv(k) }

var environmentKeys = []string{`NODE_ENV`, `CURR_ENV_PURPOSE`}

func MustGetType() string {
  envType := GetType()
  if envType == `` {
    log.Panicf("Could not determine enviroment type. Set one of %s", strings.Join(environmentKeys, ", "))
  }
  return envType
}

func GetType() string {
  for _, key := range environmentKeys {
    envType := Get(key)
    if envType != `` {
      return envType
    }
  }
  return ``
}

func IsDev() bool {
  return GetType() == `dev`
}

func IsTest() bool {
  return GetType() == `test`
}

func IsProduction() bool {
  return GetType() == `produciton`
}

func IsStandardType() bool {
  return IsDev() || IsTest() || IsProduction()
}

func RequireRecognizedType() {
  if !IsStandardType() {
    log.Panicf(`'NODE_ENV' value of '%s' is not a recognized type.`, GetType())
  }
}
