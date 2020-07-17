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

	// The current state of the file system.
	LifecycleState OceInstanceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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
}

func (m OceInstanceSummary) String() string {
	return common.PointerString(m)
}

// OceInstanceSummaryInstanceUsageTypeEnum Enum with underlying type: string
type OceInstanceSummaryInstanceUsageTypeEnum string

// Set of constants representing the allowable values for OceInstanceSummaryInstanceUsageTypeEnum
const (
	OceInstanceSummaryInstanceUsageTypePrimary    OceInstanceSummaryInstanceUsageTypeEnum = "PRIMARY"
	OceInstanceSummaryInstanceUsageTypeNonprimary OceInstanceSummaryInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingOceInstanceSummaryInstanceUsageType = map[string]OceInstanceSummaryInstanceUsageTypeEnum{
	"PRIMARY":    OceInstanceSummaryInstanceUsageTypePrimary,
	"NONPRIMARY": OceInstanceSummaryInstanceUsageTypeNonprimary,
}

// GetOceInstanceSummaryInstanceUsageTypeEnumValues Enumerates the set of values for OceInstanceSummaryInstanceUsageTypeEnum
func GetOceInstanceSummaryInstanceUsageTypeEnumValues() []OceInstanceSummaryInstanceUsageTypeEnum {
	values := make([]OceInstanceSummaryInstanceUsageTypeEnum, 0)
	for _, v := range mappingOceInstanceSummaryInstanceUsageType {
		values = append(values, v)
	}
	return values
}

// OceInstanceSummaryInstanceAccessTypeEnum Enum with underlying type: string
type OceInstanceSummaryInstanceAccessTypeEnum string

// Set of constants representing the allowable values for OceInstanceSummaryInstanceAccessTypeEnum
const (
	OceInstanceSummaryInstanceAccessTypePublic  OceInstanceSummaryInstanceAccessTypeEnum = "PUBLIC"
	OceInstanceSummaryInstanceAccessTypePrivate OceInstanceSummaryInstanceAccessTypeEnum = "PRIVATE"
)

var mappingOceInstanceSummaryInstanceAccessType = map[string]OceInstanceSummaryInstanceAccessTypeEnum{
	"PUBLIC":  OceInstanceSummaryInstanceAccessTypePublic,
	"PRIVATE": OceInstanceSummaryInstanceAccessTypePrivate,
}

// GetOceInstanceSummaryInstanceAccessTypeEnumValues Enumerates the set of values for OceInstanceSummaryInstanceAccessTypeEnum
func GetOceInstanceSummaryInstanceAccessTypeEnumValues() []OceInstanceSummaryInstanceAccessTypeEnum {
	values := make([]OceInstanceSummaryInstanceAccessTypeEnum, 0)
	for _, v := range mappingOceInstanceSummaryInstanceAccessType {
		values = append(values, v)
	}
	return values
}

// OceInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type OceInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OceInstanceSummaryLifecycleStateEnum
const (
	OceInstanceSummaryLifecycleStateCreating OceInstanceSummaryLifecycleStateEnum = "CREATING"
	OceInstanceSummaryLifecycleStateUpdating OceInstanceSummaryLifecycleStateEnum = "UPDATING"
	OceInstanceSummaryLifecycleStateActive   OceInstanceSummaryLifecycleStateEnum = "ACTIVE"
	OceInstanceSummaryLifecycleStateDeleting OceInstanceSummaryLifecycleStateEnum = "DELETING"
	OceInstanceSummaryLifecycleStateDeleted  OceInstanceSummaryLifecycleStateEnum = "DELETED"
	OceInstanceSummaryLifecycleStateFailed   OceInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingOceInstanceSummaryLifecycleState = map[string]OceInstanceSummaryLifecycleStateEnum{
	"CREATING": OceInstanceSummaryLifecycleStateCreating,
	"UPDATING": OceInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   OceInstanceSummaryLifecycleStateActive,
	"DELETING": OceInstanceSummaryLifecycleStateDeleting,
	"DELETED":  OceInstanceSummaryLifecycleStateDeleted,
	"FAILED":   OceInstanceSummaryLifecycleStateFailed,
}

// GetOceInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for OceInstanceSummaryLifecycleStateEnum
func GetOceInstanceSummaryLifecycleStateEnumValues() []OceInstanceSummaryLifecycleStateEnum {
	values := make([]OceInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingOceInstanceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
