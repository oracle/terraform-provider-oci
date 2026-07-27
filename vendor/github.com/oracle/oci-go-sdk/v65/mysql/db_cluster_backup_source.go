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

// DbClusterBackupSource Details of the source used to create the shared-storage DB cluster backup.
type DbClusterBackupSource interface {
}

type dbclusterbackupsource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *dbclusterbackupsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdbclusterbackupsource dbclusterbackupsource
	s := struct {
		Model Unmarshalerdbclusterbackupsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dbclusterbackupsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "DBCLUSTER":
		mm := DbClusterBackupSourceFromDbCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DbClusterBackupSource: %s.", m.SourceType)
		return *m, nil
	}
}

func (m dbclusterbackupsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dbclusterbackupsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbClusterBackupSourceSourceTypeEnum Enum with underlying type: string
type DbClusterBackupSourceSourceTypeEnum string

// Set of constants representing the allowable values for DbClusterBackupSourceSourceTypeEnum
const (
	DbClusterBackupSourceSourceTypeDbcluster DbClusterBackupSourceSourceTypeEnum = "DBCLUSTER"
)

var mappingDbClusterBackupSourceSourceTypeEnum = map[string]DbClusterBackupSourceSourceTypeEnum{
	"DBCLUSTER": DbClusterBackupSourceSourceTypeDbcluster,
}

var mappingDbClusterBackupSourceSourceTypeEnumLowerCase = map[string]DbClusterBackupSourceSourceTypeEnum{
	"dbcluster": DbClusterBackupSourceSourceTypeDbcluster,
}

// GetDbClusterBackupSourceSourceTypeEnumValues Enumerates the set of values for DbClusterBackupSourceSourceTypeEnum
func GetDbClusterBackupSourceSourceTypeEnumValues() []DbClusterBackupSourceSourceTypeEnum {
	values := make([]DbClusterBackupSourceSourceTypeEnum, 0)
	for _, v := range mappingDbClusterBackupSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbClusterBackupSourceSourceTypeEnumStringValues Enumerates the set of values in String for DbClusterBackupSourceSourceTypeEnum
func GetDbClusterBackupSourceSourceTypeEnumStringValues() []string {
	return []string{
		"DBCLUSTER",
	}
}

// GetMappingDbClusterBackupSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbClusterBackupSourceSourceTypeEnum(val string) (DbClusterBackupSourceSourceTypeEnum, bool) {
	enum, ok := mappingDbClusterBackupSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
