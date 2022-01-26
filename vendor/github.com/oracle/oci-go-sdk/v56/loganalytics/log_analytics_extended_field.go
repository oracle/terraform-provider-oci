// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsExtendedField LogAnalyticsExtendedField
type LogAnalyticsExtendedField struct {
	Field *LogAnalyticsField `mandatory:"false" json:"field"`

	ExtendedFieldDefinition *LogAnalyticsSourceExtendedFieldDefinition `mandatory:"false" json:"extendedFieldDefinition"`

	// The extended field unique identifier.
	ExtendedFieldDefinitionId *int64 `mandatory:"false" json:"extendedFieldDefinitionId"`

	// The field internal name
	FieldName *string `mandatory:"false" json:"fieldName"`

	// The field internal display name.
	FieldDisplayName *string `mandatory:"false" json:"fieldDisplayName"`

	// The saved regular expression name.
	SavedRegularExpressionName *string `mandatory:"false" json:"savedRegularExpressionName"`

	// The extended field unique identifier.
	ExtendedFieldId *int64 `mandatory:"false" json:"extendedFieldId"`
}

func (m LogAnalyticsExtendedField) String() string {
	return common.PointerString(m)
}
