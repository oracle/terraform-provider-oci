// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EventContent Information about the data collected as a ZIP file when the event occurred.
type EventContent struct {

	// Status of the event content
	ContentAvailability ContentAvailabilityEnum `mandatory:"false" json:"contentAvailability,omitempty"`

	// Path to the event content on the instance
	InstancePath *string `mandatory:"false" json:"instancePath"`

	// size in bytes of the event content (size of the zip file uploaded)
	Size *int `mandatory:"false" json:"size"`
}

func (m EventContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EventContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContentAvailabilityEnum(string(m.ContentAvailability)); !ok && m.ContentAvailability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentAvailability: %s. Supported values are: %s.", m.ContentAvailability, strings.Join(GetContentAvailabilityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
