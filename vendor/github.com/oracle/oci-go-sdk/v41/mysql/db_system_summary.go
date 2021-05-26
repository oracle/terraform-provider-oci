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

// DbSystemSummary A summary of a DB System.
type DbSystemSummary struct {

	// The OCID of the DB System.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the DB System.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the MySQL Version in use for the DB System.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The date and time the DB System was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DB System was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// User-provided data about the DB System.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the compartment the DB System belongs in.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// If the policy is to enable high availability of the instance, by
	// maintaining secondary/failover capacity as necessary.
	IsHighlyAvailable *bool `mandatory:"false" json:"isHighlyAvailable"`

	CurrentPlacement *DbSystemPlacement `mandatory:"false" json:"currentPlacement"`

	// DEPRECATED -- please use `isHeatWaveClusterAttached` instead.
	// If the DB System has an Analytics Cluster attached.
	IsAnalyticsClusterAttached *bool `mandatory:"false" json:"isAnalyticsClusterAttached"`

	AnalyticsCluster *AnalyticsClusterSummary `mandatory:"false" json:"analyticsCluster"`

	// If the DB System has a HeatWave Cluster attached.
	IsHeatWaveClusterAttached *bool `mandatory:"false" json:"isHeatWaveClusterAttached"`

	HeatWaveCluster *HeatWaveClusterSummary `mandatory:"false" json:"heatWaveCluster"`

	// The availability domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.
	// In a failover scenario, the Read/Write endpoint is redirected to one of the other availability domains
	// and the MySQL instance in that domain is promoted to the primary instance.
	// This redirection does not affect the IP address of the DB System in any way.
	// For a standalone DB System, this defines the availability domain in which the DB System is placed.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The fault domain on which to deploy the Read/Write endpoint. This defines the preferred primary instance.
	// In a failover scenario, the Read/Write endpoint is redirected to one of the other fault domains
	// and the MySQL instance in that domain is promoted to the primary instance.
	// This redirection does not affect the IP address of the DB System in any way.
	// For a standalone DB System, this defines the fault domain in which the DB System is placed.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The network endpoints available for this DB System.
	Endpoints []DbSystemEndpoint `mandatory:"false" json:"endpoints"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DbSystemSummary) String() string {
	return common.PointerString(m)
}
