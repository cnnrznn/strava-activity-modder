package db

import "log"

func StoreTokens(refresh, access string, expire float64) {
	log.Printf("Stored %v, %v, %v\n", refresh, access, expire)
}
