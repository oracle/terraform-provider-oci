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

// LogAnalyticsImportCustomChangeList LogAnalyticsImportCustomChangeList
type LogAnalyticsImportCustomChangeList struct {

	// An array of created parser names.
	CreatedParserNames []string `mandatory:"false" json:"createdParserNames"`

	// An array of updated parser names.
	UpdatedParserNames []string `mandatory:"false" json:"updatedParserNames"`

	// An array of created source names.
	CreatedSourceNames []string `mandatory:"false" json:"createdSourceNames"`

	// An array of updated source names.
	UpdatedSourceNames []string `mandatory:"false" json:"updatedSourceNames"`

	// An array of created field display names.
	CreatedFieldDisplayNames []string `mandatory:"false" json:"createdFieldDisplayNames"`

	// An array of updated field display names.
	UpdatedFieldDisplayNames []string `mandatory:"false" json:"updatedFieldDisplayNames"`

	// A list of parser names with conflicts.
	ConflictParserNames []string `mandatory:"false" json:"conflictParserNames"`

	// A list of source names with conflicts.
	ConflictSourceNames []string `mandatory:"false" json:"conflictSourceNames"`

	// A list of field display names with conflicts.
	ConflictFieldDisplayNames []string `mandatory:"false" json:"conflictFieldDisplayNames"`
}

func (m LogAnalyticsImportCustomChangeList) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsImportCustomChangeList) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
