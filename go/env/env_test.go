package env_test

import (
  "os"
  "reflect"
  "runtime"
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"

  // the package we're testing
  "github.com/Liquid-Labs/env/go/env"
)

const testKey string = `BLAH`

type BasicEnvSuite struct { suite.Suite }
func (s *BasicEnvSuite) TearDownTest() {
  os.Unsetenv(testKey)
}
func TestBasicEnvSuite(t *testing.T) {
  suite.Run(t, new(BasicEnvSuite))
}

func (s *BasicEnvSuite) TestGet() {
  assert.Empty(s.T(), env.Get(testKey))
  os.Setenv(testKey, `blah`)
  assert.Equal(s.T(), `blah`, env.Get(testKey))
}

func (s *BasicEnvSuite) TestMustGet() {
  assert.Panics(s.T(), func() { env.MustGet(testKey) })
  os.Setenv(testKey, `blah`)
  assert.NotPanics(s.T(), func() { env.MustGet(testKey) })
  assert.Equal(s.T(), `blah`, env.MustGet(testKey))
}

func (s *BasicEnvSuite) TestSet() {
  env.Set(testKey, `blah`)
  assert.Equal(s.T(), `blah`, os.Getenv(testKey))
}

func (s *BasicEnvSuite) TestUnset() {
  os.Setenv(testKey, `blah`)
  env.Unset(testKey)
  assert.Empty(s.T(), os.Getenv(testKey))
}

// EnvTypeSuite handles setup and teardown for the 'EnvType' series of tests.
type EnvTypeSuite struct {
  suite.Suite
  OrigValue string
}
func (s *EnvTypeSuite) SetupTest() {
  s.OrigValue = os.Getenv(env.DefaultEnvTypeKey)
  // We only capture the default, since that's all we expect to be set coming
  // into the test. But we clear everything in case the tests themselves set
  // something.
  for _, key := range env.ValidEnvTypeKeys {
    os.Unsetenv(key)
  }
}
func (s *EnvTypeSuite) TearDownTest() {
  os.Setenv(env.DefaultEnvTypeKey, s.OrigValue)
}
func TestEnvTypeSuite(t *testing.T) {
  suite.Run(t, new(EnvTypeSuite))
}

func (s *EnvTypeSuite) TestGetType() {
  assert.Empty(s.T(), env.GetType())
  os.Setenv(env.DefaultEnvTypeKey, `dev`)
  assert.Equal(s.T(), `dev`, env.GetType())
  os.Setenv(env.ValidEnvTypeKeys[1], `test`)
  assert.Equal(s.T(), `dev`, env.GetType(), "GetType appears to be disrepecting precedence.")
  os.Unsetenv(env.DefaultEnvTypeKey)
  assert.Equal(s.T(), `test`, env.GetType())
}

func (s *EnvTypeSuite) TestMustGetType() {
  assert.Panics(s.T(), func() { env.MustGetType() })
  os.Setenv(env.DefaultEnvTypeKey, `dev`)
  assert.NotPanics(s.T(), func() { env.MustGetType() })
  assert.Equal(s.T(), `dev`, env.MustGetType())
}

func (s *EnvTypeSuite) TestNoTypeSpecified() {
  assert.True(s.T(), env.NoTypeSpecified())
}

func (s *EnvTypeSuite) TestIsEnvTypes() {
  funcs := []func() bool{env.IsDev, env.IsTest, env.IsProduction}
  for i, envType := range []string{`dev`, `test`, `production`} {
    os.Setenv(env.DefaultEnvTypeKey, envType)
    s.T().Run(`Test` + runtime.FuncForPC(reflect.ValueOf(funcs[i]).Pointer()).Name(),
              func (t *testing.T) { assert.True(t, funcs[i]()) })
  }
}

func (s *EnvTypeSuite) TestIsStandardType() {
  for _, envType := range []string{`dev`, `test`, `production`} {
    os.Setenv(env.DefaultEnvTypeKey, envType)
    assert.True(s.T(), env.IsStandardType())
  }
  os.Setenv(env.DefaultEnvTypeKey, `blah`)
  assert.False(s.T(), env.IsStandardType())
}

func (s *EnvTypeSuite) TestRequireRecognizedType() {
  os.Setenv(env.DefaultEnvTypeKey, `blah`)
  assert.Panics(s.T(), func() { env.RequireRecognizedType() })
  os.Setenv(env.DefaultEnvTypeKey, `dev`)
  assert.NotPanics(s.T(), func() { env.RequireRecognizedType() })
}

func (s *EnvTypeSuite) TestGetTypeSource() {
  assert.Equal(s.T(), ``, env.GetTypeSource())
  for i := len(env.ValidEnvTypeKeys)-1; i >= 0; i-- {
    os.Setenv(env.ValidEnvTypeKeys[i], `dev`)
    assert.Equal(s.T(), env.ValidEnvTypeKeys[i], env.GetTypeSource())
  }
}
