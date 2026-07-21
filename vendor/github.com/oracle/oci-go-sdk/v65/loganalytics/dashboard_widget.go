// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DashboardWidget Dashboard widget object.
type DashboardWidget struct {

	// The widget label.
	Label *string `mandatory:"true" json:"label"`

	// The widget visualization.
	VizType *string `mandatory:"true" json:"vizType"`

	// The widget query.
	Query *string `mandatory:"true" json:"query"`

	// The widget description.
	Description *string `mandatory:"false" json:"description"`

	// List of field internal names which can be used as filters for this widget. These values must exist in the dashboard field filters.
	InputFieldFilters []string `mandatory:"false" json:"inputFieldFilters"`

	// JSON that contains user interface options for the widget.
	Options *interface{} `mandatory:"false" json:"options"`
}

func (m DashboardWidget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DashboardWidget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
