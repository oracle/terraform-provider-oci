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

// RetrieveDimensionStatesDetails The configuration details for retrieving the alarm state entries.
// Filter retrieved alarm state entries by status value and dimension key-value pairs.
type RetrieveDimensionStatesDetails struct {

	// A filter to return only alarm state entries that match the exact set of specified dimension key-value pairs.
	// If you specify `"availabilityDomain": "phx-ad-1"` but the alarm state entry corresponds to the set `"availabilityDomain": "phx-ad-1"`
	// and `"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"`, then no results are returned.
	DimensionFilters map[string]string `mandatory:"false" json:"dimensionFilters"`

	// A filter to return only alarm state entries that match the status value.
	// Example: `FIRING`
	Status AlarmDimensionStatesEntryStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m RetrieveDimensionStatesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RetrieveDimensionStatesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAlarmDimensionStatesEntryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetAlarmDimensionStatesEntryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
