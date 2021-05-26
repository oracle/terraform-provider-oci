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

// UpdateHeatWaveClusterDetails Details about the HeatWave cluster properties to be updated.
type UpdateHeatWaveClusterDetails struct {

	// A change to the shape of the nodes in the HeatWave cluster will
	// result in the entire cluster being torn down and re-created with
	// Compute instances of the new Shape. This may result in significant
	// downtime for the analytics capability while the HeatWave cluster is
	// re-provisioned.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// A change to the number of nodes in the HeatWave cluster will result
	// in the entire cluster being torn down and re-created with the new
	// cluster of nodes. This may result in a significant downtime for the
	// analytics capability while the HeatWave cluster is
	// re-provisioned.
	ClusterSize *int `mandatory:"false" json:"clusterSize"`
}

func (m UpdateHeatWaveClusterDetails) String() string {
	return common.PointerString(m)
}
