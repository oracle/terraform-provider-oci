// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// GrokPattern grok pattern object
type GrokPattern struct {

	// The grok pattern
	Pattern *string `mandatory:"true" json:"pattern"`

	// The name key to tag this grok pattern
	Name *string `mandatory:"false" json:"name"`

	// Specify time field for event time. If the event doesn't have this field, current time is used.
	FieldTimeKey *string `mandatory:"false" json:"fieldTimeKey"`

	// Process value using specified format. This is available only when time_type is string.
	FieldTimeFormat *string `mandatory:"false" json:"fieldTimeFormat"`

	// Use specified timezone. One can parse/format the time value in the specified timezone.
	FieldTimeZone *string `mandatory:"false" json:"fieldTimeZone"`
}

func (m GrokPattern) String() string {
	return common.PointerString(m)
}
