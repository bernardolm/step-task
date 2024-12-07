package cache

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type engine struct {
	instances []Cache
}

func (e *engine) add(instances ...Cache) {
	e.instances = append(e.instances, instances...)
}

func (e *engine) Get(key string, dest any) error {
	for k := range e.instances {
		err := e.instances[k].Get(key, dest)
		if errors.Is(err, ErrMiss) {
			// try next instance
		} else if err != nil {
			return errors.Join(err, fmt.Errorf("cache engine: failed to get"))
		} else {
			return nil
		}
	}

	return nil
}

func (e *engine) GetWithContext(ctx context.Context, key string, dest any) error {
	for k := range e.instances {
		err := e.instances[k].GetWithContext(ctx, key, dest)
		if errors.Is(err, ErrMiss) {
			// try next instance
		} else if err != nil {
			return errors.Join(err, fmt.Errorf("cache engine: failed to get with context"))
		} else {
			return nil
		}
	}

	return nil
}

func (e *engine) Set(key string, value any) error {
	for k := range e.instances {
		err := e.instances[k].Set(key, value)
		if errors.Is(err, ErrMiss) {
			// try next instance
		} else if err != nil {
			return errors.Join(err, fmt.Errorf("cache engine: failed to set"))
		} else {
			return nil
		}
	}

	return nil
}

func (e *engine) SetWithContext(ctx context.Context, key string, value any) error {
	for k := range e.instances {
		err := e.instances[k].SetWithContext(ctx, key, value)
		if err != nil {
			return errors.Join(err, fmt.Errorf("cache engine: failed to set with context"))
		}
	}

	return nil
}

func (e *engine) SetWithContextAndTTL(ctx context.Context, ttl time.Duration, key string, value any) error {
	for k := range e.instances {
		err := e.instances[k].SetWithContextAndTTL(ctx, ttl, key, value)
		if err != nil {
			return errors.Join(err, fmt.Errorf("cache engine: failed to set with context and TTL"))
		}
	}

	return nil
}

func (e *engine) SetWithTTL(key string, value any, ttl time.Duration) error {
	for k := range e.instances {
		err := e.instances[k].SetWithTTL(key, value, ttl)
		if err != nil {
			return errors.Join(err, fmt.Errorf("cache engine: failed to set with TTL"))
		}
	}

	return nil
}

func NewEngine(instances ...Cache) *engine {
	e := engine{}
	e.add(instances...)
	return &e
}
