// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseDetails Details of the database for the registration in Data Safe.
// To choose applicable database type and infrastructure type refer to
// https://confluence.oci.oraclecorp.com/display/DATASAFE/Target+V2+Design
type DatabaseDetails interface {

	// The infrastructure type the database is running on.
	GetInfrastructureType() InfrastructureTypeEnum
}

type databasedetails struct {
	JsonData           []byte
	InfrastructureType InfrastructureTypeEnum `mandatory:"true" json:"infrastructureType"`
	DatabaseType       string                 `json:"databaseType"`
}

// UnmarshalJSON unmarshals json
func (m *databasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasedetails databasedetails
	s := struct {
		Model Unmarshalerdatabasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InfrastructureType = s.Model.InfrastructureType
	m.DatabaseType = s.Model.DatabaseType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseType {
	case "INSTALLED_DATABASE":
		mm := InstalledDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := AutonomousDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_CLOUD_SERVICE":
		mm := DatabaseCloudServiceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetInfrastructureType returns InfrastructureType
func (m databasedetails) GetInfrastructureType() InfrastructureTypeEnum {
	return m.InfrastructureType
}

func (m databasedetails) String() string {
	return common.PointerString(m)
}
