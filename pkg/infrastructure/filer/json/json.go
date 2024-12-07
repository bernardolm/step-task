package jsonfile

import (
	"context"
	"os"

	"github.com/spf13/viper"
)

type Filer struct {
	root string
}

func (f Filer) Read(_ context.Context, path string) ([]map[string]interface{}, error) {
	path = os.ExpandEnv(path)

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	vp := viper.New()
	vp.AddConfigPath(path)

	result := make([]map[string]interface{}, 0)

	for _, e := range entries {
		v := vp
		v.SetConfigFile(path + e.Name())

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}

		result = append(result, v.AllSettings())
	}

	return result, nil
}

func NewFiler(rootPath string) (*Filer, error) {
	rootPath = os.ExpandEnv(rootPath)

	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		return nil, err
	}

	return &Filer{
		root: rootPath,
	}, nil
}
