// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Metric The properties that define a metric.
// For information about metrics, see
// Metrics Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#MetricsOverview).
type Metric struct {

	// The name of the metric.
	// Example: `CpuUtilization`
	Name *string `mandatory:"false" json:"name"`

	// The source service or application emitting the metric.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"false" json:"namespace"`

	// Resource group provided with the posted metric. A resource group is a custom string that you can match when retrieving custom metrics. Only one resource group can be applied per metric.
	// A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	// Example: `frontend-fleet`
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing
	// the resources monitored by the metric.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
	// Each dimension takes the form of a key-value pair.
	// Example: `{"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"}`
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`
}

func (m Metric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Metric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
