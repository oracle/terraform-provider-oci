// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ProblemTrendAggregation Provides the dimensions and their corresponding time and count.
type ProblemTrendAggregation struct {

	// The key-value pairs of dimensions and their names.
	DimensionsMap map[string]string `mandatory:"true" json:"dimensionsMap"`

	// Start Time in epoch seconds
	StartTimestamp *float32 `mandatory:"true" json:"startTimestamp"`

	// Duration
	DurationInSeconds *int `mandatory:"true" json:"durationInSeconds"`

	// The number of occurences with for the corresponding time range and dimensions.
	Count *int `mandatory:"true" json:"count"`
}

func (m ProblemTrendAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProblemTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
