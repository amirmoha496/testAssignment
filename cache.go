package main

import (
	"errors"
	"strconv"
	"time"
)

var database map[int64]Hash = make(map[int64]Hash)

//CacheManager class provided for utility methods for handeling Cache operations
type CacheManager struct {
}

//Hash It defines the structure that holds the SHA512 hash value ant the time at which it will be active once it is computed
type Hash struct {
	sha2   string
	active int64
}

//PutCache It stores the sequence number and the hash of the password into the in memory cache
//k -> the sequence number for the password hashing operation
//v -> the computed hash for the given sequence number
func (cm CacheManager) PutCache(k int64, v string) {
	hash := Hash{sha2: v}
	hash.active = time.Now().Unix() + defaultWait
	database[k] = hash
	getLogger().Debug("Persisted in cache for sequence:" + strconv.FormatInt(k, 10))
}

//GetCache It returns the Hash object for the password which was stored in the cache with sequence number k
//k -> is the sequence number for the password hashing operation
func (cm CacheManager) GetCache(k int64) (Hash, error) {
	hash, found := database[k]
	if found == true {
		getLogger().Debug("Found in cache for sequence:" + strconv.FormatInt(k, 10))
		return hash, nil
	}
	getLogger().Debug("Not found in cache for sequence:" + strconv.FormatInt(k, 10))
	return Hash{}, errors.New("Not Found")
}

const defaultWait int64 = 5000
