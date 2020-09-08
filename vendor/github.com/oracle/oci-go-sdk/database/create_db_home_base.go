// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbHomeBase Details for creating a Database Home.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateDbHomeBase interface {

	// The user-provided name of the Database Home.
	GetDisplayName() *string

	// The database software image OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)
	GetDatabaseSoftwareImageId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}
}

type createdbhomebase struct {
	JsonData                []byte
	DisplayName             *string                           `mandatory:"false" json:"displayName"`
	DatabaseSoftwareImageId *string                           `mandatory:"false" json:"databaseSoftwareImageId"`
	FreeformTags            map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags             map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Source                  string                            `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createdbhomebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedbhomebase createdbhomebase
	s := struct {
		Model Unmarshalercreatedbhomebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.DatabaseSoftwareImageId = s.Model.DatabaseSoftwareImageId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdbhomebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "DATABASE":
		mm := CreateDbHomeWithDbSystemIdFromDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_BACKUP":
		mm := CreateDbHomeWithDbSystemIdFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := CreateDbHomeWithDbSystemIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VM_CLUSTER_NEW":
		mm := CreateDbHomeWithVmClusterIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createdbhomebase) GetDisplayName() *string {
	return m.DisplayName
}

//GetDatabaseSoftwareImageId returns DatabaseSoftwareImageId
func (m createdbhomebase) GetDatabaseSoftwareImageId() *string {
	return m.DatabaseSoftwareImageId
}

//GetFreeformTags returns FreeformTags
func (m createdbhomebase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createdbhomebase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createdbhomebase) String() string {
	return common.PointerString(m)
}

// CreateDbHomeBaseSourceEnum Enum with underlying type: string
type CreateDbHomeBaseSourceEnum string

// Set of constants representing the allowable values for CreateDbHomeBaseSourceEnum
const (
	CreateDbHomeBaseSourceNone         CreateDbHomeBaseSourceEnum = "NONE"
	CreateDbHomeBaseSourceDbBackup     CreateDbHomeBaseSourceEnum = "DB_BACKUP"
	CreateDbHomeBaseSourceDatabase     CreateDbHomeBaseSourceEnum = "DATABASE"
	CreateDbHomeBaseSourceVmClusterNew CreateDbHomeBaseSourceEnum = "VM_CLUSTER_NEW"
)

var mappingCreateDbHomeBaseSource = map[string]CreateDbHomeBaseSourceEnum{
	"NONE":           CreateDbHomeBaseSourceNone,
	"DB_BACKUP":      CreateDbHomeBaseSourceDbBackup,
	"DATABASE":       CreateDbHomeBaseSourceDatabase,
	"VM_CLUSTER_NEW": CreateDbHomeBaseSourceVmClusterNew,
}

// GetCreateDbHomeBaseSourceEnumValues Enumerates the set of values for CreateDbHomeBaseSourceEnum
func GetCreateDbHomeBaseSourceEnumValues() []CreateDbHomeBaseSourceEnum {
	values := make([]CreateDbHomeBaseSourceEnum, 0)
	for _, v := range mappingCreateDbHomeBaseSource {
		values = append(values, v)
	}
	return values
}
