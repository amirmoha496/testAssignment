package main

import (
	"strconv"
	"sync"
)

var stats *Stats = &Stats{Total: 0, Average: 0}
var id *Sequence = &Sequence{id: 0}

//MetadaManager class to retrive and update the metadata for statistics and operations
type MetadaManager struct {
}

//Increment It generates a sequence number for a password hash operation, it is a synchronized method and returns unique sequence value
func (mdm MetadaManager) Increment() int64 {
	var ret int64
	id.mutex.Lock()
	id.id = id.id + 1
	ret = id.id
	id.mutex.Unlock()
	getLogger().Debug("Generated sequence number:" + strconv.FormatInt(ret, 10))
	return ret
}

//GetStats It returns the current statistics for number of operations performed and average time hash computation
func (mdm MetadaManager) GetStats() *Stats {
	return stats
}

//UpdateStats This methods updates the current statistics with the time provided as input
//callTime -> is hash computation time in microseconds
//id -> sequence number of the operation for which this call time is reported
func (mdm MetadaManager) UpdateStats(callTime int64, id int64) {
	stats.mutex.Lock()
	totalTime := stats.Average * stats.Total
	stats.Total = stats.Total + 1
	totalTime = totalTime + callTime
	stats.Average = totalTime / stats.Total
	stats.mutex.Unlock()
	getLogger().Debug("Update the statistics with new operation time of:" + strconv.FormatInt(callTime, 10) + " microSecs for sequence:" + strconv.FormatInt(id, 10))
	return
}

//Stats This structure stores the statistics regarding average hash computation time and number of operations performed
type Stats struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
	mutex   sync.Mutex
}

//Sequence This structure defines the next sequence number to be generated
type Sequence struct {
	id    int64
	mutex sync.Mutex
}
