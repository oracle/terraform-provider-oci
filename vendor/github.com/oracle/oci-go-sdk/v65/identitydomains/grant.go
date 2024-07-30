// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Grant Schema for Grant Resource
type Grant struct {

	// REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Schemas []string `mandatory:"true" json:"schemas"`

	// Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted.
	// A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
	//   - 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
	//   - 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
	//   - 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeNameMappings: [[defaultValue:IMPORT_GRANTS]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GrantMechanism GrantGrantMechanismEnum `mandatory:"true" json:"grantMechanism"`

	Grantee *GrantGrantee `mandatory:"true" json:"grantee"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Id *string `mandatory:"false" json:"id"`

	// Unique OCI identifier for the SCIM Resource.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: global
	Ocid *string `mandatory:"false" json:"ocid"`

	Meta *Meta `mandatory:"false" json:"meta"`

	IdcsCreatedBy *IdcsCreatedBy `mandatory:"false" json:"idcsCreatedBy"`

	IdcsLastModifiedBy *IdcsLastModifiedBy `mandatory:"false" json:"idcsLastModifiedBy"`

	// Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsPreventedOperations []IdcsPreventedOperationsEnum `mandatory:"false" json:"idcsPreventedOperations,omitempty"`

	// A list of tags on this resource.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DeleteInProgress *bool `mandatory:"false" json:"deleteInProgress"`

	// The release number when the resource was upgraded.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsLastUpgradedInRelease *string `mandatory:"false" json:"idcsLastUpgradedInRelease"`

	// OCI Domain Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DomainOcid *string `mandatory:"false" json:"domainOcid"`

	// OCI Compartment Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// OCI Tenant Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`

	// Unique key of grant, composed by combining a subset of app, entitlement, grantee, grantor and grantMechanism.  Used to prevent duplicate Grants.
	// **Added In:** 18.1.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: server
	CompositeKey *string `mandatory:"false" json:"compositeKey"`

	// If true, this Grant has been fulfilled successfully.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsFulfilled *bool `mandatory:"false" json:"isFulfilled"`

	// Store granted attribute-values as a string in Javascript Object Notation (JSON) format.
	// **Added In:** 18.3.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GrantedAttributeValuesJson *string `mandatory:"false" json:"grantedAttributeValuesJson"`

	AppEntitlementCollection *GrantAppEntitlementCollection `mandatory:"false" json:"appEntitlementCollection"`

	Grantor *GrantGrantor `mandatory:"false" json:"grantor"`

	App *GrantApp `mandatory:"false" json:"app"`

	Entitlement *GrantEntitlement `mandatory:"false" json:"entitlement"`
}

func (m Grant) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Grant) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGrantGrantMechanismEnum(string(m.GrantMechanism)); !ok && m.GrantMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantMechanism: %s. Supported values are: %s.", m.GrantMechanism, strings.Join(GetGrantGrantMechanismEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GrantGrantMechanismEnum Enum with underlying type: string
type GrantGrantMechanismEnum string

// Set of constants representing the allowable values for GrantGrantMechanismEnum
const (
	GrantGrantMechanismImportApproleMembers                GrantGrantMechanismEnum = "IMPORT_APPROLE_MEMBERS"
	GrantGrantMechanismAdministratorToUser                 GrantGrantMechanismEnum = "ADMINISTRATOR_TO_USER"
	GrantGrantMechanismAdministratorToDelegatedUser        GrantGrantMechanismEnum = "ADMINISTRATOR_TO_DELEGATED_USER"
	GrantGrantMechanismAdministratorToGroup                GrantGrantMechanismEnum = "ADMINISTRATOR_TO_GROUP"
	GrantGrantMechanismServiceManagerToUser                GrantGrantMechanismEnum = "SERVICE_MANAGER_TO_USER"
	GrantGrantMechanismAdministratorToApp                  GrantGrantMechanismEnum = "ADMINISTRATOR_TO_APP"
	GrantGrantMechanismServiceManagerToApp                 GrantGrantMechanismEnum = "SERVICE_MANAGER_TO_APP"
	GrantGrantMechanismOpcInfraToApp                       GrantGrantMechanismEnum = "OPC_INFRA_TO_APP"
	GrantGrantMechanismGroupMembership                     GrantGrantMechanismEnum = "GROUP_MEMBERSHIP"
	GrantGrantMechanismImportGrants                        GrantGrantMechanismEnum = "IMPORT_GRANTS"
	GrantGrantMechanismSyncToUser                          GrantGrantMechanismEnum = "SYNC_TO_USER"
	GrantGrantMechanismAccessRequest                       GrantGrantMechanismEnum = "ACCESS_REQUEST"
	GrantGrantMechanismAppEntitlementCollection            GrantGrantMechanismEnum = "APP_ENTITLEMENT_COLLECTION"
	GrantGrantMechanismAdministratorToDynamicResourceGroup GrantGrantMechanismEnum = "ADMINISTRATOR_TO_DYNAMIC_RESOURCE_GROUP"
)

