// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Product Catalog API
//
// Apis to manage the products used by the OCI service teams
//

package ociproductcatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdminProduct The properties of a product
type AdminProduct struct {

	// OCID of the product
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the product, defined by service teams. Unique within one service
	Name *string `mandatory:"true" json:"name"`

	// description to the product
	Description *string `mandatory:"true" json:"description"`

	// List of meters associated with this product, including SKU information
	MeterwithSKUs []MeterWithSku `mandatory:"true" json:"meterwithSKUs"`

	// Name of the metering service this product relates to.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// List of limits that this product is associated with.
	Limits []Limit `mandatory:"true" json:"limits"`

	// The status of a product following the OCI standard. ACTIVE: resources of this product can be used by the end customer. INACTIVE: resources of this product can't be used by the end customer. Product needs manual approval. NEEDS_ATTENTION: operator action is needed.
	LifecycleState AdminProductLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The displayed status of a product. lifecycleState to lifecycleDetails mapping: INACTIVE -> "Not yet launched" ACTIVE -> "Launched" NEEDS_ATTENTION -> "Pricing not set"
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Date and time when the product was created
	// Example: `2022-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time when the product was updated. Nullable
	// Example: `2022-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Date and time when the product was set to ready. Nullable
	// Example: `2022-01-25T21:10:29.600Z`
	TimeReady *common.SDKTime `mandatory:"false" json:"timeReady"`

	// Date and time when the product was launched. Nullable
	// Example: `2022-01-25T21:10:29.600Z`
	TimeLaunched *common.SDKTime `mandatory:"false" json:"timeLaunched"`

	// Whether the product is currently excluded from SKU processing flows. Nullable when the
	// exclusion flag has not been explicitly set.
	IsExcluded *bool `mandatory:"false" json:"isExcluded"`

	// Date and time when the product was marked as excluded. Nullable
	// Example: `2022-01-25T21:10:29.600Z`
	TimeExcluded *common.SDKTime `mandatory:"false" json:"timeExcluded"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AdminProduct) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdminProduct) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAdminProductLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAdminProductLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AdminProductLifecycleStateEnum Enum with underlying type: string
type AdminProductLifecycleStateEnum string

// Set of constants representing the allowable values for AdminProductLifecycleStateEnum
const (
	AdminProductLifecycleStateActive         AdminProductLifecycleStateEnum = "ACTIVE"
	AdminProductLifecycleStateInactive       AdminProductLifecycleStateEnum = "INACTIVE"
	AdminProductLifecycleStateNeedsAttention AdminProductLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingAdminProductLifecycleStateEnum = map[string]AdminProductLifecycleStateEnum{
	"ACTIVE":          AdminProductLifecycleStateActive,
	"INACTIVE":        AdminProductLifecycleStateInactive,
	"NEEDS_ATTENTION": AdminProductLifecycleStateNeedsAttention,
}

var mappingAdminProductLifecycleStateEnumLowerCase = map[string]AdminProductLifecycleStateEnum{
	"active":          AdminProductLifecycleStateActive,
	"inactive":        AdminProductLifecycleStateInactive,
	"needs_attention": AdminProductLifecycleStateNeedsAttention,
}

// GetAdminProductLifecycleStateEnumValues Enumerates the set of values for AdminProductLifecycleStateEnum
func GetAdminProductLifecycleStateEnumValues() []AdminProductLifecycleStateEnum {
	values := make([]AdminProductLifecycleStateEnum, 0)
	for _, v := range mappingAdminProductLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAdminProductLifecycleStateEnumStringValues Enumerates the set of values in String for AdminProductLifecycleStateEnum
func GetAdminProductLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
	}
}

// GetMappingAdminProductLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdminProductLifecycleStateEnum(val string) (AdminProductLifecycleStateEnum, bool) {
	enum, ok := mappingAdminProductLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
