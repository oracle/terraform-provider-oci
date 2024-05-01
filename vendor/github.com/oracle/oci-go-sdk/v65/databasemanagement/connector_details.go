// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectorDetails The connector details required to connect to an Oracle cloud database.
type ConnectorDetails interface {
}

type connectordetails struct {
	JsonData      []byte
	ConnectorType string `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *connectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectordetails connectordetails
	s := struct {
		Model Unmarshalerconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "EXTERNAL":
		mm := ExternalConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS":
		mm := MacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PE":
		mm := PrivateEndPointConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ConnectorDetails: %s.", m.ConnectorType)
		return *m, nil
	}
}

func (m connectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type ConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for ConnectorDetailsConnectorTypeEnum
const (
	ConnectorDetailsConnectorTypePe       ConnectorDetailsConnectorTypeEnum = "PE"
	ConnectorDetailsConnectorTypeMacs     ConnectorDetailsConnectorTypeEnum = "MACS"
	ConnectorDetailsConnectorTypeExternal ConnectorDetailsConnectorTypeEnum = "EXTERNAL"
)

var mappingConnectorDetailsConnectorTypeEnum = map[string]ConnectorDetailsConnectorTypeEnum{
	"PE":       ConnectorDetailsConnectorTypePe,
	"MACS":     ConnectorDetailsConnectorTypeMacs,
	"EXTERNAL": ConnectorDetailsConnectorTypeExternal,
}

var mappingConnectorDetailsConnectorTypeEnumLowerCase = map[string]ConnectorDetailsConnectorTypeEnum{
	"pe":       ConnectorDetailsConnectorTypePe,
	"macs":     ConnectorDetailsConnectorTypeMacs,
	"external": ConnectorDetailsConnectorTypeExternal,
}

// GetConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for ConnectorDetailsConnectorTypeEnum
func GetConnectorDetailsConnectorTypeEnumValues() []ConnectorDetailsConnectorTypeEnum {
	values := make([]ConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for ConnectorDetailsConnectorTypeEnum
func GetConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"PE",
		"MACS",
		"EXTERNAL",
	}
}

// GetMappingConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectorDetailsConnectorTypeEnum(val string) (ConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
