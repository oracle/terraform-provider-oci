// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Metric The properties that define a metric.
// For information about metrics, see Metrics Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#MetricsOverview).
type Metric struct {

	// The name of the metric.
	// Example: `CpuUtilization`
	Name *string `mandatory:"false" json:"name"`

	// The source service or application emitting the metric.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"false" json:"namespace"`

	// Resource group provided with the posted metric. A resource group is a custom string that can be used as a filter. Only one resource group can be applied per metric.
	// A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	// Avoid entering confidential information.
	// Example: `frontend-fleet`
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing
	// the resources monitored by the metric.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
	// Each dimension takes the form of a key-value pair.
	// Example: `"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"`
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`
}

func (m Metric) String() string {
	return common.PointerString(m)
}
