// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsagePlan A usage plan controls access of subscribers to deployments, controlling rate limits and quotas for usage.
type UsagePlan struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a usage plan
	// resource.
	Id *string `mandatory:"true" json:"id"`

	// A collection of entitlements currently assigned to the usage plan.
	Entitlements []Entitlement `mandatory:"true" json:"entitlements"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the usage plan.
	LifecycleState UsagePlanLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UsagePlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsagePlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUsagePlanLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUsagePlanLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UsagePlanLifecycleStateEnum Enum with underlying type: string
type UsagePlanLifecycleStateEnum string

// Set of constants representing the allowable values for UsagePlanLifecycleStateEnum
const (
	UsagePlanLifecycleStateCreating UsagePlanLifecycleStateEnum = "CREATING"
	UsagePlanLifecycleStateActive   UsagePlanLifecycleStateEnum = "ACTIVE"
	UsagePlanLifecycleStateUpdating UsagePlanLifecycleStateEnum = "UPDATING"
	UsagePlanLifecycleStateDeleting UsagePlanLifecycleStateEnum = "DELETING"
	UsagePlanLifecycleStateDeleted  UsagePlanLifecycleStateEnum = "DELETED"
	UsagePlanLifecycleStateFailed   UsagePlanLifecycleStateEnum = "FAILED"
)

var mappingUsagePlanLifecycleStateEnum = map[string]UsagePlanLifecycleStateEnum{
	"CREATING": UsagePlanLifecycleStateCreating,
	"ACTIVE":   UsagePlanLifecycleStateActive,
	"UPDATING": UsagePlanLifecycleStateUpdating,
	"DELETING": UsagePlanLifecycleStateDeleting,
	"DELETED":  UsagePlanLifecycleStateDeleted,
	"FAILED":   UsagePlanLifecycleStateFailed,
}

var mappingUsagePlanLifecycleStateEnumLowerCase = map[string]UsagePlanLifecycleStateEnum{
	"creating": UsagePlanLifecycleStateCreating,
	"active":   UsagePlanLifecycleStateActive,
	"updating": UsagePlanLifecycleStateUpdating,
	"deleting": UsagePlanLifecycleStateDeleting,
	"deleted":  UsagePlanLifecycleStateDeleted,
	"failed":   UsagePlanLifecycleStateFailed,
}

// GetUsagePlanLifecycleStateEnumValues Enumerates the set of values for UsagePlanLifecycleStateEnum
func GetUsagePlanLifecycleStateEnumValues() []UsagePlanLifecycleStateEnum {
	values := make([]UsagePlanLifecycleStateEnum, 0)
	for _, v := range mappingUsagePlanLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUsagePlanLifecycleStateEnumStringValues Enumerates the set of values in String for UsagePlanLifecycleStateEnum
func GetUsagePlanLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingUsagePlanLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsagePlanLifecycleStateEnum(val string) (UsagePlanLifecycleStateEnum, bool) {
	enum, ok := mappingUsagePlanLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
