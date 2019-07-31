package env

import (
  "log"
  "os"
  "strings"
)

func Get(k string) string { return os.Getenv(k) }

func MustGet(k string) string {
  v := os.Getenv(k)
  if v == `` {
    log.Panicf("'%s' environment variable not set.", k)
  }
  return v
}

func Set(key, value string) {
  os.Setenv(key, value)
}

func Unset(key string) {
  os.Unsetenv(key)
}

const DefaultEnvTypeKey = `CURR_ENV_PURPOSE`
var ValidEnvTypeKeys = []string{DefaultEnvTypeKey, `NODE_ENV`}

func GetType() string {
  for _, key := range ValidEnvTypeKeys {
    envType := Get(key)
    if envType != `` {
      return envType
    }
  }
  return ``
}

func MustGetType() string {
  envType := GetType()
  if envType == "" {
    log.Panicf("Could not determine enviroment type. Set one of %s", strings.Join(ValidEnvTypeKeys, ", "))
  }
  return envType
}

func NoTypeSpecified() bool {
  return GetType() == ``
}

func IsDev() bool {
  return GetType() == `dev`
}

func IsTest() bool {
  return GetType() == `test`
}

func IsProduction() bool {
  return GetType() == `production`
}

func IsStandardType() bool {
  return IsDev() || IsTest() || IsProduction()
}

func RequireRecognizedType() {
  if !IsStandardType() {
    log.Panicf(`Environment type value of '%s' is not a recognized type. (source: %s)`, GetType(), GetTypeSource())
  }
}

func GetTypeSource() string {
  for _, key := range ValidEnvTypeKeys {
    envType := Get(key)
    if envType != `` {
      return key
    }
  }
  return ``
}
