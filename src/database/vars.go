package database

import (
	"sync"

	"gorm.io/gorm"
)

var (
	dbSession *gorm.DB
	dbMutex   = &sync.Mutex{}
)
