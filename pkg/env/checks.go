package env

import (
	"fmt"
	"reflect"
)

func checkEnvErrors(errSlice []error) error {
	for _, err := range errSlice {
		if err != nil {
			return fmt.Errorf("environment misconfigured; error %s", err.Error())
		} else {
			continue
		}
	}
	return nil
}

func checkEnvValues(envs ...interface{}) error {
	for _, env := range envs {
		valuesSlice := reflect.ValueOf(env)
		for i := 0; i < valuesSlice.NumField(); i++ {
			if valuesSlice.Field(i).Type().Kind() == reflect.String && valuesSlice.Field(i).Interface() == "" {
				return fmt.Errorf("environment misconfigured; missing field %s", valuesSlice.Type().Field(i).Name)
			}
		}
	}
	return nil
}
