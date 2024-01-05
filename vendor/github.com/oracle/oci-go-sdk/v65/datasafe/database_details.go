// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseDetails Details of the database for the registration in Data Safe.
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
		common.Logf("Recieved unsupported enum value for DatabaseDetails: %s.", m.DatabaseType)
		return *m, nil
	}
}

// GetInfrastructureType returns InfrastructureType
func (m databasedetails) GetInfrastructureType() InfrastructureTypeEnum {
	return m.InfrastructureType
}

func (m databasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetInfrastructureTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
