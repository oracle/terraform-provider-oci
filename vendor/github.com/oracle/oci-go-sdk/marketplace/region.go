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

// Region The model for regions supported by a listing and package.
type Region struct {

	// The name of the region.
	Name *string `mandatory:"false" json:"name"`

	// The code of the region.
	Code *string `mandatory:"false" json:"code"`

	// Countries in the region.
	Countries []Item `mandatory:"false" json:"countries"`
}

func (m Region) String() string {
	return common.PointerString(m)
}
