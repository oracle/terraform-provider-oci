// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdpGroupMapping A mapping between a single group defined by the identity provider (IdP) you're federating with
// and a single IAM Service Group in Oracle Cloud Infrastructure.
// For more information about group mappings and what they're for, see
// Identity Providers and Federation (https://docs.cloud.oracle.com/Content/Identity/Concepts/federation.htm).
// A given IdP group can be mapped to zero, one, or multiple IAM Service groups, and vice versa.
// But each `IdPGroupMapping` object is between only a single IdP group and IAM Service group.
// Each `IdPGroupMapping` object has its own OCID.
// **Note:** Any users who are in more than 50 IdP groups cannot be authenticated to use the Oracle
// Cloud Infrastructure Console.
type IdpGroupMapping struct {

	// The OCID of the `IdpGroupMapping`.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the `IdentityProvider` this mapping belongs to.
	IdpId *string `mandatory:"true" json:"idpId"`

	// The name of the IdP group that is mapped to the IAM Service group.
	IdpGroupName *string `mandatory:"true" json:"idpGroupName"`

	// The OCID of the IAM Service group that is mapped to the IdP group.
	GroupId *string `mandatory:"true" json:"groupId"`

	// The OCID of the tenancy containing the `IdentityProvider`.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Date and time the mapping was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The mapping's current state.  After creating a mapping object, make sure its `lifecycleState` changes
	// from CREATING to ACTIVE before using it.
	LifecycleState IdpGroupMappingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m IdpGroupMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdpGroupMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIdpGroupMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIdpGroupMappingLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdpGroupMappingLifecycleStateEnum Enum with underlying type: string
type IdpGroupMappingLifecycleStateEnum string

// Set of constants representing the allowable values for IdpGroupMappingLifecycleStateEnum
const (
	IdpGroupMappingLifecycleStateCreating IdpGroupMappingLifecycleStateEnum = "CREATING"
	IdpGroupMappingLifecycleStateActive   IdpGroupMappingLifecycleStateEnum = "ACTIVE"
	IdpGroupMappingLifecycleStateInactive IdpGroupMappingLifecycleStateEnum = "INACTIVE"
	IdpGroupMappingLifecycleStateDeleting IdpGroupMappingLifecycleStateEnum = "DELETING"
	IdpGroupMappingLifecycleStateDeleted  IdpGroupMappingLifecycleStateEnum = "DELETED"
)

var mappingIdpGroupMappingLifecycleStateEnum = map[string]IdpGroupMappingLifecycleStateEnum{
	"CREATING": IdpGroupMappingLifecycleStateCreating,
	"ACTIVE":   IdpGroupMappingLifecycleStateActive,
	"INACTIVE": IdpGroupMappingLifecycleStateInactive,
	"DELETING": IdpGroupMappingLifecycleStateDeleting,
	"DELETED":  IdpGroupMappingLifecycleStateDeleted,
}

var mappingIdpGroupMappingLifecycleStateEnumLowerCase = map[string]IdpGroupMappingLifecycleStateEnum{
	"creating": IdpGroupMappingLifecycleStateCreating,
	"active":   IdpGroupMappingLifecycleStateActive,
	"inactive": IdpGroupMappingLifecycleStateInactive,
	"deleting": IdpGroupMappingLifecycleStateDeleting,
	"deleted":  IdpGroupMappingLifecycleStateDeleted,
}

// GetIdpGroupMappingLifecycleStateEnumValues Enumerates the set of values for IdpGroupMappingLifecycleStateEnum
func GetIdpGroupMappingLifecycleStateEnumValues() []IdpGroupMappingLifecycleStateEnum {
	values := make([]IdpGroupMappingLifecycleStateEnum, 0)
	for _, v := range mappingIdpGroupMappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIdpGroupMappingLifecycleStateEnumStringValues Enumerates the set of values in String for IdpGroupMappingLifecycleStateEnum
func GetIdpGroupMappingLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingIdpGroupMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdpGroupMappingLifecycleStateEnum(val string) (IdpGroupMappingLifecycleStateEnum, bool) {
	enum, ok := mappingIdpGroupMappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
