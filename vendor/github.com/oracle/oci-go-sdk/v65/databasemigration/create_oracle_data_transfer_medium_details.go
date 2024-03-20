// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateOracleDataTransferMediumDetails Optional additional properties for data transfer.
type CreateOracleDataTransferMediumDetails interface {
}

type createoracledatatransfermediumdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createoracledatatransfermediumdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateoracledatatransfermediumdetails createoracledatatransfermediumdetails
	s := struct {
		Model Unmarshalercreateoracledatatransfermediumdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createoracledatatransfermediumdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DBLINK":
		mm := CreateOracleDbLinkDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NFS":
		mm := CreateOracleNfsDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := CreateOracleObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWS_S3":
		mm := CreateOracleAwsS3DataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateOracleDataTransferMediumDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createoracledatatransfermediumdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createoracledatatransfermediumdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOracleDataTransferMediumDetailsTypeEnum Enum with underlying type: string
type CreateOracleDataTransferMediumDetailsTypeEnum string

// Set of constants representing the allowable values for CreateOracleDataTransferMediumDetailsTypeEnum
const (
	CreateOracleDataTransferMediumDetailsTypeDblink        CreateOracleDataTransferMediumDetailsTypeEnum = "DBLINK"
	CreateOracleDataTransferMediumDetailsTypeObjectStorage CreateOracleDataTransferMediumDetailsTypeEnum = "OBJECT_STORAGE"
	CreateOracleDataTransferMediumDetailsTypeAwsS3         CreateOracleDataTransferMediumDetailsTypeEnum = "AWS_S3"
	CreateOracleDataTransferMediumDetailsTypeNfs           CreateOracleDataTransferMediumDetailsTypeEnum = "NFS"
)

var mappingCreateOracleDataTransferMediumDetailsTypeEnum = map[string]CreateOracleDataTransferMediumDetailsTypeEnum{
	"DBLINK":         CreateOracleDataTransferMediumDetailsTypeDblink,
	"OBJECT_STORAGE": CreateOracleDataTransferMediumDetailsTypeObjectStorage,
	"AWS_S3":         CreateOracleDataTransferMediumDetailsTypeAwsS3,
	"NFS":            CreateOracleDataTransferMediumDetailsTypeNfs,
}

var mappingCreateOracleDataTransferMediumDetailsTypeEnumLowerCase = map[string]CreateOracleDataTransferMediumDetailsTypeEnum{
	"dblink":         CreateOracleDataTransferMediumDetailsTypeDblink,
	"object_storage": CreateOracleDataTransferMediumDetailsTypeObjectStorage,
	"aws_s3":         CreateOracleDataTransferMediumDetailsTypeAwsS3,
	"nfs":            CreateOracleDataTransferMediumDetailsTypeNfs,
}

// GetCreateOracleDataTransferMediumDetailsTypeEnumValues Enumerates the set of values for CreateOracleDataTransferMediumDetailsTypeEnum
func GetCreateOracleDataTransferMediumDetailsTypeEnumValues() []CreateOracleDataTransferMediumDetailsTypeEnum {
	values := make([]CreateOracleDataTransferMediumDetailsTypeEnum, 0)
	for _, v := range mappingCreateOracleDataTransferMediumDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOracleDataTransferMediumDetailsTypeEnumStringValues Enumerates the set of values in String for CreateOracleDataTransferMediumDetailsTypeEnum
func GetCreateOracleDataTransferMediumDetailsTypeEnumStringValues() []string {
	return []string{
		"DBLINK",
		"OBJECT_STORAGE",
		"AWS_S3",
		"NFS",
	}
}

// GetMappingCreateOracleDataTransferMediumDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOracleDataTransferMediumDetailsTypeEnum(val string) (CreateOracleDataTransferMediumDetailsTypeEnum, bool) {
	enum, ok := mappingCreateOracleDataTransferMediumDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
