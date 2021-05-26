// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
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

// ChannelSourceSourceTypeEnum Enum with underlying type: string
type ChannelSourceSourceTypeEnum string

// Set of constants representing the allowable values for ChannelSourceSourceTypeEnum
const (
	ChannelSourceSourceTypeMysql ChannelSourceSourceTypeEnum = "MYSQL"
)

var mappingChannelSourceSourceType = map[string]ChannelSourceSourceTypeEnum{
	"MYSQL": ChannelSourceSourceTypeMysql,
}

// GetChannelSourceSourceTypeEnumValues Enumerates the set of values for ChannelSourceSourceTypeEnum
func GetChannelSourceSourceTypeEnumValues() []ChannelSourceSourceTypeEnum {
	values := make([]ChannelSourceSourceTypeEnum, 0)
	for _, v := range mappingChannelSourceSourceType {
		values = append(values, v)
	}
	return values
}
