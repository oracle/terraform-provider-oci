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

// AppBundleConfigurationProperties ConnectorBundle configuration properties
// **SCIM++ Properties:**
//   - idcsCompositeKey: [name]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AppBundleConfigurationProperties struct {

	// Name of the bundle configuration property. This attribute maps to \"name\" attribute in \"ConfigurationProperty\" in ICF.
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

	// ICF data type of the bundle configuration property. This attribute maps to \"type\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IcfType AppBundleConfigurationPropertiesIcfTypeEnum `mandatory:"true" json:"icfType"`

	// If true, this bundle configuration property is required to connect to the target connected managed app. This attribute maps to \"isRequired\" attribute in \"ConfigurationProperty\" in ICF.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Required *bool `mandatory:"true" json:"required"`

	// Display name of the bundle configuration property. This attribute maps to \"displayName\" attribute in \"ConfigurationProperty\" in ICF.
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

	// Value of the bundle configuration property. This attribute maps to \"value\" attribute in \"ConfigurationProperty\" in ICF.
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

	// Help message of the bundle configuration property. This attribute maps to \"helpMessage\" attribute in \"ConfigurationProperty\" in ICF.
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

	// If true, this bundle configuration property value is confidential and will be encrypted in Oracle Identity Cloud Service. This attribute maps to \"isConfidential\" attribute in \"ConfigurationProperty\" in ICF.
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

