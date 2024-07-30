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

// AppExtensionFormFillAppApp This extension provides attributes for Form-Fill facet of App
type AppExtensionFormFillAppApp struct {

	// Type of the FormFill application like WebApplication, MainFrameApplication, WindowsApplication. Initially, we will support only WebApplication.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FormType AppExtensionFormFillAppAppFormTypeEnum `mandatory:"false" json:"formType,omitempty"`

	// Credential Sharing Group to which this form-fill application belongs.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FormCredentialSharingGroupID *string `mandatory:"false" json:"formCredentialSharingGroupID"`

	// If true, indicates that system is allowed to show the password in plain-text for this account after re-authentication.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	RevealPasswordOnForm *bool `mandatory:"false" json:"revealPasswordOnForm"`

	// Format for generating a username.  This value can be Username or Email Address; any other value will be treated as a custom expression.  A custom expression may combine 'concat' and 'substring' operations with literals and with any attribute of the Oracle Identity Cloud Service user.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	UserNameFormTemplate *string `mandatory:"false" json:"userNameFormTemplate"`

	// Indicates the custom expression, which can combine concat and substring operations with literals and with any attribute of the Oracle Identity Cloud Service User
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	UserNameFormExpression *string `mandatory:"false" json:"userNameFormExpression"`

	// Indicates how FormFill obtains the username and password of the account that FormFill will use to sign into the target App.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FormCredMethod AppExtensionFormFillAppAppFormCredMethodEnum `mandatory:"false" json:"formCredMethod,omitempty"`

	// FormFill Application Configuration CLOB which has to be maintained in Form-Fill APP for legacy code to do Form-Fill injection
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Configuration *string `mandatory:"false" json:"configuration"`

	// If true, indicates that each of the Form-Fill-related attributes that can be inherited from the template actually will be inherited from the template. If false, indicates that the AppTemplate on which this App is based has disabled inheritance for these Form-Fill-related attributes.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	SyncFromTemplate *bool `mandatory:"false" json:"syncFromTemplate"`

	// A list of application-formURLs that FormFill should match against any formUrl that the user-specifies when signing in to the target service.  Each item in the list also indicates how FormFill should interpret that formUrl.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [formUrl]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	FormFillUrlMatch []AppFormFillUrlMatch `mandatory:"false" json:"formFillUrlMatch"`
}

func (m AppExtensionFormFillAppApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionFormFillAppApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppExtensionFormFillAppAppFormTypeEnum(string(m.FormType)); !ok && m.FormType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FormType: %s. Supported values are: %s.", m.FormType, strings.Join(GetAppExtensionFormFillAppAppFormTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionFormFillAppAppFormCredMethodEnum(string(m.FormCredMethod)); !ok && m.FormCredMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FormCredMethod: %s. Supported values are: %s.", m.FormCredMethod, strings.Join(GetAppExtensionFormFillAppAppFormCredMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppExtensionFormFillAppAppFormTypeEnum Enum with underlying type: string
type AppExtensionFormFillAppAppFormTypeEnum string

// Set of constants representing the allowable values for AppExtensionFormFillAppAppFormTypeEnum
const (
	AppExtensionFormFillAppAppFormTypeWebapplication AppExtensionFormFillAppAppFormTypeEnum = "WebApplication"
)

var mappingAppExtensionFormFillAppAppFormTypeEnum = map[string]AppExtensionFormFillAppAppFormTypeEnum{
	"WebApplication": AppExtensionFormFillAppAppFormTypeWebapplication,
}

var mappingAppExtensionFormFillAppAppFormTypeEnumLowerCase = map[string]AppExtensionFormFillAppAppFormTypeEnum{
	"webapplication": AppExtensionFormFillAppAppFormTypeWebapplication,
}

// GetAppExtensionFormFillAppAppFormTypeEnumValues Enumerates the set of values for AppExtensionFormFillAppAppFormTypeEnum
func GetAppExtensionFormFillAppAppFormTypeEnumValues() []AppExtensionFormFillAppAppFormTypeEnum {
	values := make([]AppExtensionFormFillAppAppFormTypeEnum, 0)
	for _, v := range mappingAppExtensionFormFillAppAppFormTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionFormFillAppAppFormTypeEnumStringValues Enumerates the set of values in String for AppExtensionFormFillAppAppFormTypeEnum
func GetAppExtensionFormFillAppAppFormTypeEnumStringValues() []string {
	return []string{
		"WebApplication",
	}
}

// GetMappingAppExtensionFormFillAppAppFormTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionFormFillAppAppFormTypeEnum(val string) (AppExtensionFormFillAppAppFormTypeEnum, bool) {
	enum, ok := mappingAppExtensionFormFillAppAppFormTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionFormFillAppAppFormCredMethodEnum Enum with underlying type: string
type AppExtensionFormFillAppAppFormCredMethodEnum string

// Set of constants representing the allowable values for AppExtensionFormFillAppAppFormCredMethodEnum
const (
	AppExtensionFormFillAppAppFormCredMethodAdminSetsCredentials       AppExtensionFormFillAppAppFormCredMethodEnum = "ADMIN_SETS_CREDENTIALS"
	AppExtensionFormFillAppAppFormCredMethodAdminSetsSharedCredentials AppExtensionFormFillAppAppFormCredMethodEnum = "ADMIN_SETS_SHARED_CREDENTIALS"
	AppExtensionFormFillAppAppFormCredMethodUserSetsPasswordOnly       AppExtensionFormFillAppAppFormCredMethodEnum = "USER_SETS_PASSWORD_ONLY"
	AppExtensionFormFillAppAppFormCredMethodUserSetsCredentials        AppExtensionFormFillAppAppFormCredMethodEnum = "USER_SETS_CREDENTIALS"
)

var mappingAppExtensionFormFillAppAppFormCredMethodEnum = map[string]AppExtensionFormFillAppAppFormCredMethodEnum{
	"ADMIN_SETS_CREDENTIALS":        AppExtensionFormFillAppAppFormCredMethodAdminSetsCredentials,
	"ADMIN_SETS_SHARED_CREDENTIALS": AppExtensionFormFillAppAppFormCredMethodAdminSetsSharedCredentials,
	"USER_SETS_PASSWORD_ONLY":       AppExtensionFormFillAppAppFormCredMethodUserSetsPasswordOnly,
	"USER_SETS_CREDENTIALS":         AppExtensionFormFillAppAppFormCredMethodUserSetsCredentials,
}

var mappingAppExtensionFormFillAppAppFormCredMethodEnumLowerCase = map[string]AppExtensionFormFillAppAppFormCredMethodEnum{
	"admin_sets_credentials":        AppExtensionFormFillAppAppFormCredMethodAdminSetsCredentials,
	"admin_sets_shared_credentials": AppExtensionFormFillAppAppFormCredMethodAdminSetsSharedCredentials,
	"user_sets_password_only":       AppExtensionFormFillAppAppFormCredMethodUserSetsPasswordOnly,
	"user_sets_credentials":         AppExtensionFormFillAppAppFormCredMethodUserSetsCredentials,
}

// GetAppExtensionFormFillAppAppFormCredMethodEnumValues Enumerates the set of values for AppExtensionFormFillAppAppFormCredMethodEnum
func GetAppExtensionFormFillAppAppFormCredMethodEnumValues() []AppExtensionFormFillAppAppFormCredMethodEnum {
	values := make([]AppExtensionFormFillAppAppFormCredMethodEnum, 0)
	for _, v := range mappingAppExtensionFormFillAppAppFormCredMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionFormFillAppAppFormCredMethodEnumStringValues Enumerates the set of values in String for AppExtensionFormFillAppAppFormCredMethodEnum
func GetAppExtensionFormFillAppAppFormCredMethodEnumStringValues() []string {
	return []string{
		"ADMIN_SETS_CREDENTIALS",
		"ADMIN_SETS_SHARED_CREDENTIALS",
		"USER_SETS_PASSWORD_ONLY",
		"USER_SETS_CREDENTIALS",
	}
}

// GetMappingAppExtensionFormFillAppAppFormCredMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionFormFillAppAppFormCredMethodEnum(val string) (AppExtensionFormFillAppAppFormCredMethodEnum, bool) {
	enum, ok := mappingAppExtensionFormFillAppAppFormCredMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
