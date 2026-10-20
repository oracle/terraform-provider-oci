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

// UpdateProductDetails The properties of a product
type UpdateProductDetails struct {

	// Name of the product, defined by service teams. Unique within one service
	Name *string `mandatory:"false" json:"name"`

	// description to the product
	Description *string `mandatory:"false" json:"description"`

	// The current state of the product. READY - The product is fully configured and ready for activation. ENABLED - The product is active. DISABLED - The product is inactive and not ready for activation. DELETED - The product has been deleted. NEEDS_ATTENTION - Pricing not set.
	State UpdateProductDetailsStateEnum `mandatory:"false" json:"state,omitempty"`

	// List of meters associated with this product
	Meters []Meter `mandatory:"false" json:"meters"`

	// Name of the metering service this product relates to.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// List of limits that this product is associated with.
	Limits []Limit `mandatory:"false" json:"limits"`

	// Optional exclusion flag update. When true, marks the product as excluded. When false,
	// clears the exclusion state.
	IsExcluded *bool `mandatory:"false" json:"isExcluded"`
}

func (m UpdateProductDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateProductDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateProductDetailsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetUpdateProductDetailsStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateProductDetailsStateEnum Enum with underlying type: string
type UpdateProductDetailsStateEnum string

// Set of constants representing the allowable values for UpdateProductDetailsStateEnum
const (
	UpdateProductDetailsStateReady          UpdateProductDetailsStateEnum = "READY"
	UpdateProductDetailsStateEnabled        UpdateProductDetailsStateEnum = "ENABLED"
	UpdateProductDetailsStateDisabled       UpdateProductDetailsStateEnum = "DISABLED"
	UpdateProductDetailsStateDeleted        UpdateProductDetailsStateEnum = "DELETED"
	UpdateProductDetailsStateNeedsAttention UpdateProductDetailsStateEnum = "NEEDS_ATTENTION"
)

var mappingUpdateProductDetailsStateEnum = map[string]UpdateProductDetailsStateEnum{
	"READY":           UpdateProductDetailsStateReady,
	"ENABLED":         UpdateProductDetailsStateEnabled,
	"DISABLED":        UpdateProductDetailsStateDisabled,
	"DELETED":         UpdateProductDetailsStateDeleted,
	"NEEDS_ATTENTION": UpdateProductDetailsStateNeedsAttention,
}

var mappingUpdateProductDetailsStateEnumLowerCase = map[string]UpdateProductDetailsStateEnum{
	"ready":           UpdateProductDetailsStateReady,
	"enabled":         UpdateProductDetailsStateEnabled,
	"disabled":        UpdateProductDetailsStateDisabled,
	"deleted":         UpdateProductDetailsStateDeleted,
	"needs_attention": UpdateProductDetailsStateNeedsAttention,
}

// GetUpdateProductDetailsStateEnumValues Enumerates the set of values for UpdateProductDetailsStateEnum
func GetUpdateProductDetailsStateEnumValues() []UpdateProductDetailsStateEnum {
	values := make([]UpdateProductDetailsStateEnum, 0)
	for _, v := range mappingUpdateProductDetailsStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateProductDetailsStateEnumStringValues Enumerates the set of values in String for UpdateProductDetailsStateEnum
func GetUpdateProductDetailsStateEnumStringValues() []string {
	return []string{
		"READY",
		"ENABLED",
		"DISABLED",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingUpdateProductDetailsStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateProductDetailsStateEnum(val string) (UpdateProductDetailsStateEnum, bool) {
	enum, ok := mappingUpdateProductDetailsStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
