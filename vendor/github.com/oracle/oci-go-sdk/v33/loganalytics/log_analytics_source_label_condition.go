// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// LogAnalyticsSourceLabelCondition LogAnalyticsSourceLabelCondition
type LogAnalyticsSourceLabelCondition struct {

	// message
	Message *string `mandatory:"false" json:"message"`

	// visible flag
	IsVisible *bool `mandatory:"false" json:"isVisible"`

	// block condition field
	BlockConditionField *string `mandatory:"false" json:"blockConditionField"`

	// block condition operator
	BlockConditionOperator *string `mandatory:"false" json:"blockConditionOperator"`

	// block condition value
	BlockConditionValue *string `mandatory:"false" json:"blockConditionValue"`

	// condition value
	LabelConditionValue *string `mandatory:"false" json:"labelConditionValue"`

	// list of condition values
	LabelConditionValues []string `mandatory:"false" json:"labelConditionValues"`

	// content example
	ContentExample *string `mandatory:"false" json:"contentExample"`

	// enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// field internal name
	FieldName *string `mandatory:"false" json:"fieldName"`

	// Id
	LabelConditionId *int64 `mandatory:"false" json:"labelConditionId"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// operator
	LabelConditionOperator *string `mandatory:"false" json:"labelConditionOperator"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// label display name
	LabelDisplayName *string `mandatory:"false" json:"labelDisplayName"`

	// label storage field
	StorageField *string `mandatory:"false" json:"storageField"`

	// label name
	LabelName *string `mandatory:"false" json:"labelName"`

	// inline label exists in DB flag
	IsInlineLabelExistingInDatabase *bool `mandatory:"false" json:"isInlineLabelExistingInDatabase"`
}

func (m LogAnalyticsSourceLabelCondition) String() string {
	return common.PointerString(m)
}
