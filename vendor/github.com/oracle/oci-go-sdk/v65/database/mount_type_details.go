// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MountTypeDetails Mount type details for backup destination.
type MountTypeDetails interface {
}

type mounttypedetails struct {
	JsonData  []byte
	MountType string `json:"mountType"`
}

// UnmarshalJSON unmarshals json
func (m *mounttypedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermounttypedetails mounttypedetails
	s := struct {
		Model Unmarshalermounttypedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MountType = s.Model.MountType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *mounttypedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MountType {
	case "SELF_MOUNT":
		mm := SelfMountDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTOMATED_MOUNT":
		mm := AutomatedMountDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MountTypeDetails: %s.", m.MountType)
		return *m, nil
	}
}

func (m mounttypedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m mounttypedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MountTypeDetailsMountTypeEnum Enum with underlying type: string
type MountTypeDetailsMountTypeEnum string

// Set of constants representing the allowable values for MountTypeDetailsMountTypeEnum
const (
	MountTypeDetailsMountTypeSelfMount      MountTypeDetailsMountTypeEnum = "SELF_MOUNT"
	MountTypeDetailsMountTypeAutomatedMount MountTypeDetailsMountTypeEnum = "AUTOMATED_MOUNT"
)

var mappingMountTypeDetailsMountTypeEnum = map[string]MountTypeDetailsMountTypeEnum{
	"SELF_MOUNT":      MountTypeDetailsMountTypeSelfMount,
	"AUTOMATED_MOUNT": MountTypeDetailsMountTypeAutomatedMount,
}

var mappingMountTypeDetailsMountTypeEnumLowerCase = map[string]MountTypeDetailsMountTypeEnum{
	"self_mount":      MountTypeDetailsMountTypeSelfMount,
	"automated_mount": MountTypeDetailsMountTypeAutomatedMount,
}

// GetMountTypeDetailsMountTypeEnumValues Enumerates the set of values for MountTypeDetailsMountTypeEnum
func GetMountTypeDetailsMountTypeEnumValues() []MountTypeDetailsMountTypeEnum {
	values := make([]MountTypeDetailsMountTypeEnum, 0)
	for _, v := range mappingMountTypeDetailsMountTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMountTypeDetailsMountTypeEnumStringValues Enumerates the set of values in String for MountTypeDetailsMountTypeEnum
func GetMountTypeDetailsMountTypeEnumStringValues() []string {
	return []string{
		"SELF_MOUNT",
		"AUTOMATED_MOUNT",
	}
}

// GetMappingMountTypeDetailsMountTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMountTypeDetailsMountTypeEnum(val string) (MountTypeDetailsMountTypeEnum, bool) {
	enum, ok := mappingMountTypeDetailsMountTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
