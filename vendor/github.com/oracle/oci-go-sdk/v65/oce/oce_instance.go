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

// OceInstance Details of OceInstance.
type OceInstance struct {

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

	// Upgrade schedule type representing service to be upgraded immediately whenever latest version is released
	// or delay upgrade of the service to previous released version
	UpgradeSchedule OceInstanceUpgradeScheduleEnum `mandatory:"false" json:"upgradeSchedule,omitempty"`

	IdentityStripe *IdentityStripeDetails `mandatory:"false" json:"identityStripe"`

	// Instance type based on its usage
	InstanceUsageType OceInstanceInstanceUsageTypeEnum `mandatory:"false" json:"instanceUsageType,omitempty"`

	// a list of add-on features for the ocm instance
	AddOnFeatures []string `mandatory:"false" json:"addOnFeatures"`

	// Web Application Firewall(WAF) primary domain
	WafPrimaryDomain *string `mandatory:"false" json:"wafPrimaryDomain"`

	// Flag indicating whether the instance access is private or public
	InstanceAccessType OceInstanceInstanceAccessTypeEnum `mandatory:"false" json:"instanceAccessType,omitempty"`

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

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// SERVICE data.
	// Example: `{"service": {"IDCS": "value"}}`
	Service map[string]interface{} `mandatory:"false" json:"service"`
}

