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

// AppExtensionFormFillAppTemplateAppTemplate This extension provides attributes for Form-Fill facet of AppTemplate
type AppExtensionFormFillAppTemplateAppTemplate struct {

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
	FormType AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum `mandatory:"false" json:"formType,omitempty"`

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
	FormCredMethod AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum `mandatory:"false" json:"formCredMethod,omitempty"`

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

	// If true, indicates that each of the Form-Fill-related attributes that can be inherited from the template actually will be inherited from the template. If false, indicates that the AppTemplate disabled inheritance for these Form-Fill-related attributes.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
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

func (m AppExtensionFormFillAppTemplateAppTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionFormFillAppTemplateAppTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnum(string(m.FormType)); !ok && m.FormType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FormType: %s. Supported values are: %s.", m.FormType, strings.Join(GetAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum(string(m.FormCredMethod)); !ok && m.FormCredMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FormCredMethod: %s. Supported values are: %s.", m.FormCredMethod, strings.Join(GetAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum Enum with underlying type: string
type AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum string

// Set of constants representing the allowable values for AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum
const (
	AppExtensionFormFillAppTemplateAppTemplateFormTypeWebapplication AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum = "WebApplication"
)

var mappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnum = map[string]AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum{
	"WebApplication": AppExtensionFormFillAppTemplateAppTemplateFormTypeWebapplication,
}

var mappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumLowerCase = map[string]AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum{
	"webapplication": AppExtensionFormFillAppTemplateAppTemplateFormTypeWebapplication,
}

// GetAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumValues Enumerates the set of values for AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum
func GetAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumValues() []AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum {
	values := make([]AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum, 0)
	for _, v := range mappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumStringValues Enumerates the set of values in String for AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum
func GetAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumStringValues() []string {
	return []string{
		"WebApplication",
	}
}

// GetMappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnum(val string) (AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum, bool) {
	enum, ok := mappingAppExtensionFormFillAppTemplateAppTemplateFormTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum Enum with underlying type: string
type AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum string

// Set of constants representing the allowable values for AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum
const (
	AppExtensionFormFillAppTemplateAppTemplateFormCredMethodAdminSetsCredentials           AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum = "ADMIN_SETS_CREDENTIALS"
	AppExtensionFormFillAppTemplateAppTemplateFormCredMethodAdminSetsSharedCredentials     AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum = "ADMIN_SETS_SHARED_CREDENTIALS"
	AppExtensionFormFillAppTemplateAppTemplateFormCredMethodUserSetsPasswordOnly           AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum = "USER_SETS_PASSWORD_ONLY"
	AppExtensionFormFillAppTemplateAppTemplateFormCredMethodUserSetsCredentials            AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum = "USER_SETS_CREDENTIALS"
	AppExtensionFormFillAppTemplateAppTemplateFormCredMethodSsoCredentialsAsAppCredentials AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum = "SSO_CREDENTIALS_AS_APP_CREDENTIALS"
)

var mappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum = map[string]AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum{
	"ADMIN_SETS_CREDENTIALS":             AppExtensionFormFillAppTemplateAppTemplateFormCredMethodAdminSetsCredentials,
	"ADMIN_SETS_SHARED_CREDENTIALS":      AppExtensionFormFillAppTemplateAppTemplateFormCredMethodAdminSetsSharedCredentials,
	"USER_SETS_PASSWORD_ONLY":            AppExtensionFormFillAppTemplateAppTemplateFormCredMethodUserSetsPasswordOnly,
	"USER_SETS_CREDENTIALS":              AppExtensionFormFillAppTemplateAppTemplateFormCredMethodUserSetsCredentials,
	"SSO_CREDENTIALS_AS_APP_CREDENTIALS": AppExtensionFormFillAppTemplateAppTemplateFormCredMethodSsoCredentialsAsAppCredentials,
}

var mappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumLowerCase = map[string]AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum{
	"admin_sets_credentials":             AppExtensionFormFillAppTemplateAppTemplateFormCredMethodAdminSetsCredentials,
	"admin_sets_shared_credentials":      AppExtensionFormFillAppTemplateAppTemplateFormCredMethodAdminSetsSharedCredentials,
	"user_sets_password_only":            AppExtensionFormFillAppTemplateAppTemplateFormCredMethodUserSetsPasswordOnly,
	"user_sets_credentials":              AppExtensionFormFillAppTemplateAppTemplateFormCredMethodUserSetsCredentials,
	"sso_credentials_as_app_credentials": AppExtensionFormFillAppTemplateAppTemplateFormCredMethodSsoCredentialsAsAppCredentials,
}

// GetAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumValues Enumerates the set of values for AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum
func GetAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumValues() []AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum {
	values := make([]AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum, 0)
	for _, v := range mappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumStringValues Enumerates the set of values in String for AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum
func GetAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumStringValues() []string {
	return []string{
		"ADMIN_SETS_CREDENTIALS",
		"ADMIN_SETS_SHARED_CREDENTIALS",
		"USER_SETS_PASSWORD_ONLY",
		"USER_SETS_CREDENTIALS",
		"SSO_CREDENTIALS_AS_APP_CREDENTIALS",
	}
}

// GetMappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum(val string) (AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum, bool) {
	enum, ok := mappingAppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
