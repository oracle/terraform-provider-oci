// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ChannelTarget Details about the Channel target.
type ChannelTarget interface {
}

type channeltarget struct {
	JsonData   []byte
	TargetType string `json:"targetType"`
}

// UnmarshalJSON unmarshals json
func (m *channeltarget) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerchanneltarget channeltarget
	s := struct {
		Model Unmarshalerchanneltarget
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetType = s.Model.TargetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *channeltarget) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetType {
	case "DBSYSTEM":
		mm := ChannelTargetDbSystem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ChannelTarget: %s.", m.TargetType)
		return *m, nil
	}
}

func (m channeltarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m channeltarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChannelTargetTargetTypeEnum Enum with underlying type: string
type ChannelTargetTargetTypeEnum string

// Set of constants representing the allowable values for ChannelTargetTargetTypeEnum
const (
	ChannelTargetTargetTypeDbsystem ChannelTargetTargetTypeEnum = "DBSYSTEM"
)

var mappingChannelTargetTargetTypeEnum = map[string]ChannelTargetTargetTypeEnum{
	"DBSYSTEM": ChannelTargetTargetTypeDbsystem,
}

var mappingChannelTargetTargetTypeEnumLowerCase = map[string]ChannelTargetTargetTypeEnum{
	"dbsystem": ChannelTargetTargetTypeDbsystem,
}

// GetChannelTargetTargetTypeEnumValues Enumerates the set of values for ChannelTargetTargetTypeEnum
func GetChannelTargetTargetTypeEnumValues() []ChannelTargetTargetTypeEnum {
	values := make([]ChannelTargetTargetTypeEnum, 0)
	for _, v := range mappingChannelTargetTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelTargetTargetTypeEnumStringValues Enumerates the set of values in String for ChannelTargetTargetTypeEnum
func GetChannelTargetTargetTypeEnumStringValues() []string {
	return []string{
		"DBSYSTEM",
	}
}

// GetMappingChannelTargetTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelTargetTargetTypeEnum(val string) (ChannelTargetTargetTypeEnum, bool) {
	enum, ok := mappingChannelTargetTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
