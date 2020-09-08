// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
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

// SourceSourceTypeEnum Enum with underlying type: string
type SourceSourceTypeEnum string

// Set of constants representing the allowable values for SourceSourceTypeEnum
const (
	SourceSourceTypeOciservice SourceSourceTypeEnum = "OCISERVICE"
)

var mappingSourceSourceType = map[string]SourceSourceTypeEnum{
	"OCISERVICE": SourceSourceTypeOciservice,
}

// GetSourceSourceTypeEnumValues Enumerates the set of values for SourceSourceTypeEnum
func GetSourceSourceTypeEnumValues() []SourceSourceTypeEnum {
	values := make([]SourceSourceTypeEnum, 0)
	for _, v := range mappingSourceSourceType {
		values = append(values, v)
	}
	return values
}
