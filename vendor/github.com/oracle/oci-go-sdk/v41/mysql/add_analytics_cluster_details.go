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

// AddAnalyticsClusterDetails DEPRECATED -- please use HeatWave API instead.
// Details required to add an Analytics Cluster.
type AddAnalyticsClusterDetails struct {

	// The shape determines resources to allocate to the Analytics
	// Cluster nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing nodes provisioned for the
	// Analytics Cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`
}

func (m AddAnalyticsClusterDetails) String() string {
	return common.PointerString(m)
}
