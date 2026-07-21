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

// DashboardMetadataContentRequest Request payload for generating dashboard content from one or more log analytics sources.
type DashboardMetadataContentRequest struct {

	// The source internal names.
	Sources []string `mandatory:"true" json:"sources"`

	// List of field internal names to use as field filters in the generated dashboard.
	FieldFilters []string `mandatory:"false" json:"fieldFilters"`

	// An optional free-form user prompt included with the request for custom dashboard tailoring.
	Prompt *string `mandatory:"false" json:"prompt"`

	SearchOptions *SearchOptions `mandatory:"false" json:"searchOptions"`

	// The top-level business insight categories (e.g., performance, security) for which this dashboard is being generated.
	BusinessInsights []BusinessInsightsCategoryEnum `mandatory:"true" json:"businessInsights"`
}

func (m DashboardMetadataContentRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DashboardMetadataContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.BusinessInsights {
		if _, ok := GetMappingBusinessInsightsCategoryEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BusinessInsights: %s. Supported values are: %s.", val, strings.Join(GetBusinessInsightsCategoryEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DashboardMetadataContentRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDashboardMetadataContentRequest DashboardMetadataContentRequest
	s := struct {
		DiscriminatorParam string `json:"contentKind"`
		MarshalTypeDashboardMetadataContentRequest
	}{
		"DASHBOARD_CONTENT",
		(MarshalTypeDashboardMetadataContentRequest)(m),
	}

	return json.Marshal(&s)
}
