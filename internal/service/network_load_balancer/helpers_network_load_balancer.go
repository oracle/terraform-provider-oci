package network_load_balancer

import (
	"sync"
)

var nlbMutexes NlbSafeMutexMap

type NlbSafeMutexMap struct {
	mutexes map[string]*sync.Mutex
	m       sync.Mutex // Controls access to this map
}

// Given a load balancer ID, finds a mutex. If a mutex doesn't exist, then Create one for that NLB.
func (safeMap *NlbSafeMutexMap) GetOrCreateNlbMutex(nlbId string) *sync.Mutex {
	if nlbId == "" {
		return nil
	}

	safeMap.m.Lock()
	defer safeMap.m.Unlock()

	key := nlbId

	if safeMap.mutexes == nil {
		safeMap.mutexes = map[string]*sync.Mutex{}
	}

	m, exists := safeMap.mutexes[key]
	if !exists {
		m = &sync.Mutex{}
		safeMap.mutexes[key] = m
	}

	return m
}
