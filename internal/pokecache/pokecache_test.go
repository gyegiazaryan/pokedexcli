package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("val1"))

	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found")
	}
	if string(actual) != "val1" {
		t.Errorf("value doesn't match")
	}
}
