// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v37/common"
)

// ManagedDatabase The details of a Managed Database.
type ManagedDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the Managed Database.
	Name *string `mandatory:"true" json:"name"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"true" json:"databaseType"`

	// The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"true" json:"databaseSubType"`

	// Indicates whether the Oracle Database is part of a cluster.
	IsCluster *bool `mandatory:"true" json:"isCluster"`

	// The date and time the Managed Database was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent Container Database
	// if Managed Database is a Pluggable Database.
	ParentContainerId *string `mandatory:"false" json:"parentContainerId"`

	// A list of Managed Database Groups that the Managed Database belongs to.
	ManagedDatabaseGroups []ParentGroup `mandatory:"false" json:"managedDatabaseGroups"`

	// The status of the Oracle Database. Indicates whether the status of the database
	// is UP, DOWN, or UNKNOWN at the current time.
	DatabaseStatus DatabaseStatusEnum `mandatory:"false" json:"databaseStatus,omitempty"`

	// The additional details specific to a type of database defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`
}

func (m ManagedDatabase) String() string {
	return common.PointerString(m)
}
