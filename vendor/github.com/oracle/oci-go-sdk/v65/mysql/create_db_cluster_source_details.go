// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDbClusterSourceDetails Details that specify how the initial data for the shared-storage DB cluster is provisioned.
type CreateDbClusterSourceDetails interface {
}

type createdbclustersourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *createdbclustersourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedbclustersourcedetails createdbclustersourcedetails
	s := struct {
		Model Unmarshalercreatedbclustersourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdbclustersourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "DBCLUSTER_BACKUP":
		mm := CreateDbClusterSourceFromDbClusterBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := CreateDbClusterSourceFromNoneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMPORTURL":
		mm := CreateDbClusterSourceFromImportUrlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateDbClusterSourceDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m createdbclustersourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdbclustersourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDbClusterSourceDetailsSourceTypeEnum Enum with underlying type: string
type CreateDbClusterSourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for CreateDbClusterSourceDetailsSourceTypeEnum
const (
	CreateDbClusterSourceDetailsSourceTypeNone            CreateDbClusterSourceDetailsSourceTypeEnum = "NONE"
	CreateDbClusterSourceDetailsSourceTypeDbclusterBackup CreateDbClusterSourceDetailsSourceTypeEnum = "DBCLUSTER_BACKUP"
	CreateDbClusterSourceDetailsSourceTypeImporturl       CreateDbClusterSourceDetailsSourceTypeEnum = "IMPORTURL"
)

var mappingCreateDbClusterSourceDetailsSourceTypeEnum = map[string]CreateDbClusterSourceDetailsSourceTypeEnum{
	"NONE":             CreateDbClusterSourceDetailsSourceTypeNone,
	"DBCLUSTER_BACKUP": CreateDbClusterSourceDetailsSourceTypeDbclusterBackup,
	"IMPORTURL":        CreateDbClusterSourceDetailsSourceTypeImporturl,
}

var mappingCreateDbClusterSourceDetailsSourceTypeEnumLowerCase = map[string]CreateDbClusterSourceDetailsSourceTypeEnum{
	"none":             CreateDbClusterSourceDetailsSourceTypeNone,
	"dbcluster_backup": CreateDbClusterSourceDetailsSourceTypeDbclusterBackup,
	"importurl":        CreateDbClusterSourceDetailsSourceTypeImporturl,
}

// GetCreateDbClusterSourceDetailsSourceTypeEnumValues Enumerates the set of values for CreateDbClusterSourceDetailsSourceTypeEnum
func GetCreateDbClusterSourceDetailsSourceTypeEnumValues() []CreateDbClusterSourceDetailsSourceTypeEnum {
	values := make([]CreateDbClusterSourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingCreateDbClusterSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateDbClusterSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for CreateDbClusterSourceDetailsSourceTypeEnum
func GetCreateDbClusterSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"DBCLUSTER_BACKUP",
		"IMPORTURL",
	}
}

// GetMappingCreateDbClusterSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateDbClusterSourceDetailsSourceTypeEnum(val string) (CreateDbClusterSourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingCreateDbClusterSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
