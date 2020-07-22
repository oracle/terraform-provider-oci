// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
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

	// The current state of the file system.
	LifecycleState OceInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// SERVICE data.
	// Example: `{"service": {"IDCS": "value"}}`
	Service map[string]interface{} `mandatory:"false" json:"service"`
}

func (m OceInstance) String() string {
	return common.PointerString(m)
}

// OceInstanceUpgradeScheduleEnum Enum with underlying type: string
type OceInstanceUpgradeScheduleEnum string

// Set of constants representing the allowable values for OceInstanceUpgradeScheduleEnum
const (
	OceInstanceUpgradeScheduleUpgradeImmediately OceInstanceUpgradeScheduleEnum = "UPGRADE_IMMEDIATELY"
	OceInstanceUpgradeScheduleDelayedUpgrade     OceInstanceUpgradeScheduleEnum = "DELAYED_UPGRADE"
)

var mappingOceInstanceUpgradeSchedule = map[string]OceInstanceUpgradeScheduleEnum{
	"UPGRADE_IMMEDIATELY": OceInstanceUpgradeScheduleUpgradeImmediately,
	"DELAYED_UPGRADE":     OceInstanceUpgradeScheduleDelayedUpgrade,
}

// GetOceInstanceUpgradeScheduleEnumValues Enumerates the set of values for OceInstanceUpgradeScheduleEnum
func GetOceInstanceUpgradeScheduleEnumValues() []OceInstanceUpgradeScheduleEnum {
	values := make([]OceInstanceUpgradeScheduleEnum, 0)
	for _, v := range mappingOceInstanceUpgradeSchedule {
		values = append(values, v)
	}
	return values
}

// OceInstanceInstanceUsageTypeEnum Enum with underlying type: string
type OceInstanceInstanceUsageTypeEnum string

// Set of constants representing the allowable values for OceInstanceInstanceUsageTypeEnum
const (
	OceInstanceInstanceUsageTypePrimary    OceInstanceInstanceUsageTypeEnum = "PRIMARY"
	OceInstanceInstanceUsageTypeNonprimary OceInstanceInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingOceInstanceInstanceUsageType = map[string]OceInstanceInstanceUsageTypeEnum{
	"PRIMARY":    OceInstanceInstanceUsageTypePrimary,
	"NONPRIMARY": OceInstanceInstanceUsageTypeNonprimary,
}

// GetOceInstanceInstanceUsageTypeEnumValues Enumerates the set of values for OceInstanceInstanceUsageTypeEnum
func GetOceInstanceInstanceUsageTypeEnumValues() []OceInstanceInstanceUsageTypeEnum {
	values := make([]OceInstanceInstanceUsageTypeEnum, 0)
	for _, v := range mappingOceInstanceInstanceUsageType {
		values = append(values, v)
	}
	return values
}

// OceInstanceInstanceAccessTypeEnum Enum with underlying type: string
type OceInstanceInstanceAccessTypeEnum string

// Set of constants representing the allowable values for OceInstanceInstanceAccessTypeEnum
const (
	OceInstanceInstanceAccessTypePublic  OceInstanceInstanceAccessTypeEnum = "PUBLIC"
	OceInstanceInstanceAccessTypePrivate OceInstanceInstanceAccessTypeEnum = "PRIVATE"
)

var mappingOceInstanceInstanceAccessType = map[string]OceInstanceInstanceAccessTypeEnum{
	"PUBLIC":  OceInstanceInstanceAccessTypePublic,
	"PRIVATE": OceInstanceInstanceAccessTypePrivate,
}

// GetOceInstanceInstanceAccessTypeEnumValues Enumerates the set of values for OceInstanceInstanceAccessTypeEnum
func GetOceInstanceInstanceAccessTypeEnumValues() []OceInstanceInstanceAccessTypeEnum {
	values := make([]OceInstanceInstanceAccessTypeEnum, 0)
	for _, v := range mappingOceInstanceInstanceAccessType {
		values = append(values, v)
	}
	return values
}

// OceInstanceLifecycleStateEnum Enum with underlying type: string
type OceInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for OceInstanceLifecycleStateEnum
const (
	OceInstanceLifecycleStateCreating OceInstanceLifecycleStateEnum = "CREATING"
	OceInstanceLifecycleStateUpdating OceInstanceLifecycleStateEnum = "UPDATING"
	OceInstanceLifecycleStateActive   OceInstanceLifecycleStateEnum = "ACTIVE"
	OceInstanceLifecycleStateDeleting OceInstanceLifecycleStateEnum = "DELETING"
	OceInstanceLifecycleStateDeleted  OceInstanceLifecycleStateEnum = "DELETED"
	OceInstanceLifecycleStateFailed   OceInstanceLifecycleStateEnum = "FAILED"
)

var mappingOceInstanceLifecycleState = map[string]OceInstanceLifecycleStateEnum{
	"CREATING": OceInstanceLifecycleStateCreating,
	"UPDATING": OceInstanceLifecycleStateUpdating,
	"ACTIVE":   OceInstanceLifecycleStateActive,
	"DELETING": OceInstanceLifecycleStateDeleting,
	"DELETED":  OceInstanceLifecycleStateDeleted,
	"FAILED":   OceInstanceLifecycleStateFailed,
}

// GetOceInstanceLifecycleStateEnumValues Enumerates the set of values for OceInstanceLifecycleStateEnum
func GetOceInstanceLifecycleStateEnumValues() []OceInstanceLifecycleStateEnum {
	values := make([]OceInstanceLifecycleStateEnum, 0)
	for _, v := range mappingOceInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}
