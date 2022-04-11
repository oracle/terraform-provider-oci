// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// OptimizerDatabase The summary of the managed database resource.
type OptimizerDatabase struct {

	// Database ocid.
	Id *string `mandatory:"true" json:"id"`

	// Database name.
	Name *string `mandatory:"true" json:"name"`

	// The type of Oracle Database installation.
	DbType DatabaseTypeEnum `mandatory:"true" json:"dbType"`

	// The subtype of the Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database.
	DbSubType DatabaseSubTypeEnum `mandatory:"true" json:"dbSubType"`

	// The infrastructure used to deploy the Oracle Database.
	DbDeploymentType DeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`

	// Database version.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// Managed database resource provisioned compartment id.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m OptimizerDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OptimizerDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseTypeEnum(string(m.DbType)); !ok && m.DbType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbType: %s. Supported values are: %s.", m.DbType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DbSubType)); !ok && m.DbSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbSubType: %s. Supported values are: %s.", m.DbSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
