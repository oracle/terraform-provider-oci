// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HeatWaveFleetSummary The inventory count of HeatWave clusters in the fleet and a summary of the metrics that provide the aggregated usage of CPU, storage, and so on of all the clusters.
type HeatWaveFleetSummary struct {

	// The usage metrics for the HeatWave clusters in the fleet.
	AggregatedMetrics []HeatWaveFleetMetricSummaryDefinition `mandatory:"true" json:"aggregatedMetrics"`

	// The number of HeatWave clusters in the fleet, grouped by cluster type or other properties.
	Inventory []HeatWaveFleetByCategory `mandatory:"true" json:"inventory"`
}

func (m HeatWaveFleetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeatWaveFleetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
