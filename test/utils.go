package test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/shoshtari/divar-notifier/internal/configs"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetTestDir() (string, error) {
	return "/home/morteza/divar-notifier/test", nil
	ans, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var e bool
	for ans != "" && !e {
		fmt.Println(ans)
		e, err := exists(ans + "/go.mod")
		if err != nil {
			return "", errors.Wrap(err, "can't check exist")
		}
		if !e {
			ans = filepath.Dir(ans)
		}
	}

	if ans == "" {
		return "", errors.New("not found")
	}
	return ans + "/test", nil

}

func GetTestConfig() (configs.JarchiConfig, error) {
	path, err := GetTestDir()
	if err != nil {
		return configs.JarchiConfig{}, err
	}
	return configs.GetConfig(path)
}
