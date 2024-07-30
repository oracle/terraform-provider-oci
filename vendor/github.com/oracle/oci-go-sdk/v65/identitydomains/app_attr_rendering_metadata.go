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

// AppAttrRenderingMetadata Label for the attribute to be shown in the UI.
type AppAttrRenderingMetadata struct {

	// Name of the attribute.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// Label for the attribute to be shown in the UI.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Label *string `mandatory:"false" json:"label"`

	// Help text for the attribute. It can contain HTML tags.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Helptext *string `mandatory:"false" json:"helptext"`

	// UI widget to use for the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Widget AppAttrRenderingMetadataWidgetEnum `mandatory:"false" json:"widget,omitempty"`

	// Data type of the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Datatype *string `mandatory:"false" json:"datatype"`

	// UI widget to use for the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Section AppAttrRenderingMetadataSectionEnum `mandatory:"false" json:"section,omitempty"`

	// Data type of the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	Order *int `mandatory:"false" json:"order"`

	// Attribute is required or optional.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Required *bool `mandatory:"false" json:"required"`

	// Regular expression of the attribute for validation.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Regexp *string `mandatory:"false" json:"regexp"`

	// Is the attribute readOnly.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ReadOnly *bool `mandatory:"false" json:"readOnly"`

	// Indicates whether the attribute is to be shown on the application creation UI.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Visible *bool `mandatory:"false" json:"visible"`

	// Minimum length of the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinLength *int `mandatory:"false" json:"minLength"`

	// Maximum length of the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxLength *int `mandatory:"false" json:"maxLength"`

	// Minimum size of the attribute..
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinSize *int `mandatory:"false" json:"minSize"`

	// Maximum size of the attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxSize *int `mandatory:"false" json:"maxSize"`
}

func (m AppAttrRenderingMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppAttrRenderingMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppAttrRenderingMetadataWidgetEnum(string(m.Widget)); !ok && m.Widget != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Widget: %s. Supported values are: %s.", m.Widget, strings.Join(GetAppAttrRenderingMetadataWidgetEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppAttrRenderingMetadataSectionEnum(string(m.Section)); !ok && m.Section != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Section: %s. Supported values are: %s.", m.Section, strings.Join(GetAppAttrRenderingMetadataSectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppAttrRenderingMetadataWidgetEnum Enum with underlying type: string
type AppAttrRenderingMetadataWidgetEnum string

// Set of constants representing the allowable values for AppAttrRenderingMetadataWidgetEnum
const (
	AppAttrRenderingMetadataWidgetInputtext AppAttrRenderingMetadataWidgetEnum = "inputtext"
	AppAttrRenderingMetadataWidgetCheckbox  AppAttrRenderingMetadataWidgetEnum = "checkbox"
	AppAttrRenderingMetadataWidgetTextarea  AppAttrRenderingMetadataWidgetEnum = "textarea"
)

var mappingAppAttrRenderingMetadataWidgetEnum = map[string]AppAttrRenderingMetadataWidgetEnum{
	"inputtext": AppAttrRenderingMetadataWidgetInputtext,
	"checkbox":  AppAttrRenderingMetadataWidgetCheckbox,
	"textarea":  AppAttrRenderingMetadataWidgetTextarea,
}

var mappingAppAttrRenderingMetadataWidgetEnumLowerCase = map[string]AppAttrRenderingMetadataWidgetEnum{
	"inputtext": AppAttrRenderingMetadataWidgetInputtext,
	"checkbox":  AppAttrRenderingMetadataWidgetCheckbox,
	"textarea":  AppAttrRenderingMetadataWidgetTextarea,
}

// GetAppAttrRenderingMetadataWidgetEnumValues Enumerates the set of values for AppAttrRenderingMetadataWidgetEnum
func GetAppAttrRenderingMetadataWidgetEnumValues() []AppAttrRenderingMetadataWidgetEnum {
	values := make([]AppAttrRenderingMetadataWidgetEnum, 0)
	for _, v := range mappingAppAttrRenderingMetadataWidgetEnum {
		values = append(values, v)
	}
	return values
}

// GetAppAttrRenderingMetadataWidgetEnumStringValues Enumerates the set of values in String for AppAttrRenderingMetadataWidgetEnum
func GetAppAttrRenderingMetadataWidgetEnumStringValues() []string {
	return []string{
		"inputtext",
		"checkbox",
		"textarea",
	}
}

// GetMappingAppAttrRenderingMetadataWidgetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppAttrRenderingMetadataWidgetEnum(val string) (AppAttrRenderingMetadataWidgetEnum, bool) {
	enum, ok := mappingAppAttrRenderingMetadataWidgetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppAttrRenderingMetadataSectionEnum Enum with underlying type: string
type AppAttrRenderingMetadataSectionEnum string

// Set of constants representing the allowable values for AppAttrRenderingMetadataSectionEnum
const (
	AppAttrRenderingMetadataSectionSaml    AppAttrRenderingMetadataSectionEnum = "saml"
	AppAttrRenderingMetadataSectionGeneral AppAttrRenderingMetadataSectionEnum = "general"
)

var mappingAppAttrRenderingMetadataSectionEnum = map[string]AppAttrRenderingMetadataSectionEnum{
	"saml":    AppAttrRenderingMetadataSectionSaml,
	"general": AppAttrRenderingMetadataSectionGeneral,
}

var mappingAppAttrRenderingMetadataSectionEnumLowerCase = map[string]AppAttrRenderingMetadataSectionEnum{
	"saml":    AppAttrRenderingMetadataSectionSaml,
	"general": AppAttrRenderingMetadataSectionGeneral,
}

// GetAppAttrRenderingMetadataSectionEnumValues Enumerates the set of values for AppAttrRenderingMetadataSectionEnum
func GetAppAttrRenderingMetadataSectionEnumValues() []AppAttrRenderingMetadataSectionEnum {
	values := make([]AppAttrRenderingMetadataSectionEnum, 0)
	for _, v := range mappingAppAttrRenderingMetadataSectionEnum {
		values = append(values, v)
	}
	return values
}

// GetAppAttrRenderingMetadataSectionEnumStringValues Enumerates the set of values in String for AppAttrRenderingMetadataSectionEnum
func GetAppAttrRenderingMetadataSectionEnumStringValues() []string {
	return []string{
		"saml",
		"general",
	}
}

// GetMappingAppAttrRenderingMetadataSectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppAttrRenderingMetadataSectionEnum(val string) (AppAttrRenderingMetadataSectionEnum, bool) {
	enum, ok := mappingAppAttrRenderingMetadataSectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
