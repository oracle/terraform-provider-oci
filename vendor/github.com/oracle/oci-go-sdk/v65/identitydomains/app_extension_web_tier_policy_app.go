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

// AppExtensionWebTierPolicyApp WebTier Policy
type AppExtensionWebTierPolicyApp struct {

	// Store the web tier policy for an application as a string in Javascript Object Notification (JSON) format.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	WebTierPolicyJson *string `mandatory:"false" json:"webTierPolicyJson"`

	// Webtier policy AZ Control
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	WebTierPolicyAZControl AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum `mandatory:"false" json:"webTierPolicyAZControl,omitempty"`

	// If this Attribute is true, resource ref id and resource ref name attributes will we included in wtp json response.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ResourceRef *bool `mandatory:"false" json:"resourceRef"`
}

func (m AppExtensionWebTierPolicyApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionWebTierPolicyApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum(string(m.WebTierPolicyAZControl)); !ok && m.WebTierPolicyAZControl != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WebTierPolicyAZControl: %s. Supported values are: %s.", m.WebTierPolicyAZControl, strings.Join(GetAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum Enum with underlying type: string
type AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum string

// Set of constants representing the allowable values for AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum
const (
	AppExtensionWebTierPolicyAppWebTierPolicyAZControlServer AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum = "server"
	AppExtensionWebTierPolicyAppWebTierPolicyAZControlLocal  AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum = "local"
)

var mappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum = map[string]AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum{
	"server": AppExtensionWebTierPolicyAppWebTierPolicyAZControlServer,
	"local":  AppExtensionWebTierPolicyAppWebTierPolicyAZControlLocal,
}

var mappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumLowerCase = map[string]AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum{
	"server": AppExtensionWebTierPolicyAppWebTierPolicyAZControlServer,
	"local":  AppExtensionWebTierPolicyAppWebTierPolicyAZControlLocal,
}

// GetAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumValues Enumerates the set of values for AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum
func GetAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumValues() []AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum {
	values := make([]AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum, 0)
	for _, v := range mappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumStringValues Enumerates the set of values in String for AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum
func GetAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumStringValues() []string {
	return []string{
		"server",
		"local",
	}
}

// GetMappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum(val string) (AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum, bool) {
	enum, ok := mappingAppExtensionWebTierPolicyAppWebTierPolicyAZControlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
