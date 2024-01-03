// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomerSecretKeySummary As the name suggests, a `CustomerSecretKeySummary` object contains information about a `CustomerSecretKey`.
// A `CustomerSecretKey` is an Oracle-provided key for using the Object Storage Service's Amazon S3 compatible API.
type CustomerSecretKeySummary struct {

	// The OCID of the secret key.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the user the password belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// The displayName you assign to the secret key. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time when this password will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The secret key's current state. After creating a secret key, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState CustomerSecretKeySummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m CustomerSecretKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomerSecretKeySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomerSecretKeySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCustomerSecretKeySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomerSecretKeySummaryLifecycleStateEnum Enum with underlying type: string
type CustomerSecretKeySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for CustomerSecretKeySummaryLifecycleStateEnum
const (
	CustomerSecretKeySummaryLifecycleStateCreating CustomerSecretKeySummaryLifecycleStateEnum = "CREATING"
	CustomerSecretKeySummaryLifecycleStateActive   CustomerSecretKeySummaryLifecycleStateEnum = "ACTIVE"
	CustomerSecretKeySummaryLifecycleStateInactive CustomerSecretKeySummaryLifecycleStateEnum = "INACTIVE"
	CustomerSecretKeySummaryLifecycleStateDeleting CustomerSecretKeySummaryLifecycleStateEnum = "DELETING"
	CustomerSecretKeySummaryLifecycleStateDeleted  CustomerSecretKeySummaryLifecycleStateEnum = "DELETED"
)

var mappingCustomerSecretKeySummaryLifecycleStateEnum = map[string]CustomerSecretKeySummaryLifecycleStateEnum{
	"CREATING": CustomerSecretKeySummaryLifecycleStateCreating,
	"ACTIVE":   CustomerSecretKeySummaryLifecycleStateActive,
	"INACTIVE": CustomerSecretKeySummaryLifecycleStateInactive,
	"DELETING": CustomerSecretKeySummaryLifecycleStateDeleting,
	"DELETED":  CustomerSecretKeySummaryLifecycleStateDeleted,
}

var mappingCustomerSecretKeySummaryLifecycleStateEnumLowerCase = map[string]CustomerSecretKeySummaryLifecycleStateEnum{
	"creating": CustomerSecretKeySummaryLifecycleStateCreating,
	"active":   CustomerSecretKeySummaryLifecycleStateActive,
	"inactive": CustomerSecretKeySummaryLifecycleStateInactive,
	"deleting": CustomerSecretKeySummaryLifecycleStateDeleting,
	"deleted":  CustomerSecretKeySummaryLifecycleStateDeleted,
}

// GetCustomerSecretKeySummaryLifecycleStateEnumValues Enumerates the set of values for CustomerSecretKeySummaryLifecycleStateEnum
func GetCustomerSecretKeySummaryLifecycleStateEnumValues() []CustomerSecretKeySummaryLifecycleStateEnum {
	values := make([]CustomerSecretKeySummaryLifecycleStateEnum, 0)
	for _, v := range mappingCustomerSecretKeySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerSecretKeySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for CustomerSecretKeySummaryLifecycleStateEnum
func GetCustomerSecretKeySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingCustomerSecretKeySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerSecretKeySummaryLifecycleStateEnum(val string) (CustomerSecretKeySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingCustomerSecretKeySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
