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

// OceInstanceSummary Summary of the OceInstance.
type OceInstanceSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Unique GUID identifier that is immutable on creation
	Guid *string `mandatory:"true" json:"guid"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OceInstance Name
	Name *string `mandatory:"true" json:"name"`

	// Tenancy Identifier
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// IDCS Tenancy Identifier
	IdcsTenancy *string `mandatory:"true" json:"idcsTenancy"`

	// Tenancy Name
	TenancyName *string `mandatory:"true" json:"tenancyName"`

	// Object Storage Namespace of tenancy
	ObjectStorageNamespace *string `mandatory:"true" json:"objectStorageNamespace"`

	// Admin Email for Notification
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// OceInstance description, can be updated
	Description *string `mandatory:"false" json:"description"`

	// Instance type based on its usage
	InstanceUsageType OceInstanceSummaryInstanceUsageTypeEnum `mandatory:"false" json:"instanceUsageType,omitempty"`

	// a list of add-on features for the ocm instance
	AddOnFeatures []string `mandatory:"false" json:"addOnFeatures"`

	// Upgrade schedule type representing service to be upgraded immediately whenever latest version is released
	// or delay upgrade of the service to previous released version
	UpgradeSchedule OceInstanceUpgradeScheduleEnum `mandatory:"false" json:"upgradeSchedule,omitempty"`

	// Web Application Firewall(WAF) primary domain
	WafPrimaryDomain *string `mandatory:"false" json:"wafPrimaryDomain"`

	// Flag indicating whether the instance access is private or public
	InstanceAccessType OceInstanceSummaryInstanceAccessTypeEnum `mandatory:"false" json:"instanceAccessType,omitempty"`

	// Flag indicating whether the instance license is new cloud or bring your own license
	InstanceLicenseType LicenseTypeEnum `mandatory:"false" json:"instanceLicenseType,omitempty"`

	// The time the the OceInstance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the OceInstance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the instance lifecycle.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details of the current state of the instance lifecycle
	LifecycleDetails LifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// disaster recovery paired ragion name
	DrRegion *string `mandatory:"false" json:"drRegion"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// SERVICE data.
	// Example: `{"service": {"IDCS": "value"}}`
	Service map[string]interface{} `mandatory:"false" json:"service"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OceInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OceInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOceInstanceSummaryInstanceUsageTypeEnum(string(m.InstanceUsageType)); !ok && m.InstanceUsageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceUsageType: %s. Supported values are: %s.", m.InstanceUsageType, strings.Join(GetOceInstanceSummaryInstanceUsageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOceInstanceUpgradeScheduleEnum(string(m.UpgradeSchedule)); !ok && m.UpgradeSchedule != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeSchedule: %s. Supported values are: %s.", m.UpgradeSchedule, strings.Join(GetOceInstanceUpgradeScheduleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOceInstanceSummaryInstanceAccessTypeEnum(string(m.InstanceAccessType)); !ok && m.InstanceAccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceAccessType: %s. Supported values are: %s.", m.InstanceAccessType, strings.Join(GetOceInstanceSummaryInstanceAccessTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.InstanceLicenseType)); !ok && m.InstanceLicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceLicenseType: %s. Supported values are: %s.", m.InstanceLicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OceInstanceSummaryInstanceUsageTypeEnum Enum with underlying type: string
type OceInstanceSummaryInstanceUsageTypeEnum string

// Set of constants representing the allowable values for OceInstanceSummaryInstanceUsageTypeEnum
const (
	OceInstanceSummaryInstanceUsageTypePrimary    OceInstanceSummaryInstanceUsageTypeEnum = "PRIMARY"
	OceInstanceSummaryInstanceUsageTypeNonprimary OceInstanceSummaryInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingOceInstanceSummaryInstanceUsageTypeEnum = map[string]OceInstanceSummaryInstanceUsageTypeEnum{
	"PRIMARY":    OceInstanceSummaryInstanceUsageTypePrimary,
	"NONPRIMARY": OceInstanceSummaryInstanceUsageTypeNonprimary,
}

var mappingOceInstanceSummaryInstanceUsageTypeEnumLowerCase = map[string]OceInstanceSummaryInstanceUsageTypeEnum{
	"primary":    OceInstanceSummaryInstanceUsageTypePrimary,
	"nonprimary": OceInstanceSummaryInstanceUsageTypeNonprimary,
}

// GetOceInstanceSummaryInstanceUsageTypeEnumValues Enumerates the set of values for OceInstanceSummaryInstanceUsageTypeEnum
func GetOceInstanceSummaryInstanceUsageTypeEnumValues() []OceInstanceSummaryInstanceUsageTypeEnum {
	values := make([]OceInstanceSummaryInstanceUsageTypeEnum, 0)
	for _, v := range mappingOceInstanceSummaryInstanceUsageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOceInstanceSummaryInstanceUsageTypeEnumStringValues Enumerates the set of values in String for OceInstanceSummaryInstanceUsageTypeEnum
func GetOceInstanceSummaryInstanceUsageTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"NONPRIMARY",
	}
}

// GetMappingOceInstanceSummaryInstanceUsageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOceInstanceSummaryInstanceUsageTypeEnum(val string) (OceInstanceSummaryInstanceUsageTypeEnum, bool) {
	enum, ok := mappingOceInstanceSummaryInstanceUsageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OceInstanceSummaryInstanceAccessTypeEnum Enum with underlying type: string
type OceInstanceSummaryInstanceAccessTypeEnum string

// Set of constants representing the allowable values for OceInstanceSummaryInstanceAccessTypeEnum
const (
	OceInstanceSummaryInstanceAccessTypePublic  OceInstanceSummaryInstanceAccessTypeEnum = "PUBLIC"
	OceInstanceSummaryInstanceAccessTypePrivate OceInstanceSummaryInstanceAccessTypeEnum = "PRIVATE"
)

var mappingOceInstanceSummaryInstanceAccessTypeEnum = map[string]OceInstanceSummaryInstanceAccessTypeEnum{
	"PUBLIC":  OceInstanceSummaryInstanceAccessTypePublic,
	"PRIVATE": OceInstanceSummaryInstanceAccessTypePrivate,
}

var mappingOceInstanceSummaryInstanceAccessTypeEnumLowerCase = map[string]OceInstanceSummaryInstanceAccessTypeEnum{
	"public":  OceInstanceSummaryInstanceAccessTypePublic,
	"private": OceInstanceSummaryInstanceAccessTypePrivate,
}

// GetOceInstanceSummaryInstanceAccessTypeEnumValues Enumerates the set of values for OceInstanceSummaryInstanceAccessTypeEnum
func GetOceInstanceSummaryInstanceAccessTypeEnumValues() []OceInstanceSummaryInstanceAccessTypeEnum {
	values := make([]OceInstanceSummaryInstanceAccessTypeEnum, 0)
	for _, v := range mappingOceInstanceSummaryInstanceAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOceInstanceSummaryInstanceAccessTypeEnumStringValues Enumerates the set of values in String for OceInstanceSummaryInstanceAccessTypeEnum
func GetOceInstanceSummaryInstanceAccessTypeEnumStringValues() []string {
	return []string{
		"PUBLIC",
		"PRIVATE",
	}
}

// GetMappingOceInstanceSummaryInstanceAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOceInstanceSummaryInstanceAccessTypeEnum(val string) (OceInstanceSummaryInstanceAccessTypeEnum, bool) {
	enum, ok := mappingOceInstanceSummaryInstanceAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
