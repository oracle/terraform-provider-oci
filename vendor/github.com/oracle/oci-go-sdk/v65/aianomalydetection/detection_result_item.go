// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectionResultItem An object to hold detection result for one timestamp/row.
type DetectionResultItem struct {

	// An array of anomalies associated with a given timestamp/row.
	Anomalies []Anomaly `mandatory:"true" json:"anomalies"`

	// The time stamp associated with a list of anomaly points, format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The index number to indicate where anomaly points are located among all rows when there are no timestamps provided.
	RowIndex *int `mandatory:"false" json:"rowIndex"`

	// A significant score across multiple signals at timestamp/row level
	Score *float64 `mandatory:"false" json:"score"`
}

func (m DetectionResultItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectionResultItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
