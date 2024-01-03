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

// User An individual employee or system that needs to manage or use your company's Oracle Cloud Infrastructure
// resources. Users might need to launch instances, manage remote disks, work with your cloud network, etc. Users
// have one or more IAM Service credentials (ApiKey,
// UIPassword, SwiftPassword and
// AuthToken).
// For more information, see User Credentials (https://docs.cloud.oracle.com/Content/Identity/usercred/usercredentials.htm)). End users of your
// application are not typically IAM Service users, but for tenancies that have identity domains, they might be.
// For conceptual information about users and other IAM Service components, see Overview of IAM (https://docs.cloud.oracle.com/Content/Identity/getstarted/identity-domains.htm).
// These users are created directly within the Oracle Cloud Infrastructure system, via the IAM service.
// They are different from *federated users*, who authenticate themselves to the Oracle Cloud Infrastructure
// Console via an identity provider. For more information, see
// Identity Providers and Federation (https://docs.cloud.oracle.com/Content/Identity/Concepts/federation.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Get Started with Policies (https://docs.cloud.oracle.com/Content/Identity/policiesgs/get-started-with-policies.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values
// using the API.
type User struct {

	// The OCID of the user.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy containing the user.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the user during creation. This is the user's login for the Console.
	// The name must be unique across all users in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the user. Does not have to be unique, and it's changeable.
	// (For tenancies that support identity domains) You can have an empty description.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the user was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The user's current state. After creating a user, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState UserLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Flag indicates if MFA has been activated for the user.
	IsMfaActivated *bool `mandatory:"true" json:"isMfaActivated"`

	// The email address you assign to the user.
	// The email address must be unique across all users in the tenancy.
	// (For tenancies that support identity domains) The email address is required unless the requirement is disabled at the tenancy level.
	Email *string `mandatory:"false" json:"email"`

	// Whether the email address has been validated.
	EmailVerified *bool `mandatory:"false" json:"emailVerified"`

	// DB username of the DB credential. Has to be unique across the tenancy.
	DbUserName *string `mandatory:"false" json:"dbUserName"`

	// The OCID of the `IdentityProvider` this user belongs to.
	IdentityProviderId *string `mandatory:"false" json:"identityProviderId"`

	// Identifier of the user in the identity provider
	ExternalIdentifier *string `mandatory:"false" json:"externalIdentifier"`

	// Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user
	// is inactive:
	// - bit 0: SUSPENDED (reserved for future use)
	// - bit 1: DISABLED (reserved for future use)
	// - bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console)
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	Capabilities *UserCapabilities `mandatory:"false" json:"capabilities"`

	// The date and time of when the user most recently logged in the
	// format defined by RFC3339 (ex. `2016-08-25T21:10:29.600Z`).
	// If there is no login history, this field is null.
	// For illustrative purposes, suppose we have a user who has logged in
	// at July 1st, 2020 at 1200 PST and logged out 30 minutes later.
	// They then login again on July 2nd, 2020 at 1500 PST.
	// Their previousSuccessfulLoginTime would be `2020-07-01:19:00.000Z`.
	// Their lastSuccessfulLoginTime would be `2020-07-02:22:00.000Z`.
	LastSuccessfulLoginTime *common.SDKTime `mandatory:"false" json:"lastSuccessfulLoginTime"`

	// The date and time of when the user most recently logged in the
	// format defined by RFC3339 (ex. `2016-08-25T21:10:29.600Z`).
	// If there is no login history, this field is null.
	// For illustrative purposes, suppose we have a user who has logged in
	// at July 1st, 2020 at 1200 PST and logged out 30 minutes later.
	// They then login again on July 2nd, 2020 at 1500 PST.
	// Their previousSuccessfulLoginTime would be `2020-07-01:19:00.000Z`.
	// Their lastSuccessfulLoginTime would be `2020-07-02:22:00.000Z`.
	PreviousSuccessfulLoginTime *common.SDKTime `mandatory:"false" json:"previousSuccessfulLoginTime"`
}

func (m User) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m User) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUserLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserLifecycleStateEnum Enum with underlying type: string
type UserLifecycleStateEnum string

// Set of constants representing the allowable values for UserLifecycleStateEnum
const (
	UserLifecycleStateCreating UserLifecycleStateEnum = "CREATING"
	UserLifecycleStateActive   UserLifecycleStateEnum = "ACTIVE"
	UserLifecycleStateInactive UserLifecycleStateEnum = "INACTIVE"
	UserLifecycleStateDeleting UserLifecycleStateEnum = "DELETING"
	UserLifecycleStateDeleted  UserLifecycleStateEnum = "DELETED"
)

var mappingUserLifecycleStateEnum = map[string]UserLifecycleStateEnum{
	"CREATING": UserLifecycleStateCreating,
	"ACTIVE":   UserLifecycleStateActive,
	"INACTIVE": UserLifecycleStateInactive,
	"DELETING": UserLifecycleStateDeleting,
	"DELETED":  UserLifecycleStateDeleted,
}

var mappingUserLifecycleStateEnumLowerCase = map[string]UserLifecycleStateEnum{
	"creating": UserLifecycleStateCreating,
	"active":   UserLifecycleStateActive,
	"inactive": UserLifecycleStateInactive,
	"deleting": UserLifecycleStateDeleting,
	"deleted":  UserLifecycleStateDeleted,
}

// GetUserLifecycleStateEnumValues Enumerates the set of values for UserLifecycleStateEnum
func GetUserLifecycleStateEnumValues() []UserLifecycleStateEnum {
	values := make([]UserLifecycleStateEnum, 0)
	for _, v := range mappingUserLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUserLifecycleStateEnumStringValues Enumerates the set of values in String for UserLifecycleStateEnum
func GetUserLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingUserLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserLifecycleStateEnum(val string) (UserLifecycleStateEnum, bool) {
	enum, ok := mappingUserLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
