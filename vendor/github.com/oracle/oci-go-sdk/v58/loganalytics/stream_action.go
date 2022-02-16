// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// StreamAction Stream action for scheduled task.
type StreamAction struct {

	// The ManagementSavedSearch id [OCID] utilized in the action.
	SavedSearchId *string `mandatory:"false" json:"savedSearchId"`

	MetricExtraction *MetricExtraction `mandatory:"false" json:"metricExtraction"`

	// The duration of data to be searched for SAVED_SEARCH tasks,
	// used when the task fires to calculate the query time range.
	// Duration in ISO 8601 extended format as described in
	// https://en.wikipedia.org/wiki/ISO_8601#Durations.
	// The value should be positive.
	// The largest supported unit (as opposed to value) is D, e.g.  P14D (not P2W).
	// There are restrictions on the maximum duration value relative to the task schedule
	// value as specified in the following table.
	//    Schedule Interval Range          | Maximum Duration
	// ----------------------------------- | -----------------
	//   5 Minutes     to 30 Minutes       |   1 hour  "PT60M"
	//  31 Minutes     to  1 Hour          |  12 hours "PT720M"
	//  1 Hour+1Minute to  1 Day           |   1 day   "P1D"
	//  1 Day+1Minute  to  1 Week-1Minute  |   7 days  "P7D"
	//  1 Week         to  2 Weeks         |  14 days  "P14D"
	//  greater than 2 Weeks               |  30 days  "P30D"
	// If not specified, the duration will be based on the schedule. For example,
	// if the schedule is every 5 minutes then the savedSearchDuration will be "PT5M";
	// if the schedule is every 3 weeks then the savedSearchDuration will be "P21D".
	SavedSearchDuration *string `mandatory:"false" json:"savedSearchDuration"`
}

func (m StreamAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StreamAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamAction StreamAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStreamAction
	}{
		"STREAM",
		(MarshalTypeStreamAction)(m),
	}

	return json.Marshal(&s)
}
