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

// Link The model for links.
type Link struct {

	// Reference links to the previous page, next page, and other pages.
	Rel LinkEnumEnum `mandatory:"false" json:"rel,omitempty"`

	// The anchor tag.
	Href *string `mandatory:"false" json:"href"`
}

func (m Link) String() string {
	return common.PointerString(m)
}
