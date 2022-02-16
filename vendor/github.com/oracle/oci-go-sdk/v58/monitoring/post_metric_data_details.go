// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PostMetricDataDetails An array of metric objects containing raw metric data points to be posted to the Monitoring service.
type PostMetricDataDetails struct {

	// A metric object containing raw metric data points to be posted to the Monitoring service.
	MetricData []MetricDataDetails `mandatory:"true" json:"metricData"`

	// Batch atomicity behavior. Requires either partial or full pass of input validation for
	// metric objects in PostMetricData requests. The default value of NON_ATOMIC requires a
	// partial pass: at least one metric object in the request must pass input validation, and
	// any objects that failed validation are identified in the returned summary, along with
	// their error messages. A value of ATOMIC requires a full pass: all metric objects in
	// the request must pass input validation.
	// Example: `NON_ATOMIC`
	BatchAtomicity PostMetricDataDetailsBatchAtomicityEnum `mandatory:"false" json:"batchAtomicity,omitempty"`
}

func (m PostMetricDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PostMetricDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPostMetricDataDetailsBatchAtomicityEnum(string(m.BatchAtomicity)); !ok && m.BatchAtomicity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BatchAtomicity: %s. Supported values are: %s.", m.BatchAtomicity, strings.Join(GetPostMetricDataDetailsBatchAtomicityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PostMetricDataDetailsBatchAtomicityEnum Enum with underlying type: string
type PostMetricDataDetailsBatchAtomicityEnum string

// Set of constants representing the allowable values for PostMetricDataDetailsBatchAtomicityEnum
const (
	PostMetricDataDetailsBatchAtomicityAtomic    PostMetricDataDetailsBatchAtomicityEnum = "ATOMIC"
	PostMetricDataDetailsBatchAtomicityNonAtomic PostMetricDataDetailsBatchAtomicityEnum = "NON_ATOMIC"
)

var mappingPostMetricDataDetailsBatchAtomicityEnum = map[string]PostMetricDataDetailsBatchAtomicityEnum{
	"ATOMIC":     PostMetricDataDetailsBatchAtomicityAtomic,
	"NON_ATOMIC": PostMetricDataDetailsBatchAtomicityNonAtomic,
}

// GetPostMetricDataDetailsBatchAtomicityEnumValues Enumerates the set of values for PostMetricDataDetailsBatchAtomicityEnum
func GetPostMetricDataDetailsBatchAtomicityEnumValues() []PostMetricDataDetailsBatchAtomicityEnum {
	values := make([]PostMetricDataDetailsBatchAtomicityEnum, 0)
	for _, v := range mappingPostMetricDataDetailsBatchAtomicityEnum {
		values = append(values, v)
	}
	return values
}

// GetPostMetricDataDetailsBatchAtomicityEnumStringValues Enumerates the set of values in String for PostMetricDataDetailsBatchAtomicityEnum
func GetPostMetricDataDetailsBatchAtomicityEnumStringValues() []string {
	return []string{
		"ATOMIC",
		"NON_ATOMIC",
	}
}

// GetMappingPostMetricDataDetailsBatchAtomicityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPostMetricDataDetailsBatchAtomicityEnum(val string) (PostMetricDataDetailsBatchAtomicityEnum, bool) {
	mappingPostMetricDataDetailsBatchAtomicityEnumIgnoreCase := make(map[string]PostMetricDataDetailsBatchAtomicityEnum)
	for k, v := range mappingPostMetricDataDetailsBatchAtomicityEnum {
		mappingPostMetricDataDetailsBatchAtomicityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPostMetricDataDetailsBatchAtomicityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
