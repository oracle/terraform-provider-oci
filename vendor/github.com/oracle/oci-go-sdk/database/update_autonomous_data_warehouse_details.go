// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAutonomousDataWarehouseDetails Details to update an Oracle Autonomous Data Warehouse.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAutonomousDataWarehouseDetails struct {

	// A strong password for Admin. The password must be between 12 and 60 characters long, and must contain at least 1 uppercase, 1 lowercase and 2 numeric characters. It cannot contain the double quote symbol ("). It must be different than the last 4 passwords.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// Size, in terabytes, of the data volume that will be attached to the database.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The user-friendly name for the Autonomous Data Warehouse. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdateAutonomousDataWarehouseDetails) String() string {
	return common.PointerString(m)
}
