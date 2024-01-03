// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OdaPrivateEndpoint A private endpoint allows Digital Assistant Instance to access resources in a customer's virtual cloud network (VCN).
type OdaPrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that was assigned when the ODA private endpoint was created.
	Id *string `mandatory:"true" json:"id"`

	// User-defined name for the ODA private endpoint. Avoid entering confidential information.
	// You can change this value.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Description of the ODA private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the resource was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ODA private endpoint.
	LifecycleState OdaPrivateEndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// List of OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of network security groups (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm)
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OdaPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdaPrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOdaPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOdaPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdaPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type OdaPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for OdaPrivateEndpointLifecycleStateEnum
const (
	OdaPrivateEndpointLifecycleStateCreating OdaPrivateEndpointLifecycleStateEnum = "CREATING"
	OdaPrivateEndpointLifecycleStateUpdating OdaPrivateEndpointLifecycleStateEnum = "UPDATING"
	OdaPrivateEndpointLifecycleStateActive   OdaPrivateEndpointLifecycleStateEnum = "ACTIVE"
	OdaPrivateEndpointLifecycleStateDeleting OdaPrivateEndpointLifecycleStateEnum = "DELETING"
	OdaPrivateEndpointLifecycleStateDeleted  OdaPrivateEndpointLifecycleStateEnum = "DELETED"
	OdaPrivateEndpointLifecycleStateFailed   OdaPrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingOdaPrivateEndpointLifecycleStateEnum = map[string]OdaPrivateEndpointLifecycleStateEnum{
	"CREATING": OdaPrivateEndpointLifecycleStateCreating,
	"UPDATING": OdaPrivateEndpointLifecycleStateUpdating,
	"ACTIVE":   OdaPrivateEndpointLifecycleStateActive,
	"DELETING": OdaPrivateEndpointLifecycleStateDeleting,
	"DELETED":  OdaPrivateEndpointLifecycleStateDeleted,
	"FAILED":   OdaPrivateEndpointLifecycleStateFailed,
}

var mappingOdaPrivateEndpointLifecycleStateEnumLowerCase = map[string]OdaPrivateEndpointLifecycleStateEnum{
	"creating": OdaPrivateEndpointLifecycleStateCreating,
	"updating": OdaPrivateEndpointLifecycleStateUpdating,
	"active":   OdaPrivateEndpointLifecycleStateActive,
	"deleting": OdaPrivateEndpointLifecycleStateDeleting,
	"deleted":  OdaPrivateEndpointLifecycleStateDeleted,
	"failed":   OdaPrivateEndpointLifecycleStateFailed,
}

// GetOdaPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for OdaPrivateEndpointLifecycleStateEnum
func GetOdaPrivateEndpointLifecycleStateEnumValues() []OdaPrivateEndpointLifecycleStateEnum {
	values := make([]OdaPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingOdaPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for OdaPrivateEndpointLifecycleStateEnum
func GetOdaPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOdaPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaPrivateEndpointLifecycleStateEnum(val string) (OdaPrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingOdaPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
