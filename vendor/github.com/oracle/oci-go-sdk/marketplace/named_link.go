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

// NamedLink A link to a resource on the internet.
type NamedLink struct {

	// Text that describes the resource.
	Name *string `mandatory:"false" json:"name"`

	// The URL of the resource.
	Url *string `mandatory:"false" json:"url"`
}

func (m NamedLink) String() string {
	return common.PointerString(m)
}
