// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpoint PrivateEndpoint resource.
type PrivateEndpoint struct {

	// The identifier of the Private Endpoint.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment in which private endpoint exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Identifier of the subnet in which private endpoint exists.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Identifier of the VCN in which subnet exists.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// PrivateEndpoint display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the PrivateEndpoint was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Private Endpoint was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Lifecycle states for private endpoint.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// PrivateEndpoint description.
	Description *string `mandatory:"false" json:"description"`

	// IP address of the Private Endpoint.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCIDs of sharded databases that consumes the given private endpoint.
	ShardedDatabases []string `mandatory:"false" json:"shardedDatabases"`

	// Detailed message for the lifecycle state.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// The identifier of the proxy compute instance.
	ProxyComputeInstanceId *string `mandatory:"false" json:"proxyComputeInstanceId"`

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

func (m PrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateEndpointLifecycleStateEnum Enum with underlying type: string
type PrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateEndpointLifecycleStateEnum
const (
	PrivateEndpointLifecycleStateActive   PrivateEndpointLifecycleStateEnum = "ACTIVE"
	PrivateEndpointLifecycleStateFailed   PrivateEndpointLifecycleStateEnum = "FAILED"
	PrivateEndpointLifecycleStateInactive PrivateEndpointLifecycleStateEnum = "INACTIVE"
	PrivateEndpointLifecycleStateDeleting PrivateEndpointLifecycleStateEnum = "DELETING"
	PrivateEndpointLifecycleStateDeleted  PrivateEndpointLifecycleStateEnum = "DELETED"
	PrivateEndpointLifecycleStateUpdating PrivateEndpointLifecycleStateEnum = "UPDATING"
	PrivateEndpointLifecycleStateCreating PrivateEndpointLifecycleStateEnum = "CREATING"
)

var mappingPrivateEndpointLifecycleStateEnum = map[string]PrivateEndpointLifecycleStateEnum{
	"ACTIVE":   PrivateEndpointLifecycleStateActive,
	"FAILED":   PrivateEndpointLifecycleStateFailed,
	"INACTIVE": PrivateEndpointLifecycleStateInactive,
	"DELETING": PrivateEndpointLifecycleStateDeleting,
	"DELETED":  PrivateEndpointLifecycleStateDeleted,
	"UPDATING": PrivateEndpointLifecycleStateUpdating,
	"CREATING": PrivateEndpointLifecycleStateCreating,
}

var mappingPrivateEndpointLifecycleStateEnumLowerCase = map[string]PrivateEndpointLifecycleStateEnum{
	"active":   PrivateEndpointLifecycleStateActive,
	"failed":   PrivateEndpointLifecycleStateFailed,
	"inactive": PrivateEndpointLifecycleStateInactive,
	"deleting": PrivateEndpointLifecycleStateDeleting,
	"deleted":  PrivateEndpointLifecycleStateDeleted,
	"updating": PrivateEndpointLifecycleStateUpdating,
	"creating": PrivateEndpointLifecycleStateCreating,
}

// GetPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumValues() []PrivateEndpointLifecycleStateEnum {
	values := make([]PrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for PrivateEndpointLifecycleStateEnum
func GetPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
	}
}

// GetMappingPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateEndpointLifecycleStateEnum(val string) (PrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
