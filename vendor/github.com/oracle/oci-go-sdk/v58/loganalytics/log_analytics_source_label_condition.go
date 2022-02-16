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

// LogAnalyticsSourceLabelCondition LogAnalyticsSourceLabelCondition
type LogAnalyticsSourceLabelCondition struct {

	// The message.
	Message *string `mandatory:"false" json:"message"`

	// A flag indicating whether or not the label condition is visible.
	IsVisible *bool `mandatory:"false" json:"isVisible"`

	// The block condition field.
	BlockConditionField *string `mandatory:"false" json:"blockConditionField"`

	// The block condition operator.
	BlockConditionOperator *string `mandatory:"false" json:"blockConditionOperator"`

	// The block condition value.
	BlockConditionValue *string `mandatory:"false" json:"blockConditionValue"`

	// The condition value.
	LabelConditionValue *string `mandatory:"false" json:"labelConditionValue"`

	// A list of condition values.
	LabelConditionValues []string `mandatory:"false" json:"labelConditionValues"`

	// The content example.
	ContentExample *string `mandatory:"false" json:"contentExample"`

	// A flag inidcating whether or not the condition is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The internal field name.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// The unique identifier of the condition.
	LabelConditionId *int64 `mandatory:"false" json:"labelConditionId"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The condition operator.
	LabelConditionOperator *string `mandatory:"false" json:"labelConditionOperator"`

	// The unique identifier of the source.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The label display name.
	LabelDisplayName *string `mandatory:"false" json:"labelDisplayName"`

	// The label storage field.
	StorageField *string `mandatory:"false" json:"storageField"`

	// The label name.
	LabelName *string `mandatory:"false" json:"labelName"`

	// A flag indicating whether or not the inline label exists in the database.
	IsInlineLabelExistingInDatabase *bool `mandatory:"false" json:"isInlineLabelExistingInDatabase"`
}

func (m LogAnalyticsSourceLabelCondition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsSourceLabelCondition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
