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

// UploadSummary Summary of the Upload.
type UploadSummary struct {

	// Unique internal identifier to refer the upload container.
	Reference *string `mandatory:"true" json:"reference"`

	// The name of the upload container.
	Name *string `mandatory:"true" json:"name"`

	// The time when this upload container is created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The latest time when this upload container is modified. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// This time represents the earliest time of the log entry in this container. An RFC3339 formatted datetime string.
	TimeEarliestLogEntry *common.SDKTime `mandatory:"false" json:"timeEarliestLogEntry"`

	// This time represents the latest time of the log entry in this container. An RFC3339 formatted datetime string.
	TimeLatestLogEntry *common.SDKTime `mandatory:"false" json:"timeLatestLogEntry"`

	// Number of warnings associated to the upload.
	WarningsCount *int `mandatory:"false" json:"warningsCount"`
}

func (m UploadSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UploadSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
