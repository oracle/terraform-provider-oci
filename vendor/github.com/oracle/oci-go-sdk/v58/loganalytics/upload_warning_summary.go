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

// UploadWarningSummary Summary of Upload warnings.
type UploadWarningSummary struct {

	// Unique internal identifier to refer upload warning.
	Reference *string `mandatory:"true" json:"reference"`

	// Status of the upload. Ex - Failed.
	Status *string `mandatory:"false" json:"status"`

	// The time when the upload processing started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The details about upload processing failure.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m UploadWarningSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UploadWarningSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
