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

// DocumentationLink A link to a documentation resource on the internet.
type DocumentationLink struct {

	// Text that describes the resource.
	Name *string `mandatory:"false" json:"name"`

	// The URL of the resource.
	Url *string `mandatory:"false" json:"url"`

	// The category that the document belongs to.
	DocumentCategory *string `mandatory:"false" json:"documentCategory"`
}

func (m DocumentationLink) String() string {
	return common.PointerString(m)
}
