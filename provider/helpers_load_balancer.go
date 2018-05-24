// Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"sync"
)

var lbBackendSetMutexes SafeMutexMap

type SafeMutexMap struct {
	mutexes map[string]*sync.Mutex
	m       sync.Mutex // Controls access to this map
}

// Given a load balancer ID and backend set name, finds a mutex. If a mutex doesn't exist, then create one for that backend set.
//
// We can add more generic ways of accessing this map in the future; if other resources need to use this pattern for
// resolving concurrent resource access issues. For now, keep it specific to backend sets
func (safeMap *SafeMutexMap) GetOrCreateBackendSetMutex(lbId string, backendSetName string) *sync.Mutex {
	if lbId == "" || backendSetName == "" {
		return nil
	}

	safeMap.m.Lock()
	defer safeMap.m.Unlock()

	key := fmt.Sprintf("%s.%s", lbId, backendSetName)

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
