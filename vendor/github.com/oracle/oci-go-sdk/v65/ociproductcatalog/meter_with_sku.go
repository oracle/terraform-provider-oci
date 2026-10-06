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

// MeterWithSku Meter defined in metering service with the associated SKU
type MeterWithSku struct {

	// Name of the meter
	Name *string `mandatory:"false" json:"name"`

	// B number of the SKU associated with the meter
	SkuBNumber *string `mandatory:"false" json:"skuBNumber"`

	// Description of the SKU associated with the meter
	SkuDescription *string `mandatory:"false" json:"skuDescription"`
}

func (m MeterWithSku) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MeterWithSku) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
