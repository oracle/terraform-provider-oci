// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// Me User Account
type Me struct {

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

	// User name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: User ID
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:User Name, deprecatedColumnHeaderName:User ID]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	UserName *string `mandatory:"true" json:"userName"`

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

	// An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:External Id]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Description of the user
	// **Added In:** 2012271618
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsPii: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Display name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Display Name
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Display Name]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Nick name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Nick Name
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Nick Name]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	NickName *string `mandatory:"false" json:"nickName"`

	// A fully-qualified URL to a page representing the User's online profile
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Profile URL
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Profile Url]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	ProfileUrl *string `mandatory:"false" json:"profileUrl"`

	// Title
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Title
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Title]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Title *string `mandatory:"false" json:"title"`

	// Used to identify the organization-to-user relationship
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: User Type
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:User Type]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	UserType MeUserTypeEnum `mandatory:"false" json:"userType,omitempty"`

	// Used to indicate the User's default location for purposes of localizing items such as currency, date and time format, numerical representations, and so on.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Locale
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Locale]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Locale *string `mandatory:"false" json:"locale"`

	// User's preferred written or spoken language used for localized user interfaces
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Preferred Language
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Preferred Language]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PreferredLanguage *string `mandatory:"false" json:"preferredLanguage"`

	// User's timezone
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCanonicalValueSourceFilter: attrName eq "timezones" and attrValues.value eq "$(timezone)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - idcsCsvAttributeName: TimeZone
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Time Zone, deprecatedColumnHeaderName:TimeZone]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Timezone *string `mandatory:"false" json:"timezone"`

	// User status
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Active
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Active]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Active *bool `mandatory:"false" json:"active"`

	// Password attribute. Max length for password is controlled via Password Policy.
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: Password
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Password]]
	//  - idcsPii: true
	//  - idcsSearchable: false
	//  - idcsSensitive: hash
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: string
	//  - uniqueness: none
	Password *string `mandatory:"false" json:"password"`

	Name *MeName `mandatory:"false" json:"name"`

	// A complex attribute representing emails
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, type]
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Work Email, mapsTo:emails[work].value], [columnHeaderName:Home Email, mapsTo:emails[home].value], [columnHeaderName:Primary Email Type, mapsTo:emails[$(type)].primary], [columnHeaderName:Other Email, mapsTo:emails[other].value], [columnHeaderName:Recovery Email, mapsTo:emails[recovery].value], [columnHeaderName:Work Email Verified, mapsTo:emails[work].verified], [columnHeaderName:Home Email Verified, mapsTo:emails[home].verified], [columnHeaderName:Other Email Verified, mapsTo:emails[other].verified], [columnHeaderName:Recovery Email Verified, mapsTo:emails[recovery].verified]]
	//  - idcsPii: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Emails []MeEmails `mandatory:"false" json:"emails"`

	// Phone numbers
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, type]
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Work Phone, mapsTo:phoneNumbers[work].value], [columnHeaderName:Mobile No, mapsTo:phoneNumbers[mobile].value], [columnHeaderName:Home Phone, mapsTo:phoneNumbers[home].value], [columnHeaderName:Fax, mapsTo:phoneNumbers[fax].value], [columnHeaderName:Pager, mapsTo:phoneNumbers[pager].value], [columnHeaderName:Other Phone, mapsTo:phoneNumbers[other].value], [columnHeaderName:Recovery Phone, mapsTo:phoneNumbers[recovery].value], [columnHeaderName:Primary Phone Type, mapsTo:phoneNumbers[$(type)].primary]]
	//  - idcsPii: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	PhoneNumbers []MePhoneNumbers `mandatory:"false" json:"phoneNumbers"`

	// User's instant messaging addresses
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, type]
	//  - idcsPii: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Ims []MeIms `mandatory:"false" json:"ims"`

	// URLs of photos for the User
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, type]
	//  - idcsPii: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Photos []MePhotos `mandatory:"false" json:"photos"`

	// A physical mailing address for this User, as described in (address Element). Canonical Type Values of work, home, and other. The value attribute is a complex type with the following sub-attributes.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [type]
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Work Address Street, deprecatedColumnHeaderName:Work Street Address, mapsTo:addresses[work].streetAddress], [columnHeaderName:Work Address Locality, deprecatedColumnHeaderName:Work City, mapsTo:addresses[work].locality], [columnHeaderName:Work Address Region, deprecatedColumnHeaderName:Work State, mapsTo:addresses[work].region], [columnHeaderName:Work Address Postal Code, deprecatedColumnHeaderName:Work Postal Code, mapsTo:addresses[work].postalCode], [columnHeaderName:Work Address Country, deprecatedColumnHeaderName:Work Country, mapsTo:addresses[work].country], [columnHeaderName:Work Address Formatted, mapsTo:addresses[work].formatted], [columnHeaderName:Home Address Formatted, mapsTo:addresses[home].formatted], [columnHeaderName:Other Address Formatted, mapsTo:addresses[other].formatted], [columnHeaderName:Home Address Street, mapsTo:addresses[home].streetAddress], [columnHeaderName:Other Address Street, mapsTo:addresses[other].streetAddress], [columnHeaderName:Home Address Locality, mapsTo:addresses[home].locality], [columnHeaderName:Other Address Locality, mapsTo:addresses[other].locality], [columnHeaderName:Home Address Region, mapsTo:addresses[home].region], [columnHeaderName:Other Address Region, mapsTo:addresses[other].region], [columnHeaderName:Home Address Country, mapsTo:addresses[home].country], [columnHeaderName:Other Address Country, mapsTo:addresses[other].country], [columnHeaderName:Home Address Postal Code, mapsTo:addresses[home].postalCode], [columnHeaderName:Other Address Postal Code, mapsTo:addresses[other].postalCode], [columnHeaderName:Primary Address Type, mapsTo:addresses[$(type)].primary]]
	//  - idcsPii: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Addresses []Addresses `mandatory:"false" json:"addresses"`

	// A list of groups that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Groups []MeGroups `mandatory:"false" json:"groups"`

	// A list of entitlements for the User that represent a thing the User has.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, type]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Entitlements []MeEntitlements `mandatory:"false" json:"entitlements"`

	// A list of roles for the User that collectively represent who the User is; e.g., 'Student', 'Faculty'.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, type]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Roles []MeRoles `mandatory:"false" json:"roles"`

	// A list of certificates issued to the User.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	X509Certificates []MeX509Certificates `mandatory:"false" json:"x509Certificates"`

	UrnIetfParamsScimSchemasExtensionEnterprise2_0User *ExtensionEnterprise20User `mandatory:"false" json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionUserUser *ExtensionUserUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:user:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionPasswordStateUser *ExtensionPasswordStateUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:passwordState:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionUserStateUser *ExtensionUserStateUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:userState:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionMeUser *ExtensionMeUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:me:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionPosixUser *ExtensionPosixUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:posix:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionMfaUser *ExtensionMfaUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:mfa:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionSecurityQuestionsUser *ExtensionSecurityQuestionsUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:securityQuestions:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionSelfRegistrationUser *ExtensionSelfRegistrationUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:selfRegistration:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionTermsOfUseUser *ExtensionTermsOfUseUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:termsOfUse:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags *ExtensionOciTags `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:OCITags"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionUserCredentialsUser *ExtensionUserCredentialsUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:userCredentials:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionCapabilitiesUser *ExtensionCapabilitiesUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:capabilities:User"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionDbCredentialsUser *ExtensionDbCredentialsUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:dbCredentials:User"`
}

