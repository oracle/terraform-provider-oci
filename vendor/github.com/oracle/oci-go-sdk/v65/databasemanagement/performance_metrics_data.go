// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PerformanceMetricsData The list of Data Guard performance metrics for Managed Databases.
type PerformanceMetricsData struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The ID of the primary database.
	PrimaryDbId *string `mandatory:"false" json:"primaryDbId"`

	// The primary database unique name of the Managed Database.
	PrimaryDbUniqueName *string `mandatory:"false" json:"primaryDbUniqueName"`

	// The database ID of the Managed Database. Every database had its own ID and that value is captured here.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The database unique name of the Managed Database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The deployment type of the Managed Database.
	DeploymentType *string `mandatory:"false" json:"deploymentType"`

	// The resource name of the Managed Database.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The database role of the Managed Database.
	DbRole DbRoleEnum `mandatory:"false" json:"dbRole,omitempty"`

	// The list of Data Guard performance metrics such as ApplyLag, TransportLag and RedoApplyRate for the Managed Databases.
	Metrics []PerformanceMetrics `mandatory:"false" json:"metrics"`
}

func (m PerformanceMetricsData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PerformanceMetricsData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbRoleEnum(string(m.DbRole)); !ok && m.DbRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbRole: %s. Supported values are: %s.", m.DbRole, strings.Join(GetDbRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
