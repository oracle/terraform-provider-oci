// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StorageMountConfigurationDetails The storage mount configuration details
type StorageMountConfigurationDetails interface {

	// The local directory name to be mounted
	GetDestinationDirectoryName() *string

	// The local path of the mounted directory, excluding directory name.
	GetDestinationPath() *string
}

type storagemountconfigurationdetails struct {
	JsonData                 []byte
	DestinationPath          *string `mandatory:"false" json:"destinationPath"`
	DestinationDirectoryName *string `mandatory:"true" json:"destinationDirectoryName"`
	StorageType              string  `json:"storageType"`
}

// UnmarshalJSON unmarshals json
func (m *storagemountconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstoragemountconfigurationdetails storagemountconfigurationdetails
	s := struct {
		Model Unmarshalerstoragemountconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DestinationDirectoryName = s.Model.DestinationDirectoryName
	m.DestinationPath = s.Model.DestinationPath
	m.StorageType = s.Model.StorageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *storagemountconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StorageType {
	case "FILE_STORAGE":
		mm := FileStorageMountConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE":
		mm := ObjectStorageMountConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StorageMountConfigurationDetails: %s.", m.StorageType)
		return *m, nil
	}
}

// GetDestinationPath returns DestinationPath
func (m storagemountconfigurationdetails) GetDestinationPath() *string {
	return m.DestinationPath
}

// GetDestinationDirectoryName returns DestinationDirectoryName
func (m storagemountconfigurationdetails) GetDestinationDirectoryName() *string {
	return m.DestinationDirectoryName
}

func (m storagemountconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m storagemountconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StorageMountConfigurationDetailsStorageTypeEnum Enum with underlying type: string
type StorageMountConfigurationDetailsStorageTypeEnum string

// Set of constants representing the allowable values for StorageMountConfigurationDetailsStorageTypeEnum
const (
	StorageMountConfigurationDetailsStorageTypeFileStorage   StorageMountConfigurationDetailsStorageTypeEnum = "FILE_STORAGE"
	StorageMountConfigurationDetailsStorageTypeObjectStorage StorageMountConfigurationDetailsStorageTypeEnum = "OBJECT_STORAGE"
)

var mappingStorageMountConfigurationDetailsStorageTypeEnum = map[string]StorageMountConfigurationDetailsStorageTypeEnum{
	"FILE_STORAGE":   StorageMountConfigurationDetailsStorageTypeFileStorage,
	"OBJECT_STORAGE": StorageMountConfigurationDetailsStorageTypeObjectStorage,
}

var mappingStorageMountConfigurationDetailsStorageTypeEnumLowerCase = map[string]StorageMountConfigurationDetailsStorageTypeEnum{
	"file_storage":   StorageMountConfigurationDetailsStorageTypeFileStorage,
	"object_storage": StorageMountConfigurationDetailsStorageTypeObjectStorage,
}

// GetStorageMountConfigurationDetailsStorageTypeEnumValues Enumerates the set of values for StorageMountConfigurationDetailsStorageTypeEnum
func GetStorageMountConfigurationDetailsStorageTypeEnumValues() []StorageMountConfigurationDetailsStorageTypeEnum {
	values := make([]StorageMountConfigurationDetailsStorageTypeEnum, 0)
	for _, v := range mappingStorageMountConfigurationDetailsStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStorageMountConfigurationDetailsStorageTypeEnumStringValues Enumerates the set of values in String for StorageMountConfigurationDetailsStorageTypeEnum
func GetStorageMountConfigurationDetailsStorageTypeEnumStringValues() []string {
	return []string{
		"FILE_STORAGE",
		"OBJECT_STORAGE",
	}
}

// GetMappingStorageMountConfigurationDetailsStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStorageMountConfigurationDetailsStorageTypeEnum(val string) (StorageMountConfigurationDetailsStorageTypeEnum, bool) {
	enum, ok := mappingStorageMountConfigurationDetailsStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
