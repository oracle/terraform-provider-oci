// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SkuProducts The SKU Product Id details for a resource.
type SkuProducts struct {

	// The Sku Id for the resource.
	SkuId *string `mandatory:"false" json:"skuId"`

	// The Sku type for the resource.
	SkuType *string `mandatory:"false" json:"skuType"`

	// The cloud credit type for the resource.
	CloudCreditType *string `mandatory:"false" json:"cloudCreditType"`
}

func (m SkuProducts) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SkuProducts) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
