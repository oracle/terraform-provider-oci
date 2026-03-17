// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedDatabasePrivateEndpoint DistributedDatabasePrivateEndpoint resource.
type DistributedDatabasePrivateEndpoint struct {

	// The identifier of the Private Endpoint.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment in which private endpoint exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Identifier of the subnet in which private endpoint exists.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Identifier of the VCN in which subnet exists.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// DistributedDatabasePrivateEndpoint display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the DistributedDatabasePrivateEndpoint was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Private Endpoint was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Lifecycle states for private endpoint.
	LifecycleState DistributedDatabasePrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// DistributedDatabasePrivateEndpoint description.
	Description *string `mandatory:"false" json:"description"`

	// IP address of the Private Endpoint.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// This field is deprecated. Support for this field will be removed after Mon, 1 Mar 2027 00:00:00 GMT.
	GloballyDistributedDatabases []DistributedDatabaseAssociatedWithPrivateEndpoint `mandatory:"false" json:"globallyDistributedDatabases"`

	// This field is deprecated. Support for this field will be removed after Mon, 1 Mar 2027 00:00:00 GMT.
	GloballyDistributedAutonomousDatabases []DistributedAutonomousDatabaseAssociatedWithPrivateEndpoint `mandatory:"false" json:"globallyDistributedAutonomousDatabases"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

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

func (m DistributedDatabasePrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabasePrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabasePrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDistributedDatabasePrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabasePrivateEndpointLifecycleStateEnum Enum with underlying type: string
type DistributedDatabasePrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for DistributedDatabasePrivateEndpointLifecycleStateEnum
const (
	DistributedDatabasePrivateEndpointLifecycleStateActive   DistributedDatabasePrivateEndpointLifecycleStateEnum = "ACTIVE"
	DistributedDatabasePrivateEndpointLifecycleStateFailed   DistributedDatabasePrivateEndpointLifecycleStateEnum = "FAILED"
	DistributedDatabasePrivateEndpointLifecycleStateInactive DistributedDatabasePrivateEndpointLifecycleStateEnum = "INACTIVE"
	DistributedDatabasePrivateEndpointLifecycleStateDeleting DistributedDatabasePrivateEndpointLifecycleStateEnum = "DELETING"
	DistributedDatabasePrivateEndpointLifecycleStateDeleted  DistributedDatabasePrivateEndpointLifecycleStateEnum = "DELETED"
	DistributedDatabasePrivateEndpointLifecycleStateUpdating DistributedDatabasePrivateEndpointLifecycleStateEnum = "UPDATING"
	DistributedDatabasePrivateEndpointLifecycleStateCreating DistributedDatabasePrivateEndpointLifecycleStateEnum = "CREATING"
)

var mappingDistributedDatabasePrivateEndpointLifecycleStateEnum = map[string]DistributedDatabasePrivateEndpointLifecycleStateEnum{
	"ACTIVE":   DistributedDatabasePrivateEndpointLifecycleStateActive,
	"FAILED":   DistributedDatabasePrivateEndpointLifecycleStateFailed,
	"INACTIVE": DistributedDatabasePrivateEndpointLifecycleStateInactive,
	"DELETING": DistributedDatabasePrivateEndpointLifecycleStateDeleting,
	"DELETED":  DistributedDatabasePrivateEndpointLifecycleStateDeleted,
	"UPDATING": DistributedDatabasePrivateEndpointLifecycleStateUpdating,
	"CREATING": DistributedDatabasePrivateEndpointLifecycleStateCreating,
}

var mappingDistributedDatabasePrivateEndpointLifecycleStateEnumLowerCase = map[string]DistributedDatabasePrivateEndpointLifecycleStateEnum{
	"active":   DistributedDatabasePrivateEndpointLifecycleStateActive,
	"failed":   DistributedDatabasePrivateEndpointLifecycleStateFailed,
	"inactive": DistributedDatabasePrivateEndpointLifecycleStateInactive,
	"deleting": DistributedDatabasePrivateEndpointLifecycleStateDeleting,
	"deleted":  DistributedDatabasePrivateEndpointLifecycleStateDeleted,
	"updating": DistributedDatabasePrivateEndpointLifecycleStateUpdating,
	"creating": DistributedDatabasePrivateEndpointLifecycleStateCreating,
}

// GetDistributedDatabasePrivateEndpointLifecycleStateEnumValues Enumerates the set of values for DistributedDatabasePrivateEndpointLifecycleStateEnum
func GetDistributedDatabasePrivateEndpointLifecycleStateEnumValues() []DistributedDatabasePrivateEndpointLifecycleStateEnum {
	values := make([]DistributedDatabasePrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingDistributedDatabasePrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabasePrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for DistributedDatabasePrivateEndpointLifecycleStateEnum
func GetDistributedDatabasePrivateEndpointLifecycleStateEnumStringValues() []string {
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

// GetMappingDistributedDatabasePrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabasePrivateEndpointLifecycleStateEnum(val string) (DistributedDatabasePrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingDistributedDatabasePrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
