// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailOutboundIpSummary Summary of the Outbound IP address assigned to the tenancy.
type EmailOutboundIpSummary struct {

	// The public IP address assigned to the tenancy.
	OutboundIp *string `mandatory:"true" json:"outboundIp"`

	// The assignment state of the public IP address.
	AssignmentState EmailOutboundIpSummaryAssignmentStateEnum `mandatory:"true" json:"assignmentState"`

	// The current state of the Email Outbound Public IP.
	LifecycleState EmailOutboundIpSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a
	// resource in 'DRAINING' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m EmailOutboundIpSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailOutboundIpSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailOutboundIpSummaryAssignmentStateEnum(string(m.AssignmentState)); !ok && m.AssignmentState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssignmentState: %s. Supported values are: %s.", m.AssignmentState, strings.Join(GetEmailOutboundIpSummaryAssignmentStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEmailOutboundIpSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailOutboundIpSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailOutboundIpSummaryLifecycleStateEnum Enum with underlying type: string
type EmailOutboundIpSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for EmailOutboundIpSummaryLifecycleStateEnum
const (
	EmailOutboundIpSummaryLifecycleStateUpdating EmailOutboundIpSummaryLifecycleStateEnum = "UPDATING"
	EmailOutboundIpSummaryLifecycleStateActive   EmailOutboundIpSummaryLifecycleStateEnum = "ACTIVE"
	EmailOutboundIpSummaryLifecycleStateFailed   EmailOutboundIpSummaryLifecycleStateEnum = "FAILED"
	EmailOutboundIpSummaryLifecycleStateDraining EmailOutboundIpSummaryLifecycleStateEnum = "DRAINING"
)

var mappingEmailOutboundIpSummaryLifecycleStateEnum = map[string]EmailOutboundIpSummaryLifecycleStateEnum{
	"UPDATING": EmailOutboundIpSummaryLifecycleStateUpdating,
	"ACTIVE":   EmailOutboundIpSummaryLifecycleStateActive,
	"FAILED":   EmailOutboundIpSummaryLifecycleStateFailed,
	"DRAINING": EmailOutboundIpSummaryLifecycleStateDraining,
}

var mappingEmailOutboundIpSummaryLifecycleStateEnumLowerCase = map[string]EmailOutboundIpSummaryLifecycleStateEnum{
	"updating": EmailOutboundIpSummaryLifecycleStateUpdating,
	"active":   EmailOutboundIpSummaryLifecycleStateActive,
	"failed":   EmailOutboundIpSummaryLifecycleStateFailed,
	"draining": EmailOutboundIpSummaryLifecycleStateDraining,
}

// GetEmailOutboundIpSummaryLifecycleStateEnumValues Enumerates the set of values for EmailOutboundIpSummaryLifecycleStateEnum
func GetEmailOutboundIpSummaryLifecycleStateEnumValues() []EmailOutboundIpSummaryLifecycleStateEnum {
	values := make([]EmailOutboundIpSummaryLifecycleStateEnum, 0)
	for _, v := range mappingEmailOutboundIpSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailOutboundIpSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for EmailOutboundIpSummaryLifecycleStateEnum
func GetEmailOutboundIpSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"UPDATING",
		"ACTIVE",
		"FAILED",
		"DRAINING",
	}
}

// GetMappingEmailOutboundIpSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailOutboundIpSummaryLifecycleStateEnum(val string) (EmailOutboundIpSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingEmailOutboundIpSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailOutboundIpSummaryAssignmentStateEnum Enum with underlying type: string
type EmailOutboundIpSummaryAssignmentStateEnum string

// Set of constants representing the allowable values for EmailOutboundIpSummaryAssignmentStateEnum
const (
	EmailOutboundIpSummaryAssignmentStateAvailable EmailOutboundIpSummaryAssignmentStateEnum = "AVAILABLE"
	EmailOutboundIpSummaryAssignmentStateAssigned  EmailOutboundIpSummaryAssignmentStateEnum = "ASSIGNED"
)

var mappingEmailOutboundIpSummaryAssignmentStateEnum = map[string]EmailOutboundIpSummaryAssignmentStateEnum{
	"AVAILABLE": EmailOutboundIpSummaryAssignmentStateAvailable,
	"ASSIGNED":  EmailOutboundIpSummaryAssignmentStateAssigned,
}

var mappingEmailOutboundIpSummaryAssignmentStateEnumLowerCase = map[string]EmailOutboundIpSummaryAssignmentStateEnum{
	"available": EmailOutboundIpSummaryAssignmentStateAvailable,
	"assigned":  EmailOutboundIpSummaryAssignmentStateAssigned,
}

// GetEmailOutboundIpSummaryAssignmentStateEnumValues Enumerates the set of values for EmailOutboundIpSummaryAssignmentStateEnum
func GetEmailOutboundIpSummaryAssignmentStateEnumValues() []EmailOutboundIpSummaryAssignmentStateEnum {
	values := make([]EmailOutboundIpSummaryAssignmentStateEnum, 0)
	for _, v := range mappingEmailOutboundIpSummaryAssignmentStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailOutboundIpSummaryAssignmentStateEnumStringValues Enumerates the set of values in String for EmailOutboundIpSummaryAssignmentStateEnum
func GetEmailOutboundIpSummaryAssignmentStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"ASSIGNED",
	}
}

// GetMappingEmailOutboundIpSummaryAssignmentStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailOutboundIpSummaryAssignmentStateEnum(val string) (EmailOutboundIpSummaryAssignmentStateEnum, bool) {
	enum, ok := mappingEmailOutboundIpSummaryAssignmentStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
