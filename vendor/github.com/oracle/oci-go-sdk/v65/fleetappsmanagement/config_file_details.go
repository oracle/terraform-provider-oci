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

// ConfigFileDetails Content Source details.
type ConfigFileDetails interface {
}

type configfiledetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *configfiledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigfiledetails configfiledetails
	s := struct {
		Model Unmarshalerconfigfiledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configfiledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OBJECT_STORAGE_BUCKET":
		mm := ObjectStorageBucketConfigFileDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ConfigFileDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m configfiledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configfiledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigFileDetailsSourceTypeEnum Enum with underlying type: string
type ConfigFileDetailsSourceTypeEnum string

// Set of constants representing the allowable values for ConfigFileDetailsSourceTypeEnum
const (
	ConfigFileDetailsSourceTypeObjectStorageBucket ConfigFileDetailsSourceTypeEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingConfigFileDetailsSourceTypeEnum = map[string]ConfigFileDetailsSourceTypeEnum{
	"OBJECT_STORAGE_BUCKET": ConfigFileDetailsSourceTypeObjectStorageBucket,
}

var mappingConfigFileDetailsSourceTypeEnumLowerCase = map[string]ConfigFileDetailsSourceTypeEnum{
	"object_storage_bucket": ConfigFileDetailsSourceTypeObjectStorageBucket,
}

// GetConfigFileDetailsSourceTypeEnumValues Enumerates the set of values for ConfigFileDetailsSourceTypeEnum
func GetConfigFileDetailsSourceTypeEnumValues() []ConfigFileDetailsSourceTypeEnum {
	values := make([]ConfigFileDetailsSourceTypeEnum, 0)
	for _, v := range mappingConfigFileDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigFileDetailsSourceTypeEnumStringValues Enumerates the set of values in String for ConfigFileDetailsSourceTypeEnum
func GetConfigFileDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingConfigFileDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigFileDetailsSourceTypeEnum(val string) (ConfigFileDetailsSourceTypeEnum, bool) {
	enum, ok := mappingConfigFileDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
