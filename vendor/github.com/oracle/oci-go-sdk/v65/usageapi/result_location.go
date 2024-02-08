// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResultLocation The location where usage or cost CSVs will be uploaded defined by `locationType`,
// which corresponds with type-specific characteristics.
type ResultLocation interface {
}

type resultlocation struct {
	JsonData     []byte
	LocationType string `json:"locationType"`
}

// UnmarshalJSON unmarshals json
func (m *resultlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresultlocation resultlocation
	s := struct {
		Model Unmarshalerresultlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LocationType = s.Model.LocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resultlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LocationType {
	case "OBJECT_STORAGE":
		mm := ObjectStorageLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResultLocation: %s.", m.LocationType)
		return *m, nil
	}
}

func (m resultlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resultlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResultLocationLocationTypeEnum Enum with underlying type: string
type ResultLocationLocationTypeEnum string

// Set of constants representing the allowable values for ResultLocationLocationTypeEnum
const (
	ResultLocationLocationTypeObjectStorage ResultLocationLocationTypeEnum = "OBJECT_STORAGE"
)

var mappingResultLocationLocationTypeEnum = map[string]ResultLocationLocationTypeEnum{
	"OBJECT_STORAGE": ResultLocationLocationTypeObjectStorage,
}

var mappingResultLocationLocationTypeEnumLowerCase = map[string]ResultLocationLocationTypeEnum{
	"object_storage": ResultLocationLocationTypeObjectStorage,
}

// GetResultLocationLocationTypeEnumValues Enumerates the set of values for ResultLocationLocationTypeEnum
func GetResultLocationLocationTypeEnumValues() []ResultLocationLocationTypeEnum {
	values := make([]ResultLocationLocationTypeEnum, 0)
	for _, v := range mappingResultLocationLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResultLocationLocationTypeEnumStringValues Enumerates the set of values in String for ResultLocationLocationTypeEnum
func GetResultLocationLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE",
	}
}

// GetMappingResultLocationLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResultLocationLocationTypeEnum(val string) (ResultLocationLocationTypeEnum, bool) {
	enum, ok := mappingResultLocationLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
