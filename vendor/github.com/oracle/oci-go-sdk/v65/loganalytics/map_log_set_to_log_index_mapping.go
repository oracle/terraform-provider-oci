// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MapLogSetToLogIndexMapping Mapping for a log set(s) to a log index.
// - If logSet is specified, maps a single log set.
// - If logSets is specified, maps all in the array.
type MapLogSetToLogIndexMapping struct {

	// The log index that logSets should be mapped to.
	LogIndex *int `mandatory:"true" json:"logIndex"`

	// A single log set name.
	LogSet *string `mandatory:"false" json:"logSet"`

	// List of log set names to be mapped to the same logIndex.
	LogSets []string `mandatory:"false" json:"logSets"`
}

func (m MapLogSetToLogIndexMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MapLogSetToLogIndexMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
