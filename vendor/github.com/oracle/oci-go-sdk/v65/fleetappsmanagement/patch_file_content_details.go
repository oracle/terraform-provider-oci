// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchFileContentDetails Content Source details.
type PatchFileContentDetails interface {
}

type patchfilecontentdetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *patchfilecontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpatchfilecontentdetails patchfilecontentdetails
	s := struct {
		Model Unmarshalerpatchfilecontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *patchfilecontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_BUCKET":
		mm := PatchFileObjectStorageBucketContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PatchFileContentDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m patchfilecontentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m patchfilecontentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchFileContentDetailsSourceTypeEnum Enum with underlying type: string
type PatchFileContentDetailsSourceTypeEnum string

// Set of constants representing the allowable values for PatchFileContentDetailsSourceTypeEnum
const (
	PatchFileContentDetailsSourceTypeObjectStorageBucket PatchFileContentDetailsSourceTypeEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingPatchFileContentDetailsSourceTypeEnum = map[string]PatchFileContentDetailsSourceTypeEnum{
	"OBJECT_STORAGE_BUCKET": PatchFileContentDetailsSourceTypeObjectStorageBucket,
}

var mappingPatchFileContentDetailsSourceTypeEnumLowerCase = map[string]PatchFileContentDetailsSourceTypeEnum{
	"object_storage_bucket": PatchFileContentDetailsSourceTypeObjectStorageBucket,
}

// GetPatchFileContentDetailsSourceTypeEnumValues Enumerates the set of values for PatchFileContentDetailsSourceTypeEnum
func GetPatchFileContentDetailsSourceTypeEnumValues() []PatchFileContentDetailsSourceTypeEnum {
	values := make([]PatchFileContentDetailsSourceTypeEnum, 0)
	for _, v := range mappingPatchFileContentDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchFileContentDetailsSourceTypeEnumStringValues Enumerates the set of values in String for PatchFileContentDetailsSourceTypeEnum
func GetPatchFileContentDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingPatchFileContentDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchFileContentDetailsSourceTypeEnum(val string) (PatchFileContentDetailsSourceTypeEnum, bool) {
	enum, ok := mappingPatchFileContentDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
