package memory

import "sync"

type user struct {
	id           int
	expires      int
	refreshToken string
	accessToken  string
}

type MemoryDB struct {
	table map[int]user
	sync.Mutex
}

func New() *MemoryDB {
	return &MemoryDB{
		table: make(map[int]user),
	}
}

func (db *MemoryDB) StoreTokens(refresh, access string, id, expires int) error {
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
