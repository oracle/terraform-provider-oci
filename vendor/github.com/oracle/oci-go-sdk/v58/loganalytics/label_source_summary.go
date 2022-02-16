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

// LabelSourceSummary source summary
type LabelSourceSummary struct {

	// The source display name.
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// The source internal name.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The label operator.
	LabelOperatorName *string `mandatory:"false" json:"labelOperatorName"`

	// The label condition.
	LabelCondition *string `mandatory:"false" json:"labelCondition"`

	// The label field display name.
	LabelFieldDisplayname *string `mandatory:"false" json:"labelFieldDisplayname"`

	// The label field name.
	LabelFieldName *string `mandatory:"false" json:"labelFieldName"`
}

func (m LabelSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LabelSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
