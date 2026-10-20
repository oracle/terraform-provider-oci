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

// Product The fields of a product to be displayed to the operator in OCI Console
type Product struct {

	// OCID of the product
	Id *string `mandatory:"true" json:"id"`

	// Name of the product, defined by service teams. Unique within one service
	Name *string `mandatory:"true" json:"name"`

	// List of skus associated with this product.
	Skus []SkuSummary `mandatory:"true" json:"skus"`

	// Name of the metering service this product relates to.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The status of a product following the OCI standard. ACTIVE: resources of this product can be used by the end customer. INACTIVE: resources of this product can't be used by the end customer. Product needs manual approval. NEEDS_ATTENTION: operator action is needed.
	LifecycleState ProductLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// description to the product
	Description *string `mandatory:"true" json:"description"`

	// The displayed status of a product. lifecycleState to lifecycleDetails mapping: INACTIVE -> "Not yet launched" ACTIVE -> "Launched" NEEDS_ATTENTION -> "Pricing not set"
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Date and time when the product is created
	// Example: `2022-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time when the product is launched. Nullable
	// Example: `2022-01-25T21:10:29.600Z`
	TimeLaunched *common.SDKTime `mandatory:"false" json:"timeLaunched"`

	// Date and time when the product is ready to be launched. Nullable
	// Example: `2022-01-25T21:10:29.600Z`
	TimeReady *common.SDKTime `mandatory:"false" json:"timeReady"`

	// Indicates whether the product is currently excluded from SKU processing flows. This
	// field is nullable when the exclusion flag has not been explicitly set.
	IsExcluded *bool `mandatory:"false" json:"isExcluded"`

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

func (m Product) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Product) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProductLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProductLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProductLifecycleStateEnum Enum with underlying type: string
type ProductLifecycleStateEnum string

// Set of constants representing the allowable values for ProductLifecycleStateEnum
const (
	ProductLifecycleStateActive         ProductLifecycleStateEnum = "ACTIVE"
	ProductLifecycleStateInactive       ProductLifecycleStateEnum = "INACTIVE"
	ProductLifecycleStateNeedsAttention ProductLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingProductLifecycleStateEnum = map[string]ProductLifecycleStateEnum{
	"ACTIVE":          ProductLifecycleStateActive,
	"INACTIVE":        ProductLifecycleStateInactive,
	"NEEDS_ATTENTION": ProductLifecycleStateNeedsAttention,
}

var mappingProductLifecycleStateEnumLowerCase = map[string]ProductLifecycleStateEnum{
	"active":          ProductLifecycleStateActive,
	"inactive":        ProductLifecycleStateInactive,
	"needs_attention": ProductLifecycleStateNeedsAttention,
}

// GetProductLifecycleStateEnumValues Enumerates the set of values for ProductLifecycleStateEnum
func GetProductLifecycleStateEnumValues() []ProductLifecycleStateEnum {
	values := make([]ProductLifecycleStateEnum, 0)
	for _, v := range mappingProductLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProductLifecycleStateEnumStringValues Enumerates the set of values in String for ProductLifecycleStateEnum
func GetProductLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
	}
}

// GetMappingProductLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProductLifecycleStateEnum(val string) (ProductLifecycleStateEnum, bool) {
	enum, ok := mappingProductLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
