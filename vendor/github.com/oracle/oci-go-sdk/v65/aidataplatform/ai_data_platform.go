// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AiDataPlatform Control Plane API
//
// Use the AiDataPlatform Control Plane API to manage Data Lakes.
//

package aidataplatform

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AiDataPlatform An AiDataPlatform is a unified platform for lifecycle management and governance of data and AI objects.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type AiDataPlatform struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the AiDataPlatform.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The AiDataPlatform type.
	AiDataPlatformType *string `mandatory:"true" json:"aiDataPlatformType"`

	// The date and time the AiDataPlatform was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the AiDataPlatform.
	LifecycleState AiDataPlatformLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IAM user.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The date and time the AiDataPlatform was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The alias Id of the AiDataPlatform which is the short form of OCID.
	AliasKey *string `mandatory:"false" json:"aliasKey"`

	// The WebSocket URL of the AiDataPlatform.
	WebSocketEndpoint *string `mandatory:"false" json:"webSocketEndpoint"`

	// A message that describes the current state of the AiDataPlatform in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AiDataPlatform) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AiDataPlatform) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAiDataPlatformLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAiDataPlatformLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AiDataPlatformLifecycleStateEnum Enum with underlying type: string
type AiDataPlatformLifecycleStateEnum string

// Set of constants representing the allowable values for AiDataPlatformLifecycleStateEnum
const (
	AiDataPlatformLifecycleStateCreating AiDataPlatformLifecycleStateEnum = "CREATING"
	AiDataPlatformLifecycleStateUpdating AiDataPlatformLifecycleStateEnum = "UPDATING"
	AiDataPlatformLifecycleStateActive   AiDataPlatformLifecycleStateEnum = "ACTIVE"
	AiDataPlatformLifecycleStateDeleting AiDataPlatformLifecycleStateEnum = "DELETING"
	AiDataPlatformLifecycleStateDeleted  AiDataPlatformLifecycleStateEnum = "DELETED"
	AiDataPlatformLifecycleStateFailed   AiDataPlatformLifecycleStateEnum = "FAILED"
)

var mappingAiDataPlatformLifecycleStateEnum = map[string]AiDataPlatformLifecycleStateEnum{
	"CREATING": AiDataPlatformLifecycleStateCreating,
	"UPDATING": AiDataPlatformLifecycleStateUpdating,
	"ACTIVE":   AiDataPlatformLifecycleStateActive,
	"DELETING": AiDataPlatformLifecycleStateDeleting,
	"DELETED":  AiDataPlatformLifecycleStateDeleted,
	"FAILED":   AiDataPlatformLifecycleStateFailed,
}

var mappingAiDataPlatformLifecycleStateEnumLowerCase = map[string]AiDataPlatformLifecycleStateEnum{
	"creating": AiDataPlatformLifecycleStateCreating,
	"updating": AiDataPlatformLifecycleStateUpdating,
	"active":   AiDataPlatformLifecycleStateActive,
	"deleting": AiDataPlatformLifecycleStateDeleting,
	"deleted":  AiDataPlatformLifecycleStateDeleted,
	"failed":   AiDataPlatformLifecycleStateFailed,
}

// GetAiDataPlatformLifecycleStateEnumValues Enumerates the set of values for AiDataPlatformLifecycleStateEnum
func GetAiDataPlatformLifecycleStateEnumValues() []AiDataPlatformLifecycleStateEnum {
	values := make([]AiDataPlatformLifecycleStateEnum, 0)
	for _, v := range mappingAiDataPlatformLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAiDataPlatformLifecycleStateEnumStringValues Enumerates the set of values in String for AiDataPlatformLifecycleStateEnum
func GetAiDataPlatformLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAiDataPlatformLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAiDataPlatformLifecycleStateEnum(val string) (AiDataPlatformLifecycleStateEnum, bool) {
	enum, ok := mappingAiDataPlatformLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
