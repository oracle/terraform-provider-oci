// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// IngestTimeRuleMetricExtractionAction Details of metric to post to OCI Monitoring if ingest time rule condition(s) are satisfied.
type IngestTimeRuleMetricExtractionAction struct {

	// The compartment OCID (/iaas/Content/General/Concepts/identifiers.htm) of the extracted metric.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The namespace of the extracted metric.
	// A valid value starts with an alphabetical character and includes only
	// alphanumeric characters and underscores (_).
	Namespace *string `mandatory:"true" json:"namespace"`

	// The metric name of the extracted metric.
	// A valid value starts with an alphabetical character and includes only
	// alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	MetricName *string `mandatory:"true" json:"metricName"`

	// The resourceGroup of the extracted metric.
	// A valid value starts with an alphabetical character and includes only
	// alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// Additional dimensions to publish for the extracted metric.
	// A valid list contains the source field names whose values are to be published as dimensions.
	// The source name itself is specified using a special macro SOURCE_NAME
	Dimensions []string `mandatory:"false" json:"dimensions"`
}

func (m IngestTimeRuleMetricExtractionAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngestTimeRuleMetricExtractionAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IngestTimeRuleMetricExtractionAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIngestTimeRuleMetricExtractionAction IngestTimeRuleMetricExtractionAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIngestTimeRuleMetricExtractionAction
	}{
		"METRIC_EXTRACTION",
		(MarshalTypeIngestTimeRuleMetricExtractionAction)(m),
	}

	return json.Marshal(&s)
}
