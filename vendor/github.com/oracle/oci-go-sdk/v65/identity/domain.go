// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Domain (For tenancies that support identity domains) Properties for an identity domain. An identity domain is used to manage users and groups, integration standards, external identities, and secure application integration through Oracle Single Sign-on (SSO) configuration.
type Domain struct {

	// The OCID of the identity domain.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the identity domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The mutable display name of the identity domain.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The identity domain description. You can have an empty description.
	Description *string `mandatory:"true" json:"description"`

	// Region-agnostic identity domain URL.
	Url *string `mandatory:"true" json:"url"`

	// Region-specific identity domain URL.
	HomeRegionUrl *string `mandatory:"true" json:"homeRegionUrl"`

	// The home region for the identity domain.
	// See Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm)
	// for the full list of supported region names.
	// Example: `us-phoenix-1`
	HomeRegion *string `mandatory:"true" json:"homeRegion"`

	// The regions where replicas of the identity domain exist.
	ReplicaRegions []ReplicatedRegionDetails `mandatory:"true" json:"replicaRegions"`

	// The type of the domain.
	Type DomainTypeEnum `mandatory:"true" json:"type"`

	// The license type of the identity domain.
	LicenseType *string `mandatory:"true" json:"licenseType"`

	// Indicates whether the identity domain is hidden on the sign-in screen or not.
	IsHiddenOnLogin *bool `mandatory:"true" json:"isHiddenOnLogin"`

	// Date and time the identity domain was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state.
	LifecycleState DomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Any additional details about the current state of the identity domain.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Domain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDomainTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDomainTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDomainLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDomainLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetDomainLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DomainTypeEnum Enum with underlying type: string
type DomainTypeEnum string

// Set of constants representing the allowable values for DomainTypeEnum
const (
	DomainTypeDefault   DomainTypeEnum = "DEFAULT"
	DomainTypeSecondary DomainTypeEnum = "SECONDARY"
)

var mappingDomainTypeEnum = map[string]DomainTypeEnum{
	"DEFAULT":   DomainTypeDefault,
	"SECONDARY": DomainTypeSecondary,
}

var mappingDomainTypeEnumLowerCase = map[string]DomainTypeEnum{
	"default":   DomainTypeDefault,
	"secondary": DomainTypeSecondary,
}

// GetDomainTypeEnumValues Enumerates the set of values for DomainTypeEnum
func GetDomainTypeEnumValues() []DomainTypeEnum {
	values := make([]DomainTypeEnum, 0)
	for _, v := range mappingDomainTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDomainTypeEnumStringValues Enumerates the set of values in String for DomainTypeEnum
func GetDomainTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SECONDARY",
	}
}

// GetMappingDomainTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDomainTypeEnum(val string) (DomainTypeEnum, bool) {
	enum, ok := mappingDomainTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingDomainLifecycleStateEnum = map[string]DomainLifecycleStateEnum{
	"CREATING": DomainLifecycleStateCreating,
	"ACTIVE":   DomainLifecycleStateActive,
	"DELETING": DomainLifecycleStateDeleting,
	"INACTIVE": DomainLifecycleStateInactive,
}

var mappingDomainLifecycleStateEnumLowerCase = map[string]DomainLifecycleStateEnum{
	"creating": DomainLifecycleStateCreating,
	"active":   DomainLifecycleStateActive,
	"deleting": DomainLifecycleStateDeleting,
	"inactive": DomainLifecycleStateInactive,
}

// GetDomainLifecycleStateEnumValues Enumerates the set of values for DomainLifecycleStateEnum
func GetDomainLifecycleStateEnumValues() []DomainLifecycleStateEnum {
	values := make([]DomainLifecycleStateEnum, 0)
	for _, v := range mappingDomainLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDomainLifecycleStateEnumStringValues Enumerates the set of values in String for DomainLifecycleStateEnum
func GetDomainLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"INACTIVE",
	}
}

// GetMappingDomainLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDomainLifecycleStateEnum(val string) (DomainLifecycleStateEnum, bool) {
	enum, ok := mappingDomainLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DomainLifecycleDetailsEnum Enum with underlying type: string
type DomainLifecycleDetailsEnum string

// Set of constants representing the allowable values for DomainLifecycleDetailsEnum
const (
	DomainLifecycleDetailsDeactivating DomainLifecycleDetailsEnum = "DEACTIVATING"
	DomainLifecycleDetailsActivating   DomainLifecycleDetailsEnum = "ACTIVATING"
	DomainLifecycleDetailsUpdating     DomainLifecycleDetailsEnum = "UPDATING"
)

var mappingDomainLifecycleDetailsEnum = map[string]DomainLifecycleDetailsEnum{
	"DEACTIVATING": DomainLifecycleDetailsDeactivating,
	"ACTIVATING":   DomainLifecycleDetailsActivating,
	"UPDATING":     DomainLifecycleDetailsUpdating,
}

var mappingDomainLifecycleDetailsEnumLowerCase = map[string]DomainLifecycleDetailsEnum{
	"deactivating": DomainLifecycleDetailsDeactivating,
	"activating":   DomainLifecycleDetailsActivating,
	"updating":     DomainLifecycleDetailsUpdating,
}

// GetDomainLifecycleDetailsEnumValues Enumerates the set of values for DomainLifecycleDetailsEnum
func GetDomainLifecycleDetailsEnumValues() []DomainLifecycleDetailsEnum {
	values := make([]DomainLifecycleDetailsEnum, 0)
	for _, v := range mappingDomainLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetDomainLifecycleDetailsEnumStringValues Enumerates the set of values in String for DomainLifecycleDetailsEnum
func GetDomainLifecycleDetailsEnumStringValues() []string {
	return []string{
		"DEACTIVATING",
		"ACTIVATING",
		"UPDATING",
	}
}

// GetMappingDomainLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDomainLifecycleDetailsEnum(val string) (DomainLifecycleDetailsEnum, bool) {
	enum, ok := mappingDomainLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
