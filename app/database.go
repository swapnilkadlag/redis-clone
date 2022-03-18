package main

import (
	"fmt"
	"sync"
	"time"
)

var db = struct {
	sync.RWMutex
	data map[any]any
}{data: make(map[any]any)}

func getKey(key any) any {
	db.RLock()
	value := db.data[key]
	db.RUnlock()
	return value
}

func setKey(key any, value any, px int) {
	db.Lock()
	db.data[key] = value
	db.Unlock()
	if px > 0 {
		go expireKey(key, px)
	}
}

func expireKey(key any, ms int) {
	var duration = time.Duration(ms * int(time.Millisecond))
	time.Sleep(duration)
	deleteKey(key)
}

func deleteKey(key any) {
	db.Lock()
	delete(db.data, key)
	fmt.Println("Deleted key -> " + fmt.Sprint(key))
	db.Unlock()
}
