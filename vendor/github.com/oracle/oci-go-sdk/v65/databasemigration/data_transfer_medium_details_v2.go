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

// DataTransferMediumDetailsV2 Optional additional properties for dump transfer in source or target host. Default kind is CURL
type DataTransferMediumDetailsV2 interface {
}

type datatransfermediumdetailsv2 struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *datatransfermediumdetailsv2) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatatransfermediumdetailsv2 datatransfermediumdetailsv2
	s := struct {
		Model Unmarshalerdatatransfermediumdetailsv2
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datatransfermediumdetailsv2) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NFS":
		mm := NfsDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := ObjectStorageDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DBLINK":
		mm := DbLinkDataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWS_S3":
		mm := AwsS3DataTransferMediumDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataTransferMediumDetailsV2: %s.", m.Type)
		return *m, nil
	}
}

func (m datatransfermediumdetailsv2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m datatransfermediumdetailsv2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataTransferMediumDetailsV2TypeEnum Enum with underlying type: string
type DataTransferMediumDetailsV2TypeEnum string

// Set of constants representing the allowable values for DataTransferMediumDetailsV2TypeEnum
const (
	DataTransferMediumDetailsV2TypeDblink        DataTransferMediumDetailsV2TypeEnum = "DBLINK"
	DataTransferMediumDetailsV2TypeObjectStorage DataTransferMediumDetailsV2TypeEnum = "OBJECT_STORAGE"
	DataTransferMediumDetailsV2TypeAwsS3         DataTransferMediumDetailsV2TypeEnum = "AWS_S3"
	DataTransferMediumDetailsV2TypeNfs           DataTransferMediumDetailsV2TypeEnum = "NFS"
)

var mappingDataTransferMediumDetailsV2TypeEnum = map[string]DataTransferMediumDetailsV2TypeEnum{
	"DBLINK":         DataTransferMediumDetailsV2TypeDblink,
	"OBJECT_STORAGE": DataTransferMediumDetailsV2TypeObjectStorage,
	"AWS_S3":         DataTransferMediumDetailsV2TypeAwsS3,
	"NFS":            DataTransferMediumDetailsV2TypeNfs,
}

var mappingDataTransferMediumDetailsV2TypeEnumLowerCase = map[string]DataTransferMediumDetailsV2TypeEnum{
	"dblink":         DataTransferMediumDetailsV2TypeDblink,
	"object_storage": DataTransferMediumDetailsV2TypeObjectStorage,
	"aws_s3":         DataTransferMediumDetailsV2TypeAwsS3,
	"nfs":            DataTransferMediumDetailsV2TypeNfs,
}

// GetDataTransferMediumDetailsV2TypeEnumValues Enumerates the set of values for DataTransferMediumDetailsV2TypeEnum
func GetDataTransferMediumDetailsV2TypeEnumValues() []DataTransferMediumDetailsV2TypeEnum {
	values := make([]DataTransferMediumDetailsV2TypeEnum, 0)
	for _, v := range mappingDataTransferMediumDetailsV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataTransferMediumDetailsV2TypeEnumStringValues Enumerates the set of values in String for DataTransferMediumDetailsV2TypeEnum
func GetDataTransferMediumDetailsV2TypeEnumStringValues() []string {
	return []string{
		"DBLINK",
		"OBJECT_STORAGE",
		"AWS_S3",
		"NFS",
	}
}

// GetMappingDataTransferMediumDetailsV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataTransferMediumDetailsV2TypeEnum(val string) (DataTransferMediumDetailsV2TypeEnum, bool) {
	enum, ok := mappingDataTransferMediumDetailsV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
