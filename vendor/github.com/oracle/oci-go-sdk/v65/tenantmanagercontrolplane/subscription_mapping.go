// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscriptionMapping Subscription mapping information.
type SubscriptionMapping struct {

	// OCID of the mapping between subscription and compartment identified by the tenancy.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the subscription.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// OCID of the compartment. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Denotes if the subscription is explicity assigned to the root compartment or tenancy.
	IsExplicitlyAssigned *bool `mandatory:"true" json:"isExplicitlyAssigned"`

	// Lifecycle state of the subscriptionMapping.
	LifecycleState SubscriptionMappingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Date-time when subscription mapping was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date-time when subscription mapping was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Date-time when subscription mapping was terminated.
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`
}

func (m SubscriptionMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubscriptionMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubscriptionMappingLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubscriptionMappingLifecycleStateEnum Enum with underlying type: string
type SubscriptionMappingLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionMappingLifecycleStateEnum
const (
	SubscriptionMappingLifecycleStateCreating SubscriptionMappingLifecycleStateEnum = "CREATING"
	SubscriptionMappingLifecycleStateActive   SubscriptionMappingLifecycleStateEnum = "ACTIVE"
	SubscriptionMappingLifecycleStateInactive SubscriptionMappingLifecycleStateEnum = "INACTIVE"
	SubscriptionMappingLifecycleStateUpdating SubscriptionMappingLifecycleStateEnum = "UPDATING"
	SubscriptionMappingLifecycleStateDeleting SubscriptionMappingLifecycleStateEnum = "DELETING"
	SubscriptionMappingLifecycleStateDeleted  SubscriptionMappingLifecycleStateEnum = "DELETED"
	SubscriptionMappingLifecycleStateFailed   SubscriptionMappingLifecycleStateEnum = "FAILED"
)

var mappingSubscriptionMappingLifecycleStateEnum = map[string]SubscriptionMappingLifecycleStateEnum{
	"CREATING": SubscriptionMappingLifecycleStateCreating,
	"ACTIVE":   SubscriptionMappingLifecycleStateActive,
	"INACTIVE": SubscriptionMappingLifecycleStateInactive,
	"UPDATING": SubscriptionMappingLifecycleStateUpdating,
	"DELETING": SubscriptionMappingLifecycleStateDeleting,
	"DELETED":  SubscriptionMappingLifecycleStateDeleted,
	"FAILED":   SubscriptionMappingLifecycleStateFailed,
}

var mappingSubscriptionMappingLifecycleStateEnumLowerCase = map[string]SubscriptionMappingLifecycleStateEnum{
	"creating": SubscriptionMappingLifecycleStateCreating,
	"active":   SubscriptionMappingLifecycleStateActive,
	"inactive": SubscriptionMappingLifecycleStateInactive,
	"updating": SubscriptionMappingLifecycleStateUpdating,
	"deleting": SubscriptionMappingLifecycleStateDeleting,
	"deleted":  SubscriptionMappingLifecycleStateDeleted,
	"failed":   SubscriptionMappingLifecycleStateFailed,
}

// GetSubscriptionMappingLifecycleStateEnumValues Enumerates the set of values for SubscriptionMappingLifecycleStateEnum
func GetSubscriptionMappingLifecycleStateEnumValues() []SubscriptionMappingLifecycleStateEnum {
	values := make([]SubscriptionMappingLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionMappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionMappingLifecycleStateEnumStringValues Enumerates the set of values in String for SubscriptionMappingLifecycleStateEnum
func GetSubscriptionMappingLifecycleStateEnumStringValues() []string {
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

// GetMappingSubscriptionMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionMappingLifecycleStateEnum(val string) (SubscriptionMappingLifecycleStateEnum, bool) {
	enum, ok := mappingSubscriptionMappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
