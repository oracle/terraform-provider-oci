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

// ExternalAsmServicedDatabase The details of a database serviced by an external ASM.
type ExternalAsmServicedDatabase struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique name of the external database.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the external database resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// The subtype of Oracle Database. Indicates whether the database is a Container Database,
	// Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// Indicates whether the database is a Managed Database or not.
	IsManaged *bool `mandatory:"false" json:"isManaged"`

	// The list of ASM disk groups used by the database.
	DiskGroups []string `mandatory:"false" json:"diskGroups"`
}

func (m ExternalAsmServicedDatabase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalAsmServicedDatabase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
