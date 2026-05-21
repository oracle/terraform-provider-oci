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

// OciFssMountTarget An OCI File Storage Service (FSS) Mount Target.
// Check https://docs.oracle.com/en-us/iaas/api/#/en/filestorage/20171215/MountTarget for more details.
type OciFssMountTarget interface {
}

type ocifssmounttarget struct {
	JsonData              []byte
	OciFssMountTargetType string `json:"ociFssMountTargetType"`
}

// UnmarshalJSON unmarshals json
func (m *ocifssmounttarget) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerocifssmounttarget ocifssmounttarget
	s := struct {
		Model Unmarshalerocifssmounttarget
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OciFssMountTargetType = s.Model.OciFssMountTargetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ocifssmounttarget) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OciFssMountTargetType {
	case "OCID":
		mm := OciFssMountTargetId{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for OciFssMountTarget: %s.", m.OciFssMountTargetType)
		return *m, nil
	}
}

func (m ocifssmounttarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ocifssmounttarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciFssMountTargetOciFssMountTargetTypeEnum Enum with underlying type: string
type OciFssMountTargetOciFssMountTargetTypeEnum string

// Set of constants representing the allowable values for OciFssMountTargetOciFssMountTargetTypeEnum
const (
	OciFssMountTargetOciFssMountTargetTypeOcid OciFssMountTargetOciFssMountTargetTypeEnum = "OCID"
)

var mappingOciFssMountTargetOciFssMountTargetTypeEnum = map[string]OciFssMountTargetOciFssMountTargetTypeEnum{
	"OCID": OciFssMountTargetOciFssMountTargetTypeOcid,
}

var mappingOciFssMountTargetOciFssMountTargetTypeEnumLowerCase = map[string]OciFssMountTargetOciFssMountTargetTypeEnum{
	"ocid": OciFssMountTargetOciFssMountTargetTypeOcid,
}

// GetOciFssMountTargetOciFssMountTargetTypeEnumValues Enumerates the set of values for OciFssMountTargetOciFssMountTargetTypeEnum
func GetOciFssMountTargetOciFssMountTargetTypeEnumValues() []OciFssMountTargetOciFssMountTargetTypeEnum {
	values := make([]OciFssMountTargetOciFssMountTargetTypeEnum, 0)
	for _, v := range mappingOciFssMountTargetOciFssMountTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciFssMountTargetOciFssMountTargetTypeEnumStringValues Enumerates the set of values in String for OciFssMountTargetOciFssMountTargetTypeEnum
func GetOciFssMountTargetOciFssMountTargetTypeEnumStringValues() []string {
	return []string{
		"OCID",
	}
}

// GetMappingOciFssMountTargetOciFssMountTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciFssMountTargetOciFssMountTargetTypeEnum(val string) (OciFssMountTargetOciFssMountTargetTypeEnum, bool) {
	enum, ok := mappingOciFssMountTargetOciFssMountTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
