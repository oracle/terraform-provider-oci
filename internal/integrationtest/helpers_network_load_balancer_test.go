package integrationtest

import (
	"sync"
	"testing"

	tf_network_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/network_load_balancer"
)

// issue-routing-tag: terraform/default
func TestUnitSafeMutexMap_GetOrCreateNlbBackendSetMutex(t *testing.T) {
	testMap := tf_network_load_balancer.NlbSafeMutexMap{}

	mutex1 := testMap.GetOrCreateNlbBackendSetMutex("nlbocid1", "bes1")
	if mutex1 == nil {
		t.Errorf("Did not get a mutex with new network load balancer id and backend set combination")
		return
	}

	mutex2 := testMap.GetOrCreateNlbBackendSetMutex("nlbocid1", "bes2")
	if mutex2 == mutex1 {
		t.Errorf("Expected a new mutex but got an existing mutex with a new network load balancer and backend set name combination")
		return
	}

	mutex2 = testMap.GetOrCreateNlbBackendSetMutex("nlbocid1", "bes1")
	if mutex2 != mutex1 {
		t.Errorf("Expected an existing mutex, but got a new mutex with an existing network load balancer and backend set name combination")
		return
	}

	mutex2 = testMap.GetOrCreateNlbBackendSetMutex("", "bes2")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with an empty network load balancer id")
		return
	}

	mutex2 = testMap.GetOrCreateNlbBackendSetMutex("nlbocid1", "")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with an empty backend set name")
		return
	}

	mutex2 = testMap.GetOrCreateNlbBackendSetMutex("", "")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with empty network load balancer id and backend set name")
		return
	}
}

// Test that concurrently accessing an empty map with the same nlb and backend set name always gives back the same mutex
// i.e. A mutex should never be overwritten by concurrent accesses
// issue-routing-tag: terraform/default
func TestUnitSafeMutexMap_ConcurrentGetOrCreateNlbBackendSetMutex(t *testing.T) {
	testMap := tf_network_load_balancer.NlbSafeMutexMap{}
	numConcurrentAccesses := 100

	getMutexFn := func(lbId string, backendSetName string, c chan *sync.Mutex) {
		result := testMap.GetOrCreateNlbBackendSetMutex(lbId, backendSetName)
		c <- result
	}

	channel := make(chan *sync.Mutex, numConcurrentAccesses)

	for i := 0; i < numConcurrentAccesses; i++ {
		go getMutexFn("nlbocid1", "bes1", channel)
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
