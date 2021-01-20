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

// LogAnalyticsSourceMetadataField LogAnalyticsSourceMetadataField
type LogAnalyticsSourceMetadataField struct {

	// field internal name
	FieldName *string `mandatory:"false" json:"fieldName"`

	// is enabled flag
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// key
	Key *string `mandatory:"false" json:"key"`

	// source internal name
	SourceName *string `mandatory:"false" json:"sourceName"`
}

func (m LogAnalyticsSourceMetadataField) String() string {
	return common.PointerString(m)
}
