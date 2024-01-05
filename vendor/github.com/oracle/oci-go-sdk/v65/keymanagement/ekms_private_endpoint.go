// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EkmsPrivateEndpoint EKMS private endpoint created in customer subnet used to connect to external key manager system
type EkmsPrivateEndpoint struct {

	// Unique identifier that is immutable
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// EKMS Private Endpoint display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the EKMS private endpoint was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the EKMS private endpoint resource.
	LifecycleState EkmsPrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Private IP of the external key manager system to connect to from the EKMS private endpoint
	ExternalKeyManagerIp *string `mandatory:"true" json:"externalKeyManagerIp"`

	// The time the EKMS private endpoint was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The port of the external key manager system
	Port *int `mandatory:"false" json:"port"`

	// CABundle to validate TLS certificate of the external key manager system in PEM format
	CaBundle *string `mandatory:"false" json:"caBundle"`

	// The IP address in the customer's VCN for the EKMS private endpoint. This is taken from subnet
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`
}

func (m EkmsPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EkmsPrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEkmsPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEkmsPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EkmsPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type EkmsPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for EkmsPrivateEndpointLifecycleStateEnum
const (
	EkmsPrivateEndpointLifecycleStateCreating EkmsPrivateEndpointLifecycleStateEnum = "CREATING"
	EkmsPrivateEndpointLifecycleStateActive   EkmsPrivateEndpointLifecycleStateEnum = "ACTIVE"
	EkmsPrivateEndpointLifecycleStateDeleting EkmsPrivateEndpointLifecycleStateEnum = "DELETING"
	EkmsPrivateEndpointLifecycleStateDeleted  EkmsPrivateEndpointLifecycleStateEnum = "DELETED"
	EkmsPrivateEndpointLifecycleStateFailed   EkmsPrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingEkmsPrivateEndpointLifecycleStateEnum = map[string]EkmsPrivateEndpointLifecycleStateEnum{
	"CREATING": EkmsPrivateEndpointLifecycleStateCreating,
	"ACTIVE":   EkmsPrivateEndpointLifecycleStateActive,
	"DELETING": EkmsPrivateEndpointLifecycleStateDeleting,
	"DELETED":  EkmsPrivateEndpointLifecycleStateDeleted,
	"FAILED":   EkmsPrivateEndpointLifecycleStateFailed,
}

var mappingEkmsPrivateEndpointLifecycleStateEnumLowerCase = map[string]EkmsPrivateEndpointLifecycleStateEnum{
	"creating": EkmsPrivateEndpointLifecycleStateCreating,
	"active":   EkmsPrivateEndpointLifecycleStateActive,
	"deleting": EkmsPrivateEndpointLifecycleStateDeleting,
	"deleted":  EkmsPrivateEndpointLifecycleStateDeleted,
	"failed":   EkmsPrivateEndpointLifecycleStateFailed,
}

// GetEkmsPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for EkmsPrivateEndpointLifecycleStateEnum
func GetEkmsPrivateEndpointLifecycleStateEnumValues() []EkmsPrivateEndpointLifecycleStateEnum {
	values := make([]EkmsPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingEkmsPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEkmsPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for EkmsPrivateEndpointLifecycleStateEnum
func GetEkmsPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEkmsPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEkmsPrivateEndpointLifecycleStateEnum(val string) (EkmsPrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingEkmsPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
