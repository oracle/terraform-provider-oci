// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddHeatWaveClusterDetails Details required to add a HeatWave cluster.
type AddHeatWaveClusterDetails struct {

	// The shape determines resources to allocate to the HeatWave
	// nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing nodes provisioned for the
	// HeatWave cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	// Enable/disable Lakehouse for the HeatWave cluster.
	IsLakehouseEnabled *bool `mandatory:"false" json:"isLakehouseEnabled"`
}

func (m AddHeatWaveClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddHeatWaveClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
