// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LocationDetails Possible object storage location types
type LocationDetails interface {
}

type locationdetails struct {
	JsonData     []byte
	LocationType string `json:"locationType"`
}

// UnmarshalJSON unmarshals json
func (m *locationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlocationdetails locationdetails
	s := struct {
		Model Unmarshalerlocationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LocationType = s.Model.LocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *locationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LocationType {
	case "OBJECT_LIST":
		mm := ObjectListDataset{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LocationDetails: %s.", m.LocationType)
		return *m, nil
	}
}

func (m locationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m locationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LocationDetailsLocationTypeEnum Enum with underlying type: string
type LocationDetailsLocationTypeEnum string

// Set of constants representing the allowable values for LocationDetailsLocationTypeEnum
const (
	LocationDetailsLocationTypeObjectList LocationDetailsLocationTypeEnum = "OBJECT_LIST"
)

var mappingLocationDetailsLocationTypeEnum = map[string]LocationDetailsLocationTypeEnum{
	"OBJECT_LIST": LocationDetailsLocationTypeObjectList,
}

var mappingLocationDetailsLocationTypeEnumLowerCase = map[string]LocationDetailsLocationTypeEnum{
	"object_list": LocationDetailsLocationTypeObjectList,
}

// GetLocationDetailsLocationTypeEnumValues Enumerates the set of values for LocationDetailsLocationTypeEnum
func GetLocationDetailsLocationTypeEnumValues() []LocationDetailsLocationTypeEnum {
	values := make([]LocationDetailsLocationTypeEnum, 0)
	for _, v := range mappingLocationDetailsLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLocationDetailsLocationTypeEnumStringValues Enumerates the set of values in String for LocationDetailsLocationTypeEnum
func GetLocationDetailsLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_LIST",
	}
}

// GetMappingLocationDetailsLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLocationDetailsLocationTypeEnum(val string) (LocationDetailsLocationTypeEnum, bool) {
	enum, ok := mappingLocationDetailsLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
