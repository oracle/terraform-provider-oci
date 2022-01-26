// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TaxSummary Tax implication that current tenant may be eligible while using specific listing
type TaxSummary struct {

	// Unique code for the tax.
	Code *string `mandatory:"true" json:"code"`

	// Name of the tax code.
	Name *string `mandatory:"false" json:"name"`

	// Country, which imposes the tax.
	Country *string `mandatory:"false" json:"country"`

	// The URL with more details about this tax.
	Url *string `mandatory:"false" json:"url"`
}

func (m TaxSummary) String() string {
	return common.PointerString(m)
}
