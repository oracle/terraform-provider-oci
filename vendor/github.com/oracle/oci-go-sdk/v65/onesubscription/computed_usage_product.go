// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputedUsageProduct Product description
type ComputedUsageProduct struct {

	// Product part number
	PartNumber *string `mandatory:"true" json:"partNumber"`

	// Product name
	Name *string `mandatory:"true" json:"name"`

	// Unit of Measure
	UnitOfMeasure *string `mandatory:"false" json:"unitOfMeasure"`

	// Product provisioning group
	ProvisioningGroup *string `mandatory:"false" json:"provisioningGroup"`

	// Metered service billing category
	BillingCategory *string `mandatory:"false" json:"billingCategory"`

	// Product category
	ProductCategory *string `mandatory:"false" json:"productCategory"`

	// Rate card part type of Product
	UcmRateCardPartType *string `mandatory:"false" json:"ucmRateCardPartType"`
}

func (m ComputedUsageProduct) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputedUsageProduct) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
