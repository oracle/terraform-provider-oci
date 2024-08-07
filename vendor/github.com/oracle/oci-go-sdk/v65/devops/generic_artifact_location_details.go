// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenericArtifactLocationDetails Location where artifact is uploaded for user access.
type GenericArtifactLocationDetails interface {
}

type genericartifactlocationdetails struct {
	JsonData    []byte
	StorageType string `json:"storageType"`
}

// UnmarshalJSON unmarshals json
func (m *genericartifactlocationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalergenericartifactlocationdetails genericartifactlocationdetails
	s := struct {
		Model Unmarshalergenericartifactlocationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.StorageType = s.Model.StorageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *genericartifactlocationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.StorageType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageGenericArtifactLocationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for GenericArtifactLocationDetails: %s.", m.StorageType)
		return *m, nil
	}
}

func (m genericartifactlocationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m genericartifactlocationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenericArtifactLocationDetailsStorageTypeEnum Enum with underlying type: string
type GenericArtifactLocationDetailsStorageTypeEnum string

// Set of constants representing the allowable values for GenericArtifactLocationDetailsStorageTypeEnum
const (
	GenericArtifactLocationDetailsStorageTypeObjectStorage GenericArtifactLocationDetailsStorageTypeEnum = "OBJECT_STORAGE"
)

var mappingGenericArtifactLocationDetailsStorageTypeEnum = map[string]GenericArtifactLocationDetailsStorageTypeEnum{
	"OBJECT_STORAGE": GenericArtifactLocationDetailsStorageTypeObjectStorage,
}

var mappingGenericArtifactLocationDetailsStorageTypeEnumLowerCase = map[string]GenericArtifactLocationDetailsStorageTypeEnum{
	"object_storage": GenericArtifactLocationDetailsStorageTypeObjectStorage,
}

// GetGenericArtifactLocationDetailsStorageTypeEnumValues Enumerates the set of values for GenericArtifactLocationDetailsStorageTypeEnum
func GetGenericArtifactLocationDetailsStorageTypeEnumValues() []GenericArtifactLocationDetailsStorageTypeEnum {
	values := make([]GenericArtifactLocationDetailsStorageTypeEnum, 0)
	for _, v := range mappingGenericArtifactLocationDetailsStorageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericArtifactLocationDetailsStorageTypeEnumStringValues Enumerates the set of values in String for GenericArtifactLocationDetailsStorageTypeEnum
func GetGenericArtifactLocationDetailsStorageTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingGenericArtifactLocationDetailsStorageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericArtifactLocationDetailsStorageTypeEnum(val string) (GenericArtifactLocationDetailsStorageTypeEnum, bool) {
	enum, ok := mappingGenericArtifactLocationDetailsStorageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
