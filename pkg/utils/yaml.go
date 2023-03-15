package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func YamlToStruct(file string, out interface{}) (err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, out)
	return
}
