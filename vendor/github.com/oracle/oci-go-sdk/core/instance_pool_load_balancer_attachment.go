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

// InstancePoolLoadBalancerAttachment Represents a load balancer that is attached to an instance pool.
type InstancePoolLoadBalancerAttachment struct {

	// The OCID of the load balancer attachment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the instance pool of the load balancer attachment.
	InstancePoolId *string `mandatory:"true" json:"instancePoolId"`

	// The OCID of the load balancer attached to the pool.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// The name of the backend set on the load balancer.
	BackendSetName *string `mandatory:"true" json:"backendSetName"`

	// The port value used for the backends.
	Port *int `mandatory:"true" json:"port"`

	// Indicates which vnic on each instance in the pool should be used to associate with the load balancer. possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated to the instance pool.
	VnicSelection *string `mandatory:"true" json:"vnicSelection"`

	// The status of the interaction between the pool and the load balancer.
	LifecycleState InstancePoolLoadBalancerAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m InstancePoolLoadBalancerAttachment) String() string {
	return common.PointerString(m)
}

// InstancePoolLoadBalancerAttachmentLifecycleStateEnum Enum with underlying type: string
type InstancePoolLoadBalancerAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for InstancePoolLoadBalancerAttachmentLifecycleStateEnum
const (
	InstancePoolLoadBalancerAttachmentLifecycleStateAttaching InstancePoolLoadBalancerAttachmentLifecycleStateEnum = "ATTACHING"
	InstancePoolLoadBalancerAttachmentLifecycleStateAttached  InstancePoolLoadBalancerAttachmentLifecycleStateEnum = "ATTACHED"
	InstancePoolLoadBalancerAttachmentLifecycleStateDetaching InstancePoolLoadBalancerAttachmentLifecycleStateEnum = "DETACHING"
	InstancePoolLoadBalancerAttachmentLifecycleStateDetached  InstancePoolLoadBalancerAttachmentLifecycleStateEnum = "DETACHED"
)

var mappingInstancePoolLoadBalancerAttachmentLifecycleState = map[string]InstancePoolLoadBalancerAttachmentLifecycleStateEnum{
	"ATTACHING": InstancePoolLoadBalancerAttachmentLifecycleStateAttaching,
	"ATTACHED":  InstancePoolLoadBalancerAttachmentLifecycleStateAttached,
	"DETACHING": InstancePoolLoadBalancerAttachmentLifecycleStateDetaching,
	"DETACHED":  InstancePoolLoadBalancerAttachmentLifecycleStateDetached,
}

// GetInstancePoolLoadBalancerAttachmentLifecycleStateEnumValues Enumerates the set of values for InstancePoolLoadBalancerAttachmentLifecycleStateEnum
func GetInstancePoolLoadBalancerAttachmentLifecycleStateEnumValues() []InstancePoolLoadBalancerAttachmentLifecycleStateEnum {
	values := make([]InstancePoolLoadBalancerAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingInstancePoolLoadBalancerAttachmentLifecycleState {
		values = append(values, v)
	}
	return values
}
