// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AddHeatWaveClusterDetails Details required to add a HeatWave cluster.
type AddHeatWaveClusterDetails struct {

	// The shape determines resources to allocate to the HeatWave
	// nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing nodes provisioned for the
	// HeatWave cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`
}

func (m AddHeatWaveClusterDetails) String() string {
	return common.PointerString(m)
}
