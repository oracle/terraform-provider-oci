// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DatabaseUpgradeSourceBase Details for the database upgrade source.
type DatabaseUpgradeSourceBase interface {

	// Additional upgrade options supported by DBUA(Database Upgrade Assistant).
	// Example: "-upgradeTimezone false -keepEvents"
	GetOptions() *string
}

type databaseupgradesourcebase struct {
	JsonData []byte
	Options  *string `mandatory:"false" json:"options"`
	Source   string  `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *databaseupgradesourcebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseupgradesourcebase databaseupgradesourcebase
	s := struct {
		Model Unmarshalerdatabaseupgradesourcebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Options = s.Model.Options
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databaseupgradesourcebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "DB_HOME":
		mm := DatabaseUpgradeWithDbHomeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_SOFTWARE_IMAGE":
		mm := DatabaseUpgradeWithDatabaseSoftwareImageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_VERSION":
		mm := DatabaseUpgradeWithDbVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetOptions returns Options
func (m databaseupgradesourcebase) GetOptions() *string {
	return m.Options
}

func (m databaseupgradesourcebase) String() string {
	return common.PointerString(m)
}

// DatabaseUpgradeSourceBaseSourceEnum Enum with underlying type: string
type DatabaseUpgradeSourceBaseSourceEnum string

// Set of constants representing the allowable values for DatabaseUpgradeSourceBaseSourceEnum
const (
	DatabaseUpgradeSourceBaseSourceHome          DatabaseUpgradeSourceBaseSourceEnum = "DB_HOME"
	DatabaseUpgradeSourceBaseSourceVersion       DatabaseUpgradeSourceBaseSourceEnum = "DB_VERSION"
	DatabaseUpgradeSourceBaseSourceSoftwareImage DatabaseUpgradeSourceBaseSourceEnum = "DB_SOFTWARE_IMAGE"
)

var mappingDatabaseUpgradeSourceBaseSource = map[string]DatabaseUpgradeSourceBaseSourceEnum{
	"DB_HOME":           DatabaseUpgradeSourceBaseSourceHome,
	"DB_VERSION":        DatabaseUpgradeSourceBaseSourceVersion,
	"DB_SOFTWARE_IMAGE": DatabaseUpgradeSourceBaseSourceSoftwareImage,
}

// GetDatabaseUpgradeSourceBaseSourceEnumValues Enumerates the set of values for DatabaseUpgradeSourceBaseSourceEnum
func GetDatabaseUpgradeSourceBaseSourceEnumValues() []DatabaseUpgradeSourceBaseSourceEnum {
	values := make([]DatabaseUpgradeSourceBaseSourceEnum, 0)
	for _, v := range mappingDatabaseUpgradeSourceBaseSource {
		values = append(values, v)
	}
	return values
}
