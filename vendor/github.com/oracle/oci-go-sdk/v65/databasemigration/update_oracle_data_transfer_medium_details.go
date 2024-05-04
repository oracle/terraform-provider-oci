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

// UpdateOracleDataTransferMediumDetails Optional additional properties for data transfer.
type UpdateOracleDataTransferMediumDetails interface {
}

type updateoracledatatransfermediumdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updateoracledatatransfermediumdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateoracledatatransfermediumdetails updateoracledatatransfermediumdetails
	s := struct {
		Model Unmarshalerupdateoracledatatransfermediumdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateoracledatatransfermediumdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "AWS_S3":
		mm := UpdateOracleAwsS3DataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NFS":
		mm := UpdateOracleNfsDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := UpdateOracleObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DBLINK":
		mm := UpdateOracleDbLinkDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateOracleDataTransferMediumDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m updateoracledatatransfermediumdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateoracledatatransfermediumdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOracleDataTransferMediumDetailsTypeEnum Enum with underlying type: string
type UpdateOracleDataTransferMediumDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateOracleDataTransferMediumDetailsTypeEnum
const (
	UpdateOracleDataTransferMediumDetailsTypeDblink        UpdateOracleDataTransferMediumDetailsTypeEnum = "DBLINK"
	UpdateOracleDataTransferMediumDetailsTypeObjectStorage UpdateOracleDataTransferMediumDetailsTypeEnum = "OBJECT_STORAGE"
	UpdateOracleDataTransferMediumDetailsTypeAwsS3         UpdateOracleDataTransferMediumDetailsTypeEnum = "AWS_S3"
	UpdateOracleDataTransferMediumDetailsTypeNfs           UpdateOracleDataTransferMediumDetailsTypeEnum = "NFS"
)

var mappingUpdateOracleDataTransferMediumDetailsTypeEnum = map[string]UpdateOracleDataTransferMediumDetailsTypeEnum{
	"DBLINK":         UpdateOracleDataTransferMediumDetailsTypeDblink,
	"OBJECT_STORAGE": UpdateOracleDataTransferMediumDetailsTypeObjectStorage,
	"AWS_S3":         UpdateOracleDataTransferMediumDetailsTypeAwsS3,
	"NFS":            UpdateOracleDataTransferMediumDetailsTypeNfs,
}

var mappingUpdateOracleDataTransferMediumDetailsTypeEnumLowerCase = map[string]UpdateOracleDataTransferMediumDetailsTypeEnum{
	"dblink":         UpdateOracleDataTransferMediumDetailsTypeDblink,
	"object_storage": UpdateOracleDataTransferMediumDetailsTypeObjectStorage,
	"aws_s3":         UpdateOracleDataTransferMediumDetailsTypeAwsS3,
	"nfs":            UpdateOracleDataTransferMediumDetailsTypeNfs,
}

// GetUpdateOracleDataTransferMediumDetailsTypeEnumValues Enumerates the set of values for UpdateOracleDataTransferMediumDetailsTypeEnum
func GetUpdateOracleDataTransferMediumDetailsTypeEnumValues() []UpdateOracleDataTransferMediumDetailsTypeEnum {
	values := make([]UpdateOracleDataTransferMediumDetailsTypeEnum, 0)
	for _, v := range mappingUpdateOracleDataTransferMediumDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOracleDataTransferMediumDetailsTypeEnumStringValues Enumerates the set of values in String for UpdateOracleDataTransferMediumDetailsTypeEnum
func GetUpdateOracleDataTransferMediumDetailsTypeEnumStringValues() []string {
	return []string{
		"DBLINK",
		"OBJECT_STORAGE",
		"AWS_S3",
		"NFS",
	}
}

// GetMappingUpdateOracleDataTransferMediumDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOracleDataTransferMediumDetailsTypeEnum(val string) (UpdateOracleDataTransferMediumDetailsTypeEnum, bool) {
	enum, ok := mappingUpdateOracleDataTransferMediumDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
