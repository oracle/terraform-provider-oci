// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DashboardMetadataResponse Response structure for GenAI generated dashboard metadata for one or more log analytics sources.
type DashboardMetadataResponse struct {

	// The dashboard name.
	Name *string `mandatory:"true" json:"name"`

	// The field filters available for use across widgets in the dashboard.
	FieldFilters []DashboardFieldFilter `mandatory:"true" json:"fieldFilters"`

	// The grouped dashboard widget content.
	WidgetGroups []DashboardWidgetGroup `mandatory:"true" json:"widgetGroups"`

	// The dashboard description.
	Description *string `mandatory:"false" json:"description"`

	// JSON that contains user interface options for the dashboard.
	Options *interface{} `mandatory:"false" json:"options"`
}

func (m DashboardMetadataResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DashboardMetadataResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DashboardMetadataResponse) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDashboardMetadataResponse DashboardMetadataResponse
	s := struct {
		DiscriminatorParam string `json:"contentKind"`
		MarshalTypeDashboardMetadataResponse
	}{
		"DASHBOARD_CONTENT",
		(MarshalTypeDashboardMetadataResponse)(m),
	}

	return json.Marshal(&s)
}
