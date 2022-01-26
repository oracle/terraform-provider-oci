// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TargetDatabaseSummary Summary of a Data Safe target database.
type TargetDatabaseSummary struct {

	// The OCID of the compartment that contains the Data Safe target database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Data Safe target database.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the target database in Data Safe.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The infrastructure type the database is running on.
	InfrastructureType InfrastructureTypeEnum `mandatory:"true" json:"infrastructureType"`

	// The database type.
	DatabaseType DatabaseTypeEnum `mandatory:"true" json:"databaseType"`

	// The current state of the target database in Data Safe.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the database was registered in Data Safe and created as a target database in Data Safe.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The description of the target database in Data Safe.
	Description *string `mandatory:"false" json:"description"`

	// Details about the current state of the target database in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TargetDatabaseSummary) String() string {
	return common.PointerString(m)
}
