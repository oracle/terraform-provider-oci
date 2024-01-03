// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceTypeMetadataDetails The metadata details for resource type.
type ResourceTypeMetadataDetails interface {
}

type resourcetypemetadatadetails struct {
	JsonData []byte
	Format   string `json:"format"`
}

// UnmarshalJSON unmarshals json
func (m *resourcetypemetadatadetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresourcetypemetadatadetails resourcetypemetadatadetails
	s := struct {
		Model Unmarshalerresourcetypemetadatadetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Format = s.Model.Format

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resourcetypemetadatadetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Format {
	case "SYSTEM_FORMAT":
		mm := SystemFormatResourceTypeMetadataDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResourceTypeMetadataDetails: %s.", m.Format)
		return *m, nil
	}
}

func (m resourcetypemetadatadetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resourcetypemetadatadetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceTypeMetadataDetailsFormatEnum Enum with underlying type: string
type ResourceTypeMetadataDetailsFormatEnum string

// Set of constants representing the allowable values for ResourceTypeMetadataDetailsFormatEnum
const (
	ResourceTypeMetadataDetailsFormatSystemFormat ResourceTypeMetadataDetailsFormatEnum = "SYSTEM_FORMAT"
)

var mappingResourceTypeMetadataDetailsFormatEnum = map[string]ResourceTypeMetadataDetailsFormatEnum{
	"SYSTEM_FORMAT": ResourceTypeMetadataDetailsFormatSystemFormat,
}

var mappingResourceTypeMetadataDetailsFormatEnumLowerCase = map[string]ResourceTypeMetadataDetailsFormatEnum{
	"system_format": ResourceTypeMetadataDetailsFormatSystemFormat,
}

// GetResourceTypeMetadataDetailsFormatEnumValues Enumerates the set of values for ResourceTypeMetadataDetailsFormatEnum
func GetResourceTypeMetadataDetailsFormatEnumValues() []ResourceTypeMetadataDetailsFormatEnum {
	values := make([]ResourceTypeMetadataDetailsFormatEnum, 0)
	for _, v := range mappingResourceTypeMetadataDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeMetadataDetailsFormatEnumStringValues Enumerates the set of values in String for ResourceTypeMetadataDetailsFormatEnum
func GetResourceTypeMetadataDetailsFormatEnumStringValues() []string {
	return []string{
		"SYSTEM_FORMAT",
	}
}

// GetMappingResourceTypeMetadataDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeMetadataDetailsFormatEnum(val string) (ResourceTypeMetadataDetailsFormatEnum, bool) {
	enum, ok := mappingResourceTypeMetadataDetailsFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
