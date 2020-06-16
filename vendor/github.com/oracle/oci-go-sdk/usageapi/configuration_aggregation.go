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

// ConfigurationAggregation The available configurations
type ConfigurationAggregation struct {

	// The list of available configurations
	Items []Configuration `mandatory:"true" json:"items"`
}

func (m ConfigurationAggregation) String() string {
	return common.PointerString(m)
}
