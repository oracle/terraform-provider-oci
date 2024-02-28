// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDatabaseBase Details for creating a database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDatabaseBase interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	GetDbHomeId() *string

	// A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	GetDbVersion() *string

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	GetKmsKeyId() *string

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. Autonomous Database Serverless does not use key versions, hence is not applicable for Autonomous Database Serverless instances.
	GetKmsKeyVersionId() *string
}

type createdatabasebase struct {
	JsonData        []byte
	DbVersion       *string `mandatory:"false" json:"dbVersion"`
	KmsKeyId        *string `mandatory:"false" json:"kmsKeyId"`
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`
	DbHomeId        *string `mandatory:"true" json:"dbHomeId"`
	Source          string  `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdatabasebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedatabasebase createdatabasebase
	s := struct {
		Model Unmarshalercreatedatabasebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DbHomeId = s.Model.DbHomeId
	m.DbVersion = s.Model.DbVersion
	m.KmsKeyId = s.Model.KmsKeyId
	m.KmsKeyVersionId = s.Model.KmsKeyVersionId
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdatabasebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "NONE":
		mm := CreateNewDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_BACKUP":
		mm := CreateDatabaseFromBackup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateDatabaseBase: %s.", m.Source)
		return *m, nil
	}
}

// GetDbVersion returns DbVersion
func (m createdatabasebase) GetDbVersion() *string {
	return m.DbVersion
}

// GetKmsKeyId returns KmsKeyId
func (m createdatabasebase) GetKmsKeyId() *string {
	return m.KmsKeyId
}

// GetKmsKeyVersionId returns KmsKeyVersionId
func (m createdatabasebase) GetKmsKeyVersionId() *string {
	return m.KmsKeyVersionId
}

// GetDbHomeId returns DbHomeId
func (m createdatabasebase) GetDbHomeId() *string {
	return m.DbHomeId
}

func (m createdatabasebase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdatabasebase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseBaseSourceEnum Enum with underlying type: string
type CreateDatabaseBaseSourceEnum string

// Set of constants representing the allowable values for CreateDatabaseBaseSourceEnum
const (
	CreateDatabaseBaseSourceNone     CreateDatabaseBaseSourceEnum = "NONE"
	CreateDatabaseBaseSourceDbBackup CreateDatabaseBaseSourceEnum = "DB_BACKUP"
)

var mappingCreateDatabaseBaseSourceEnum = map[string]CreateDatabaseBaseSourceEnum{
	"NONE":      CreateDatabaseBaseSourceNone,
	"DB_BACKUP": CreateDatabaseBaseSourceDbBackup,
}

var mappingCreateDatabaseBaseSourceEnumLowerCase = map[string]CreateDatabaseBaseSourceEnum{
	"none":      CreateDatabaseBaseSourceNone,
	"db_backup": CreateDatabaseBaseSourceDbBackup,
}

// GetCreateDatabaseBaseSourceEnumValues Enumerates the set of values for CreateDatabaseBaseSourceEnum
func GetCreateDatabaseBaseSourceEnumValues() []CreateDatabaseBaseSourceEnum {
	values := make([]CreateDatabaseBaseSourceEnum, 0)
	for _, v := range mappingCreateDatabaseBaseSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDatabaseBaseSourceEnumStringValues Enumerates the set of values in String for CreateDatabaseBaseSourceEnum
func GetCreateDatabaseBaseSourceEnumStringValues() []string {
	return []string{
		"NONE",
		"DB_BACKUP",
	}
}

// GetMappingCreateDatabaseBaseSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDatabaseBaseSourceEnum(val string) (CreateDatabaseBaseSourceEnum, bool) {
	enum, ok := mappingCreateDatabaseBaseSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
