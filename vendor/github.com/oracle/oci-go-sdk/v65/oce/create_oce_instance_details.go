// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content Management API
//
// Oracle Content Management is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOceInstanceDetails The information about new OceInstance.
type CreateOceInstanceDetails struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OceInstance Name
	Name *string `mandatory:"true" json:"name"`

	// Tenancy Identifier
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// Identity Cloud Service access token identifying a stripe and service administrator user
	IdcsAccessToken *string `mandatory:"true" json:"idcsAccessToken"`

	// Tenancy Name
	TenancyName *string `mandatory:"true" json:"tenancyName"`

	// Object Storage Namespace of Tenancy
	ObjectStorageNamespace *string `mandatory:"true" json:"objectStorageNamespace"`

	// Admin Email for Notification
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// OceInstance description
	Description *string `mandatory:"false" json:"description"`

	IdentityStripe *IdentityStripeDetails `mandatory:"false" json:"identityStripe"`

	// Instance type based on its usage
	InstanceUsageType CreateOceInstanceDetailsInstanceUsageTypeEnum `mandatory:"false" json:"instanceUsageType,omitempty"`

	// a list of add-on features for the ocm instance
	AddOnFeatures []string `mandatory:"false" json:"addOnFeatures"`

	// Upgrade schedule type representing service to be upgraded immediately whenever latest version is released
	// or delay upgrade of the service to previous released version
	UpgradeSchedule OceInstanceUpgradeScheduleEnum `mandatory:"false" json:"upgradeSchedule,omitempty"`

	// Web Application Firewall(WAF) primary domain
	WafPrimaryDomain *string `mandatory:"false" json:"wafPrimaryDomain"`

	// Flag indicating whether the instance access is private or public
	InstanceAccessType CreateOceInstanceDetailsInstanceAccessTypeEnum `mandatory:"false" json:"instanceAccessType,omitempty"`

	// Flag indicating whether the instance license is new cloud or bring your own license
	InstanceLicenseType LicenseTypeEnum `mandatory:"false" json:"instanceLicenseType,omitempty"`

	// disaster recovery paired ragion name
	DrRegion *string `mandatory:"false" json:"drRegion"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOceInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOceInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateOceInstanceDetailsInstanceUsageTypeEnum(string(m.InstanceUsageType)); !ok && m.InstanceUsageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceUsageType: %s. Supported values are: %s.", m.InstanceUsageType, strings.Join(GetCreateOceInstanceDetailsInstanceUsageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOceInstanceUpgradeScheduleEnum(string(m.UpgradeSchedule)); !ok && m.UpgradeSchedule != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeSchedule: %s. Supported values are: %s.", m.UpgradeSchedule, strings.Join(GetOceInstanceUpgradeScheduleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateOceInstanceDetailsInstanceAccessTypeEnum(string(m.InstanceAccessType)); !ok && m.InstanceAccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceAccessType: %s. Supported values are: %s.", m.InstanceAccessType, strings.Join(GetCreateOceInstanceDetailsInstanceAccessTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.InstanceLicenseType)); !ok && m.InstanceLicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceLicenseType: %s. Supported values are: %s.", m.InstanceLicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOceInstanceDetailsInstanceUsageTypeEnum Enum with underlying type: string
type CreateOceInstanceDetailsInstanceUsageTypeEnum string

// Set of constants representing the allowable values for CreateOceInstanceDetailsInstanceUsageTypeEnum
const (
	CreateOceInstanceDetailsInstanceUsageTypePrimary    CreateOceInstanceDetailsInstanceUsageTypeEnum = "PRIMARY"
	CreateOceInstanceDetailsInstanceUsageTypeNonprimary CreateOceInstanceDetailsInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingCreateOceInstanceDetailsInstanceUsageTypeEnum = map[string]CreateOceInstanceDetailsInstanceUsageTypeEnum{
	"PRIMARY":    CreateOceInstanceDetailsInstanceUsageTypePrimary,
	"NONPRIMARY": CreateOceInstanceDetailsInstanceUsageTypeNonprimary,
}

var mappingCreateOceInstanceDetailsInstanceUsageTypeEnumLowerCase = map[string]CreateOceInstanceDetailsInstanceUsageTypeEnum{
	"primary":    CreateOceInstanceDetailsInstanceUsageTypePrimary,
	"nonprimary": CreateOceInstanceDetailsInstanceUsageTypeNonprimary,
}

// GetCreateOceInstanceDetailsInstanceUsageTypeEnumValues Enumerates the set of values for CreateOceInstanceDetailsInstanceUsageTypeEnum
func GetCreateOceInstanceDetailsInstanceUsageTypeEnumValues() []CreateOceInstanceDetailsInstanceUsageTypeEnum {
	values := make([]CreateOceInstanceDetailsInstanceUsageTypeEnum, 0)
	for _, v := range mappingCreateOceInstanceDetailsInstanceUsageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOceInstanceDetailsInstanceUsageTypeEnumStringValues Enumerates the set of values in String for CreateOceInstanceDetailsInstanceUsageTypeEnum
func GetCreateOceInstanceDetailsInstanceUsageTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"NONPRIMARY",
	}
}

// GetMappingCreateOceInstanceDetailsInstanceUsageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOceInstanceDetailsInstanceUsageTypeEnum(val string) (CreateOceInstanceDetailsInstanceUsageTypeEnum, bool) {
	enum, ok := mappingCreateOceInstanceDetailsInstanceUsageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateOceInstanceDetailsInstanceAccessTypeEnum Enum with underlying type: string
type CreateOceInstanceDetailsInstanceAccessTypeEnum string

// Set of constants representing the allowable values for CreateOceInstanceDetailsInstanceAccessTypeEnum
const (
	CreateOceInstanceDetailsInstanceAccessTypePublic  CreateOceInstanceDetailsInstanceAccessTypeEnum = "PUBLIC"
	CreateOceInstanceDetailsInstanceAccessTypePrivate CreateOceInstanceDetailsInstanceAccessTypeEnum = "PRIVATE"
)

var mappingCreateOceInstanceDetailsInstanceAccessTypeEnum = map[string]CreateOceInstanceDetailsInstanceAccessTypeEnum{
	"PUBLIC":  CreateOceInstanceDetailsInstanceAccessTypePublic,
	"PRIVATE": CreateOceInstanceDetailsInstanceAccessTypePrivate,
}

var mappingCreateOceInstanceDetailsInstanceAccessTypeEnumLowerCase = map[string]CreateOceInstanceDetailsInstanceAccessTypeEnum{
	"public":  CreateOceInstanceDetailsInstanceAccessTypePublic,
	"private": CreateOceInstanceDetailsInstanceAccessTypePrivate,
}

// GetCreateOceInstanceDetailsInstanceAccessTypeEnumValues Enumerates the set of values for CreateOceInstanceDetailsInstanceAccessTypeEnum
func GetCreateOceInstanceDetailsInstanceAccessTypeEnumValues() []CreateOceInstanceDetailsInstanceAccessTypeEnum {
	values := make([]CreateOceInstanceDetailsInstanceAccessTypeEnum, 0)
	for _, v := range mappingCreateOceInstanceDetailsInstanceAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOceInstanceDetailsInstanceAccessTypeEnumStringValues Enumerates the set of values in String for CreateOceInstanceDetailsInstanceAccessTypeEnum
func GetCreateOceInstanceDetailsInstanceAccessTypeEnumStringValues() []string {
	return []string{
		"PUBLIC",
		"PRIVATE",
	}
}

// GetMappingCreateOceInstanceDetailsInstanceAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOceInstanceDetailsInstanceAccessTypeEnum(val string) (CreateOceInstanceDetailsInstanceAccessTypeEnum, bool) {
	enum, ok := mappingCreateOceInstanceDetailsInstanceAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
