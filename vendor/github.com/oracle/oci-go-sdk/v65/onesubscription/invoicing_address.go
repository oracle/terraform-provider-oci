// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// InvoicingAddress Address.
type InvoicingAddress struct {
	Location *InvoicingLocation `mandatory:"false" json:"location"`

	// Address name identifier.
	Name *string `mandatory:"false" json:"name"`

	// Phone.
	Phone *string `mandatory:"false" json:"phone"`

	// Identify as the customer's billing address.
	IsBillTo *bool `mandatory:"false" json:"isBillTo"`

	// Identify as the customer's shipping address.
	IsShipTo *bool `mandatory:"false" json:"isShipTo"`

	// Bill to site use Id.
	BillSiteUseId *int64 `mandatory:"false" json:"billSiteUseId"`

	// Service to site use Id.
	Service2SiteUseId *int64 `mandatory:"false" json:"service2SiteUseId"`

	// TCA customer account site Id.
	TcaCustAcctSiteId *int64 `mandatory:"false" json:"tcaCustAcctSiteId"`

	// Party site number.
	TcaPartySiteNumber *string `mandatory:"false" json:"tcaPartySiteNumber"`
}

func (m InvoicingAddress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoicingAddress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
