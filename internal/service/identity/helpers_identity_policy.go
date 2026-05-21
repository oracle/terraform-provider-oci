// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import "sync"

var identityPolicyMutexes IdentityPolicySafeMutexMap

type IdentityPolicySafeMutexMap struct {
	mutexes map[string]*sync.Mutex
	m       sync.Mutex
}

// GetOrCreatePolicyMutex returns a mutex for policy mutations within a tenancy.
func (safeMap *IdentityPolicySafeMutexMap) GetOrCreatePolicyMutex(tenancyId string) *sync.Mutex {
	if tenancyId == "" {
		tenancyId = "unknown-tenancy"
	}

	safeMap.m.Lock()
	defer safeMap.m.Unlock()

	if safeMap.mutexes == nil {
		safeMap.mutexes = map[string]*sync.Mutex{}
	}

	m, exists := safeMap.mutexes[tenancyId]
	if !exists {
		m = &sync.Mutex{}
		safeMap.mutexes[tenancyId] = m
	}

	return m
}
