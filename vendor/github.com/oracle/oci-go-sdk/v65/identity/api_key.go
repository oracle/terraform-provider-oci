// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ApiKey A PEM-format RSA credential for securing requests to the Oracle Cloud Infrastructure REST API. Also known
// as an *API signing key*. Specifically, this is the public key from the key pair. The private key remains with
// the user calling the API. For information about generating a key pair
// in the required PEM format, see Required Keys and OCIDs (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm).
// **Important:** This is **not** the SSH key for accessing compute instances.
// Each user can have a maximum of three API signing keys.
// For more information about user credentials, see User Credentials (https://docs.cloud.oracle.com/Content/Identity/Concepts/usercredentials.htm).
type ApiKey struct {

	// An Oracle-assigned identifier for the key, in this format:
	// TENANCY_OCID/USER_OCID/KEY_FINGERPRINT.
	KeyId *string `mandatory:"false" json:"keyId"`

	// The key's value.
	KeyValue *string `mandatory:"false" json:"keyValue"`

	// The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
	Fingerprint *string `mandatory:"false" json:"fingerprint"`

	// The OCID of the user the key belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// Date and time the `ApiKey` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The API key's current state. After creating an `ApiKey` object, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState ApiKeyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m ApiKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApiKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingApiKeyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApiKeyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApiKeyLifecycleStateEnum Enum with underlying type: string
type ApiKeyLifecycleStateEnum string

// Set of constants representing the allowable values for ApiKeyLifecycleStateEnum
const (
	ApiKeyLifecycleStateCreating ApiKeyLifecycleStateEnum = "CREATING"
	ApiKeyLifecycleStateActive   ApiKeyLifecycleStateEnum = "ACTIVE"
	ApiKeyLifecycleStateInactive ApiKeyLifecycleStateEnum = "INACTIVE"
	ApiKeyLifecycleStateDeleting ApiKeyLifecycleStateEnum = "DELETING"
	ApiKeyLifecycleStateDeleted  ApiKeyLifecycleStateEnum = "DELETED"
)

var mappingApiKeyLifecycleStateEnum = map[string]ApiKeyLifecycleStateEnum{
	"CREATING": ApiKeyLifecycleStateCreating,
	"ACTIVE":   ApiKeyLifecycleStateActive,
	"INACTIVE": ApiKeyLifecycleStateInactive,
	"DELETING": ApiKeyLifecycleStateDeleting,
	"DELETED":  ApiKeyLifecycleStateDeleted,
}

var mappingApiKeyLifecycleStateEnumLowerCase = map[string]ApiKeyLifecycleStateEnum{
	"creating": ApiKeyLifecycleStateCreating,
	"active":   ApiKeyLifecycleStateActive,
	"inactive": ApiKeyLifecycleStateInactive,
	"deleting": ApiKeyLifecycleStateDeleting,
	"deleted":  ApiKeyLifecycleStateDeleted,
}

// GetApiKeyLifecycleStateEnumValues Enumerates the set of values for ApiKeyLifecycleStateEnum
func GetApiKeyLifecycleStateEnumValues() []ApiKeyLifecycleStateEnum {
	values := make([]ApiKeyLifecycleStateEnum, 0)
	for _, v := range mappingApiKeyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApiKeyLifecycleStateEnumStringValues Enumerates the set of values in String for ApiKeyLifecycleStateEnum
func GetApiKeyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingApiKeyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiKeyLifecycleStateEnum(val string) (ApiKeyLifecycleStateEnum, bool) {
	enum, ok := mappingApiKeyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
