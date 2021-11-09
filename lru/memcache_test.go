package main

import "testing"

type simpleStruct struct {
	int
	string
}

type complexStruct struct {
	int
	simpleStruct
}

var getTests = []struct {
	name       string
	keyToAdd   string
	keyToGet   string
	expectedOk bool
}{
	{"string_hit", "myKey", "myKey", true},
	{"string_miss", "myKey", "nonsense", false},
}

func TestSet(t *testing.T) {
	var cache Cache
	cache = NewMemcache(0)
	values := []string{"test1", "test2", "test3"}
	key := "key1"
	for _, v := range values {
		cache.Set(key, v)
		val, ok := cache.Get(key)
		if !ok {
			t.Fatalf("expect key:%v ,value:%v", key, v)
		} else if ok && val != v {
			t.Fatalf("expect value:%v, get value:%v", v, val)
		}
		t.Logf("value:%v ", val)
	}
}

func TestGet(t *testing.T) {
	var cache Cache
	cache = NewMemcache(0)
	for _, tt := range getTests {
		cache.Set(tt.keyToAdd, 1234)
		val, ok := cache.Get(tt.keyToGet)

		if ok != tt.expectedOk {
			t.Fatalf("%s: val:%v cache hit = %v; want %v", tt.name, val, ok, !ok)
		} else if ok && val != 1234 {
			t.Fatalf("%s expected get to return 1234 but got %v", tt.name, val)
		}

	}
}

func TestLRU(t *testing.T) {
	keys := []string{"1", "2", "3", "4", "2", "1", "3", "5", "6", "5", "6"}
	maxSize := 3
	var cache Cache
	cache = NewMemcache(maxSize)
	for i, key := range keys {
		cache.Set(key, 1234)
		if i == 3 {
			status := cache.Status()
			if status.CurrentSize != maxSize {
				t.Fatalf("expected maxSize %v,currentSize:%v", maxSize, status.CurrentSize)
			}
			_, ok1 := cache.Get("2")
			_, ok2 := cache.Get("3")
			if !(ok1 && ok2) {
				t.Fatalf("expected remains key 2:%v,3:%v", ok1, ok2)
			}
		}

		if i == 5 {
			_, ok1 := cache.Get("1")
			_, ok2 := cache.Get("2")
			_, ok3 := cache.Get("3")

			if !(ok1 && ok2 && ok3) {
				t.Fatalf("expected remains key 1:%v 2:%v,3:%v", ok1, ok2, ok3)
			}
		}
	}

	status := cache.Status()
	if status.CurrentSize != maxSize {
		t.Fatalf("expected maxSize %v,currentSize:%v", maxSize, status.CurrentSize)
	}

	_, ok1 := cache.Get("3")
	_, ok2 := cache.Get("5")
	_, ok3 := cache.Get("6")
	if !(ok1 && ok2 && ok3) {
		t.Fatalf("expected remains key 3:%v,5:%v, 6:%v", ok1, ok2, ok3)
	}
}