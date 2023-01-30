// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Endpoint An endpoint is an organizational construct to keep multiple data connectivity management solutions and their resources (pe-id, dnsProxyIp, dnsZones, and so on) separate from each other, helping you to stay organized. For example, you could have separate registries for development, testing, and production.
type Endpoint struct {

	// VCN OCID where the subnet resides.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Subnet OCID of the customer connected network where, for example, the databases reside.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Data Connectivity Management Registry display name; registries can be renamed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// List of DNS zones to be used by the data assets to be harvested.
	// Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com
	DnsZones []string `mandatory:"false" json:"dnsZones"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Registry description
	Description *string `mandatory:"false" json:"description"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Time when the Data Connectivity Management registry was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Data Connectivity Management registry was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Lifecycle states for registries in the Data Connectivity Management Service.
	// CREATING - The resource is being created and may not be usable until the entire metadata is defined.
	// UPDATING - The resource is being updated and may not be usable until all changes are commited.
	// DELETING - The resource is being deleted and might require deep cleanup of children.
	// ACTIVE   - The resource is valid and available for access.
	// INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for
	//          administrative reasons.
	// DELETED  - The resource has been deleted and isn't available.
	// FAILED   - The resource is in a failed state due to validation or other errors.
	LifecycleState EndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Endpoint size for reverse connection capacity.
	EndpointSize *int `mandatory:"false" json:"endpointSize"`

	// The list of NSGs to which the private endpoint VNIC must be added.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m Endpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Endpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EndpointLifecycleStateEnum Enum with underlying type: string
type EndpointLifecycleStateEnum string

// Set of constants representing the allowable values for EndpointLifecycleStateEnum
const (
	EndpointLifecycleStateCreating EndpointLifecycleStateEnum = "CREATING"
	EndpointLifecycleStateActive   EndpointLifecycleStateEnum = "ACTIVE"
	EndpointLifecycleStateInactive EndpointLifecycleStateEnum = "INACTIVE"
	EndpointLifecycleStateUpdating EndpointLifecycleStateEnum = "UPDATING"
	EndpointLifecycleStateDeleting EndpointLifecycleStateEnum = "DELETING"
	EndpointLifecycleStateDeleted  EndpointLifecycleStateEnum = "DELETED"
	EndpointLifecycleStateFailed   EndpointLifecycleStateEnum = "FAILED"
)

var mappingEndpointLifecycleStateEnum = map[string]EndpointLifecycleStateEnum{
	"CREATING": EndpointLifecycleStateCreating,
	"ACTIVE":   EndpointLifecycleStateActive,
	"INACTIVE": EndpointLifecycleStateInactive,
	"UPDATING": EndpointLifecycleStateUpdating,
	"DELETING": EndpointLifecycleStateDeleting,
	"DELETED":  EndpointLifecycleStateDeleted,
	"FAILED":   EndpointLifecycleStateFailed,
}

var mappingEndpointLifecycleStateEnumLowerCase = map[string]EndpointLifecycleStateEnum{
	"creating": EndpointLifecycleStateCreating,
	"active":   EndpointLifecycleStateActive,
	"inactive": EndpointLifecycleStateInactive,
	"updating": EndpointLifecycleStateUpdating,
	"deleting": EndpointLifecycleStateDeleting,
	"deleted":  EndpointLifecycleStateDeleted,
	"failed":   EndpointLifecycleStateFailed,
}

// GetEndpointLifecycleStateEnumValues Enumerates the set of values for EndpointLifecycleStateEnum
func GetEndpointLifecycleStateEnumValues() []EndpointLifecycleStateEnum {
	values := make([]EndpointLifecycleStateEnum, 0)
	for _, v := range mappingEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for EndpointLifecycleStateEnum
func GetEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEndpointLifecycleStateEnum(val string) (EndpointLifecycleStateEnum, bool) {
	enum, ok := mappingEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