func (m AppBundleConfigurationProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppBundleConfigurationProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAppBundleConfigurationPropertiesIcfTypeEnum(string(m.IcfType)); !ok && m.IcfType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IcfType: %s. Supported values are: %s.", m.IcfType, strings.Join(GetAppBundleConfigurationPropertiesIcfTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppBundleConfigurationPropertiesIcfTypeEnum Enum with underlying type: string
type AppBundleConfigurationPropertiesIcfTypeEnum string

// Set of constants representing the allowable values for AppBundleConfigurationPropertiesIcfTypeEnum
const (
	AppBundleConfigurationPropertiesIcfTypeLong                    AppBundleConfigurationPropertiesIcfTypeEnum = "Long"
	AppBundleConfigurationPropertiesIcfTypeString                  AppBundleConfigurationPropertiesIcfTypeEnum = "String"
	AppBundleConfigurationPropertiesIcfTypeCharacter               AppBundleConfigurationPropertiesIcfTypeEnum = "Character"
	AppBundleConfigurationPropertiesIcfTypeDouble                  AppBundleConfigurationPropertiesIcfTypeEnum = "Double"
	AppBundleConfigurationPropertiesIcfTypeFloat                   AppBundleConfigurationPropertiesIcfTypeEnum = "Float"
	AppBundleConfigurationPropertiesIcfTypeInteger                 AppBundleConfigurationPropertiesIcfTypeEnum = "Integer"
	AppBundleConfigurationPropertiesIcfTypeBoolean                 AppBundleConfigurationPropertiesIcfTypeEnum = "Boolean"
	AppBundleConfigurationPropertiesIcfTypeUri                     AppBundleConfigurationPropertiesIcfTypeEnum = "URI"
	AppBundleConfigurationPropertiesIcfTypeFile                    AppBundleConfigurationPropertiesIcfTypeEnum = "File"
	AppBundleConfigurationPropertiesIcfTypeGuardedbytearray        AppBundleConfigurationPropertiesIcfTypeEnum = "GuardedByteArray"
	AppBundleConfigurationPropertiesIcfTypeGuardedstring           AppBundleConfigurationPropertiesIcfTypeEnum = "GuardedString"
	AppBundleConfigurationPropertiesIcfTypeArrayoflong             AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfLong"
	AppBundleConfigurationPropertiesIcfTypeArrayofstring           AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfString"
	AppBundleConfigurationPropertiesIcfTypeArrayofcharacter        AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfCharacter"
	AppBundleConfigurationPropertiesIcfTypeArrayofdouble           AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfDouble"
	AppBundleConfigurationPropertiesIcfTypeArrayoffloat            AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfFloat"
	AppBundleConfigurationPropertiesIcfTypeArrayofinteger          AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfInteger"
	AppBundleConfigurationPropertiesIcfTypeArrayofboolean          AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfBoolean"
	AppBundleConfigurationPropertiesIcfTypeArrayofuri              AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfURI"
	AppBundleConfigurationPropertiesIcfTypeArrayoffile             AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfFile"
	AppBundleConfigurationPropertiesIcfTypeArrayofguardedbytearray AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfGuardedByteArray"
	AppBundleConfigurationPropertiesIcfTypeArrayofguardedstring    AppBundleConfigurationPropertiesIcfTypeEnum = "ArrayOfGuardedString"
)

var mappingAppBundleConfigurationPropertiesIcfTypeEnum = map[string]AppBundleConfigurationPropertiesIcfTypeEnum{
	"Long":                    AppBundleConfigurationPropertiesIcfTypeLong,
	"String":                  AppBundleConfigurationPropertiesIcfTypeString,
	"Character":               AppBundleConfigurationPropertiesIcfTypeCharacter,
	"Double":                  AppBundleConfigurationPropertiesIcfTypeDouble,
	"Float":                   AppBundleConfigurationPropertiesIcfTypeFloat,
	"Integer":                 AppBundleConfigurationPropertiesIcfTypeInteger,
	"Boolean":                 AppBundleConfigurationPropertiesIcfTypeBoolean,
	"URI":                     AppBundleConfigurationPropertiesIcfTypeUri,
	"File":                    AppBundleConfigurationPropertiesIcfTypeFile,
	"GuardedByteArray":        AppBundleConfigurationPropertiesIcfTypeGuardedbytearray,
	"GuardedString":           AppBundleConfigurationPropertiesIcfTypeGuardedstring,
	"ArrayOfLong":             AppBundleConfigurationPropertiesIcfTypeArrayoflong,
	"ArrayOfString":           AppBundleConfigurationPropertiesIcfTypeArrayofstring,
	"ArrayOfCharacter":        AppBundleConfigurationPropertiesIcfTypeArrayofcharacter,
	"ArrayOfDouble":           AppBundleConfigurationPropertiesIcfTypeArrayofdouble,
	"ArrayOfFloat":            AppBundleConfigurationPropertiesIcfTypeArrayoffloat,
	"ArrayOfInteger":          AppBundleConfigurationPropertiesIcfTypeArrayofinteger,
	"ArrayOfBoolean":          AppBundleConfigurationPropertiesIcfTypeArrayofboolean,
	"ArrayOfURI":              AppBundleConfigurationPropertiesIcfTypeArrayofuri,
	"ArrayOfFile":             AppBundleConfigurationPropertiesIcfTypeArrayoffile,
	"ArrayOfGuardedByteArray": AppBundleConfigurationPropertiesIcfTypeArrayofguardedbytearray,
	"ArrayOfGuardedString":    AppBundleConfigurationPropertiesIcfTypeArrayofguardedstring,
}

var mappingAppBundleConfigurationPropertiesIcfTypeEnumLowerCase = map[string]AppBundleConfigurationPropertiesIcfTypeEnum{
	"long":                    AppBundleConfigurationPropertiesIcfTypeLong,
	"string":                  AppBundleConfigurationPropertiesIcfTypeString,
	"character":               AppBundleConfigurationPropertiesIcfTypeCharacter,
	"double":                  AppBundleConfigurationPropertiesIcfTypeDouble,
	"float":                   AppBundleConfigurationPropertiesIcfTypeFloat,
	"integer":                 AppBundleConfigurationPropertiesIcfTypeInteger,
	"boolean":                 AppBundleConfigurationPropertiesIcfTypeBoolean,
	"uri":                     AppBundleConfigurationPropertiesIcfTypeUri,
	"file":                    AppBundleConfigurationPropertiesIcfTypeFile,
	"guardedbytearray":        AppBundleConfigurationPropertiesIcfTypeGuardedbytearray,
	"guardedstring":           AppBundleConfigurationPropertiesIcfTypeGuardedstring,
	"arrayoflong":             AppBundleConfigurationPropertiesIcfTypeArrayoflong,
	"arrayofstring":           AppBundleConfigurationPropertiesIcfTypeArrayofstring,
	"arrayofcharacter":        AppBundleConfigurationPropertiesIcfTypeArrayofcharacter,
	"arrayofdouble":           AppBundleConfigurationPropertiesIcfTypeArrayofdouble,
	"arrayoffloat":            AppBundleConfigurationPropertiesIcfTypeArrayoffloat,
	"arrayofinteger":          AppBundleConfigurationPropertiesIcfTypeArrayofinteger,
	"arrayofboolean":          AppBundleConfigurationPropertiesIcfTypeArrayofboolean,
	"arrayofuri":              AppBundleConfigurationPropertiesIcfTypeArrayofuri,
	"arrayoffile":             AppBundleConfigurationPropertiesIcfTypeArrayoffile,
	"arrayofguardedbytearray": AppBundleConfigurationPropertiesIcfTypeArrayofguardedbytearray,
	"arrayofguardedstring":    AppBundleConfigurationPropertiesIcfTypeArrayofguardedstring,
}

// GetAppBundleConfigurationPropertiesIcfTypeEnumValues Enumerates the set of values for AppBundleConfigurationPropertiesIcfTypeEnum
func GetAppBundleConfigurationPropertiesIcfTypeEnumValues() []AppBundleConfigurationPropertiesIcfTypeEnum {
	values := make([]AppBundleConfigurationPropertiesIcfTypeEnum, 0)
	for _, v := range mappingAppBundleConfigurationPropertiesIcfTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppBundleConfigurationPropertiesIcfTypeEnumStringValues Enumerates the set of values in String for AppBundleConfigurationPropertiesIcfTypeEnum
func GetAppBundleConfigurationPropertiesIcfTypeEnumStringValues() []string {
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

// GetMappingAppBundleConfigurationPropertiesIcfTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppBundleConfigurationPropertiesIcfTypeEnum(val string) (AppBundleConfigurationPropertiesIcfTypeEnum, bool) {
	enum, ok := mappingAppBundleConfigurationPropertiesIcfTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
