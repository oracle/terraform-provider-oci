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

// AppExtensionMulticloudServiceAppApp This extension defines attributes specific to Apps that represent instances of Multicloud Service App
type AppExtensionMulticloudServiceAppApp struct {

	// Specifies the service type for which the application is configured for multicloud integration. For applicable external service types, app will invoke multicloud service for runtime operations
	// **Added In:** 2301202328
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	MulticloudServiceType AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum `mandatory:"true" json:"multicloudServiceType"`

	// The multicloud platform service URL which the application will invoke for runtime operations such as AWSCredentials api invocation
	// **Added In:** 2301202328
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	MulticloudPlatformUrl *string `mandatory:"false" json:"multicloudPlatformUrl"`
}

func (m AppExtensionMulticloudServiceAppApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionMulticloudServiceAppApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum(string(m.MulticloudServiceType)); !ok && m.MulticloudServiceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MulticloudServiceType: %s. Supported values are: %s.", m.MulticloudServiceType, strings.Join(GetAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum Enum with underlying type: string
type AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum string

// Set of constants representing the allowable values for AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum
const (
	AppExtensionMulticloudServiceAppAppMulticloudServiceTypeAwscognito AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum = "AWSCognito"
)

var mappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum = map[string]AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum{
	"AWSCognito": AppExtensionMulticloudServiceAppAppMulticloudServiceTypeAwscognito,
}

var mappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumLowerCase = map[string]AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum{
	"awscognito": AppExtensionMulticloudServiceAppAppMulticloudServiceTypeAwscognito,
}

// GetAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumValues Enumerates the set of values for AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum
func GetAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumValues() []AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum {
	values := make([]AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum, 0)
	for _, v := range mappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumStringValues Enumerates the set of values in String for AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum
func GetAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumStringValues() []string {
	return []string{
		"AWSCognito",
	}
}

// GetMappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum(val string) (AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum, bool) {
	enum, ok := mappingAppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
