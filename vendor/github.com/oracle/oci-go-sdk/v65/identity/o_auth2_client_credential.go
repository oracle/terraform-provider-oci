// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OAuth2ClientCredential User can define Oauth clients in IAM, then use it to generate a token to grant access to app resources.
type OAuth2ClientCredential struct {

	// Allowed scopes for the given oauth credential.
	Scopes []FullyQualifiedScope `mandatory:"false" json:"scopes"`

	// Returned during create and update with password reset requests.
	Password *string `mandatory:"false" json:"password"`

	// The OCID of the user the Oauth credential belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// Date and time when this credential will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	ExpiresOn *common.SDKTime `mandatory:"false" json:"expiresOn"`

	// The OCID of the Oauth credential.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the compartment containing the Oauth credential.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the Oauth credential.
	Name *string `mandatory:"false" json:"name"`

	// The description of the Oauth credential.
	Description *string `mandatory:"false" json:"description"`

	// The credential's current state. After creating a Oauth credential, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState OAuth2ClientCredentialLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Date and time the `OAuth2ClientCredential` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m OAuth2ClientCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuth2ClientCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOAuth2ClientCredentialLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOAuth2ClientCredentialLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OAuth2ClientCredentialLifecycleStateEnum Enum with underlying type: string
type OAuth2ClientCredentialLifecycleStateEnum string

// Set of constants representing the allowable values for OAuth2ClientCredentialLifecycleStateEnum
const (
	OAuth2ClientCredentialLifecycleStateCreating OAuth2ClientCredentialLifecycleStateEnum = "CREATING"
	OAuth2ClientCredentialLifecycleStateActive   OAuth2ClientCredentialLifecycleStateEnum = "ACTIVE"
	OAuth2ClientCredentialLifecycleStateInactive OAuth2ClientCredentialLifecycleStateEnum = "INACTIVE"
	OAuth2ClientCredentialLifecycleStateDeleting OAuth2ClientCredentialLifecycleStateEnum = "DELETING"
	OAuth2ClientCredentialLifecycleStateDeleted  OAuth2ClientCredentialLifecycleStateEnum = "DELETED"
)

var mappingOAuth2ClientCredentialLifecycleStateEnum = map[string]OAuth2ClientCredentialLifecycleStateEnum{
	"CREATING": OAuth2ClientCredentialLifecycleStateCreating,
	"ACTIVE":   OAuth2ClientCredentialLifecycleStateActive,
	"INACTIVE": OAuth2ClientCredentialLifecycleStateInactive,
	"DELETING": OAuth2ClientCredentialLifecycleStateDeleting,
	"DELETED":  OAuth2ClientCredentialLifecycleStateDeleted,
}

var mappingOAuth2ClientCredentialLifecycleStateEnumLowerCase = map[string]OAuth2ClientCredentialLifecycleStateEnum{
	"creating": OAuth2ClientCredentialLifecycleStateCreating,
	"active":   OAuth2ClientCredentialLifecycleStateActive,
	"inactive": OAuth2ClientCredentialLifecycleStateInactive,
	"deleting": OAuth2ClientCredentialLifecycleStateDeleting,
	"deleted":  OAuth2ClientCredentialLifecycleStateDeleted,
}

// GetOAuth2ClientCredentialLifecycleStateEnumValues Enumerates the set of values for OAuth2ClientCredentialLifecycleStateEnum
func GetOAuth2ClientCredentialLifecycleStateEnumValues() []OAuth2ClientCredentialLifecycleStateEnum {
	values := make([]OAuth2ClientCredentialLifecycleStateEnum, 0)
	for _, v := range mappingOAuth2ClientCredentialLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOAuth2ClientCredentialLifecycleStateEnumStringValues Enumerates the set of values in String for OAuth2ClientCredentialLifecycleStateEnum
func GetOAuth2ClientCredentialLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingOAuth2ClientCredentialLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOAuth2ClientCredentialLifecycleStateEnum(val string) (OAuth2ClientCredentialLifecycleStateEnum, bool) {
	enum, ok := mappingOAuth2ClientCredentialLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
