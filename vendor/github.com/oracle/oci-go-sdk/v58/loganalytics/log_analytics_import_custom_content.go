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

// LogAnalyticsImportCustomContent LogAnalyticsImportCustomContent
type LogAnalyticsImportCustomContent struct {

	// The parser names.
	ParserNames []string `mandatory:"false" json:"parserNames"`

	// The source names.
	SourceNames []string `mandatory:"false" json:"sourceNames"`

	// The field names.
	FieldNames []string `mandatory:"false" json:"fieldNames"`

	ChangeList *LogAnalyticsImportCustomChangeList `mandatory:"false" json:"changeList"`

	// The content name.
	ContentName *string `mandatory:"false" json:"contentName"`
}

func (m LogAnalyticsImportCustomContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsImportCustomContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
