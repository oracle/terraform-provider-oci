// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyValueDetectionLabelMetricsReport Label Metrics report for Key Value Detection Model.
type KeyValueDetectionLabelMetricsReport struct {

	// Mean average precision under different thresholds
	MeanAveragePrecision *float32 `mandatory:"true" json:"meanAveragePrecision"`

	// List of key value detection confidence report.
	ConfidenceEntries []KeyValueDetectionConfidenceEntry `mandatory:"true" json:"confidenceEntries"`

	// Total test documents in the label.
	DocumentCount *int `mandatory:"false" json:"documentCount"`

	// Label name
	Label *string `mandatory:"false" json:"label"`
}

func (m KeyValueDetectionLabelMetricsReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KeyValueDetectionLabelMetricsReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
