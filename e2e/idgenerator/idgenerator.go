package idgenerator

import "sync"

var (
	lastID int = 10000
	mutex  sync.Mutex
)

func GenerateID() int {
	mutex.Lock()
	defer mutex.Unlock()
	lastID++
	return lastID
}
