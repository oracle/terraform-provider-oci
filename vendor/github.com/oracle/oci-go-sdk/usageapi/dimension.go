// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// A description of the UsageApi API.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Dimension The dimension use for filtering.
// example:
// `[{value: "COMPUTE", key: "service"}]`
type Dimension struct {

	// The key of the dimension.
	Key *string `mandatory:"true" json:"key"`

	// The value of the dimension.
	Value *string `mandatory:"true" json:"value"`
}

func (m Dimension) String() string {
	return common.PointerString(m)
}
