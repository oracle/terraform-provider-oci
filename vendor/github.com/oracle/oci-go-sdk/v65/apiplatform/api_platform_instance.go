// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APIP Control Plane API
//
// Control Plane designed to manage lifecycle of APIP Instances
//

package apiplatform

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApiPlatformInstance A Logical Entity that can be used to create, manage, secure, and advertise APIs to connect to new or existing services
type ApiPlatformInstance struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance
	Id *string `mandatory:"true" json:"id"`

	// A regionally unique, non-changeable instance name provided by the user during instance creation
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the instance was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the instance
	LifecycleState ApiPlatformInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// User-provided changeable and non-unique description of the instance
	Description *string `mandatory:"false" json:"description"`

	IdcsApp *IdcsApp `mandatory:"false" json:"idcsApp"`

	Uris *Uris `mandatory:"false" json:"uris"`

	// The date and time the instance was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the instance in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ApiPlatformInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiPlatformInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApiPlatformInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApiPlatformInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiPlatformInstanceLifecycleStateEnum Enum with underlying type: string
type ApiPlatformInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for ApiPlatformInstanceLifecycleStateEnum
const (
	ApiPlatformInstanceLifecycleStateCreating ApiPlatformInstanceLifecycleStateEnum = "CREATING"
	ApiPlatformInstanceLifecycleStateUpdating ApiPlatformInstanceLifecycleStateEnum = "UPDATING"
	ApiPlatformInstanceLifecycleStateActive   ApiPlatformInstanceLifecycleStateEnum = "ACTIVE"
	ApiPlatformInstanceLifecycleStateDeleting ApiPlatformInstanceLifecycleStateEnum = "DELETING"
	ApiPlatformInstanceLifecycleStateDeleted  ApiPlatformInstanceLifecycleStateEnum = "DELETED"
	ApiPlatformInstanceLifecycleStateFailed   ApiPlatformInstanceLifecycleStateEnum = "FAILED"
)

var mappingApiPlatformInstanceLifecycleStateEnum = map[string]ApiPlatformInstanceLifecycleStateEnum{
	"CREATING": ApiPlatformInstanceLifecycleStateCreating,
	"UPDATING": ApiPlatformInstanceLifecycleStateUpdating,
	"ACTIVE":   ApiPlatformInstanceLifecycleStateActive,
	"DELETING": ApiPlatformInstanceLifecycleStateDeleting,
	"DELETED":  ApiPlatformInstanceLifecycleStateDeleted,
	"FAILED":   ApiPlatformInstanceLifecycleStateFailed,
}

var mappingApiPlatformInstanceLifecycleStateEnumLowerCase = map[string]ApiPlatformInstanceLifecycleStateEnum{
	"creating": ApiPlatformInstanceLifecycleStateCreating,
	"updating": ApiPlatformInstanceLifecycleStateUpdating,
	"active":   ApiPlatformInstanceLifecycleStateActive,
	"deleting": ApiPlatformInstanceLifecycleStateDeleting,
	"deleted":  ApiPlatformInstanceLifecycleStateDeleted,
	"failed":   ApiPlatformInstanceLifecycleStateFailed,
}

// GetApiPlatformInstanceLifecycleStateEnumValues Enumerates the set of values for ApiPlatformInstanceLifecycleStateEnum
func GetApiPlatformInstanceLifecycleStateEnumValues() []ApiPlatformInstanceLifecycleStateEnum {
	values := make([]ApiPlatformInstanceLifecycleStateEnum, 0)
	for _, v := range mappingApiPlatformInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApiPlatformInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for ApiPlatformInstanceLifecycleStateEnum
func GetApiPlatformInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingApiPlatformInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiPlatformInstanceLifecycleStateEnum(val string) (ApiPlatformInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingApiPlatformInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
