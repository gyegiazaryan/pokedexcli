package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key1", []byte("val1"))

	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found")
	}
	if string(actual) != "val1" {
		t.Errorf("value doesn't match")
	}
}

func TestPurgeCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cache.Add("key1", []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")
	if ok {
		t.Error("key1 should have been purged")
	}
}
