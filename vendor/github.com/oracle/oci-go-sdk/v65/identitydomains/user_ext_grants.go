// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserExtGrants Grants to this User. Each value of this attribute refers to a Grant to this User of some App (and optionally of some entitlement). Therefore, this attribute is a convenience that allows one to see on each User all of the Grants to that User.
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type UserExtGrants struct {

	// The ID of this Grant to this User.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"false" json:"value"`

	// The URI of this Grant to this User.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// The ID of the App in this Grant.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppId *string `mandatory:"false" json:"appId"`

	// Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted.
	// A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
	//   - 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
	//   - 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
	//   - 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GrantMechanism UserExtGrantsGrantMechanismEnum `mandatory:"false" json:"grantMechanism,omitempty"`

	// Grantor identifier
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GrantorId *string `mandatory:"false" json:"grantorId"`
}

func (m UserExtGrants) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtGrants) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUserExtGrantsGrantMechanismEnum(string(m.GrantMechanism)); !ok && m.GrantMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantMechanism: %s. Supported values are: %s.", m.GrantMechanism, strings.Join(GetUserExtGrantsGrantMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserExtGrantsGrantMechanismEnum Enum with underlying type: string
type UserExtGrantsGrantMechanismEnum string

// Set of constants representing the allowable values for UserExtGrantsGrantMechanismEnum
const (
	UserExtGrantsGrantMechanismImportApproleMembers UserExtGrantsGrantMechanismEnum = "IMPORT_APPROLE_MEMBERS"
	UserExtGrantsGrantMechanismAdministratorToUser  UserExtGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_USER"
	UserExtGrantsGrantMechanismAdministratorToGroup UserExtGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_GROUP"
	UserExtGrantsGrantMechanismServiceManagerToUser UserExtGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_USER"
	UserExtGrantsGrantMechanismAdministratorToApp   UserExtGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_APP"
	UserExtGrantsGrantMechanismServiceManagerToApp  UserExtGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_APP"
	UserExtGrantsGrantMechanismOpcInfraToApp        UserExtGrantsGrantMechanismEnum = "OPC_INFRA_TO_APP"
	UserExtGrantsGrantMechanismGroupMembership      UserExtGrantsGrantMechanismEnum = "GROUP_MEMBERSHIP"
)

var mappingUserExtGrantsGrantMechanismEnum = map[string]UserExtGrantsGrantMechanismEnum{
	"IMPORT_APPROLE_MEMBERS":  UserExtGrantsGrantMechanismImportApproleMembers,
	"ADMINISTRATOR_TO_USER":   UserExtGrantsGrantMechanismAdministratorToUser,
	"ADMINISTRATOR_TO_GROUP":  UserExtGrantsGrantMechanismAdministratorToGroup,
	"SERVICE_MANAGER_TO_USER": UserExtGrantsGrantMechanismServiceManagerToUser,
	"ADMINISTRATOR_TO_APP":    UserExtGrantsGrantMechanismAdministratorToApp,
	"SERVICE_MANAGER_TO_APP":  UserExtGrantsGrantMechanismServiceManagerToApp,
	"OPC_INFRA_TO_APP":        UserExtGrantsGrantMechanismOpcInfraToApp,
	"GROUP_MEMBERSHIP":        UserExtGrantsGrantMechanismGroupMembership,
}

var mappingUserExtGrantsGrantMechanismEnumLowerCase = map[string]UserExtGrantsGrantMechanismEnum{
	"import_approle_members":  UserExtGrantsGrantMechanismImportApproleMembers,
	"administrator_to_user":   UserExtGrantsGrantMechanismAdministratorToUser,
	"administrator_to_group":  UserExtGrantsGrantMechanismAdministratorToGroup,
	"service_manager_to_user": UserExtGrantsGrantMechanismServiceManagerToUser,
	"administrator_to_app":    UserExtGrantsGrantMechanismAdministratorToApp,
	"service_manager_to_app":  UserExtGrantsGrantMechanismServiceManagerToApp,
	"opc_infra_to_app":        UserExtGrantsGrantMechanismOpcInfraToApp,
	"group_membership":        UserExtGrantsGrantMechanismGroupMembership,
}

// GetUserExtGrantsGrantMechanismEnumValues Enumerates the set of values for UserExtGrantsGrantMechanismEnum
func GetUserExtGrantsGrantMechanismEnumValues() []UserExtGrantsGrantMechanismEnum {
	values := make([]UserExtGrantsGrantMechanismEnum, 0)
	for _, v := range mappingUserExtGrantsGrantMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetUserExtGrantsGrantMechanismEnumStringValues Enumerates the set of values in String for UserExtGrantsGrantMechanismEnum
func GetUserExtGrantsGrantMechanismEnumStringValues() []string {
	return []string{
		"IMPORT_APPROLE_MEMBERS",
		"ADMINISTRATOR_TO_USER",
		"ADMINISTRATOR_TO_GROUP",
		"SERVICE_MANAGER_TO_USER",
		"ADMINISTRATOR_TO_APP",
		"SERVICE_MANAGER_TO_APP",
		"OPC_INFRA_TO_APP",
		"GROUP_MEMBERSHIP",
	}
}

// GetMappingUserExtGrantsGrantMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserExtGrantsGrantMechanismEnum(val string) (UserExtGrantsGrantMechanismEnum, bool) {
	enum, ok := mappingUserExtGrantsGrantMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
