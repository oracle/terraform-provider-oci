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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EventType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
