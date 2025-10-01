// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GiGoalVersionDetails Details of goal 'GI' software version.
type GiGoalVersionDetails interface {
}

type gigoalversiondetails struct {
	JsonData []byte
	GoalType string `json:"goalType"`
}

// UnmarshalJSON unmarshals json
func (m *gigoalversiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalergigoalversiondetails gigoalversiondetails
	s := struct {
		Model Unmarshalergigoalversiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.GoalType = s.Model.GoalType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *gigoalversiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.GoalType {
	case "GI_ORACLE_IMAGE":
		mm := OracleGiGoalVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GI_CUSTOM_IMAGE":
		mm := CustomGiGoalVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for GiGoalVersionDetails: %s.", m.GoalType)
		return *m, nil
	}
}

func (m gigoalversiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m gigoalversiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GiGoalVersionDetailsGoalTypeEnum Enum with underlying type: string
type GiGoalVersionDetailsGoalTypeEnum string

// Set of constants representing the allowable values for GiGoalVersionDetailsGoalTypeEnum
const (
	GiGoalVersionDetailsGoalTypeOracleImage GiGoalVersionDetailsGoalTypeEnum = "GI_ORACLE_IMAGE"
	GiGoalVersionDetailsGoalTypeCustomImage GiGoalVersionDetailsGoalTypeEnum = "GI_CUSTOM_IMAGE"
)

var mappingGiGoalVersionDetailsGoalTypeEnum = map[string]GiGoalVersionDetailsGoalTypeEnum{
	"GI_ORACLE_IMAGE": GiGoalVersionDetailsGoalTypeOracleImage,
	"GI_CUSTOM_IMAGE": GiGoalVersionDetailsGoalTypeCustomImage,
}

var mappingGiGoalVersionDetailsGoalTypeEnumLowerCase = map[string]GiGoalVersionDetailsGoalTypeEnum{
	"gi_oracle_image": GiGoalVersionDetailsGoalTypeOracleImage,
	"gi_custom_image": GiGoalVersionDetailsGoalTypeCustomImage,
}

// GetGiGoalVersionDetailsGoalTypeEnumValues Enumerates the set of values for GiGoalVersionDetailsGoalTypeEnum
func GetGiGoalVersionDetailsGoalTypeEnumValues() []GiGoalVersionDetailsGoalTypeEnum {
	values := make([]GiGoalVersionDetailsGoalTypeEnum, 0)
	for _, v := range mappingGiGoalVersionDetailsGoalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGiGoalVersionDetailsGoalTypeEnumStringValues Enumerates the set of values in String for GiGoalVersionDetailsGoalTypeEnum
func GetGiGoalVersionDetailsGoalTypeEnumStringValues() []string {
	return []string{
		"GI_ORACLE_IMAGE",
		"GI_CUSTOM_IMAGE",
	}
}

// GetMappingGiGoalVersionDetailsGoalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiGoalVersionDetailsGoalTypeEnum(val string) (GiGoalVersionDetailsGoalTypeEnum, bool) {
	enum, ok := mappingGiGoalVersionDetailsGoalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
