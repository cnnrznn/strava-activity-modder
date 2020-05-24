package memory

import "sync"

type user struct {
	id           string
	refreshToken string
	accessToken  string
	expires      float64
}

type MemoryDB struct {
	table map[string]user
	sync.Mutex
}

func New() *MemoryDB {
	return &MemoryDB{
		table: make(map[string]user),
	}
}

func (db *MemoryDB) StoreTokens(id, refresh, access string, expires float64) error {
	db.Lock()
	defer db.Unlock()

	db.table[id] = user{
		id:           id,
		refreshToken: refresh,
		accessToken:  access,
		expires:      expires,
	}

	return nil
}
