package database

import "sync"

var SyncMap sync.Map

func Start() {
	SyncMap = sync.Map{}
}
