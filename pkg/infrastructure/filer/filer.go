package filer

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/step-task/pkg/infrastructure/cache"
)

type filer struct {
	fileExt  string
	rootPath string
}

func (f *filer) path(key string) string {
	key = slug.Make(key)
	return fmt.Sprintf("%s/%s.%s", f.rootPath, key, f.fileExt)
}

func (f *filer) Get(key string, dest any) error {
	ctx := context.Background()
	return f.GetWithContext(ctx, key, dest)
}

func (f *filer) GetWithContext(ctx context.Context, key string, dest any) error {
	filePath := f.path(key)

	data, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.
				WithField("key", key).
				WithField("file_path", filePath).
				Debug("filer: cache miss")
			return cache.ErrMiss
		}
		return err
	}

	log.
		WithField("key", key).
		WithField("file_path", filePath).
		Debug("filer: cache hit")

	value := reflect.ValueOf(dest).Elem()
	value.SetBytes(data)

	return nil
}

func (f *filer) Set(key string, value any) error {
	ctx := context.Background()
	return f.SetWithContextAndTTL(ctx, 0, key, value)
}

func (f *filer) SetWithContext(ctx context.Context, key string, value any) error {
	return f.SetWithContextAndTTL(ctx, 0, key, value)
}

func (f *filer) SetWithContextAndTTL(ctx context.Context, ttl time.Duration, key string, value any) error {
	if ttl > 0 {
		log.Warn("filer: TTL parameter isn't used here")
	}

	filePath := f.path(key)

	data, ok := value.([]byte)
	if !ok {
		return errors.New("filer: can't parse value to byte array")
	}

	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return errors.Join(err, fmt.Errorf("filer: can't write file %s", filePath))
	}

	return nil
}

func (f *filer) SetWithTTL(key string, value any, ttl time.Duration) error {
	ctx := context.Background()
	return f.SetWithContextAndTTL(ctx, ttl, key, value)
}

func New(_ context.Context, rootPath, fileExt string) (*filer, error) {
	return &filer{
		// rootPath: "./tmp",
		fileExt:  fileExt,
		rootPath: rootPath,
	}, nil
}
