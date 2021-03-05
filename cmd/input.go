package cmd

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
)

func GetAndConvertInput(fileFlag string) ([]byte, error) {
	var file *os.File
	var err error
	if fileFlag != "" {
		file, err = os.Open(fileFlag)
		if err != nil {
			return nil, err
		}
		defer func() { _ = file.Close()}()
	} else {
		file = os.Stdin
	}

	var object = map[interface{}]interface{} {}
	err = yaml.NewDecoder(file).Decode(&object)
	if err != nil {
		return nil, err
	}

	o := decodeType(object)

	return json.Marshal(o)
}

func decodeType(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = decodeType(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = decodeType(v)
		}
	}
	return i
}
