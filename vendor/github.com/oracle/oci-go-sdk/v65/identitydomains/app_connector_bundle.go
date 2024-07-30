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

// AppConnectorBundle ConnectorBundle
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AppConnectorBundle struct {

	// ConnectorBundle identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Connector Bundle type. Allowed values are ConnectorBundle, LocalConnectorBundle.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsDefaultValue: ConnectorBundle
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type AppConnectorBundleTypeEnum `mandatory:"true" json:"type"`

	// ConnectorBundle URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// ConnectorBundle display name
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// Unique Well-known identifier used to reference connector bundle.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	WellKnownId *string `mandatory:"false" json:"wellKnownId"`
}

func (m AppConnectorBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppConnectorBundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAppConnectorBundleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAppConnectorBundleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppConnectorBundleTypeEnum Enum with underlying type: string
type AppConnectorBundleTypeEnum string

// Set of constants representing the allowable values for AppConnectorBundleTypeEnum
const (
	AppConnectorBundleTypeConnectorbundle      AppConnectorBundleTypeEnum = "ConnectorBundle"
	AppConnectorBundleTypeLocalconnectorbundle AppConnectorBundleTypeEnum = "LocalConnectorBundle"
)

var mappingAppConnectorBundleTypeEnum = map[string]AppConnectorBundleTypeEnum{
	"ConnectorBundle":      AppConnectorBundleTypeConnectorbundle,
	"LocalConnectorBundle": AppConnectorBundleTypeLocalconnectorbundle,
}

var mappingAppConnectorBundleTypeEnumLowerCase = map[string]AppConnectorBundleTypeEnum{
	"connectorbundle":      AppConnectorBundleTypeConnectorbundle,
	"localconnectorbundle": AppConnectorBundleTypeLocalconnectorbundle,
}

// GetAppConnectorBundleTypeEnumValues Enumerates the set of values for AppConnectorBundleTypeEnum
func GetAppConnectorBundleTypeEnumValues() []AppConnectorBundleTypeEnum {
	values := make([]AppConnectorBundleTypeEnum, 0)
	for _, v := range mappingAppConnectorBundleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppConnectorBundleTypeEnumStringValues Enumerates the set of values in String for AppConnectorBundleTypeEnum
func GetAppConnectorBundleTypeEnumStringValues() []string {
	return []string{
		"ConnectorBundle",
		"LocalConnectorBundle",
	}
}

// GetMappingAppConnectorBundleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppConnectorBundleTypeEnum(val string) (AppConnectorBundleTypeEnum, bool) {
	enum, ok := mappingAppConnectorBundleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
