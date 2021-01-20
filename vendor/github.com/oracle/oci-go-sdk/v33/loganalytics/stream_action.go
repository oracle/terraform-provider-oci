// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v33/common"
)

// StreamAction Stream action for scheduled task.
type StreamAction struct {

	// The ManagementSavedSearch id [OCID] utilized in the action.
	SavedSearchId *string `mandatory:"false" json:"savedSearchId"`

	MetricExtraction *MetricExtraction `mandatory:"false" json:"metricExtraction"`
}

func (m StreamAction) String() string {
	return common.PointerString(m)
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
