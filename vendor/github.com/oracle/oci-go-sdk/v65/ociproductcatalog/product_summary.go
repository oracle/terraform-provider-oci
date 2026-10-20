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

// ProductSummary The fields of a product to be displayed to the operator in OCI Console
type ProductSummary struct {

	// OCID of the product
	Id *string `mandatory:"true" json:"id"`

	// Name of the product, defined by service teams. Unique within one service
	Name *string `mandatory:"true" json:"name"`

	// List of skus associated with this product.
	Skus []SkuSummary `mandatory:"true" json:"skus"`

	// Name of the metering service this product relates to.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The status of a product following the OCI standard. ACTIVE: resources of this product can be used by the end customer. INACTIVE: resources of this product can't be used by the end customer. Product needs manual approval. NEEDS_ATTENTION: operator action is needed.
	LifecycleState ProductSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m ProductSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProductSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProductSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProductSummaryLifecycleStateEnum Enum with underlying type: string
type ProductSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ProductSummaryLifecycleStateEnum
const (
	ProductSummaryLifecycleStateActive         ProductSummaryLifecycleStateEnum = "ACTIVE"
	ProductSummaryLifecycleStateInactive       ProductSummaryLifecycleStateEnum = "INACTIVE"
	ProductSummaryLifecycleStateNeedsAttention ProductSummaryLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingProductSummaryLifecycleStateEnum = map[string]ProductSummaryLifecycleStateEnum{
	"ACTIVE":          ProductSummaryLifecycleStateActive,
	"INACTIVE":        ProductSummaryLifecycleStateInactive,
	"NEEDS_ATTENTION": ProductSummaryLifecycleStateNeedsAttention,
}

var mappingProductSummaryLifecycleStateEnumLowerCase = map[string]ProductSummaryLifecycleStateEnum{
	"active":          ProductSummaryLifecycleStateActive,
	"inactive":        ProductSummaryLifecycleStateInactive,
	"needs_attention": ProductSummaryLifecycleStateNeedsAttention,
}

// GetProductSummaryLifecycleStateEnumValues Enumerates the set of values for ProductSummaryLifecycleStateEnum
func GetProductSummaryLifecycleStateEnumValues() []ProductSummaryLifecycleStateEnum {
	values := make([]ProductSummaryLifecycleStateEnum, 0)
	for _, v := range mappingProductSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProductSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ProductSummaryLifecycleStateEnum
func GetProductSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
	}
}

// GetMappingProductSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProductSummaryLifecycleStateEnum(val string) (ProductSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingProductSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
