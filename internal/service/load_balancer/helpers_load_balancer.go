// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var lbBackendSetMutexes SafeMutexMap

type SafeMutexMap struct {
	mutexes map[string]*sync.Mutex
	m       sync.Mutex // Controls access to this map
}

// Given a load balancer ID and backend set name, finds a mutex. If a mutex doesn't exist, then Create one for that backend set.
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

func loadBalancerResourceID(res interface{}, workReq *oci_load_balancer.WorkRequest) (id *string, workReqSucceeded bool) {
	v := reflect.ValueOf(res).Elem()
	if v.IsValid() {
		// This is super fugly. It's this way because the LB API has no convention for ID formats.

		// Load balancer
		id := v.FieldByName("Id")
		if id.IsValid() && !id.IsNil() {
			s := id.Elem().String()
			return &s, false
		}
		// backendset, listener
		name := v.FieldByName("Name")
		if name.IsValid() && !name.IsNil() {
			s := name.Elem().String()
			return &s, false
		}
		// certificate
		certName := v.FieldByName("CertificateName")
		if certName.IsValid() && !certName.IsNil() {
			s := certName.Elem().String()
			return &s, false
		}
		// backend TODO The following can probably be removed because the Backend object has a Name parameter)
		ip := v.FieldByName("IpAddress")
		port := v.FieldByName("Port")
		if ip.IsValid() && !ip.IsNil() && port.IsValid() && !port.IsNil() {
			s := ip.Elem().String() + ":" + strconv.Itoa(int(int(port.Elem().Int())))
			return &s, false
		}
	}
	if workReq != nil {
		if workReq.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return nil, true
		} else {
			return workReq.Id, false
		}
	}
	return nil, false
}

func loadBalancerResourceGet(client *oci_load_balancer.LoadBalancerClient, d *schema.ResourceData, wr *oci_load_balancer.WorkRequest, retryPolicy *oci_common.RetryPolicy) (id string, stillWorking bool, err error) {
	// NOTE: if the id is for a work request, refresh its state and loadBalancerID.
	if wr != nil && wr.Id != nil {
		getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
		getWorkRequestRequest.WorkRequestId = wr.Id
		getWorkRequestRequest.RequestMetadata.RetryPolicy = retryPolicy
		updatedWorkRes, err := client.GetWorkRequest(context.Background(), getWorkRequestRequest)
		if err != nil {
			return "", false, err
		}
		if wr != nil {
			*wr = updatedWorkRes.WorkRequest
			d.Set("state", wr.LifecycleState)
			if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
				return "", false, nil
			}
			if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateFailed {
				return "", false, fmt.Errorf("WorkRequest FAILED: %+v", wr.ErrorDetails)
			}
		}
		return "", true, nil
	}
	return id, false, nil
}

func loadBalancerWaitForWorkRequest(client *oci_load_balancer.LoadBalancerClient, d *schema.ResourceData, wr *oci_load_balancer.WorkRequest, retryPolicy *oci_common.RetryPolicy) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
			string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
		},
		Target: []string{
			string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
			string(oci_load_balancer.WorkRequestLifecycleStateFailed),
		},
		Refresh: func() (interface{}, string, error) {
			getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
			getWorkRequestRequest.WorkRequestId = wr.Id
			getWorkRequestRequest.RequestMetadata.RetryPolicy = retryPolicy
			workRequestResponse, err := client.GetWorkRequest(context.Background(), getWorkRequestRequest)
			wr = &workRequestResponse.WorkRequest
			return wr, string(wr.LifecycleState), err
		},
		Timeout: d.Timeout(schema.TimeoutCreate),
	}

	// Should not wait when in replay mode
	if httpreplay.ShouldRetryImmediately() {
		stateConf.PollInterval = 1
	}

	if _, e := stateConf.WaitForState(); e != nil {
		return e
	}

	if wr.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateFailed {
		return fmt.Errorf("WorkRequest FAILED: %+v", wr.ErrorDetails)
	}
	return nil
}
