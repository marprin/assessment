package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadYamlYmlConfig(src string, dest interface{}) error {
	file, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, dest)
	if err != nil {
		return err
	}

	return nil
}

func ReadYMLConfig(path, filename string, dest interface{}) error {
	src := fmt.Sprintf("%s/%s.yml", path, filename)

	return ReadYamlYmlConfig(src, dest)

}
