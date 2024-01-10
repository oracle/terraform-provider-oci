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

// CreatePluggableDatabaseCreationTypeDetails The Pluggable Database creation type.
// Use `LOCAL_CLONE_PDB` for creating a new PDB using Local Clone on Source Pluggable Database. This will Clone and starts a
// pluggable database (PDB) in the same database (CDB) as the source PDB. The source PDB must be in the `READ_WRITE` openMode to
// perform the clone operation.
// Use `REMOTE_CLONE_PDB` for creating a new PDB using Remote Clone on Source Pluggable Database. This will Clone a pluggable
// database (PDB) to a different database from the source PDB. The cloned PDB will be started upon completion of the clone
// operation. The source PDB must be in the `READ_WRITE` openMode when performing the clone.
// For Exadata Cloud@Customer instances, the source pluggable database (PDB) must be on the same Exadata Infrastructure as the
// target container database (CDB) to create a remote clone.
// Use `RELOCATE_PDB` for relocating the Pluggable Database from Source CDB and creating it in target CDB. This will relocate a
// pluggable database (PDB) to a different database from the source PDB. The source PDB must be in the `READ_WRITE` openMode when
// performing the relocate.
type CreatePluggableDatabaseCreationTypeDetails interface {
}

type createpluggabledatabasecreationtypedetails struct {
	JsonData     []byte
	CreationType string `json:"creationType"`
}

// UnmarshalJSON unmarshals json
func (m *createpluggabledatabasecreationtypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatepluggabledatabasecreationtypedetails createpluggabledatabasecreationtypedetails
	s := struct {
		Model Unmarshalercreatepluggabledatabasecreationtypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CreationType = s.Model.CreationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createpluggabledatabasecreationtypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CreationType {
	case "RELOCATE_PDB":
		mm := CreatePluggableDatabaseFromRelocateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REMOTE_CLONE_PDB":
		mm := CreatePluggableDatabaseFromRemoteCloneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOCAL_CLONE_PDB":
		mm := CreatePluggableDatabaseFromLocalCloneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreatePluggableDatabaseCreationTypeDetails: %s.", m.CreationType)
		return *m, nil
	}
}

func (m createpluggabledatabasecreationtypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createpluggabledatabasecreationtypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum Enum with underlying type: string
type CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum string

// Set of constants representing the allowable values for CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum
const (
	CreatePluggableDatabaseCreationTypeDetailsCreationTypeLocalClonePdb  CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum = "LOCAL_CLONE_PDB"
	CreatePluggableDatabaseCreationTypeDetailsCreationTypeRemoteClonePdb CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum = "REMOTE_CLONE_PDB"
	CreatePluggableDatabaseCreationTypeDetailsCreationTypeRelocatePdb    CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum = "RELOCATE_PDB"
)

var mappingCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum = map[string]CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum{
	"LOCAL_CLONE_PDB":  CreatePluggableDatabaseCreationTypeDetailsCreationTypeLocalClonePdb,
	"REMOTE_CLONE_PDB": CreatePluggableDatabaseCreationTypeDetailsCreationTypeRemoteClonePdb,
	"RELOCATE_PDB":     CreatePluggableDatabaseCreationTypeDetailsCreationTypeRelocatePdb,
}

var mappingCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnumLowerCase = map[string]CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum{
	"local_clone_pdb":  CreatePluggableDatabaseCreationTypeDetailsCreationTypeLocalClonePdb,
	"remote_clone_pdb": CreatePluggableDatabaseCreationTypeDetailsCreationTypeRemoteClonePdb,
	"relocate_pdb":     CreatePluggableDatabaseCreationTypeDetailsCreationTypeRelocatePdb,
}

// GetCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnumValues Enumerates the set of values for CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum
func GetCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnumValues() []CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum {
	values := make([]CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum, 0)
	for _, v := range mappingCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnumStringValues Enumerates the set of values in String for CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum
func GetCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnumStringValues() []string {
	return []string{
		"LOCAL_CLONE_PDB",
		"REMOTE_CLONE_PDB",
		"RELOCATE_PDB",
	}
}

// GetMappingCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum(val string) (CreatePluggableDatabaseCreationTypeDetailsCreationTypeEnum, bool) {
	enum, ok := mappingCreatePluggableDatabaseCreationTypeDetailsCreationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
