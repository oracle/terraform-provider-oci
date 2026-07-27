// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOracleMigrationSettings Discriminator-only base update settings for Oracle migration methods.
type UpdateOracleMigrationSettings interface {
}

type updateoraclemigrationsettings struct {
	JsonData        []byte
	MigrationMethod string `json:"migrationMethod"`
}

// UnmarshalJSON unmarshals json
func (m *updateoraclemigrationsettings) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateoraclemigrationsettings updateoraclemigrationsettings
	s := struct {
		Model Unmarshalerupdateoraclemigrationsettings
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MigrationMethod = s.Model.MigrationMethod

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateoraclemigrationsettings) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MigrationMethod {
	case "ONLINE_STANDBY":
		mm := UpdateOracleOnlineStandbyMigrationSettings{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONLINE_PHYSICAL":
		mm := UpdateOracleOnlinePhysicalMigrationSettings{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONLINE_LOGICAL":
		mm := UpdateOracleOnlineLogicalMigrationSettings{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OFFLINE_LOGICAL":
		mm := UpdateOracleOfflineLogicalMigrationSettings{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateOracleMigrationSettings: %s.", m.MigrationMethod)
		return *m, nil
	}
}

func (m updateoraclemigrationsettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateoraclemigrationsettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
