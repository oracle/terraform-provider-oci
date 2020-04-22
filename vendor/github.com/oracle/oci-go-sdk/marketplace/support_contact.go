// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SupportContact Contact information to use to get support.
type SupportContact struct {

	// The name of the contact.
	Name *string `mandatory:"false" json:"name"`

	// The phone number of the contact.
	Phone *string `mandatory:"false" json:"phone"`

	// The email of the contact.
	Email *string `mandatory:"false" json:"email"`

	// The email subject line to use when contacting support.
	Subject *string `mandatory:"false" json:"subject"`
}

func (m SupportContact) String() string {
	return common.PointerString(m)
}
