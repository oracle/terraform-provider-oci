// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChannelSource Parameters detailing how to provision the source for the given Channel.
type ChannelSource interface {
}

type channelsource struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *channelsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerchannelsource channelsource
	s := struct {
		Model Unmarshalerchannelsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *channelsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "MYSQL":
		mm := ChannelSourceMysql{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m channelsource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m channelsource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChannelSourceSourceTypeEnum Enum with underlying type: string
type ChannelSourceSourceTypeEnum string

// Set of constants representing the allowable values for ChannelSourceSourceTypeEnum
const (
	ChannelSourceSourceTypeMysql ChannelSourceSourceTypeEnum = "MYSQL"
)

var mappingChannelSourceSourceTypeEnum = map[string]ChannelSourceSourceTypeEnum{
	"MYSQL": ChannelSourceSourceTypeMysql,
}

var mappingChannelSourceSourceTypeEnumLowerCase = map[string]ChannelSourceSourceTypeEnum{
	"mysql": ChannelSourceSourceTypeMysql,
}

// GetChannelSourceSourceTypeEnumValues Enumerates the set of values for ChannelSourceSourceTypeEnum
func GetChannelSourceSourceTypeEnumValues() []ChannelSourceSourceTypeEnum {
	values := make([]ChannelSourceSourceTypeEnum, 0)
	for _, v := range mappingChannelSourceSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelSourceSourceTypeEnumStringValues Enumerates the set of values in String for ChannelSourceSourceTypeEnum
func GetChannelSourceSourceTypeEnumStringValues() []string {
	return []string{
		"MYSQL",
	}
}

// GetMappingChannelSourceSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelSourceSourceTypeEnum(val string) (ChannelSourceSourceTypeEnum, bool) {
	enum, ok := mappingChannelSourceSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
