package db

type Database interface {
	StoreTokens(id, refresh, access string, expires float64) error
}
