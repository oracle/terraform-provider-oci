// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceDetails The source used to restore the database system.
type SourceDetails interface {
}

type sourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *sourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcedetails sourcedetails
	s := struct {
		Model Unmarshalersourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "BACKUP":
		mm := BackupSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SourceDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m sourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceDetailsSourceTypeEnum Enum with underlying type: string
type SourceDetailsSourceTypeEnum string

// Set of constants representing the allowable values for SourceDetailsSourceTypeEnum
const (
	SourceDetailsSourceTypeBackup SourceDetailsSourceTypeEnum = "BACKUP"
	SourceDetailsSourceTypeNone   SourceDetailsSourceTypeEnum = "NONE"
)

var mappingSourceDetailsSourceTypeEnum = map[string]SourceDetailsSourceTypeEnum{
	"BACKUP": SourceDetailsSourceTypeBackup,
	"NONE":   SourceDetailsSourceTypeNone,
}

var mappingSourceDetailsSourceTypeEnumLowerCase = map[string]SourceDetailsSourceTypeEnum{
	"backup": SourceDetailsSourceTypeBackup,
	"none":   SourceDetailsSourceTypeNone,
}

// GetSourceDetailsSourceTypeEnumValues Enumerates the set of values for SourceDetailsSourceTypeEnum
func GetSourceDetailsSourceTypeEnumValues() []SourceDetailsSourceTypeEnum {
	values := make([]SourceDetailsSourceTypeEnum, 0)
	for _, v := range mappingSourceDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceDetailsSourceTypeEnumStringValues Enumerates the set of values in String for SourceDetailsSourceTypeEnum
func GetSourceDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"BACKUP",
		"NONE",
	}
}

// GetMappingSourceDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceDetailsSourceTypeEnum(val string) (SourceDetailsSourceTypeEnum, bool) {
	enum, ok := mappingSourceDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
