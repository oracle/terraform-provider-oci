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

// AppExtensionOpcServiceApp This extension defines attributes specific to Apps that represent instances of an Oracle Public Cloud (OPC) service.
type AppExtensionOpcServiceApp struct {

	// This value specifies the unique identifier assigned to an instance of an Oracle Public Cloud service app.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: server
	ServiceInstanceIdentifier *string `mandatory:"false" json:"serviceInstanceIdentifier"`

	// This value identifies the OPC region in which the service is running.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Region *string `mandatory:"false" json:"region"`

	// Current Federation Mode
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	CurrentFederationMode AppExtensionOpcServiceAppCurrentFederationModeEnum `mandatory:"false" json:"currentFederationMode,omitempty"`

	// Current Synchronization Mode
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	CurrentSynchronizationMode AppExtensionOpcServiceAppCurrentSynchronizationModeEnum `mandatory:"false" json:"currentSynchronizationMode,omitempty"`

	// Next Federation Mode
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	NextFederationMode AppExtensionOpcServiceAppNextFederationModeEnum `mandatory:"false" json:"nextFederationMode,omitempty"`

	// Next Synchronization Mode
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	NextSynchronizationMode AppExtensionOpcServiceAppNextSynchronizationModeEnum `mandatory:"false" json:"nextSynchronizationMode,omitempty"`

	// If true, indicates that enablement is in progress started but not completed
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	EnablingNextFedSyncModes *bool `mandatory:"false" json:"enablingNextFedSyncModes"`
}

func (m AppExtensionOpcServiceApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionOpcServiceApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppExtensionOpcServiceAppCurrentFederationModeEnum(string(m.CurrentFederationMode)); !ok && m.CurrentFederationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentFederationMode: %s. Supported values are: %s.", m.CurrentFederationMode, strings.Join(GetAppExtensionOpcServiceAppCurrentFederationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnum(string(m.CurrentSynchronizationMode)); !ok && m.CurrentSynchronizationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentSynchronizationMode: %s. Supported values are: %s.", m.CurrentSynchronizationMode, strings.Join(GetAppExtensionOpcServiceAppCurrentSynchronizationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionOpcServiceAppNextFederationModeEnum(string(m.NextFederationMode)); !ok && m.NextFederationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextFederationMode: %s. Supported values are: %s.", m.NextFederationMode, strings.Join(GetAppExtensionOpcServiceAppNextFederationModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionOpcServiceAppNextSynchronizationModeEnum(string(m.NextSynchronizationMode)); !ok && m.NextSynchronizationMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextSynchronizationMode: %s. Supported values are: %s.", m.NextSynchronizationMode, strings.Join(GetAppExtensionOpcServiceAppNextSynchronizationModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppExtensionOpcServiceAppCurrentFederationModeEnum Enum with underlying type: string
type AppExtensionOpcServiceAppCurrentFederationModeEnum string

// Set of constants representing the allowable values for AppExtensionOpcServiceAppCurrentFederationModeEnum
const (
	AppExtensionOpcServiceAppCurrentFederationModeNone                  AppExtensionOpcServiceAppCurrentFederationModeEnum = "None"
	AppExtensionOpcServiceAppCurrentFederationModeAppasserviceprovider  AppExtensionOpcServiceAppCurrentFederationModeEnum = "AppAsServiceProvider"
	AppExtensionOpcServiceAppCurrentFederationModeAppasidentityprovider AppExtensionOpcServiceAppCurrentFederationModeEnum = "AppAsIdentityProvider"
)

var mappingAppExtensionOpcServiceAppCurrentFederationModeEnum = map[string]AppExtensionOpcServiceAppCurrentFederationModeEnum{
	"None":                  AppExtensionOpcServiceAppCurrentFederationModeNone,
	"AppAsServiceProvider":  AppExtensionOpcServiceAppCurrentFederationModeAppasserviceprovider,
	"AppAsIdentityProvider": AppExtensionOpcServiceAppCurrentFederationModeAppasidentityprovider,
}

var mappingAppExtensionOpcServiceAppCurrentFederationModeEnumLowerCase = map[string]AppExtensionOpcServiceAppCurrentFederationModeEnum{
	"none":                  AppExtensionOpcServiceAppCurrentFederationModeNone,
	"appasserviceprovider":  AppExtensionOpcServiceAppCurrentFederationModeAppasserviceprovider,
	"appasidentityprovider": AppExtensionOpcServiceAppCurrentFederationModeAppasidentityprovider,
}

// GetAppExtensionOpcServiceAppCurrentFederationModeEnumValues Enumerates the set of values for AppExtensionOpcServiceAppCurrentFederationModeEnum
func GetAppExtensionOpcServiceAppCurrentFederationModeEnumValues() []AppExtensionOpcServiceAppCurrentFederationModeEnum {
	values := make([]AppExtensionOpcServiceAppCurrentFederationModeEnum, 0)
	for _, v := range mappingAppExtensionOpcServiceAppCurrentFederationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionOpcServiceAppCurrentFederationModeEnumStringValues Enumerates the set of values in String for AppExtensionOpcServiceAppCurrentFederationModeEnum
func GetAppExtensionOpcServiceAppCurrentFederationModeEnumStringValues() []string {
	return []string{
		"None",
		"AppAsServiceProvider",
		"AppAsIdentityProvider",
	}
}

// GetMappingAppExtensionOpcServiceAppCurrentFederationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionOpcServiceAppCurrentFederationModeEnum(val string) (AppExtensionOpcServiceAppCurrentFederationModeEnum, bool) {
	enum, ok := mappingAppExtensionOpcServiceAppCurrentFederationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionOpcServiceAppCurrentSynchronizationModeEnum Enum with underlying type: string
type AppExtensionOpcServiceAppCurrentSynchronizationModeEnum string

// Set of constants representing the allowable values for AppExtensionOpcServiceAppCurrentSynchronizationModeEnum
const (
	AppExtensionOpcServiceAppCurrentSynchronizationModeNone        AppExtensionOpcServiceAppCurrentSynchronizationModeEnum = "None"
	AppExtensionOpcServiceAppCurrentSynchronizationModeAppastarget AppExtensionOpcServiceAppCurrentSynchronizationModeEnum = "AppAsTarget"
	AppExtensionOpcServiceAppCurrentSynchronizationModeAppassource AppExtensionOpcServiceAppCurrentSynchronizationModeEnum = "AppAsSource"
)

var mappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnum = map[string]AppExtensionOpcServiceAppCurrentSynchronizationModeEnum{
	"None":        AppExtensionOpcServiceAppCurrentSynchronizationModeNone,
	"AppAsTarget": AppExtensionOpcServiceAppCurrentSynchronizationModeAppastarget,
	"AppAsSource": AppExtensionOpcServiceAppCurrentSynchronizationModeAppassource,
}

var mappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnumLowerCase = map[string]AppExtensionOpcServiceAppCurrentSynchronizationModeEnum{
	"none":        AppExtensionOpcServiceAppCurrentSynchronizationModeNone,
	"appastarget": AppExtensionOpcServiceAppCurrentSynchronizationModeAppastarget,
	"appassource": AppExtensionOpcServiceAppCurrentSynchronizationModeAppassource,
}

// GetAppExtensionOpcServiceAppCurrentSynchronizationModeEnumValues Enumerates the set of values for AppExtensionOpcServiceAppCurrentSynchronizationModeEnum
func GetAppExtensionOpcServiceAppCurrentSynchronizationModeEnumValues() []AppExtensionOpcServiceAppCurrentSynchronizationModeEnum {
	values := make([]AppExtensionOpcServiceAppCurrentSynchronizationModeEnum, 0)
	for _, v := range mappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionOpcServiceAppCurrentSynchronizationModeEnumStringValues Enumerates the set of values in String for AppExtensionOpcServiceAppCurrentSynchronizationModeEnum
func GetAppExtensionOpcServiceAppCurrentSynchronizationModeEnumStringValues() []string {
	return []string{
		"None",
		"AppAsTarget",
		"AppAsSource",
	}
}

// GetMappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnum(val string) (AppExtensionOpcServiceAppCurrentSynchronizationModeEnum, bool) {
	enum, ok := mappingAppExtensionOpcServiceAppCurrentSynchronizationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionOpcServiceAppNextFederationModeEnum Enum with underlying type: string
type AppExtensionOpcServiceAppNextFederationModeEnum string

// Set of constants representing the allowable values for AppExtensionOpcServiceAppNextFederationModeEnum
const (
	AppExtensionOpcServiceAppNextFederationModeNone                  AppExtensionOpcServiceAppNextFederationModeEnum = "None"
	AppExtensionOpcServiceAppNextFederationModeAppasserviceprovider  AppExtensionOpcServiceAppNextFederationModeEnum = "AppAsServiceProvider"
	AppExtensionOpcServiceAppNextFederationModeAppasidentityprovider AppExtensionOpcServiceAppNextFederationModeEnum = "AppAsIdentityProvider"
)

var mappingAppExtensionOpcServiceAppNextFederationModeEnum = map[string]AppExtensionOpcServiceAppNextFederationModeEnum{
	"None":                  AppExtensionOpcServiceAppNextFederationModeNone,
	"AppAsServiceProvider":  AppExtensionOpcServiceAppNextFederationModeAppasserviceprovider,
	"AppAsIdentityProvider": AppExtensionOpcServiceAppNextFederationModeAppasidentityprovider,
}

var mappingAppExtensionOpcServiceAppNextFederationModeEnumLowerCase = map[string]AppExtensionOpcServiceAppNextFederationModeEnum{
	"none":                  AppExtensionOpcServiceAppNextFederationModeNone,
	"appasserviceprovider":  AppExtensionOpcServiceAppNextFederationModeAppasserviceprovider,
	"appasidentityprovider": AppExtensionOpcServiceAppNextFederationModeAppasidentityprovider,
}

// GetAppExtensionOpcServiceAppNextFederationModeEnumValues Enumerates the set of values for AppExtensionOpcServiceAppNextFederationModeEnum
func GetAppExtensionOpcServiceAppNextFederationModeEnumValues() []AppExtensionOpcServiceAppNextFederationModeEnum {
	values := make([]AppExtensionOpcServiceAppNextFederationModeEnum, 0)
	for _, v := range mappingAppExtensionOpcServiceAppNextFederationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionOpcServiceAppNextFederationModeEnumStringValues Enumerates the set of values in String for AppExtensionOpcServiceAppNextFederationModeEnum
func GetAppExtensionOpcServiceAppNextFederationModeEnumStringValues() []string {
	return []string{
		"None",
		"AppAsServiceProvider",
		"AppAsIdentityProvider",
	}
}

// GetMappingAppExtensionOpcServiceAppNextFederationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionOpcServiceAppNextFederationModeEnum(val string) (AppExtensionOpcServiceAppNextFederationModeEnum, bool) {
	enum, ok := mappingAppExtensionOpcServiceAppNextFederationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionOpcServiceAppNextSynchronizationModeEnum Enum with underlying type: string
type AppExtensionOpcServiceAppNextSynchronizationModeEnum string

// Set of constants representing the allowable values for AppExtensionOpcServiceAppNextSynchronizationModeEnum
const (
	AppExtensionOpcServiceAppNextSynchronizationModeNone        AppExtensionOpcServiceAppNextSynchronizationModeEnum = "None"
	AppExtensionOpcServiceAppNextSynchronizationModeAppastarget AppExtensionOpcServiceAppNextSynchronizationModeEnum = "AppAsTarget"
	AppExtensionOpcServiceAppNextSynchronizationModeAppassource AppExtensionOpcServiceAppNextSynchronizationModeEnum = "AppAsSource"
)

var mappingAppExtensionOpcServiceAppNextSynchronizationModeEnum = map[string]AppExtensionOpcServiceAppNextSynchronizationModeEnum{
	"None":        AppExtensionOpcServiceAppNextSynchronizationModeNone,
	"AppAsTarget": AppExtensionOpcServiceAppNextSynchronizationModeAppastarget,
	"AppAsSource": AppExtensionOpcServiceAppNextSynchronizationModeAppassource,
}

var mappingAppExtensionOpcServiceAppNextSynchronizationModeEnumLowerCase = map[string]AppExtensionOpcServiceAppNextSynchronizationModeEnum{
	"none":        AppExtensionOpcServiceAppNextSynchronizationModeNone,
	"appastarget": AppExtensionOpcServiceAppNextSynchronizationModeAppastarget,
	"appassource": AppExtensionOpcServiceAppNextSynchronizationModeAppassource,
}

// GetAppExtensionOpcServiceAppNextSynchronizationModeEnumValues Enumerates the set of values for AppExtensionOpcServiceAppNextSynchronizationModeEnum
func GetAppExtensionOpcServiceAppNextSynchronizationModeEnumValues() []AppExtensionOpcServiceAppNextSynchronizationModeEnum {
	values := make([]AppExtensionOpcServiceAppNextSynchronizationModeEnum, 0)
	for _, v := range mappingAppExtensionOpcServiceAppNextSynchronizationModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionOpcServiceAppNextSynchronizationModeEnumStringValues Enumerates the set of values in String for AppExtensionOpcServiceAppNextSynchronizationModeEnum
func GetAppExtensionOpcServiceAppNextSynchronizationModeEnumStringValues() []string {
	return []string{
		"None",
		"AppAsTarget",
		"AppAsSource",
	}
}

// GetMappingAppExtensionOpcServiceAppNextSynchronizationModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionOpcServiceAppNextSynchronizationModeEnum(val string) (AppExtensionOpcServiceAppNextSynchronizationModeEnum, bool) {
	enum, ok := mappingAppExtensionOpcServiceAppNextSynchronizationModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
