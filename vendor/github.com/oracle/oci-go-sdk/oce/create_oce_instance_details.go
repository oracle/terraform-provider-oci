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

	// Upgrade schedule type representing service to be upgraded immediately whenever latest version is released
	// or delay upgrade of the service to previous released version
	UpgradeSchedule OceInstanceUpgradeScheduleEnum `mandatory:"false" json:"upgradeSchedule,omitempty"`

	// Web Application Firewall(WAF) primary domain
	WafPrimaryDomain *string `mandatory:"false" json:"wafPrimaryDomain"`

	// Flag indicating whether the instance access is private or public
	InstanceAccessType CreateOceInstanceDetailsInstanceAccessTypeEnum `mandatory:"false" json:"instanceAccessType,omitempty"`

	// Flag indicating whether the instance license is new cloud or bring your own license
	InstanceLicenseType LicenseTypeEnum `mandatory:"false" json:"instanceLicenseType,omitempty"`

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

// CreateOceInstanceDetailsInstanceUsageTypeEnum Enum with underlying type: string
type CreateOceInstanceDetailsInstanceUsageTypeEnum string

// Set of constants representing the allowable values for CreateOceInstanceDetailsInstanceUsageTypeEnum
const (
	CreateOceInstanceDetailsInstanceUsageTypePrimary    CreateOceInstanceDetailsInstanceUsageTypeEnum = "PRIMARY"
	CreateOceInstanceDetailsInstanceUsageTypeNonprimary CreateOceInstanceDetailsInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingCreateOceInstanceDetailsInstanceUsageType = map[string]CreateOceInstanceDetailsInstanceUsageTypeEnum{
	"PRIMARY":    CreateOceInstanceDetailsInstanceUsageTypePrimary,
	"NONPRIMARY": CreateOceInstanceDetailsInstanceUsageTypeNonprimary,
}

// GetCreateOceInstanceDetailsInstanceUsageTypeEnumValues Enumerates the set of values for CreateOceInstanceDetailsInstanceUsageTypeEnum
func GetCreateOceInstanceDetailsInstanceUsageTypeEnumValues() []CreateOceInstanceDetailsInstanceUsageTypeEnum {
	values := make([]CreateOceInstanceDetailsInstanceUsageTypeEnum, 0)
	for _, v := range mappingCreateOceInstanceDetailsInstanceUsageType {
		values = append(values, v)
	}
	return values
}

// CreateOceInstanceDetailsInstanceAccessTypeEnum Enum with underlying type: string
type CreateOceInstanceDetailsInstanceAccessTypeEnum string

// Set of constants representing the allowable values for CreateOceInstanceDetailsInstanceAccessTypeEnum
const (
	CreateOceInstanceDetailsInstanceAccessTypePublic  CreateOceInstanceDetailsInstanceAccessTypeEnum = "PUBLIC"
	CreateOceInstanceDetailsInstanceAccessTypePrivate CreateOceInstanceDetailsInstanceAccessTypeEnum = "PRIVATE"
)

var mappingCreateOceInstanceDetailsInstanceAccessType = map[string]CreateOceInstanceDetailsInstanceAccessTypeEnum{
	"PUBLIC":  CreateOceInstanceDetailsInstanceAccessTypePublic,
	"PRIVATE": CreateOceInstanceDetailsInstanceAccessTypePrivate,
}

// GetCreateOceInstanceDetailsInstanceAccessTypeEnumValues Enumerates the set of values for CreateOceInstanceDetailsInstanceAccessTypeEnum
func GetCreateOceInstanceDetailsInstanceAccessTypeEnumValues() []CreateOceInstanceDetailsInstanceAccessTypeEnum {
	values := make([]CreateOceInstanceDetailsInstanceAccessTypeEnum, 0)
	for _, v := range mappingCreateOceInstanceDetailsInstanceAccessType {
		values = append(values, v)
	}
	return values
}
