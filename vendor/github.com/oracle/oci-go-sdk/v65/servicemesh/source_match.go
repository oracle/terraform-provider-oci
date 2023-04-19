// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceMatch Rate limits to be applied based on source type.
type SourceMatch interface {
}

type sourcematch struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *sourcematch) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcematch sourcematch
	s := struct {
		Model Unmarshalersourcematch
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcematch) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "ALL":
		mm := AllSourceMatch{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SourceMatch: %s.", m.SourceType)
		return *m, nil
	}
}

func (m sourcematch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourcematch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceMatchSourceTypeEnum Enum with underlying type: string
type SourceMatchSourceTypeEnum string

// Set of constants representing the allowable values for SourceMatchSourceTypeEnum
const (
	SourceMatchSourceTypeAll SourceMatchSourceTypeEnum = "ALL"
)

var mappingSourceMatchSourceTypeEnum = map[string]SourceMatchSourceTypeEnum{
	"ALL": SourceMatchSourceTypeAll,
}

var mappingSourceMatchSourceTypeEnumLowerCase = map[string]SourceMatchSourceTypeEnum{
	"all": SourceMatchSourceTypeAll,
}

// GetSourceMatchSourceTypeEnumValues Enumerates the set of values for SourceMatchSourceTypeEnum
func GetSourceMatchSourceTypeEnumValues() []SourceMatchSourceTypeEnum {
	values := make([]SourceMatchSourceTypeEnum, 0)
	for _, v := range mappingSourceMatchSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceMatchSourceTypeEnumStringValues Enumerates the set of values in String for SourceMatchSourceTypeEnum
func GetSourceMatchSourceTypeEnumStringValues() []string {
	return []string{
		"ALL",
	}
}

// GetMappingSourceMatchSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceMatchSourceTypeEnum(val string) (SourceMatchSourceTypeEnum, bool) {
	enum, ok := mappingSourceMatchSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
