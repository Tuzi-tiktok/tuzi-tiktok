package cfg

import (
	"github.com/spf13/viper"
	"sync"
	"time"
)

type Getter interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint16(key string) uint16
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetSizeInBytes(key string) uint
}

// vConfig  并发安全的封装类
type vConfig struct {
	viper *viper.Viper
	rLock sync.RWMutex
}

var VConfig *vConfig

func SetDefault(v *viper.Viper) {
	VConfig = &vConfig{
		viper: v,
		rLock: sync.RWMutex{},
	}
}

func (c *vConfig) GetViper() *viper.Viper {
	return c.viper
}

func (c *vConfig) Get(key string) interface{} {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.Get(key)
}

func (c *vConfig) GetString(key string) string {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetString(key)
}

func (c *vConfig) GetBool(key string) bool {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetBool(key)
}

func (c *vConfig) GetInt(key string) int {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetInt(key)
}

func (c *vConfig) GetInt32(key string) int32 {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetInt32(key)
}

func (c *vConfig) GetInt64(key string) int64 {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetInt64(key)
}

func (c *vConfig) GetUint(key string) uint {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetUint(key)
}

func (c *vConfig) GetUint16(key string) uint16 {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetUint16(key)
}

func (c *vConfig) GetUint32(key string) uint32 {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetUint32(key)
}

func (c *vConfig) GetUint64(key string) uint64 {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetUint64(key)
}

func (c *vConfig) GetFloat64(key string) float64 {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetFloat64(key)
}

func (c *vConfig) GetTime(key string) time.Time {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetTime(key)
}

func (c *vConfig) GetDuration(key string) time.Duration {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetDuration(key)
}

func (c *vConfig) GetIntSlice(key string) []int {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetIntSlice(key)
}

func (c *vConfig) GetStringSlice(key string) []string {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetStringSlice(key)
}

func (c *vConfig) GetStringMap(key string) map[string]interface{} {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetStringMap(key)
}

func (c *vConfig) GetStringMapString(key string) map[string]string {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetStringMapString(key)
}

func (c *vConfig) GetStringMapStringSlice(key string) map[string][]string {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetStringMapStringSlice(key)
}

func (c *vConfig) GetSizeInBytes(key string) uint {
	c.rLock.RLock()
	defer c.rLock.RUnlock()
	return c.viper.GetSizeInBytes(key)
}
