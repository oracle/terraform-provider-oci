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

// GuestOsGoalVersionDetails Details of goal 'GUEST_OS' software version.
type GuestOsGoalVersionDetails interface {
}

type guestosgoalversiondetails struct {
	JsonData []byte
	GoalType string `json:"goalType"`
}

// UnmarshalJSON unmarshals json
func (m *guestosgoalversiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerguestosgoalversiondetails guestosgoalversiondetails
	s := struct {
		Model Unmarshalerguestosgoalversiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.GoalType = s.Model.GoalType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *guestosgoalversiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.GoalType {
	case "GUEST_OS_ORACLE_IMAGE":
		mm := OracleGuestOsGoalVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for GuestOsGoalVersionDetails: %s.", m.GoalType)
		return *m, nil
	}
}

func (m guestosgoalversiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m guestosgoalversiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GuestOsGoalVersionDetailsGoalTypeEnum Enum with underlying type: string
type GuestOsGoalVersionDetailsGoalTypeEnum string

// Set of constants representing the allowable values for GuestOsGoalVersionDetailsGoalTypeEnum
const (
	GuestOsGoalVersionDetailsGoalTypeGuestOsOracleImage GuestOsGoalVersionDetailsGoalTypeEnum = "GUEST_OS_ORACLE_IMAGE"
)

var mappingGuestOsGoalVersionDetailsGoalTypeEnum = map[string]GuestOsGoalVersionDetailsGoalTypeEnum{
	"GUEST_OS_ORACLE_IMAGE": GuestOsGoalVersionDetailsGoalTypeGuestOsOracleImage,
}

var mappingGuestOsGoalVersionDetailsGoalTypeEnumLowerCase = map[string]GuestOsGoalVersionDetailsGoalTypeEnum{
	"guest_os_oracle_image": GuestOsGoalVersionDetailsGoalTypeGuestOsOracleImage,
}

// GetGuestOsGoalVersionDetailsGoalTypeEnumValues Enumerates the set of values for GuestOsGoalVersionDetailsGoalTypeEnum
func GetGuestOsGoalVersionDetailsGoalTypeEnumValues() []GuestOsGoalVersionDetailsGoalTypeEnum {
	values := make([]GuestOsGoalVersionDetailsGoalTypeEnum, 0)
	for _, v := range mappingGuestOsGoalVersionDetailsGoalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGuestOsGoalVersionDetailsGoalTypeEnumStringValues Enumerates the set of values in String for GuestOsGoalVersionDetailsGoalTypeEnum
func GetGuestOsGoalVersionDetailsGoalTypeEnumStringValues() []string {
	return []string{
		"GUEST_OS_ORACLE_IMAGE",
	}
}

// GetMappingGuestOsGoalVersionDetailsGoalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGuestOsGoalVersionDetailsGoalTypeEnum(val string) (GuestOsGoalVersionDetailsGoalTypeEnum, bool) {
	enum, ok := mappingGuestOsGoalVersionDetailsGoalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
