package cache

import (
	"bufio"
	"os"
)

type (
	cache struct {
	}
)

func New() Cache {
	return &cache{}
}

func (c *cache) Get(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	resp := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resp = append(resp, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cache) Set(path string, value string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(value)
	if err != nil {
		return err
	}

	return nil
}
