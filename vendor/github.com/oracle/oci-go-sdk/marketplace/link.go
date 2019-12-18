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

// Link The model for links.
type Link struct {

	// The anchor tag.
	Href *string `mandatory:"false" json:"href"`
}

func (m Link) String() string {
	return common.PointerString(m)
}
