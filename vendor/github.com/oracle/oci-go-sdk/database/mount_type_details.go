// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
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
		return *m, nil
	}
}

func (m mounttypedetails) String() string {
	return common.PointerString(m)
}

// MountTypeDetailsMountTypeEnum Enum with underlying type: string
type MountTypeDetailsMountTypeEnum string

// Set of constants representing the allowable values for MountTypeDetailsMountTypeEnum
const (
	MountTypeDetailsMountTypeSelfMount      MountTypeDetailsMountTypeEnum = "SELF_MOUNT"
	MountTypeDetailsMountTypeAutomatedMount MountTypeDetailsMountTypeEnum = "AUTOMATED_MOUNT"
)

var mappingMountTypeDetailsMountType = map[string]MountTypeDetailsMountTypeEnum{
	"SELF_MOUNT":      MountTypeDetailsMountTypeSelfMount,
	"AUTOMATED_MOUNT": MountTypeDetailsMountTypeAutomatedMount,
}

// GetMountTypeDetailsMountTypeEnumValues Enumerates the set of values for MountTypeDetailsMountTypeEnum
func GetMountTypeDetailsMountTypeEnumValues() []MountTypeDetailsMountTypeEnum {
	values := make([]MountTypeDetailsMountTypeEnum, 0)
	for _, v := range mappingMountTypeDetailsMountType {
		values = append(values, v)
	}
	return values
}
