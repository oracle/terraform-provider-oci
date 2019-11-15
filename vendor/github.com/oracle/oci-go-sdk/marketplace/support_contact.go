// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
