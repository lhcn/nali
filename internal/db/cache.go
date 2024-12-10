package db

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zu1k/nali/pkg/dbif"
)

var (
	dbNameCache = make(map[string]dbif.DB)
	dbTypeCache = make(map[dbif.QueryType]dbif.DB)
	queryCache  = cache.New(24*time.Hour, 10*time.Minute)
)

var (
	NameDBMap = make(NameMap)
	TypeDBMap = make(TypeMap)
)
