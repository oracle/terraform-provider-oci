// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MarketingBrand marketing brand details for fusion environment
type MarketingBrand struct {

	// The unique identifier (OCID) of marketing brand. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// marketing brand name for fusion environment
	Name *string `mandatory:"true" json:"name"`

	// lifecycle state of marketing brand
	LifecycleState MarketingBrandLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// fusion environment id
	FusionEnvironmentId *string `mandatory:"true" json:"fusionEnvironmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Marketing Brand intermediate states
	LifecycleDetails MarketingBrandLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// timeCreated
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m MarketingBrand) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MarketingBrand) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMarketingBrandLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMarketingBrandLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMarketingBrandLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetMarketingBrandLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarketingBrandLifecycleStateEnum Enum with underlying type: string
type MarketingBrandLifecycleStateEnum string

// Set of constants representing the allowable values for MarketingBrandLifecycleStateEnum
const (
	MarketingBrandLifecycleStateActive   MarketingBrandLifecycleStateEnum = "ACTIVE"
	MarketingBrandLifecycleStateInactive MarketingBrandLifecycleStateEnum = "INACTIVE"
)

var mappingMarketingBrandLifecycleStateEnum = map[string]MarketingBrandLifecycleStateEnum{
	"ACTIVE":   MarketingBrandLifecycleStateActive,
	"INACTIVE": MarketingBrandLifecycleStateInactive,
}

var mappingMarketingBrandLifecycleStateEnumLowerCase = map[string]MarketingBrandLifecycleStateEnum{
	"active":   MarketingBrandLifecycleStateActive,
	"inactive": MarketingBrandLifecycleStateInactive,
}

// GetMarketingBrandLifecycleStateEnumValues Enumerates the set of values for MarketingBrandLifecycleStateEnum
func GetMarketingBrandLifecycleStateEnumValues() []MarketingBrandLifecycleStateEnum {
	values := make([]MarketingBrandLifecycleStateEnum, 0)
	for _, v := range mappingMarketingBrandLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMarketingBrandLifecycleStateEnumStringValues Enumerates the set of values in String for MarketingBrandLifecycleStateEnum
func GetMarketingBrandLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingMarketingBrandLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMarketingBrandLifecycleStateEnum(val string) (MarketingBrandLifecycleStateEnum, bool) {
	enum, ok := mappingMarketingBrandLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MarketingBrandLifecycleDetailsEnum Enum with underlying type: string
type MarketingBrandLifecycleDetailsEnum string

// Set of constants representing the allowable values for MarketingBrandLifecycleDetailsEnum
const (
	MarketingBrandLifecycleDetailsNone     MarketingBrandLifecycleDetailsEnum = "NONE"
	MarketingBrandLifecycleDetailsDisabled MarketingBrandLifecycleDetailsEnum = "DISABLED"
)

var mappingMarketingBrandLifecycleDetailsEnum = map[string]MarketingBrandLifecycleDetailsEnum{
	"NONE":     MarketingBrandLifecycleDetailsNone,
	"DISABLED": MarketingBrandLifecycleDetailsDisabled,
}

var mappingMarketingBrandLifecycleDetailsEnumLowerCase = map[string]MarketingBrandLifecycleDetailsEnum{
	"none":     MarketingBrandLifecycleDetailsNone,
	"disabled": MarketingBrandLifecycleDetailsDisabled,
}

// GetMarketingBrandLifecycleDetailsEnumValues Enumerates the set of values for MarketingBrandLifecycleDetailsEnum
func GetMarketingBrandLifecycleDetailsEnumValues() []MarketingBrandLifecycleDetailsEnum {
	values := make([]MarketingBrandLifecycleDetailsEnum, 0)
	for _, v := range mappingMarketingBrandLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetMarketingBrandLifecycleDetailsEnumStringValues Enumerates the set of values in String for MarketingBrandLifecycleDetailsEnum
func GetMarketingBrandLifecycleDetailsEnumStringValues() []string {
	return []string{
		"NONE",
		"DISABLED",
	}
}

// GetMappingMarketingBrandLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMarketingBrandLifecycleDetailsEnum(val string) (MarketingBrandLifecycleDetailsEnum, bool) {
	enum, ok := mappingMarketingBrandLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
