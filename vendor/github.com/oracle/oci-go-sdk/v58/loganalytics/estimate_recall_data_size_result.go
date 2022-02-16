// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// EstimateRecallDataSizeResult This is the size and time range of data to be recalled
type EstimateRecallDataSizeResult struct {

	// This is the end of the time range of data to be recalled.  timeDataStarted and timeDataEnded delineate
	// the time range of the archived data to be recalled.  They may not be exact the same as the
	// parameters in the request input (EstimateRecallDataSizeDetails).
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the start of the time range of data to be recalled
	TimeDataStarted *common.SDKTime `mandatory:"true" json:"timeDataStarted"`

	// This is the size in bytes
	SizeInBytes *int64 `mandatory:"true" json:"sizeInBytes"`

	// This indicates if the time range of data to be recalled overlaps with existing recalled data
	IsOverlappingWithExistingRecalls *bool `mandatory:"false" json:"isOverlappingWithExistingRecalls"`
}

func (m EstimateRecallDataSizeResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EstimateRecallDataSizeResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
