package db

type Database interface {
	StoreTokens(refresh, access string, id, expires int) error
}
