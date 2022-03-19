package main

import (
	"sync"
	"time"
)

type Database struct {
	sync.RWMutex
	data map[any]any
}

func (db *Database) write(f func(db map[any]any)) {
	db.Lock()
	f(db.data)
	db.Unlock()
}

func (db *Database) read(f func(db map[any]any) any) any {
	db.RLock()
	value := f(db.data)
	db.RUnlock()
	return value
}

var mainDb = Database{sync.RWMutex{}, map[any]any{}}

func expireKey(key any, ms int) {
	var duration = time.Duration(ms * int(time.Millisecond))
	time.Sleep(duration)
	deleteKey(key)
}

func deleteKey(key any) {
	mainDb.write(func(db map[any]any) {
		delete(db, key)
	})
}

func getKey(key any) any {
	value := mainDb.read(func(db map[any]any) any {
		return db[key]
	})
	return value
}

func setKey(key any, value any, px int) {
	mainDb.write(func(db map[any]any) {
		db[key] = value
	})
	if px > 0 {
		go expireKey(key, px)
	}
}
