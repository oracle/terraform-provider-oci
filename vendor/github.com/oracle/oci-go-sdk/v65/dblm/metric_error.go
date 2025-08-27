// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricError Metric error content.
type MetricError struct {

	// Time the error record was generated
	TimeGenerated *common.SDKTime `mandatory:"true" json:"timeGenerated"`

	// Error type
	ErrorType *string `mandatory:"true" json:"errorType"`

	// Content type
	ContentType *string `mandatory:"true" json:"contentType"`

	// Data
	Data *string `mandatory:"true" json:"data"`

	// Subject
	Subject *string `mandatory:"true" json:"subject"`

	// Level
	Level *string `mandatory:"true" json:"level"`
}

func (m MetricError) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricError) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
