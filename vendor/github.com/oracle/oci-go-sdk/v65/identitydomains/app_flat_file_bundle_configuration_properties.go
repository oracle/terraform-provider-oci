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

// AppFlatFileBundleConfigurationProperties Flat file connector bundle configuration properties
// **SCIM++ Properties:**
//   - idcsCompositeKey: [name]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AppFlatFileBundleConfigurationProperties struct {

	// Name of the flatfile bundle configuration property. This attribute maps to \"name\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// ICF data type of flatfile the bundle configuration property. This attribute maps to \"type\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IcfType AppFlatFileBundleConfigurationPropertiesIcfTypeEnum `mandatory:"true" json:"icfType"`

	// If true, this flatfile bundle configuration property is required to connect to the target connected managed app. This attribute maps to \"isRequired\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Required *bool `mandatory:"true" json:"required"`

	// Display name of the flatfile bundle configuration property. This attribute maps to \"displayName\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Value of the flatfile bundle configuration property. This attribute maps to \"value\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - idcsSensitive: encrypt
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value []string `mandatory:"false" json:"value"`

	// Display sequence of the bundle configuration property.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	Order *int `mandatory:"false" json:"order"`

	// Help message of the flatfile bundle configuration property. This attribute maps to \"helpMessage\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HelpMessage *string `mandatory:"false" json:"helpMessage"`

	// If true, this flatfile bundle configuration property value is confidential and will be encrypted in Oracle Identity Cloud Service. This attribute maps to \"isConfidential\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Confidential *bool `mandatory:"false" json:"confidential"`
}

func (m AppFlatFileBundleConfigurationProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppFlatFileBundleConfigurationProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnum(string(m.IcfType)); !ok && m.IcfType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IcfType: %s. Supported values are: %s.", m.IcfType, strings.Join(GetAppFlatFileBundleConfigurationPropertiesIcfTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppFlatFileBundleConfigurationPropertiesIcfTypeEnum Enum with underlying type: string
type AppFlatFileBundleConfigurationPropertiesIcfTypeEnum string

// Set of constants representing the allowable values for AppFlatFileBundleConfigurationPropertiesIcfTypeEnum
const (
	AppFlatFileBundleConfigurationPropertiesIcfTypeLong                    AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "Long"
	AppFlatFileBundleConfigurationPropertiesIcfTypeString                  AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "String"
	AppFlatFileBundleConfigurationPropertiesIcfTypeCharacter               AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "Character"
	AppFlatFileBundleConfigurationPropertiesIcfTypeDouble                  AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "Double"
	AppFlatFileBundleConfigurationPropertiesIcfTypeFloat                   AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "Float"
	AppFlatFileBundleConfigurationPropertiesIcfTypeInteger                 AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "Integer"
	AppFlatFileBundleConfigurationPropertiesIcfTypeBoolean                 AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "Boolean"
	AppFlatFileBundleConfigurationPropertiesIcfTypeUri                     AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "URI"
	AppFlatFileBundleConfigurationPropertiesIcfTypeFile                    AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "File"
	AppFlatFileBundleConfigurationPropertiesIcfTypeGuardedbytearray        AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "GuardedByteArray"
	AppFlatFileBundleConfigurationPropertiesIcfTypeGuardedstring           AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "GuardedString"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoflong             AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfLong"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofstring           AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfString"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofcharacter        AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfCharacter"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofdouble           AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfDouble"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoffloat            AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfFloat"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofinteger          AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfInteger"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofboolean          AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfBoolean"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofuri              AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfURI"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoffile             AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfFile"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofguardedbytearray AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfGuardedByteArray"
	AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofguardedstring    AppFlatFileBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfGuardedString"
)

var mappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnum = map[string]AppFlatFileBundleConfigurationPropertiesIcfTypeEnum{
	"Long":                    AppFlatFileBundleConfigurationPropertiesIcfTypeLong,
	"String":                  AppFlatFileBundleConfigurationPropertiesIcfTypeString,
	"Character":               AppFlatFileBundleConfigurationPropertiesIcfTypeCharacter,
	"Double":                  AppFlatFileBundleConfigurationPropertiesIcfTypeDouble,
	"Float":                   AppFlatFileBundleConfigurationPropertiesIcfTypeFloat,
	"Integer":                 AppFlatFileBundleConfigurationPropertiesIcfTypeInteger,
	"Boolean":                 AppFlatFileBundleConfigurationPropertiesIcfTypeBoolean,
	"URI":                     AppFlatFileBundleConfigurationPropertiesIcfTypeUri,
	"File":                    AppFlatFileBundleConfigurationPropertiesIcfTypeFile,
	"GuardedByteArray":        AppFlatFileBundleConfigurationPropertiesIcfTypeGuardedbytearray,
	"GuardedString":           AppFlatFileBundleConfigurationPropertiesIcfTypeGuardedstring,
	"ArrayOfLong":             AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoflong,
	"ArrayOfString":           AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofstring,
	"ArrayOfCharacter":        AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofcharacter,
	"ArrayOfDouble":           AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofdouble,
	"ArrayOfFloat":            AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoffloat,
	"ArrayOfInteger":          AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofinteger,
	"ArrayOfBoolean":          AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofboolean,
	"ArrayOfURI":              AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofuri,
	"ArrayOfFile":             AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoffile,
	"ArrayOfGuardedByteArray": AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofguardedbytearray,
	"ArrayOfGuardedString":    AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofguardedstring,
}

var mappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnumLowerCase = map[string]AppFlatFileBundleConfigurationPropertiesIcfTypeEnum{
	"long":                    AppFlatFileBundleConfigurationPropertiesIcfTypeLong,
	"string":                  AppFlatFileBundleConfigurationPropertiesIcfTypeString,
	"character":               AppFlatFileBundleConfigurationPropertiesIcfTypeCharacter,
	"double":                  AppFlatFileBundleConfigurationPropertiesIcfTypeDouble,
	"float":                   AppFlatFileBundleConfigurationPropertiesIcfTypeFloat,
	"integer":                 AppFlatFileBundleConfigurationPropertiesIcfTypeInteger,
	"boolean":                 AppFlatFileBundleConfigurationPropertiesIcfTypeBoolean,
	"uri":                     AppFlatFileBundleConfigurationPropertiesIcfTypeUri,
	"file":                    AppFlatFileBundleConfigurationPropertiesIcfTypeFile,
	"guardedbytearray":        AppFlatFileBundleConfigurationPropertiesIcfTypeGuardedbytearray,
	"guardedstring":           AppFlatFileBundleConfigurationPropertiesIcfTypeGuardedstring,
	"arrayoflong":             AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoflong,
	"arrayofstring":           AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofstring,
	"arrayofcharacter":        AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofcharacter,
	"arrayofdouble":           AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofdouble,
	"arrayoffloat":            AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoffloat,
	"arrayofinteger":          AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofinteger,
	"arrayofboolean":          AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofboolean,
	"arrayofuri":              AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofuri,
	"arrayoffile":             AppFlatFileBundleConfigurationPropertiesIcfTypeArrayoffile,
	"arrayofguardedbytearray": AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofguardedbytearray,
	"arrayofguardedstring":    AppFlatFileBundleConfigurationPropertiesIcfTypeArrayofguardedstring,
}

// GetAppFlatFileBundleConfigurationPropertiesIcfTypeEnumValues Enumerates the set of values for AppFlatFileBundleConfigurationPropertiesIcfTypeEnum
func GetAppFlatFileBundleConfigurationPropertiesIcfTypeEnumValues() []AppFlatFileBundleConfigurationPropertiesIcfTypeEnum {
	values := make([]AppFlatFileBundleConfigurationPropertiesIcfTypeEnum, 0)
	for _, v := range mappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppFlatFileBundleConfigurationPropertiesIcfTypeEnumStringValues Enumerates the set of values in String for AppFlatFileBundleConfigurationPropertiesIcfTypeEnum
func GetAppFlatFileBundleConfigurationPropertiesIcfTypeEnumStringValues() []string {
	return []string{
		"Long",
		"String",
		"Character",
		"Double",
		"Float",
		"Integer",
		"Boolean",
		"URI",
		"File",
		"GuardedByteArray",
		"GuardedString",
		"ArrayOfLong",
		"ArrayOfString",
		"ArrayOfCharacter",
		"ArrayOfDouble",
		"ArrayOfFloat",
		"ArrayOfInteger",
		"ArrayOfBoolean",
		"ArrayOfURI",
		"ArrayOfFile",
		"ArrayOfGuardedByteArray",
		"ArrayOfGuardedString",
	}
}

// GetMappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnum(val string) (AppFlatFileBundleConfigurationPropertiesIcfTypeEnum, bool) {
	enum, ok := mappingAppFlatFileBundleConfigurationPropertiesIcfTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
