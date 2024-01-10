// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlCollectionLogAggregation The details of SQL collection log aggregation items.
type SqlCollectionLogAggregation struct {

	// Name of the aggregation.
	MetricName *string `mandatory:"true" json:"metricName"`

	// Total count of aggregated value.
	Count *int64 `mandatory:"true" json:"count"`

	// The time at which the aggregation started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time at which the aggregation ended.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	Dimensions *SqlCollectionLogDimensions `mandatory:"false" json:"dimensions"`
}

func (m SqlCollectionLogAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlCollectionLogAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
