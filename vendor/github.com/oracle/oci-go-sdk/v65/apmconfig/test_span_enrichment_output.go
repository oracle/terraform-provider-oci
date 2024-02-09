// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestSpanEnrichmentOutput Output of running a set of span enrichment rules against a span.
type TestSpanEnrichmentOutput struct {

	// The span after applying enrichment rules.
	Span *interface{} `mandatory:"false" json:"span"`

	// A list of booleans indicating whether the corresponding filter in the input matched the input span.
	Filters []bool `mandatory:"false" json:"filters"`
}

func (m TestSpanEnrichmentOutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestSpanEnrichmentOutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TestSpanEnrichmentOutput) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTestSpanEnrichmentOutput TestSpanEnrichmentOutput
	s := struct {
		DiscriminatorParam string `json:"testType"`
		MarshalTypeTestSpanEnrichmentOutput
	}{
		"SPAN_ENRICHMENT",
		(MarshalTypeTestSpanEnrichmentOutput)(m),
	}

	return json.Marshal(&s)
}
