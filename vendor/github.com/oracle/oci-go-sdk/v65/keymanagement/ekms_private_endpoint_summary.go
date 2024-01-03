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

// EkmsPrivateEndpointSummary EKMS private endpoints summary
type EkmsPrivateEndpointSummary struct {

	// Unique identifier that is immutable
	Id *string `mandatory:"true" json:"id"`

	// Subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Identifier of the compartment this EKMS private endpoint belongs to
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the EKMS private endpoint was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Mutable name of the EKMS private endpoint
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the EKMS private endpoint resource.
	LifecycleState EkmsPrivateEndpointSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the EKMS private endpoint was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m EkmsPrivateEndpointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EkmsPrivateEndpointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEkmsPrivateEndpointSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEkmsPrivateEndpointSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EkmsPrivateEndpointSummaryLifecycleStateEnum Enum with underlying type: string
type EkmsPrivateEndpointSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for EkmsPrivateEndpointSummaryLifecycleStateEnum
const (
	EkmsPrivateEndpointSummaryLifecycleStateCreating EkmsPrivateEndpointSummaryLifecycleStateEnum = "CREATING"
	EkmsPrivateEndpointSummaryLifecycleStateActive   EkmsPrivateEndpointSummaryLifecycleStateEnum = "ACTIVE"
	EkmsPrivateEndpointSummaryLifecycleStateDeleting EkmsPrivateEndpointSummaryLifecycleStateEnum = "DELETING"
	EkmsPrivateEndpointSummaryLifecycleStateDeleted  EkmsPrivateEndpointSummaryLifecycleStateEnum = "DELETED"
	EkmsPrivateEndpointSummaryLifecycleStateFailed   EkmsPrivateEndpointSummaryLifecycleStateEnum = "FAILED"
)

var mappingEkmsPrivateEndpointSummaryLifecycleStateEnum = map[string]EkmsPrivateEndpointSummaryLifecycleStateEnum{
	"CREATING": EkmsPrivateEndpointSummaryLifecycleStateCreating,
	"ACTIVE":   EkmsPrivateEndpointSummaryLifecycleStateActive,
	"DELETING": EkmsPrivateEndpointSummaryLifecycleStateDeleting,
	"DELETED":  EkmsPrivateEndpointSummaryLifecycleStateDeleted,
	"FAILED":   EkmsPrivateEndpointSummaryLifecycleStateFailed,
}

var mappingEkmsPrivateEndpointSummaryLifecycleStateEnumLowerCase = map[string]EkmsPrivateEndpointSummaryLifecycleStateEnum{
	"creating": EkmsPrivateEndpointSummaryLifecycleStateCreating,
	"active":   EkmsPrivateEndpointSummaryLifecycleStateActive,
	"deleting": EkmsPrivateEndpointSummaryLifecycleStateDeleting,
	"deleted":  EkmsPrivateEndpointSummaryLifecycleStateDeleted,
	"failed":   EkmsPrivateEndpointSummaryLifecycleStateFailed,
}

// GetEkmsPrivateEndpointSummaryLifecycleStateEnumValues Enumerates the set of values for EkmsPrivateEndpointSummaryLifecycleStateEnum
func GetEkmsPrivateEndpointSummaryLifecycleStateEnumValues() []EkmsPrivateEndpointSummaryLifecycleStateEnum {
	values := make([]EkmsPrivateEndpointSummaryLifecycleStateEnum, 0)
	for _, v := range mappingEkmsPrivateEndpointSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEkmsPrivateEndpointSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for EkmsPrivateEndpointSummaryLifecycleStateEnum
func GetEkmsPrivateEndpointSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEkmsPrivateEndpointSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEkmsPrivateEndpointSummaryLifecycleStateEnum(val string) (EkmsPrivateEndpointSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingEkmsPrivateEndpointSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
