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

func (db *MemoryDB) StoreToken(id, refresh, access string, expires float64) error {
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
