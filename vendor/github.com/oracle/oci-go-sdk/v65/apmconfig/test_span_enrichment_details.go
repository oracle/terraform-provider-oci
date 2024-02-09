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

// TestSpanEnrichmentDetails Run a set of span enrichment rules on a given span to see the result.
type TestSpanEnrichmentDetails struct {

	// The span enrichment rules to test in the format of an Options resource.
	Options *interface{} `mandatory:"true" json:"options"`

	// The span to test the rules on. This should be a valid JSON object that follows one
	// of the formats used by distributed tracing frameworks, such as OpenTelemetry, Zipkin, or
	// Oracle Application Performance Monitoring.
	Span *interface{} `mandatory:"true" json:"span"`

	// A list of filters to try against the given span.
	Filters []FilterTextOrId `mandatory:"false" json:"filters"`
}

func (m TestSpanEnrichmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestSpanEnrichmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TestSpanEnrichmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTestSpanEnrichmentDetails TestSpanEnrichmentDetails
	s := struct {
		DiscriminatorParam string `json:"testType"`
		MarshalTypeTestSpanEnrichmentDetails
	}{
		"SPAN_ENRICHMENT",
		(MarshalTypeTestSpanEnrichmentDetails)(m),
	}

	return json.Marshal(&s)
}
