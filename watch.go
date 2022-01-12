//go:build !js
// +build !js

package viper

import (
	"time"

	"github.com/meandrewdev/viper/batcher"
)

type watcher = batcher.Batcher

func newWatcher() (*watcher, error) {
	return batcher.New(250 * time.Millisecond)
}
