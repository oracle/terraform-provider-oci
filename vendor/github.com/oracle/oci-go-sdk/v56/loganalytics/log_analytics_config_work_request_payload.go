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

// LogAnalyticsConfigWorkRequestPayload LogAnalyticsConfigWorkRequestPayload
type LogAnalyticsConfigWorkRequestPayload struct {

	// The source name.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The entity unique identifier.
	EntityId *string `mandatory:"false" json:"entityId"`

	// The lookup reference as an integer.
	LookupReference *int64 `mandatory:"false" json:"lookupReference"`

	// lookupReference
	LookupReferenceString *string `mandatory:"false" json:"lookupReferenceString"`
}

func (m LogAnalyticsConfigWorkRequestPayload) String() string {
	return common.PointerString(m)
}
