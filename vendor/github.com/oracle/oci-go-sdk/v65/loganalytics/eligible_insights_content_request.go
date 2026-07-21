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

// EligibleInsightsContentRequest Request structure for generating applicable insights from one or more log analytics sources.
type EligibleInsightsContentRequest struct {

	// The source internal names.
	Sources []string `mandatory:"true" json:"sources"`

	SearchOptions *SearchOptions `mandatory:"false" json:"searchOptions"`
}

func (m EligibleInsightsContentRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EligibleInsightsContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EligibleInsightsContentRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEligibleInsightsContentRequest EligibleInsightsContentRequest
	s := struct {
		DiscriminatorParam string `json:"contentKind"`
		MarshalTypeEligibleInsightsContentRequest
	}{
		"ELIGIBLE_BUSINESS_INSIGHTS",
		(MarshalTypeEligibleInsightsContentRequest)(m),
	}

	return json.Marshal(&s)
}
