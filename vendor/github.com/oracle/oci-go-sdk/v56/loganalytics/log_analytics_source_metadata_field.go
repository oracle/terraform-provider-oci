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

// LogAnalyticsSourceMetadataField LogAnalyticsSourceMetadataField
type LogAnalyticsSourceMetadataField struct {

	// The field internal name.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// A flag inidcating whether or not the source metadata field is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The key.
	Key *string `mandatory:"false" json:"key"`

	// The source internal name.
	SourceName *string `mandatory:"false" json:"sourceName"`
}

func (m LogAnalyticsSourceMetadataField) String() string {
	return common.PointerString(m)
}
