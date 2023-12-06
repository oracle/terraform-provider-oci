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

// VirtualService This resource represents a customer-managed service in the Service Mesh. Each virtual service declares multiple running versions of the service and maps to a group of instances/pods running a specific version of the actual service.
type VirtualService struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the service mesh in which this virtual service is created.
	MeshId *string `mandatory:"true" json:"meshId"`

	// A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation.
	// Avoid entering confidential information.
	// Example: `My unique resource name`
	Name *string `mandatory:"true" json:"name"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Resource.
	LifecycleState VirtualServiceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	DefaultRoutingPolicy *DefaultVirtualServiceRoutingPolicy `mandatory:"false" json:"defaultRoutingPolicy"`

	// The DNS hostnames of the virtual service that is used by its callers.
	// Wildcard hostnames are supported in the prefix form.
	// Examples of valid hostnames are "www.example.com", "*.example.com", "*.com".
	// Can be omitted if the virtual service will only have TCP virtual deployments.
	Hosts []string `mandatory:"false" json:"hosts"`

	Mtls *MutualTransportLayerSecurity `mandatory:"false" json:"mtls"`

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

func (m VirtualService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VirtualService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVirtualServiceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVirtualServiceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VirtualServiceLifecycleStateEnum Enum with underlying type: string
type VirtualServiceLifecycleStateEnum string

// Set of constants representing the allowable values for VirtualServiceLifecycleStateEnum
const (
	VirtualServiceLifecycleStateCreating VirtualServiceLifecycleStateEnum = "CREATING"
	VirtualServiceLifecycleStateUpdating VirtualServiceLifecycleStateEnum = "UPDATING"
	VirtualServiceLifecycleStateActive   VirtualServiceLifecycleStateEnum = "ACTIVE"
	VirtualServiceLifecycleStateDeleting VirtualServiceLifecycleStateEnum = "DELETING"
	VirtualServiceLifecycleStateDeleted  VirtualServiceLifecycleStateEnum = "DELETED"
	VirtualServiceLifecycleStateFailed   VirtualServiceLifecycleStateEnum = "FAILED"
)

var mappingVirtualServiceLifecycleStateEnum = map[string]VirtualServiceLifecycleStateEnum{
	"CREATING": VirtualServiceLifecycleStateCreating,
	"UPDATING": VirtualServiceLifecycleStateUpdating,
	"ACTIVE":   VirtualServiceLifecycleStateActive,
	"DELETING": VirtualServiceLifecycleStateDeleting,
	"DELETED":  VirtualServiceLifecycleStateDeleted,
	"FAILED":   VirtualServiceLifecycleStateFailed,
}

var mappingVirtualServiceLifecycleStateEnumLowerCase = map[string]VirtualServiceLifecycleStateEnum{
	"creating": VirtualServiceLifecycleStateCreating,
	"updating": VirtualServiceLifecycleStateUpdating,
	"active":   VirtualServiceLifecycleStateActive,
	"deleting": VirtualServiceLifecycleStateDeleting,
	"deleted":  VirtualServiceLifecycleStateDeleted,
	"failed":   VirtualServiceLifecycleStateFailed,
}

// GetVirtualServiceLifecycleStateEnumValues Enumerates the set of values for VirtualServiceLifecycleStateEnum
func GetVirtualServiceLifecycleStateEnumValues() []VirtualServiceLifecycleStateEnum {
	values := make([]VirtualServiceLifecycleStateEnum, 0)
	for _, v := range mappingVirtualServiceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVirtualServiceLifecycleStateEnumStringValues Enumerates the set of values in String for VirtualServiceLifecycleStateEnum
func GetVirtualServiceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingVirtualServiceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVirtualServiceLifecycleStateEnum(val string) (VirtualServiceLifecycleStateEnum, bool) {
	enum, ok := mappingVirtualServiceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
