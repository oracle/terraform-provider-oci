// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequest Many of the API requests you use to create and configure load balancing do not take effect immediately.
// In these cases, the request spawns an asynchronous work flow to fulfill the request. WorkRequest objects provide visibility
// for in-progress work flows.
// For more information about work requests, see Viewing the State of a Work Request (https://docs.oracle.com/iaas/Content/Balance/Tasks/viewingworkrequest.htm).
type WorkRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer with which the work request
	// is associated.
	LoadBalancerId *string `mandatory:"true" json:"loadBalancerId"`

	// The type of action the work request represents.
	// Example: `CreateListener`
	Type *string `mandatory:"true" json:"type"`

	// The current state of the work request.
	LifecycleState WorkRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A collection of data, related to the load balancer provisioning process, that helps with debugging in the event of failure.
	// Possible data elements include:
	// - workflow name
	// - event ID
	// - work request ID
	// - load balancer ID
	// - workflow completion message
	Message *string `mandatory:"true" json:"message"`

	// The date and time the work request was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	ErrorDetails []WorkRequestError `mandatory:"true" json:"errorDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The date and time the work request was completed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetWorkRequestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestLifecycleStateEnum Enum with underlying type: string
type WorkRequestLifecycleStateEnum string

// Set of constants representing the allowable values for WorkRequestLifecycleStateEnum
const (
	WorkRequestLifecycleStateAccepted   WorkRequestLifecycleStateEnum = "ACCEPTED"
	WorkRequestLifecycleStateInProgress WorkRequestLifecycleStateEnum = "IN_PROGRESS"
	WorkRequestLifecycleStateFailed     WorkRequestLifecycleStateEnum = "FAILED"
	WorkRequestLifecycleStateSucceeded  WorkRequestLifecycleStateEnum = "SUCCEEDED"
)

var mappingWorkRequestLifecycleStateEnum = map[string]WorkRequestLifecycleStateEnum{
	"ACCEPTED":    WorkRequestLifecycleStateAccepted,
	"IN_PROGRESS": WorkRequestLifecycleStateInProgress,
	"FAILED":      WorkRequestLifecycleStateFailed,
	"SUCCEEDED":   WorkRequestLifecycleStateSucceeded,
}

var mappingWorkRequestLifecycleStateEnumLowerCase = map[string]WorkRequestLifecycleStateEnum{
	"accepted":    WorkRequestLifecycleStateAccepted,
	"in_progress": WorkRequestLifecycleStateInProgress,
	"failed":      WorkRequestLifecycleStateFailed,
	"succeeded":   WorkRequestLifecycleStateSucceeded,
}

// GetWorkRequestLifecycleStateEnumValues Enumerates the set of values for WorkRequestLifecycleStateEnum
func GetWorkRequestLifecycleStateEnumValues() []WorkRequestLifecycleStateEnum {
	values := make([]WorkRequestLifecycleStateEnum, 0)
	for _, v := range mappingWorkRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestLifecycleStateEnumStringValues Enumerates the set of values in String for WorkRequestLifecycleStateEnum
func GetWorkRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingWorkRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestLifecycleStateEnum(val string) (WorkRequestLifecycleStateEnum, bool) {
	enum, ok := mappingWorkRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
