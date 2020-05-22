package db

type Database interface {
	StoreTokens(refresh, access string, expires float64) error
}
