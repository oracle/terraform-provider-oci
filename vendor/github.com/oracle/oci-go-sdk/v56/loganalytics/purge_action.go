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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PurgeAction Purge action for scheduled task.
type PurgeAction struct {

	// Purge query string.
	QueryString *string `mandatory:"true" json:"queryString"`

	// The duration of data to be retained, which is used to
	// calculate the timeDataEnded when the task fires.
	// The value should be negative.
	// Purge duration in ISO 8601 extended format as described in
	// https://en.wikipedia.org/wiki/ISO_8601#Durations.
	// The largest supported unit is D, e.g. -P365D (not -P1Y) or -P14D (not -P2W).
	PurgeDuration *string `mandatory:"true" json:"purgeDuration"`

	// the compartment OCID under which the data will be purged
	PurgeCompartmentId *string `mandatory:"true" json:"purgeCompartmentId"`

	// if true, purge child compartments data
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`

	// the type of the log data to be purged
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`
}

func (m PurgeAction) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PurgeAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePurgeAction PurgeAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePurgeAction
	}{
		"PURGE",
		(MarshalTypePurgeAction)(m),
	}

	return json.Marshal(&s)
}
