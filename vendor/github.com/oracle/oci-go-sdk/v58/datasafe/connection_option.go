// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ConnectionOption Types of connection supported by Data Safe.
type ConnectionOption interface {
}

type connectionoption struct {
	JsonData       []byte
	ConnectionType string `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *connectionoption) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectionoption connectionoption
	s := struct {
		Model Unmarshalerconnectionoption
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectionoption) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "PRIVATE_ENDPOINT":
		mm := PrivateEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONPREM_CONNECTOR":
		mm := OnPremiseConnector{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m connectionoption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectionoption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectionOptionConnectionTypeEnum Enum with underlying type: string
type ConnectionOptionConnectionTypeEnum string

// Set of constants representing the allowable values for ConnectionOptionConnectionTypeEnum
const (
	ConnectionOptionConnectionTypePrivateEndpoint ConnectionOptionConnectionTypeEnum = "PRIVATE_ENDPOINT"
	ConnectionOptionConnectionTypeOnpremConnector ConnectionOptionConnectionTypeEnum = "ONPREM_CONNECTOR"
)

var mappingConnectionOptionConnectionTypeEnum = map[string]ConnectionOptionConnectionTypeEnum{
	"PRIVATE_ENDPOINT": ConnectionOptionConnectionTypePrivateEndpoint,
	"ONPREM_CONNECTOR": ConnectionOptionConnectionTypeOnpremConnector,
}

// GetConnectionOptionConnectionTypeEnumValues Enumerates the set of values for ConnectionOptionConnectionTypeEnum
func GetConnectionOptionConnectionTypeEnumValues() []ConnectionOptionConnectionTypeEnum {
	values := make([]ConnectionOptionConnectionTypeEnum, 0)
	for _, v := range mappingConnectionOptionConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectionOptionConnectionTypeEnumStringValues Enumerates the set of values in String for ConnectionOptionConnectionTypeEnum
func GetConnectionOptionConnectionTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT",
		"ONPREM_CONNECTOR",
	}
}

// GetMappingConnectionOptionConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectionOptionConnectionTypeEnum(val string) (ConnectionOptionConnectionTypeEnum, bool) {
	mappingConnectionOptionConnectionTypeEnumIgnoreCase := make(map[string]ConnectionOptionConnectionTypeEnum)
	for k, v := range mappingConnectionOptionConnectionTypeEnum {
		mappingConnectionOptionConnectionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConnectionOptionConnectionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
