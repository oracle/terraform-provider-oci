// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricDataDetails A metric object containing raw metric data points to be posted to the Monitoring service.
type MetricDataDetails struct {

	// The source service or application emitting the metric.
	// A valid namespace value starts with an alphabetical character and includes only alphanumeric characters and underscores. Custom metrics can't use the following reserved prefixes: `oci_` and `oracle_`
	// Avoid entering confidential information.
	// Example: `my_namespace`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to use for metrics.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the metric.
	// A valid name value starts with an alphabetical character and includes only alphanumeric characters, dots, underscores, hyphens, and dollar signs.
	// Avoid entering confidential information.
	// Example: `my_app.success_rate`
	Name *string `mandatory:"true" json:"name"`

	// Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
	// Each dimension takes the form of a key-value pair.
	// A valid dimension key includes only printable ASCII, excluding spaces. The character limit for a dimension key is 256.
	// A valid dimension value includes only Unicode characters. The character limit for a dimension value is 512.
	// Empty strings are not allowed for keys or values. Avoid entering confidential information.
	// Example: `{"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"}`
	Dimensions map[string]string `mandatory:"true" json:"dimensions"`

	// A list of metric values with timestamps. At least one data point is required per call.
	// For a data point to be posted, its timestamp must be near current time (less than two hours in the past and less than 10 minutes in the future).
	Datapoints []Datapoint `mandatory:"true" json:"datapoints"`

	// Resource group to assign to the metric. A resource group is a custom string that you can match when retrieving custom metrics. Only one resource group can be applied per metric.
	// A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	// Avoid entering confidential information.
	// Example: `frontend-fleet`
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// Properties describing metrics. These are not part of the unique fields identifying the metric.
	// Each metadata item takes the form of a key-value pair. The character limit for a metadata key is 256. The character limit for a metadata value is 256.
	// Example: `"unit": "bytes"`
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m MetricDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
