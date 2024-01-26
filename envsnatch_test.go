package envsnatch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnv(key, value string) {
	_ = os.Setenv(key, value)
}

func unsetEnv(key string) {
	_ = os.Unsetenv(key)
}

func TestStringFieldSuccess(t *testing.T) {
	setEnv("STRING_VALUE", "testString")
	defer unsetEnv("STRING_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		StringValue string `env:"STRING_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, "testString", conf.StringValue)
}

func TestStringFieldFailure(t *testing.T) {
	unsetEnv("STRING_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		StringValue string `env:"STRING_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.NotNil(t, err)
	assert.NotNil(t, errs)
	assert.Len(t, *errs, 1)
	assert.Equal(t, "STRING_VALUE", (*errs)[0].Field)
}

func TestStringFieldOptional(t *testing.T) {
	unsetEnv("STRING_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		StringValue string `env:"STRING_VALUE,optional"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, "", conf.StringValue)
}

func TestIntFieldSuccess(t *testing.T) {
	setEnv("INT_VALUE", "123")
	defer unsetEnv("INT_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		IntValue int `env:"INT_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, 123, conf.IntValue)
}

func TestIntFieldFailure(t *testing.T) {
	unsetEnv("INT_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		IntValue int `env:"INT_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.NotNil(t, err)
	assert.NotNil(t, errs)
	assert.Len(t, *errs, 1)
	assert.Equal(t, "INT_VALUE", (*errs)[0].Field)
}

func TestIntFieldOptional(t *testing.T) {
	unsetEnv("INT_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		IntValue int `env:"INT_VALUE,optional"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, 0, conf.IntValue)
}

func TestBoolFieldSuccess(t *testing.T) {
	setEnv("BOOL_VALUE", "true")
	defer unsetEnv("BOOL_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		BoolValue bool `env:"BOOL_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, true, conf.BoolValue)
}

func TestBoolFieldFailure(t *testing.T) {
	unsetEnv("BOOL_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		BoolValue bool `env:"BOOL_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.NotNil(t, err)
	assert.NotNil(t, errs)
	assert.Len(t, *errs, 1)
	assert.Equal(t, "BOOL_VALUE", (*errs)[0].Field)
}

func TestBoolFieldOptional(t *testing.T) {
	unsetEnv("BOOL_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		BoolValue bool `env:"BOOL_VALUE,optional"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, false, conf.BoolValue)
}

func TestFloatFieldSuccess(t *testing.T) {
	setEnv("FLOAT_VALUE", "123.456")
	defer unsetEnv("FLOAT_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		FloatValue float64 `env:"FLOAT_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, 123.456, conf.FloatValue)
}

func TestFloatFieldFailure(t *testing.T) {
	unsetEnv("FLOAT_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		FloatValue float64 `env:"FLOAT_VALUE"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.NotNil(t, err)
	assert.NotNil(t, errs)
	assert.Len(t, *errs, 1)
	assert.Equal(t, "FLOAT_VALUE", (*errs)[0].Field)
}

func TestFloatFieldOptional(t *testing.T) {
	unsetEnv("FLOAT_VALUE")

	e, _ := NewEnvSnatch()
	var conf struct {
		FloatValue float64 `env:"FLOAT_VALUE,optional"`
	}

	errs, err := e.Unmarshal(&conf)
	assert.Nil(t, err)
	assert.Nil(t, errs)
	assert.Equal(t, 0.0, conf.FloatValue)
}
