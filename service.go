package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

var cm CacheManager = CacheManager{}
var mdm MetadaManager = MetadaManager{}
var crm CryptoManager = CryptoManager{}

//ServiceManager class defines the service layer to provide business methods for required operation
type ServiceManager struct {
}

//HashPassword It computes the SHA256 hash for a given input and returns a unique sequence number for the operation
//pwd -> password for which hash to be generated
func (sm ServiceManager) HashPassword(pwd string) int64 {
	start := time.Now().UnixNano()
	sha2 := crm.SHA2Hash(pwd)
	k := mdm.Increment()
	cm.PutCache(k, sha2)
	end := time.Now().UnixNano()
	callTime := (end - start) / 1000
	getLogger().Debug("Calltime for sequence number:" + strconv.FormatInt(k, 10) + " microSecs:" + strconv.FormatInt(callTime, 10))
	mdm.UpdateStats(callTime, k)
	return k
}

//GetHashForID It returns the SHA256 hash for a operation with sequence number provided as imput
//id -> sequence number of the operation
func (sm ServiceManager) GetHashForID(id int64) (string, error) {
	hash, error := cm.GetCache(id)
	if error != nil {
		return "", error
	}
	if hash.active >= time.Now().Unix() {
		return hash.sha2, nil
	}

	getLogger().Debug("5 seconds have not passed since the operation invoked for sequence number:" + strconv.FormatInt(id, 10))
	return "", errors.New("Not Found")
}

//GetStatistics It returns the Serialized JSON for statistics
func (sm ServiceManager) GetStatistics() (string, error) {
	stats := mdm.GetStats()
	if stats == nil {
		return "", errors.New("No statistics available")
	}

	ret, err := json.Marshal(stats)
	if err != nil {
		return "", err
	}

	out := string(ret)
	return out, nil
}
