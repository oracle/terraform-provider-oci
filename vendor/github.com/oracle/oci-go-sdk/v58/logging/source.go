// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Source The source the log object comes from.
type Source interface {
}

type source struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *source) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersource source
	s := struct {
		Model Unmarshalersource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *source) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OCISERVICE":
		mm := OciService{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m source) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m source) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceSourceTypeEnum Enum with underlying type: string
type SourceSourceTypeEnum string

// Set of constants representing the allowable values for SourceSourceTypeEnum
const (
	SourceSourceTypeOciservice SourceSourceTypeEnum = "OCISERVICE"
)

var mappingSourceSourceTypeEnum = map[string]SourceSourceTypeEnum{
	"OCISERVICE": SourceSourceTypeOciservice,
}

// GetSourceSourceTypeEnumValues Enumerates the set of values for SourceSourceTypeEnum
func GetSourceSourceTypeEnumValues() []SourceSourceTypeEnum {
	values := make([]SourceSourceTypeEnum, 0)
	for _, v := range mappingSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceSourceTypeEnumStringValues Enumerates the set of values in String for SourceSourceTypeEnum
func GetSourceSourceTypeEnumStringValues() []string {
	return []string{
		"OCISERVICE",
	}
}

// GetMappingSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceSourceTypeEnum(val string) (SourceSourceTypeEnum, bool) {
	mappingSourceSourceTypeEnumIgnoreCase := make(map[string]SourceSourceTypeEnum)
	for k, v := range mappingSourceSourceTypeEnum {
		mappingSourceSourceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSourceSourceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
