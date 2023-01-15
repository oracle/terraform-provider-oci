// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessPolicy Access policies enable administrators to restrict the access of certain services.
type AccessPolicy struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation.
	// Avoid entering confidential information.
	// Example: `My unique resource name`
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Resource.
	LifecycleState AccessPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the service mesh in which this access policy is created.
	MeshId *string `mandatory:"false" json:"meshId"`

	// List of applicable rules.
	Rules []AccessPolicyRule `mandatory:"false" json:"rules"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AccessPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAccessPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAccessPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AccessPolicyLifecycleStateEnum Enum with underlying type: string
type AccessPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for AccessPolicyLifecycleStateEnum
const (
	AccessPolicyLifecycleStateCreating AccessPolicyLifecycleStateEnum = "CREATING"
	AccessPolicyLifecycleStateUpdating AccessPolicyLifecycleStateEnum = "UPDATING"
	AccessPolicyLifecycleStateActive   AccessPolicyLifecycleStateEnum = "ACTIVE"
	AccessPolicyLifecycleStateDeleting AccessPolicyLifecycleStateEnum = "DELETING"
	AccessPolicyLifecycleStateDeleted  AccessPolicyLifecycleStateEnum = "DELETED"
	AccessPolicyLifecycleStateFailed   AccessPolicyLifecycleStateEnum = "FAILED"
)

var mappingAccessPolicyLifecycleStateEnum = map[string]AccessPolicyLifecycleStateEnum{
	"CREATING": AccessPolicyLifecycleStateCreating,
	"UPDATING": AccessPolicyLifecycleStateUpdating,
	"ACTIVE":   AccessPolicyLifecycleStateActive,
	"DELETING": AccessPolicyLifecycleStateDeleting,
	"DELETED":  AccessPolicyLifecycleStateDeleted,
	"FAILED":   AccessPolicyLifecycleStateFailed,
}

var mappingAccessPolicyLifecycleStateEnumLowerCase = map[string]AccessPolicyLifecycleStateEnum{
	"creating": AccessPolicyLifecycleStateCreating,
	"updating": AccessPolicyLifecycleStateUpdating,
	"active":   AccessPolicyLifecycleStateActive,
	"deleting": AccessPolicyLifecycleStateDeleting,
	"deleted":  AccessPolicyLifecycleStateDeleted,
	"failed":   AccessPolicyLifecycleStateFailed,
}

// GetAccessPolicyLifecycleStateEnumValues Enumerates the set of values for AccessPolicyLifecycleStateEnum
func GetAccessPolicyLifecycleStateEnumValues() []AccessPolicyLifecycleStateEnum {
	values := make([]AccessPolicyLifecycleStateEnum, 0)
	for _, v := range mappingAccessPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for AccessPolicyLifecycleStateEnum
func GetAccessPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAccessPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessPolicyLifecycleStateEnum(val string) (AccessPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingAccessPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
