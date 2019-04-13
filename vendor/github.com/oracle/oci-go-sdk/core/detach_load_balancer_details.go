// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DetachLoadBalancerDetails Represents a load balancer that is to be detached from an instance pool.
type DetachLoadBalancerDetails struct {

	// The OCID of the load balancer to detach from the instance pool.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// The name of the backend set on the load balancer to detach from the instance pool.
	BackendSetName *string `mandatory:"true" json:"backendSetName"`
}

func (m DetachLoadBalancerDetails) String() string {
	return common.PointerString(m)
}