func (m OceInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OceInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOceInstanceUpgradeScheduleEnum(string(m.UpgradeSchedule)); !ok && m.UpgradeSchedule != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeSchedule: %s. Supported values are: %s.", m.UpgradeSchedule, strings.Join(GetOceInstanceUpgradeScheduleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOceInstanceInstanceUsageTypeEnum(string(m.InstanceUsageType)); !ok && m.InstanceUsageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceUsageType: %s. Supported values are: %s.", m.InstanceUsageType, strings.Join(GetOceInstanceInstanceUsageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOceInstanceInstanceAccessTypeEnum(string(m.InstanceAccessType)); !ok && m.InstanceAccessType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InstanceAccessType: %s. Supported values are: %s.", m.InstanceAccessType, strings.Join(GetOceInstanceInstanceAccessTypeEnumStringValues(), ",")))
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

// OceInstanceUpgradeScheduleEnum Enum with underlying type: string
type OceInstanceUpgradeScheduleEnum string

// Set of constants representing the allowable values for OceInstanceUpgradeScheduleEnum
const (
	OceInstanceUpgradeScheduleUpgradeImmediately OceInstanceUpgradeScheduleEnum = "UPGRADE_IMMEDIATELY"
	OceInstanceUpgradeScheduleDelayedUpgrade     OceInstanceUpgradeScheduleEnum = "DELAYED_UPGRADE"
)

var mappingOceInstanceUpgradeScheduleEnum = map[string]OceInstanceUpgradeScheduleEnum{
	"UPGRADE_IMMEDIATELY": OceInstanceUpgradeScheduleUpgradeImmediately,
	"DELAYED_UPGRADE":     OceInstanceUpgradeScheduleDelayedUpgrade,
}

var mappingOceInstanceUpgradeScheduleEnumLowerCase = map[string]OceInstanceUpgradeScheduleEnum{
	"upgrade_immediately": OceInstanceUpgradeScheduleUpgradeImmediately,
	"delayed_upgrade":     OceInstanceUpgradeScheduleDelayedUpgrade,
}

// GetOceInstanceUpgradeScheduleEnumValues Enumerates the set of values for OceInstanceUpgradeScheduleEnum
func GetOceInstanceUpgradeScheduleEnumValues() []OceInstanceUpgradeScheduleEnum {
	values := make([]OceInstanceUpgradeScheduleEnum, 0)
	for _, v := range mappingOceInstanceUpgradeScheduleEnum {
		values = append(values, v)
	}
	return values
}

// GetOceInstanceUpgradeScheduleEnumStringValues Enumerates the set of values in String for OceInstanceUpgradeScheduleEnum
func GetOceInstanceUpgradeScheduleEnumStringValues() []string {
	return []string{
		"UPGRADE_IMMEDIATELY",
		"DELAYED_UPGRADE",
	}
}

// GetMappingOceInstanceUpgradeScheduleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOceInstanceUpgradeScheduleEnum(val string) (OceInstanceUpgradeScheduleEnum, bool) {
	enum, ok := mappingOceInstanceUpgradeScheduleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OceInstanceInstanceUsageTypeEnum Enum with underlying type: string
type OceInstanceInstanceUsageTypeEnum string

// Set of constants representing the allowable values for OceInstanceInstanceUsageTypeEnum
const (
	OceInstanceInstanceUsageTypePrimary    OceInstanceInstanceUsageTypeEnum = "PRIMARY"
	OceInstanceInstanceUsageTypeNonprimary OceInstanceInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingOceInstanceInstanceUsageTypeEnum = map[string]OceInstanceInstanceUsageTypeEnum{
	"PRIMARY":    OceInstanceInstanceUsageTypePrimary,
	"NONPRIMARY": OceInstanceInstanceUsageTypeNonprimary,
}

var mappingOceInstanceInstanceUsageTypeEnumLowerCase = map[string]OceInstanceInstanceUsageTypeEnum{
	"primary":    OceInstanceInstanceUsageTypePrimary,
	"nonprimary": OceInstanceInstanceUsageTypeNonprimary,
}

// GetOceInstanceInstanceUsageTypeEnumValues Enumerates the set of values for OceInstanceInstanceUsageTypeEnum
func GetOceInstanceInstanceUsageTypeEnumValues() []OceInstanceInstanceUsageTypeEnum {
	values := make([]OceInstanceInstanceUsageTypeEnum, 0)
	for _, v := range mappingOceInstanceInstanceUsageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOceInstanceInstanceUsageTypeEnumStringValues Enumerates the set of values in String for OceInstanceInstanceUsageTypeEnum
func GetOceInstanceInstanceUsageTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"NONPRIMARY",
	}
}

// GetMappingOceInstanceInstanceUsageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOceInstanceInstanceUsageTypeEnum(val string) (OceInstanceInstanceUsageTypeEnum, bool) {
	enum, ok := mappingOceInstanceInstanceUsageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OceInstanceInstanceAccessTypeEnum Enum with underlying type: string
type OceInstanceInstanceAccessTypeEnum string

// Set of constants representing the allowable values for OceInstanceInstanceAccessTypeEnum
const (
	OceInstanceInstanceAccessTypePublic  OceInstanceInstanceAccessTypeEnum = "PUBLIC"
	OceInstanceInstanceAccessTypePrivate OceInstanceInstanceAccessTypeEnum = "PRIVATE"
)

var mappingOceInstanceInstanceAccessTypeEnum = map[string]OceInstanceInstanceAccessTypeEnum{
	"PUBLIC":  OceInstanceInstanceAccessTypePublic,
	"PRIVATE": OceInstanceInstanceAccessTypePrivate,
}

var mappingOceInstanceInstanceAccessTypeEnumLowerCase = map[string]OceInstanceInstanceAccessTypeEnum{
	"public":  OceInstanceInstanceAccessTypePublic,
	"private": OceInstanceInstanceAccessTypePrivate,
}

// GetOceInstanceInstanceAccessTypeEnumValues Enumerates the set of values for OceInstanceInstanceAccessTypeEnum
func GetOceInstanceInstanceAccessTypeEnumValues() []OceInstanceInstanceAccessTypeEnum {
	values := make([]OceInstanceInstanceAccessTypeEnum, 0)
	for _, v := range mappingOceInstanceInstanceAccessTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOceInstanceInstanceAccessTypeEnumStringValues Enumerates the set of values in String for OceInstanceInstanceAccessTypeEnum
func GetOceInstanceInstanceAccessTypeEnumStringValues() []string {
	return []string{
		"PUBLIC",
		"PRIVATE",
	}
}

// GetMappingOceInstanceInstanceAccessTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOceInstanceInstanceAccessTypeEnum(val string) (OceInstanceInstanceAccessTypeEnum, bool) {
	enum, ok := mappingOceInstanceInstanceAccessTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
