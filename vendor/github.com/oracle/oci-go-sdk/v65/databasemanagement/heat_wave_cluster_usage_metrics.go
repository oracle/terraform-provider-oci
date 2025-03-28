// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HeatWaveClusterUsageMetrics The list of aggregated metrics for a HeatWave cluster in the fleet.
type HeatWaveClusterUsageMetrics struct {

	// The status of the HeatWave cluster. Indicates whether the status of the cluster is UP, DOWN, or UNKNOWN at the current time.
	Status HeatWaveClusterStatusEnum `mandatory:"true" json:"status"`

	// The OCID of the MySQL Database System associated with the HeatWave cluster.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The name of the MySQL Database System associated with the HeatWave cluster.
	DbSystemName *string `mandatory:"true" json:"dbSystemName"`

	// The name of the HeatWave cluster.
	HeatWaveClusterDisplayName *string `mandatory:"true" json:"heatWaveClusterDisplayName"`

	// The number of nodes in the HeatWave cluster.
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// Indicates whether Lakehouse is enabled for the HeatWave cluster or not.
	IsLakehouseEnabled *bool `mandatory:"true" json:"isLakehouseEnabled"`

	// The shape of the nodes in the HeatWave cluster.
	HeatWaveNodeShape *string `mandatory:"true" json:"heatWaveNodeShape"`

	// The total memory belonging to the HeatWave cluster in GBs.
	MemorySize *int `mandatory:"true" json:"memorySize"`

	// A list of the HeatWave cluster health metrics like CPU and Memory.
	Metrics []HeatWaveFleetMetricDefinition `mandatory:"true" json:"metrics"`
}

func (m HeatWaveClusterUsageMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeatWaveClusterUsageMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHeatWaveClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetHeatWaveClusterStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
