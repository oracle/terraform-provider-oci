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

// EventType The event type.
type EventType struct {

	// The name of the event type.
	EventTypeName *string `mandatory:"false" json:"eventTypeName"`

	// The version.
	SpecVersion *string `mandatory:"false" json:"specVersion"`

	// A flag indicating whether or not the event type is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// A flag indicating whether or not the event type is user defined.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The last updated time.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m EventType) String() string {
	return common.PointerString(m)
}