var mappingGrantGrantMechanismEnum = map[string]GrantGrantMechanismEnum{
	"IMPORT_APPROLE_MEMBERS":                  GrantGrantMechanismImportApproleMembers,
	"ADMINISTRATOR_TO_USER":                   GrantGrantMechanismAdministratorToUser,
	"ADMINISTRATOR_TO_DELEGATED_USER":         GrantGrantMechanismAdministratorToDelegatedUser,
	"ADMINISTRATOR_TO_GROUP":                  GrantGrantMechanismAdministratorToGroup,
	"SERVICE_MANAGER_TO_USER":                 GrantGrantMechanismServiceManagerToUser,
	"ADMINISTRATOR_TO_APP":                    GrantGrantMechanismAdministratorToApp,
	"SERVICE_MANAGER_TO_APP":                  GrantGrantMechanismServiceManagerToApp,
	"OPC_INFRA_TO_APP":                        GrantGrantMechanismOpcInfraToApp,
	"GROUP_MEMBERSHIP":                        GrantGrantMechanismGroupMembership,
	"IMPORT_GRANTS":                           GrantGrantMechanismImportGrants,
	"SYNC_TO_USER":                            GrantGrantMechanismSyncToUser,
	"ACCESS_REQUEST":                          GrantGrantMechanismAccessRequest,
	"APP_ENTITLEMENT_COLLECTION":              GrantGrantMechanismAppEntitlementCollection,
	"ADMINISTRATOR_TO_DYNAMIC_RESOURCE_GROUP": GrantGrantMechanismAdministratorToDynamicResourceGroup,
}

var mappingGrantGrantMechanismEnumLowerCase = map[string]GrantGrantMechanismEnum{
	"import_approle_members":                  GrantGrantMechanismImportApproleMembers,
	"administrator_to_user":                   GrantGrantMechanismAdministratorToUser,
	"administrator_to_delegated_user":         GrantGrantMechanismAdministratorToDelegatedUser,
	"administrator_to_group":                  GrantGrantMechanismAdministratorToGroup,
	"service_manager_to_user":                 GrantGrantMechanismServiceManagerToUser,
	"administrator_to_app":                    GrantGrantMechanismAdministratorToApp,
	"service_manager_to_app":                  GrantGrantMechanismServiceManagerToApp,
	"opc_infra_to_app":                        GrantGrantMechanismOpcInfraToApp,
	"group_membership":                        GrantGrantMechanismGroupMembership,
	"import_grants":                           GrantGrantMechanismImportGrants,
	"sync_to_user":                            GrantGrantMechanismSyncToUser,
	"access_request":                          GrantGrantMechanismAccessRequest,
	"app_entitlement_collection":              GrantGrantMechanismAppEntitlementCollection,
	"administrator_to_dynamic_resource_group": GrantGrantMechanismAdministratorToDynamicResourceGroup,
}

// GetGrantGrantMechanismEnumValues Enumerates the set of values for GrantGrantMechanismEnum
func GetGrantGrantMechanismEnumValues() []GrantGrantMechanismEnum {
	values := make([]GrantGrantMechanismEnum, 0)
	for _, v := range mappingGrantGrantMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetGrantGrantMechanismEnumStringValues Enumerates the set of values in String for GrantGrantMechanismEnum
func GetGrantGrantMechanismEnumStringValues() []string {
	return []string{
		"IMPORT_APPROLE_MEMBERS",
		"ADMINISTRATOR_TO_USER",
		"ADMINISTRATOR_TO_DELEGATED_USER",
		"ADMINISTRATOR_TO_GROUP",
		"SERVICE_MANAGER_TO_USER",
		"ADMINISTRATOR_TO_APP",
		"SERVICE_MANAGER_TO_APP",
		"OPC_INFRA_TO_APP",
		"GROUP_MEMBERSHIP",
		"IMPORT_GRANTS",
		"SYNC_TO_USER",
		"ACCESS_REQUEST",
		"APP_ENTITLEMENT_COLLECTION",
		"ADMINISTRATOR_TO_DYNAMIC_RESOURCE_GROUP",
	}
}

// GetMappingGrantGrantMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGrantGrantMechanismEnum(val string) (GrantGrantMechanismEnum, bool) {
	enum, ok := mappingGrantGrantMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
