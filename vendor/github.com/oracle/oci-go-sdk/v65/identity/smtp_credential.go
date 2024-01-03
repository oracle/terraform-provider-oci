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

// SmtpCredential Simple Mail Transfer Protocol (SMTP) credentials are needed to send email through Email Delivery.
// The SMTP credentials are used for SMTP authentication with the service. The credentials never expire.
// A user can have up to 2 SMTP credentials at a time.
// **Note:** The credential set is always an Oracle-generated SMTP user name and password pair;
// you cannot designate the SMTP user name or the SMTP password.
// For more information, see Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/access/managing-user-credentials.htm#SMTP).
type SmtpCredential struct {

	// The SMTP user name.
	Username *string `mandatory:"false" json:"username"`

	// The SMTP password.
	Password *string `mandatory:"false" json:"password"`

	// The OCID of the SMTP credential.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the user the SMTP credential belongs to.
	UserId *string `mandatory:"false" json:"userId"`

	// The description you assign to the SMTP credential. Does not have to be unique, and it's changeable.
	// (For tenancies that support identity domains) You can have an empty description.
	Description *string `mandatory:"false" json:"description"`

	// Date and time the `SmtpCredential` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time when this credential will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// The credential's current state. After creating a SMTP credential, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState SmtpCredentialLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`
}

func (m SmtpCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SmtpCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSmtpCredentialLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSmtpCredentialLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SmtpCredentialLifecycleStateEnum Enum with underlying type: string
type SmtpCredentialLifecycleStateEnum string

// Set of constants representing the allowable values for SmtpCredentialLifecycleStateEnum
const (
	SmtpCredentialLifecycleStateCreating SmtpCredentialLifecycleStateEnum = "CREATING"
	SmtpCredentialLifecycleStateActive   SmtpCredentialLifecycleStateEnum = "ACTIVE"
	SmtpCredentialLifecycleStateInactive SmtpCredentialLifecycleStateEnum = "INACTIVE"
	SmtpCredentialLifecycleStateDeleting SmtpCredentialLifecycleStateEnum = "DELETING"
	SmtpCredentialLifecycleStateDeleted  SmtpCredentialLifecycleStateEnum = "DELETED"
)

var mappingSmtpCredentialLifecycleStateEnum = map[string]SmtpCredentialLifecycleStateEnum{
	"CREATING": SmtpCredentialLifecycleStateCreating,
	"ACTIVE":   SmtpCredentialLifecycleStateActive,
	"INACTIVE": SmtpCredentialLifecycleStateInactive,
	"DELETING": SmtpCredentialLifecycleStateDeleting,
	"DELETED":  SmtpCredentialLifecycleStateDeleted,
}

var mappingSmtpCredentialLifecycleStateEnumLowerCase = map[string]SmtpCredentialLifecycleStateEnum{
	"creating": SmtpCredentialLifecycleStateCreating,
	"active":   SmtpCredentialLifecycleStateActive,
	"inactive": SmtpCredentialLifecycleStateInactive,
	"deleting": SmtpCredentialLifecycleStateDeleting,
	"deleted":  SmtpCredentialLifecycleStateDeleted,
}

// GetSmtpCredentialLifecycleStateEnumValues Enumerates the set of values for SmtpCredentialLifecycleStateEnum
func GetSmtpCredentialLifecycleStateEnumValues() []SmtpCredentialLifecycleStateEnum {
	values := make([]SmtpCredentialLifecycleStateEnum, 0)
	for _, v := range mappingSmtpCredentialLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSmtpCredentialLifecycleStateEnumStringValues Enumerates the set of values in String for SmtpCredentialLifecycleStateEnum
func GetSmtpCredentialLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSmtpCredentialLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSmtpCredentialLifecycleStateEnum(val string) (SmtpCredentialLifecycleStateEnum, bool) {
	enum, ok := mappingSmtpCredentialLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
