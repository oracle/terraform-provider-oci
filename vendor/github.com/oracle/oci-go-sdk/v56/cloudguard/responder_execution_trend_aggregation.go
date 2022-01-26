// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResponderExecutionTrendAggregation Provides the timestamps and their corresponding number of remediations.
type ResponderExecutionTrendAggregation struct {

	// The key-value pairs of dimensions and their names.
	DimensionsMap map[string]string `mandatory:"true" json:"dimensionsMap"`

	// Start Time in epoch seconds
	StartTimestamp *float32 `mandatory:"true" json:"startTimestamp"`

	// Duration
	DurationInSeconds *int `mandatory:"true" json:"durationInSeconds"`

	// The number of remediations for a given time.
	Count *int `mandatory:"true" json:"count"`
}

func (m ResponderExecutionTrendAggregation) String() string {
	return common.PointerString(m)
}
