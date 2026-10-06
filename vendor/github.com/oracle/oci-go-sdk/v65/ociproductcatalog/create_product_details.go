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

// CreateProductDetails The properties of a product
type CreateProductDetails struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the product, defined by service teams. Unique within one service
	Name *string `mandatory:"true" json:"name"`

	// description to the product
	Description *string `mandatory:"true" json:"description"`

	// Name of the metering service this product relates to.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// List of meters associated with this product
	Meters []Meter `mandatory:"false" json:"meters"`

	// List of limits that this product is associated with.
	Limits []Limit `mandatory:"false" json:"limits"`

	// Optional create-time signal that marks the product as eligible for the backfill flow in
	// existing Alloy regions. Omit this field to preserve the current create behavior. Final
	// runtime backfill and auto-enable decisions remain worker-side.
	BackfillEligibility CreateProductDetailsBackfillEligibilityEnum `mandatory:"false" json:"backfillEligibility,omitempty"`

	// Optional create-time exclusion flag. When true, the product is created in an excluded
	// state. Omit this field to preserve backward-compatible create behavior.
	IsExcluded *bool `mandatory:"false" json:"isExcluded"`
}

func (m CreateProductDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateProductDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateProductDetailsBackfillEligibilityEnum(string(m.BackfillEligibility)); !ok && m.BackfillEligibility != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackfillEligibility: %s. Supported values are: %s.", m.BackfillEligibility, strings.Join(GetCreateProductDetailsBackfillEligibilityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateProductDetailsBackfillEligibilityEnum Enum with underlying type: string
type CreateProductDetailsBackfillEligibilityEnum string

// Set of constants representing the allowable values for CreateProductDetailsBackfillEligibilityEnum
const (
	CreateProductDetailsBackfillEligibilityBackfillEligible CreateProductDetailsBackfillEligibilityEnum = "BACKFILL_ELIGIBLE"
)

var mappingCreateProductDetailsBackfillEligibilityEnum = map[string]CreateProductDetailsBackfillEligibilityEnum{
	"BACKFILL_ELIGIBLE": CreateProductDetailsBackfillEligibilityBackfillEligible,
}

var mappingCreateProductDetailsBackfillEligibilityEnumLowerCase = map[string]CreateProductDetailsBackfillEligibilityEnum{
	"backfill_eligible": CreateProductDetailsBackfillEligibilityBackfillEligible,
}

// GetCreateProductDetailsBackfillEligibilityEnumValues Enumerates the set of values for CreateProductDetailsBackfillEligibilityEnum
func GetCreateProductDetailsBackfillEligibilityEnumValues() []CreateProductDetailsBackfillEligibilityEnum {
	values := make([]CreateProductDetailsBackfillEligibilityEnum, 0)
	for _, v := range mappingCreateProductDetailsBackfillEligibilityEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateProductDetailsBackfillEligibilityEnumStringValues Enumerates the set of values in String for CreateProductDetailsBackfillEligibilityEnum
func GetCreateProductDetailsBackfillEligibilityEnumStringValues() []string {
	return []string{
		"BACKFILL_ELIGIBLE",
	}
}

// GetMappingCreateProductDetailsBackfillEligibilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateProductDetailsBackfillEligibilityEnum(val string) (CreateProductDetailsBackfillEligibilityEnum, bool) {
	enum, ok := mappingCreateProductDetailsBackfillEligibilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
