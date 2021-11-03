// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Domain Properties for a Domain
type Domain struct {

	// The OCID of the domain
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The mutable display name of the domain
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The domain descripition
	Description *string `mandatory:"true" json:"description"`

	// Region agnostic domain URL.
	Url *string `mandatory:"true" json:"url"`

	// Region specific domain URL.
	HomeRegionUrl *string `mandatory:"true" json:"homeRegionUrl"`

	// The home region for the domain.
	// See Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm)
	// for the full list of supported region names.
	// Example: `us-phoenix-1`
	HomeRegion *string `mandatory:"true" json:"homeRegion"`

	// The regions domain is replication to.
	ReplicaRegions []ReplicatedRegionDetails `mandatory:"true" json:"replicaRegions"`

	// The type of the domain.
	Type DomainTypeEnum `mandatory:"true" json:"type"`

	// The License type of Domain
	LicenseType *string `mandatory:"true" json:"licenseType"`

	// Indicates whether domain is hidden on login screen or not.
	IsHiddenOnLogin *bool `mandatory:"true" json:"isHiddenOnLogin"`

	// Date and time the domain was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state.
	LifecycleState DomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Any additional details about the current state of the Domain.
	LifecycleDetails DomainLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Domain) String() string {
	return common.PointerString(m)
}

// DomainTypeEnum Enum with underlying type: string
type DomainTypeEnum string

// Set of constants representing the allowable values for DomainTypeEnum
const (
	DomainTypeDefault   DomainTypeEnum = "DEFAULT"
	DomainTypeSecondary DomainTypeEnum = "SECONDARY"
)

var mappingDomainType = map[string]DomainTypeEnum{
	"DEFAULT":   DomainTypeDefault,
	"SECONDARY": DomainTypeSecondary,
}

// GetDomainTypeEnumValues Enumerates the set of values for DomainTypeEnum
func GetDomainTypeEnumValues() []DomainTypeEnum {
	values := make([]DomainTypeEnum, 0)
	for _, v := range mappingDomainType {
		values = append(values, v)
	}
	return values
}

// DomainLifecycleStateEnum Enum with underlying type: string
type DomainLifecycleStateEnum string

// Set of constants representing the allowable values for DomainLifecycleStateEnum
const (
	DomainLifecycleStateCreating DomainLifecycleStateEnum = "CREATING"
	DomainLifecycleStateActive   DomainLifecycleStateEnum = "ACTIVE"
	DomainLifecycleStateDeleting DomainLifecycleStateEnum = "DELETING"
	DomainLifecycleStateInactive DomainLifecycleStateEnum = "INACTIVE"
)

var mappingDomainLifecycleState = map[string]DomainLifecycleStateEnum{
	"CREATING": DomainLifecycleStateCreating,
	"ACTIVE":   DomainLifecycleStateActive,
	"DELETING": DomainLifecycleStateDeleting,
	"INACTIVE": DomainLifecycleStateInactive,
}

// GetDomainLifecycleStateEnumValues Enumerates the set of values for DomainLifecycleStateEnum
func GetDomainLifecycleStateEnumValues() []DomainLifecycleStateEnum {
	values := make([]DomainLifecycleStateEnum, 0)
	for _, v := range mappingDomainLifecycleState {
		values = append(values, v)
	}
	return values
}

// DomainLifecycleDetailsEnum Enum with underlying type: string
type DomainLifecycleDetailsEnum string

// Set of constants representing the allowable values for DomainLifecycleDetailsEnum
const (
	DomainLifecycleDetailsDeactivating DomainLifecycleDetailsEnum = "DEACTIVATING"
	DomainLifecycleDetailsActivating   DomainLifecycleDetailsEnum = "ACTIVATING"
	DomainLifecycleDetailsUpdating     DomainLifecycleDetailsEnum = "UPDATING"
)

var mappingDomainLifecycleDetails = map[string]DomainLifecycleDetailsEnum{
	"DEACTIVATING": DomainLifecycleDetailsDeactivating,
	"ACTIVATING":   DomainLifecycleDetailsActivating,
	"UPDATING":     DomainLifecycleDetailsUpdating,
}

// GetDomainLifecycleDetailsEnumValues Enumerates the set of values for DomainLifecycleDetailsEnum
func GetDomainLifecycleDetailsEnumValues() []DomainLifecycleDetailsEnum {
	values := make([]DomainLifecycleDetailsEnum, 0)
	for _, v := range mappingDomainLifecycleDetails {
		values = append(values, v)
	}
	return values
}
