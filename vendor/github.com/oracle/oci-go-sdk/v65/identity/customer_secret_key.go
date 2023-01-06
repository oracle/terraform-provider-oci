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

// CustomerSecretKey A `CustomerSecretKey` is an Oracle-provided key for using the Object Storage Service's
// Amazon S3 compatible API (https://docs.cloud.oracle.com/Content/Object/Tasks/s3compatibleapi.htm). The key consists of a
// secret key/access key pair. A user can have up to two secret keys at a time.
// **Note:** The secret key is always an Oracle-generated string; you can't change it to a string of your choice.
// For more information, see Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/access/managing-user-credentials.htm).
type CustomerSecretKey struct {

	// The secret key.
	Key *string `mandatory:"false" json:"key"`

	// The access key portion of the key pair.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the user the password belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// The display name you assign to the secret key. Does not have to be unique, and it's changeable.
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
	LifecycleState CustomerSecretKeyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m CustomerSecretKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomerSecretKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomerSecretKeyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCustomerSecretKeyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomerSecretKeyLifecycleStateEnum Enum with underlying type: string
type CustomerSecretKeyLifecycleStateEnum string

// Set of constants representing the allowable values for CustomerSecretKeyLifecycleStateEnum
const (
	CustomerSecretKeyLifecycleStateCreating CustomerSecretKeyLifecycleStateEnum = "CREATING"
	CustomerSecretKeyLifecycleStateActive   CustomerSecretKeyLifecycleStateEnum = "ACTIVE"
	CustomerSecretKeyLifecycleStateInactive CustomerSecretKeyLifecycleStateEnum = "INACTIVE"
	CustomerSecretKeyLifecycleStateDeleting CustomerSecretKeyLifecycleStateEnum = "DELETING"
	CustomerSecretKeyLifecycleStateDeleted  CustomerSecretKeyLifecycleStateEnum = "DELETED"
)

var mappingCustomerSecretKeyLifecycleStateEnum = map[string]CustomerSecretKeyLifecycleStateEnum{
	"CREATING": CustomerSecretKeyLifecycleStateCreating,
	"ACTIVE":   CustomerSecretKeyLifecycleStateActive,
	"INACTIVE": CustomerSecretKeyLifecycleStateInactive,
	"DELETING": CustomerSecretKeyLifecycleStateDeleting,
	"DELETED":  CustomerSecretKeyLifecycleStateDeleted,
}

var mappingCustomerSecretKeyLifecycleStateEnumLowerCase = map[string]CustomerSecretKeyLifecycleStateEnum{
	"creating": CustomerSecretKeyLifecycleStateCreating,
	"active":   CustomerSecretKeyLifecycleStateActive,
	"inactive": CustomerSecretKeyLifecycleStateInactive,
	"deleting": CustomerSecretKeyLifecycleStateDeleting,
	"deleted":  CustomerSecretKeyLifecycleStateDeleted,
}

// GetCustomerSecretKeyLifecycleStateEnumValues Enumerates the set of values for CustomerSecretKeyLifecycleStateEnum
func GetCustomerSecretKeyLifecycleStateEnumValues() []CustomerSecretKeyLifecycleStateEnum {
	values := make([]CustomerSecretKeyLifecycleStateEnum, 0)
	for _, v := range mappingCustomerSecretKeyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomerSecretKeyLifecycleStateEnumStringValues Enumerates the set of values in String for CustomerSecretKeyLifecycleStateEnum
func GetCustomerSecretKeyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingCustomerSecretKeyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomerSecretKeyLifecycleStateEnum(val string) (CustomerSecretKeyLifecycleStateEnum, bool) {
	enum, ok := mappingCustomerSecretKeyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
