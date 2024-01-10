// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssetAggregation The result of an analytics aggregation on a set of assets.
type AssetAggregation struct {

	// Aggregated property.
	AggregatedProperty *string `mandatory:"true" json:"aggregatedProperty"`

	// The dimensions along which assets can be aggregated for analytics.
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// Returns the total number of observations from the group of assets.
	Count *int64 `mandatory:"false" json:"count"`

	// Returns the highest value from all the assets.
	Max *float64 `mandatory:"false" json:"max"`

	// Returns the value of sum divided by count from the group of assets.
	Mean *float64 `mandatory:"false" json:"mean"`

	// Returns the lowest value from the group of assets.
	Min *float64 `mandatory:"false" json:"min"`

	// Returns all values added together from the group of assets.
	Sum *float64 `mandatory:"false" json:"sum"`
}

func (m AssetAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssetAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
