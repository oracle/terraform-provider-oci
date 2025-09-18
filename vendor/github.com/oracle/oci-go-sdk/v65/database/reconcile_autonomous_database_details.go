// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReconcileAutonomousDatabaseDetails Request payload for reconciling attributes in the Autonomous AI Database.
type ReconcileAutonomousDatabaseDetails struct {

	// A list of objects to be reconciled in the Autonomous AI Database.
	ReconciliationObjects []ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum `mandatory:"false" json:"reconciliationObjects,omitempty"`
}

func (m ReconcileAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReconcileAutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ReconciliationObjects {
		if _, ok := GetMappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReconciliationObjects: %s. Supported values are: %s.", val, strings.Join(GetReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum Enum with underlying type: string
type ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum string

// Set of constants representing the allowable values for ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum
const (
	ReconcileAutonomousDatabaseDetailsReconciliationObjectsDbTools ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum = "DB_TOOLS"
)

var mappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum = map[string]ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum{
	"DB_TOOLS": ReconcileAutonomousDatabaseDetailsReconciliationObjectsDbTools,
}

var mappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumLowerCase = map[string]ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum{
	"db_tools": ReconcileAutonomousDatabaseDetailsReconciliationObjectsDbTools,
}

// GetReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumValues Enumerates the set of values for ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum
func GetReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumValues() []ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum {
	values := make([]ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum, 0)
	for _, v := range mappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum {
		values = append(values, v)
	}
	return values
}

// GetReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumStringValues Enumerates the set of values in String for ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum
func GetReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumStringValues() []string {
	return []string{
		"DB_TOOLS",
	}
}

// GetMappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum(val string) (ReconcileAutonomousDatabaseDetailsReconciliationObjectsEnum, bool) {
	enum, ok := mappingReconcileAutonomousDatabaseDetailsReconciliationObjectsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
