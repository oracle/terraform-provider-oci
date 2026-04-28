package integrationtest

import (
	"sync"
	"testing"

	tf_network_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/network_load_balancer"
)

// issue-routing-tag: terraform/default
func TestUnitSafeMutexMap_GetOrCreateNlbMutex(t *testing.T) {
	testMap := tf_network_load_balancer.NlbSafeMutexMap{}

	mutex1 := testMap.GetOrCreateNlbMutex("nlbocid1")
	if mutex1 == nil {
		t.Errorf("Did not get a mutex for new network load balancer id")
		return
	}

	mutex2 := testMap.GetOrCreateNlbMutex("nlbocid1")
	if mutex2 != mutex1 {
		t.Errorf("Expected an existing mutex, but got a new mutex with an existing network load balancer and backend set name combination")
		return
	}

	mutex2 = testMap.GetOrCreateNlbMutex("")
	if mutex2 != nil {
		t.Errorf("Expected a nil mutex but got a valid one with an empty network load balancer id")
		return
	}
}

// Test that concurrently accessing an empty map with the same nlb and backend set name always gives back the same mutex
// i.e. A mutex should never be overwritten by concurrent accesses
// issue-routing-tag: terraform/default
func TestUnitSafeMutexMap_ConcurrentGetOrCreateNlbMutex(t *testing.T) {
	testMap := tf_network_load_balancer.NlbSafeMutexMap{}
	numConcurrentAccesses := 100

	getMutexFn := func(lbId string, c chan *sync.Mutex) {
		result := testMap.GetOrCreateNlbMutex(lbId)
		c <- result
	}

	channel := make(chan *sync.Mutex, numConcurrentAccesses)

	for i := 0; i < numConcurrentAccesses; i++ {
		go getMutexFn("nlbocid1", channel)
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
