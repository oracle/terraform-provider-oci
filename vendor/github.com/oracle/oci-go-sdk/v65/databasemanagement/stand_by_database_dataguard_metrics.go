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

// StandByDatabaseDataguardMetrics The standby database details.
type StandByDatabaseDataguardMetrics struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	DbId *string `mandatory:"true" json:"dbId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Managed Database resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Managed Database.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The database role of the Managed Database.
	DbRole DbRoleEnum `mandatory:"true" json:"dbRole"`

	// A list of standby databases with latest values for metrics such as ApplyLag, TransportLag, and RedoApplyRate.
	Metrics []HaMetricDefinition `mandatory:"true" json:"metrics"`

	// The database ID of the Managed Database. Every database had its own ID and that value is captured here.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The database unique name of the Managed Database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`
}

func (m StandByDatabaseDataguardMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StandByDatabaseDataguardMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbRoleEnum(string(m.DbRole)); !ok && m.DbRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbRole: %s. Supported values are: %s.", m.DbRole, strings.Join(GetDbRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
