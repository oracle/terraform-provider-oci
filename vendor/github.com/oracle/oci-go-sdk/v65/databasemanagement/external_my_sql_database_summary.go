// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalMySqlDatabaseSummary External database summary record.
type ExternalMySqlDatabaseSummary struct {

	// OCID of compartment for the External MySQL Database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Display name of the External MySQL Database.
	DbName *string `mandatory:"true" json:"dbName"`

	// OCID of External MySQL Database.
	ExternalDatabaseId *string `mandatory:"false" json:"externalDatabaseId"`

	// The OCID of the enabled MySQL Database Connector.
	ConnectorId *string `mandatory:"false" json:"connectorId"`

	// The deployment type of the Mysql Database.
	DeploymentType MySqlDeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// Indicates database management state.
	ManagementState ManagementStateEnum `mandatory:"false" json:"managementState,omitempty"`

	// Indicates lifecycle  state of the resource.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ExternalMySqlDatabaseSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalMySqlDatabaseSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMySqlDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetMySqlDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagementStateEnum(string(m.ManagementState)); !ok && m.ManagementState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementState: %s. Supported values are: %s.", m.ManagementState, strings.Join(GetManagementStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
