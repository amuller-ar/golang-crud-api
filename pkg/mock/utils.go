package mock

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

const (
	ResourceFolder = "%s/resources/%s/%s"

	ResourceTypeRequest = "requests"
)

func GetBasePath() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	return filepath.Dir(b)
}

func GetResource(name string, rType string) ([]byte, error) {
	basePath := GetBasePath()

	bytes, err := ioutil.ReadFile(fmt.Sprintf(ResourceFolder, basePath, rType, name))

	return bytes, err
}
