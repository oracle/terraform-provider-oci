// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// GrokPattern grok pattern object.
type GrokPattern struct {

	// The grok pattern.
	Pattern *string `mandatory:"true" json:"pattern"`

	// The name key to tag this grok pattern.
	Name *string `mandatory:"false" json:"name"`

	// Specify the time field for the event time. If the event doesn't have this field, the current time is used.
	FieldTimeKey *string `mandatory:"false" json:"fieldTimeKey"`

	// Process value using the specified format. This is available only when time_type is a string.
	FieldTimeFormat *string `mandatory:"false" json:"fieldTimeFormat"`

	// Use the specified time zone. The time value can be parsed or formatted in the specified time zone.
	FieldTimeZone *string `mandatory:"false" json:"fieldTimeZone"`
}

func (m GrokPattern) String() string {
	return common.PointerString(m)
}
