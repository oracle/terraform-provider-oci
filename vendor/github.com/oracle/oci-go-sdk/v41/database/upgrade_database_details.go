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

// UpgradeDatabaseDetails Details for upgrading a database to a specific Oracle Database version.
type UpgradeDatabaseDetails struct {

	// The database upgrade action.
	Action UpgradeDatabaseDetailsActionEnum `mandatory:"true" json:"action"`

	DatabaseUpgradeSourceDetails DatabaseUpgradeSourceBase `mandatory:"false" json:"databaseUpgradeSourceDetails"`
}

func (m UpgradeDatabaseDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpgradeDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DatabaseUpgradeSourceDetails databaseupgradesourcebase        `json:"databaseUpgradeSourceDetails"`
		Action                       UpgradeDatabaseDetailsActionEnum `json:"action"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DatabaseUpgradeSourceDetails.UnmarshalPolymorphicJSON(model.DatabaseUpgradeSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseUpgradeSourceDetails = nn.(DatabaseUpgradeSourceBase)
	} else {
		m.DatabaseUpgradeSourceDetails = nil
	}

	m.Action = model.Action

	return
}

// UpgradeDatabaseDetailsActionEnum Enum with underlying type: string
type UpgradeDatabaseDetailsActionEnum string

// Set of constants representing the allowable values for UpgradeDatabaseDetailsActionEnum
const (
	UpgradeDatabaseDetailsActionPrecheck UpgradeDatabaseDetailsActionEnum = "PRECHECK"
	UpgradeDatabaseDetailsActionUpgrade  UpgradeDatabaseDetailsActionEnum = "UPGRADE"
	UpgradeDatabaseDetailsActionRollback UpgradeDatabaseDetailsActionEnum = "ROLLBACK"
)

var mappingUpgradeDatabaseDetailsAction = map[string]UpgradeDatabaseDetailsActionEnum{
	"PRECHECK": UpgradeDatabaseDetailsActionPrecheck,
	"UPGRADE":  UpgradeDatabaseDetailsActionUpgrade,
	"ROLLBACK": UpgradeDatabaseDetailsActionRollback,
}

// GetUpgradeDatabaseDetailsActionEnumValues Enumerates the set of values for UpgradeDatabaseDetailsActionEnum
func GetUpgradeDatabaseDetailsActionEnumValues() []UpgradeDatabaseDetailsActionEnum {
	values := make([]UpgradeDatabaseDetailsActionEnum, 0)
	for _, v := range mappingUpgradeDatabaseDetailsAction {
		values = append(values, v)
	}
	return values
}
