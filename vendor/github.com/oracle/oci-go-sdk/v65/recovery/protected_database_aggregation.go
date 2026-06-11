// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProtectedDatabaseAggregation The aggregated metrics for a protected database based on the specified `groupBy` dimensions.
// For example:
//
//	```
//
//	{
//	  "dimensions": { "health": "PROTECTED", "isRedoLogsEnabled": "true" },
//	  "metricName": "count",
//	  "value": 14
//	}
//
// ```
// In this example, the requested `groupBy` dimension is [`health`, `isRedoLogsEnabled`].
// The API returns aggregated metrics for multiple protected databases each with a unique combination of `health` and `isRedoLogsEnabled` values.
type ProtectedDatabaseAggregation struct {

	// The name of the protected database metric.
	// The `metricName` parameter defaults to `COUNT` which indicates the total number of protected databases that qualify for the respective dimension.
	// `BACKUP_SPACE_USED_IN_GBS` indicates the total storage space (in GB) used by protected databases that qualify for the respective dimension.
	MetricName ProtectedDatabaseAnalyticsMetricEnum `mandatory:"true" json:"metricName"`

	// The aggregated value of the requested metricName.
	Value *float64 `mandatory:"true" json:"value"`

	// A key value pair that maps the requested `groupBy` dimensions and values.
	// For example:
	// ```
	// { "health" : "PROTECTED", "isRedoLogsEnabled": "true" }
	// ```
	// A dimension defines the protected database group properties for which the API returns aggregated metrics.
	// Note: Protected database metrics can be grouped by the protected database health status (`health`),
	// the real-time data protection status (`isRedoLogsEnabled`), the associated protection policy (`protectionPolicyId`), and the life cycle state (`lifecycleState`).
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// If the requested `groupBy` dimension includes `protectionPolicyId`, then the aggregated metrics returns additional information about the protection policy, such as the policy `displayName` and `retentionDays`.
	// For example:
	// ```
	// "metadata": {
	//   "displayName": "Silver",
	//   "retentionDays": 35,
	//   "isPredefinedPolicy": true
	// }
	// ```
	Metadata *interface{} `mandatory:"false" json:"metadata"`
}

func (m ProtectedDatabaseAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProtectedDatabaseAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProtectedDatabaseAnalyticsMetricEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetProtectedDatabaseAnalyticsMetricEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