func (m Me) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Me) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingMeUserTypeEnum(string(m.UserType)); !ok && m.UserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserType: %s. Supported values are: %s.", m.UserType, strings.Join(GetMeUserTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MeUserTypeEnum Enum with underlying type: string
type MeUserTypeEnum string

// Set of constants representing the allowable values for MeUserTypeEnum
const (
	MeUserTypeContractor MeUserTypeEnum = "Contractor"
	MeUserTypeEmployee   MeUserTypeEnum = "Employee"
	MeUserTypeIntern     MeUserTypeEnum = "Intern"
	MeUserTypeTemp       MeUserTypeEnum = "Temp"
	MeUserTypeExternal   MeUserTypeEnum = "External"
	MeUserTypeService    MeUserTypeEnum = "Service"
	MeUserTypeGeneric    MeUserTypeEnum = "Generic"
)

var mappingMeUserTypeEnum = map[string]MeUserTypeEnum{
	"Contractor": MeUserTypeContractor,
	"Employee":   MeUserTypeEmployee,
	"Intern":     MeUserTypeIntern,
	"Temp":       MeUserTypeTemp,
	"External":   MeUserTypeExternal,
	"Service":    MeUserTypeService,
	"Generic":    MeUserTypeGeneric,
}

var mappingMeUserTypeEnumLowerCase = map[string]MeUserTypeEnum{
	"contractor": MeUserTypeContractor,
	"employee":   MeUserTypeEmployee,
	"intern":     MeUserTypeIntern,
	"temp":       MeUserTypeTemp,
	"external":   MeUserTypeExternal,
	"service":    MeUserTypeService,
	"generic":    MeUserTypeGeneric,
}

// GetMeUserTypeEnumValues Enumerates the set of values for MeUserTypeEnum
func GetMeUserTypeEnumValues() []MeUserTypeEnum {
	values := make([]MeUserTypeEnum, 0)
	for _, v := range mappingMeUserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMeUserTypeEnumStringValues Enumerates the set of values in String for MeUserTypeEnum
func GetMeUserTypeEnumStringValues() []string {
	return []string{
		"Contractor",
		"Employee",
		"Intern",
		"Temp",
		"External",
		"Service",
		"Generic",
	}
}

// GetMappingMeUserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMeUserTypeEnum(val string) (MeUserTypeEnum, bool) {
	enum, ok := mappingMeUserTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
