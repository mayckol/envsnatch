package envsnatch

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type UnmarshalingErr struct {
	Field  string
	Reason string
}

type EnvSnatch struct {
	envMap           map[string]string
	UnmarshalingErrs []UnmarshalingErr
	path             string
	fileName         string
}

func NewEnvSnatch() (*EnvSnatch, error) {
	return &EnvSnatch{}, nil
}

func (e *EnvSnatch) AddPath(path string) {
	e.path = path
}

func (e *EnvSnatch) AddFileName(fileName string) {
	e.fileName = fileName
}

func (e *EnvSnatch) Unmarshal(config interface{}) (*[]UnmarshalingErr, error) {
	// Attempt to load .env file if path and fileName are provided
	if e.path != "" && e.fileName != "" {
		err := godotenv.Load(e.path + "/" + e.fileName)
		if err != nil {
			fmt.Println(".env not found, using environment variables instead")
		}
	} else {
		// load from the system environment variables
		e.loadEnvVars()
	}

	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanSet() {
			tag := typ.Field(i).Tag
			tagParts := strings.Split(tag.Get("env"), ",")
			envVar := tagParts[0]

			// Check if 'optional' is part of the tag
			optional := false
			if len(tagParts) > 1 && tagParts[1] == "optional" {
				optional = true
			}

			value, exists := e.envMap[envVar]
			if !exists && !optional {
				e.UnmarshalingErrs = append(e.UnmarshalingErrs, UnmarshalingErr{
					Field:  envVar,
					Reason: "required",
				})
				continue
			}

			// Reflective type setting
			if exists {
				switch field.Kind() {
				case reflect.String:
					field.SetString(value)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
						field.SetInt(intValue)
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if uintValue, err := strconv.ParseUint(value, 10, 64); err == nil {
						field.SetUint(uintValue)
					}
				case reflect.Float32, reflect.Float64:
					if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
						field.SetFloat(floatValue)
					}
				case reflect.Bool:
					if boolValue, err := strconv.ParseBool(value); err == nil {
						field.SetBool(boolValue)
					}
				default:
					panic(fmt.Sprintf("Unhandled type: %s", field.Type()))
				}
			}
		}
	}

	if len(e.UnmarshalingErrs) > 0 {
		return &e.UnmarshalingErrs, fmt.Errorf("failed to unmarshal")
	}

	return nil, nil
}

func (e *EnvSnatch) loadEnvVars() {
	e.envMap = make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			e.envMap[pair[0]] = pair[1]
		}
	}
}
