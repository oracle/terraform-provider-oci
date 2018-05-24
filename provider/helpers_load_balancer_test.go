package provider

import (
	"sync"
	"testing"
)

func TestSafeMutexMap_GetOrCreateBackendSetMutex(t *testing.T) {
	testMap := SafeMutexMap{}

	mutex1 := testMap.GetOrCreateBackendSetMutex("lbocid1", "bes1")
	if mutex1 == nil {
		t.Errorf("Did not get a mutex with new load balancer id and backend set combination")
		return
	}

	mutex2 := testMap.GetOrCreateBackendSetMutex("lbocid1", "bes2")
	if mutex2 == mutex1 {
		t.Errorf("Expected a new mutex but got an existing mutex with a new load balancer and backend set name combination")
		return
	}

	mutex2 = testMap.GetOrCreateBackendSetMutex("lbocid1", "bes1")
	if mutex2 != mutex1 {
		t.Errorf("Expected an existing mutex, but got a new mutex with an existing load balancer and backend set name combination")
		return
	}

	mutex2 = testMap.GetOrCreateBackendSetMutex("", "bes2")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with an empty load balancer id")
		return
	}

	mutex2 = testMap.GetOrCreateBackendSetMutex("lbocid1", "")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with an empty backend set name")
		return
	}

	mutex2 = testMap.GetOrCreateBackendSetMutex("", "")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with empty load balancer id and backend set name")
		return
	}
}

// Test that concurrently accessing an empty map with the same lb and backend set name always gives back the same mutex
// i.e. A mutex should never be overwritten by concurrent accesses
func TestSafeMutexMap_ConcurrentGetOrCreateBackendSetMutex(t *testing.T) {
	testMap := SafeMutexMap{}
	numConcurrentAccesses := 100

	getMutexFn := func(lbId string, backendSetName string, c chan *sync.Mutex) {
		result := testMap.GetOrCreateBackendSetMutex(lbId, backendSetName)
		c <- result
	}

	channel := make(chan *sync.Mutex, numConcurrentAccesses)

	for i := 0; i < numConcurrentAccesses; i++ {
		go getMutexFn("lbocid1", "bes1", channel)
	}

	mutex1 := <-channel
	for i := 0; i < numConcurrentAccesses-1; i++ {
		mutex2 := <-channel
		if mutex1 != mutex2 {
			t.Errorf("Expected the same mutex but got back a different one")
			return
		}
	}
}
