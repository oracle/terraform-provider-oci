// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OverlappingRecallSummary This is the information about overlapping recall requests
type OverlappingRecallSummary struct {

	// This is the start of the time range of the archival data
	TimeDataStarted *common.SDKTime `mandatory:"true" json:"timeDataStarted"`

	// This is the end of the time range of the archival data
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// This is the time when the recall operation was started for this recall request
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// This is the status of the recall
	Status RecallStatusEnum `mandatory:"true" json:"status"`

	// This is the purpose of the recall
	Purpose *string `mandatory:"true" json:"purpose"`

	// This is the query associated with the recall
	QueryString *string `mandatory:"true" json:"queryString"`

	// This is the list of logsets associated with this recall
	LogSets *string `mandatory:"true" json:"logSets"`

	// This is the user who initiated the recall request
	CreatedBy *string `mandatory:"true" json:"createdBy"`
}

func (m OverlappingRecallSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OverlappingRecallSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecallStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRecallStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
