// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedMySqlDatabase The details of the Managed MySQL Database.
type ManagedMySqlDatabase struct {

	// The OCID of the Managed MySQL Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the MySQL Database.
	DbName *string `mandatory:"true" json:"dbName"`

	// The version of the MySQL Database.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The date and time the Managed MySQL Database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The name of the Managed MySQL Database.
	Name *string `mandatory:"true" json:"name"`

	// The name of the HeatWave cluster.
	HeatWaveClusterDisplayName *string `mandatory:"false" json:"heatWaveClusterDisplayName"`

	// If HeatWave is enabled for this db system or not.
	IsHeatWaveEnabled *bool `mandatory:"false" json:"isHeatWaveEnabled"`

	// If HeatWave Lakehouse is enabled for the db system or not.
	IsLakehouseEnabled *bool `mandatory:"false" json:"isLakehouseEnabled"`

	// Shape of the nodes in the HeatWave cluster.
	HeatWaveNodeShape *string `mandatory:"false" json:"heatWaveNodeShape"`

	// The total memory belonging to the HeatWave cluster in GBs.
	HeatWaveMemorySize *int `mandatory:"false" json:"heatWaveMemorySize"`

	// The information about an individual HeatWave nodes in the cluster.
	HeatWaveNodes []HeatWaveNode `mandatory:"false" json:"heatWaveNodes"`

	// If the HeatWave cluster is active or not.
	IsHeatWaveActive *bool `mandatory:"false" json:"isHeatWaveActive"`

	// The date and time the Managed MySQL Database was created.
	TimeCreatedHeatWave *common.SDKTime `mandatory:"false" json:"timeCreatedHeatWave"`
}

func (m ManagedMySqlDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
