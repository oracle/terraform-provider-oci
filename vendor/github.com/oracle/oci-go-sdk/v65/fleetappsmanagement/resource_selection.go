// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceSelection Resource Selection Type
type ResourceSelection interface {
}

type resourceselection struct {
	JsonData              []byte
	ResourceSelectionType string `json:"resourceSelectionType"`
}

// UnmarshalJSON unmarshals json
func (m *resourceselection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresourceselection resourceselection
	s := struct {
		Model Unmarshalerresourceselection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ResourceSelectionType = s.Model.ResourceSelectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resourceselection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ResourceSelectionType {
	case "DYNAMIC":
		mm := DynamicResourceSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL":
		mm := ManualResourceSelection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ResourceSelection: %s.", m.ResourceSelectionType)
		return *m, nil
	}
}

func (m resourceselection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resourceselection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceSelectionResourceSelectionTypeEnum Enum with underlying type: string
type ResourceSelectionResourceSelectionTypeEnum string

// Set of constants representing the allowable values for ResourceSelectionResourceSelectionTypeEnum
const (
	ResourceSelectionResourceSelectionTypeDynamic ResourceSelectionResourceSelectionTypeEnum = "DYNAMIC"
	ResourceSelectionResourceSelectionTypeManual  ResourceSelectionResourceSelectionTypeEnum = "MANUAL"
)

var mappingResourceSelectionResourceSelectionTypeEnum = map[string]ResourceSelectionResourceSelectionTypeEnum{
	"DYNAMIC": ResourceSelectionResourceSelectionTypeDynamic,
	"MANUAL":  ResourceSelectionResourceSelectionTypeManual,
}

var mappingResourceSelectionResourceSelectionTypeEnumLowerCase = map[string]ResourceSelectionResourceSelectionTypeEnum{
	"dynamic": ResourceSelectionResourceSelectionTypeDynamic,
	"manual":  ResourceSelectionResourceSelectionTypeManual,
}

// GetResourceSelectionResourceSelectionTypeEnumValues Enumerates the set of values for ResourceSelectionResourceSelectionTypeEnum
func GetResourceSelectionResourceSelectionTypeEnumValues() []ResourceSelectionResourceSelectionTypeEnum {
	values := make([]ResourceSelectionResourceSelectionTypeEnum, 0)
	for _, v := range mappingResourceSelectionResourceSelectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceSelectionResourceSelectionTypeEnumStringValues Enumerates the set of values in String for ResourceSelectionResourceSelectionTypeEnum
func GetResourceSelectionResourceSelectionTypeEnumStringValues() []string {
	return []string{
		"DYNAMIC",
		"MANUAL",
	}
}

// GetMappingResourceSelectionResourceSelectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceSelectionResourceSelectionTypeEnum(val string) (ResourceSelectionResourceSelectionTypeEnum, bool) {
	enum, ok := mappingResourceSelectionResourceSelectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
