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

// OracleDataTransferMediumDetails Optional additional properties for data transfer.
type OracleDataTransferMediumDetails interface {
}

type oracledatatransfermediumdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *oracledatatransfermediumdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroracledatatransfermediumdetails oracledatatransfermediumdetails
	s := struct {
		Model Unmarshaleroracledatatransfermediumdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *oracledatatransfermediumdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DBLINK":
		mm := OracleDbLinkDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := OracleObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWS_S3":
		mm := OracleAwsS3DataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NFS":
		mm := OracleNfsDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OracleDataTransferMediumDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m oracledatatransfermediumdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m oracledatatransfermediumdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OracleDataTransferMediumDetailsTypeEnum Enum with underlying type: string
type OracleDataTransferMediumDetailsTypeEnum string

// Set of constants representing the allowable values for OracleDataTransferMediumDetailsTypeEnum
const (
	OracleDataTransferMediumDetailsTypeDblink        OracleDataTransferMediumDetailsTypeEnum = "DBLINK"
	OracleDataTransferMediumDetailsTypeObjectStorage OracleDataTransferMediumDetailsTypeEnum = "OBJECT_STORAGE"
	OracleDataTransferMediumDetailsTypeAwsS3         OracleDataTransferMediumDetailsTypeEnum = "AWS_S3"
	OracleDataTransferMediumDetailsTypeNfs           OracleDataTransferMediumDetailsTypeEnum = "NFS"
)

var mappingOracleDataTransferMediumDetailsTypeEnum = map[string]OracleDataTransferMediumDetailsTypeEnum{
	"DBLINK":         OracleDataTransferMediumDetailsTypeDblink,
	"OBJECT_STORAGE": OracleDataTransferMediumDetailsTypeObjectStorage,
	"AWS_S3":         OracleDataTransferMediumDetailsTypeAwsS3,
	"NFS":            OracleDataTransferMediumDetailsTypeNfs,
}

var mappingOracleDataTransferMediumDetailsTypeEnumLowerCase = map[string]OracleDataTransferMediumDetailsTypeEnum{
	"dblink":         OracleDataTransferMediumDetailsTypeDblink,
	"object_storage": OracleDataTransferMediumDetailsTypeObjectStorage,
	"aws_s3":         OracleDataTransferMediumDetailsTypeAwsS3,
	"nfs":            OracleDataTransferMediumDetailsTypeNfs,
}

// GetOracleDataTransferMediumDetailsTypeEnumValues Enumerates the set of values for OracleDataTransferMediumDetailsTypeEnum
func GetOracleDataTransferMediumDetailsTypeEnumValues() []OracleDataTransferMediumDetailsTypeEnum {
	values := make([]OracleDataTransferMediumDetailsTypeEnum, 0)
	for _, v := range mappingOracleDataTransferMediumDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOracleDataTransferMediumDetailsTypeEnumStringValues Enumerates the set of values in String for OracleDataTransferMediumDetailsTypeEnum
func GetOracleDataTransferMediumDetailsTypeEnumStringValues() []string {
	return []string{
		"DBLINK",
		"OBJECT_STORAGE",
		"AWS_S3",
		"NFS",
	}
}

// GetMappingOracleDataTransferMediumDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOracleDataTransferMediumDetailsTypeEnum(val string) (OracleDataTransferMediumDetailsTypeEnum, bool) {
	enum, ok := mappingOracleDataTransferMediumDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
