// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetSelected Specification of how targets are to be selected (select ALL, or select by TargetResourceType or TargetId).
type TargetSelected interface {
}

type targetselected struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *targetselected) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetselected targetselected
	s := struct {
		Model Unmarshalertargetselected
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetselected) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "ALL":
		mm := AllTargetsSelected{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGETTYPES":
		mm := TargetResourceTypesSelected{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGETIDS":
		mm := TargetIdsSelected{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetSelected: %s.", m.Kind)
		return *m, nil
	}
}

func (m targetselected) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetselected) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetSelectedKindEnum Enum with underlying type: string
type TargetSelectedKindEnum string

// Set of constants representing the allowable values for TargetSelectedKindEnum
const (
	TargetSelectedKindAll         TargetSelectedKindEnum = "ALL"
	TargetSelectedKindTargettypes TargetSelectedKindEnum = "TARGETTYPES"
	TargetSelectedKindTargetids   TargetSelectedKindEnum = "TARGETIDS"
)

var mappingTargetSelectedKindEnum = map[string]TargetSelectedKindEnum{
	"ALL":         TargetSelectedKindAll,
	"TARGETTYPES": TargetSelectedKindTargettypes,
	"TARGETIDS":   TargetSelectedKindTargetids,
}

var mappingTargetSelectedKindEnumLowerCase = map[string]TargetSelectedKindEnum{
	"all":         TargetSelectedKindAll,
	"targettypes": TargetSelectedKindTargettypes,
	"targetids":   TargetSelectedKindTargetids,
}

// GetTargetSelectedKindEnumValues Enumerates the set of values for TargetSelectedKindEnum
func GetTargetSelectedKindEnumValues() []TargetSelectedKindEnum {
	values := make([]TargetSelectedKindEnum, 0)
	for _, v := range mappingTargetSelectedKindEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetSelectedKindEnumStringValues Enumerates the set of values in String for TargetSelectedKindEnum
func GetTargetSelectedKindEnumStringValues() []string {
	return []string{
		"ALL",
		"TARGETTYPES",
		"TARGETIDS",
	}
}

// GetMappingTargetSelectedKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetSelectedKindEnum(val string) (TargetSelectedKindEnum, bool) {
	enum, ok := mappingTargetSelectedKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
