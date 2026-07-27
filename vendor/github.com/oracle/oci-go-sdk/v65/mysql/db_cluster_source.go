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

// DbClusterSource Details of the source from which the shared-storage DB cluster was created.
type DbClusterSource interface {
}

type dbclustersource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *dbclustersource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdbclustersource dbclustersource
	s := struct {
		Model Unmarshalerdbclustersource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dbclustersource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "NONE":
		mm := DbClusterSourceFromNone{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMPORTURL":
		mm := DbClusterSourceFromImportUrl{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DBCLUSTER_BACKUP":
		mm := DbClusterSourceFromDbClusterBackup{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DbClusterSource: %s.", m.SourceType)
		return *m, nil
	}
}

func (m dbclustersource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dbclustersource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbClusterSourceSourceTypeEnum Enum with underlying type: string
type DbClusterSourceSourceTypeEnum string

// Set of constants representing the allowable values for DbClusterSourceSourceTypeEnum
const (
	DbClusterSourceSourceTypeNone            DbClusterSourceSourceTypeEnum = "NONE"
	DbClusterSourceSourceTypeDbclusterBackup DbClusterSourceSourceTypeEnum = "DBCLUSTER_BACKUP"
	DbClusterSourceSourceTypeImporturl       DbClusterSourceSourceTypeEnum = "IMPORTURL"
)

var mappingDbClusterSourceSourceTypeEnum = map[string]DbClusterSourceSourceTypeEnum{
	"NONE":             DbClusterSourceSourceTypeNone,
	"DBCLUSTER_BACKUP": DbClusterSourceSourceTypeDbclusterBackup,
	"IMPORTURL":        DbClusterSourceSourceTypeImporturl,
}

var mappingDbClusterSourceSourceTypeEnumLowerCase = map[string]DbClusterSourceSourceTypeEnum{
	"none":             DbClusterSourceSourceTypeNone,
	"dbcluster_backup": DbClusterSourceSourceTypeDbclusterBackup,
	"importurl":        DbClusterSourceSourceTypeImporturl,
}

// GetDbClusterSourceSourceTypeEnumValues Enumerates the set of values for DbClusterSourceSourceTypeEnum
func GetDbClusterSourceSourceTypeEnumValues() []DbClusterSourceSourceTypeEnum {
	values := make([]DbClusterSourceSourceTypeEnum, 0)
	for _, v := range mappingDbClusterSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbClusterSourceSourceTypeEnumStringValues Enumerates the set of values in String for DbClusterSourceSourceTypeEnum
func GetDbClusterSourceSourceTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"DBCLUSTER_BACKUP",
		"IMPORTURL",
	}
}

// GetMappingDbClusterSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbClusterSourceSourceTypeEnum(val string) (DbClusterSourceSourceTypeEnum, bool) {
	enum, ok := mappingDbClusterSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
