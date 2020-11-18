// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// Dimension The dimension used for filtering. availabe dimension are "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"
// example:
// `[{value: "COMPUTE", key: "service"}]`
type Dimension struct {

	// The dimension key.
	Key *string `mandatory:"true" json:"key"`

	// The dimension value.
	Value *string `mandatory:"true" json:"value"`
}

func (m Dimension) String() string {
	return common.PointerString(m)
}
