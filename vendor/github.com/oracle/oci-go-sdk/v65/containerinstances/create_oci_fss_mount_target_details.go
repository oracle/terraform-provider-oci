// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciFssMountTargetDetails An OCI File Storage Service (FSS) Mount Target.
// Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/MountTarget for more details.
type CreateOciFssMountTargetDetails interface {
}

type createocifssmounttargetdetails struct {
	JsonData              []byte
	OciFssMountTargetType string `json:"ociFssMountTargetType"`
}

// UnmarshalJSON unmarshals json
func (m *createocifssmounttargetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateocifssmounttargetdetails createocifssmounttargetdetails
	s := struct {
		Model Unmarshalercreateocifssmounttargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OciFssMountTargetType = s.Model.OciFssMountTargetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createocifssmounttargetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OciFssMountTargetType {
	case "OCID":
		mm := CreateOciFssMountTargetIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateOciFssMountTargetDetails: %s.", m.OciFssMountTargetType)
		return *m, nil
	}
}

func (m createocifssmounttargetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createocifssmounttargetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum Enum with underlying type: string
type CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum string

// Set of constants representing the allowable values for CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum
const (
	CreateOciFssMountTargetDetailsOciFssMountTargetTypeOcid CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum = "OCID"
)

var mappingCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum = map[string]CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum{
	"OCID": CreateOciFssMountTargetDetailsOciFssMountTargetTypeOcid,
}

var mappingCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnumLowerCase = map[string]CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum{
	"ocid": CreateOciFssMountTargetDetailsOciFssMountTargetTypeOcid,
}

// GetCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnumValues Enumerates the set of values for CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum
func GetCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnumValues() []CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum {
	values := make([]CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum, 0)
	for _, v := range mappingCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnumStringValues Enumerates the set of values in String for CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum
func GetCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnumStringValues() []string {
	return []string{
		"OCID",
	}
}

// GetMappingCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum(val string) (CreateOciFssMountTargetDetailsOciFssMountTargetTypeEnum, bool) {
	enum, ok := mappingCreateOciFssMountTargetDetailsOciFssMountTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
