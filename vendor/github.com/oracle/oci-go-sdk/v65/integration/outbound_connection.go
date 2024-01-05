// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OutboundConnection Base representation for Outbound Connection (Reverse Connection).
type OutboundConnection interface {
}

type outboundconnection struct {
	JsonData               []byte
	OutboundConnectionType string `json:"outboundConnectionType"`
}

// UnmarshalJSON unmarshals json
func (m *outboundconnection) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroutboundconnection outboundconnection
	s := struct {
		Model Unmarshaleroutboundconnection
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutboundConnectionType = s.Model.OutboundConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *outboundconnection) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutboundConnectionType {
	case "PRIVATE_ENDPOINT":
		mm := PrivateEndpointOutboundConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneOutboundConnection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OutboundConnection: %s.", m.OutboundConnectionType)
		return *m, nil
	}
}

func (m outboundconnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m outboundconnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OutboundConnectionOutboundConnectionTypeEnum Enum with underlying type: string
type OutboundConnectionOutboundConnectionTypeEnum string

// Set of constants representing the allowable values for OutboundConnectionOutboundConnectionTypeEnum
const (
	OutboundConnectionOutboundConnectionTypePrivateEndpoint OutboundConnectionOutboundConnectionTypeEnum = "PRIVATE_ENDPOINT"
	OutboundConnectionOutboundConnectionTypeNone            OutboundConnectionOutboundConnectionTypeEnum = "NONE"
)

var mappingOutboundConnectionOutboundConnectionTypeEnum = map[string]OutboundConnectionOutboundConnectionTypeEnum{
	"PRIVATE_ENDPOINT": OutboundConnectionOutboundConnectionTypePrivateEndpoint,
	"NONE":             OutboundConnectionOutboundConnectionTypeNone,
}

var mappingOutboundConnectionOutboundConnectionTypeEnumLowerCase = map[string]OutboundConnectionOutboundConnectionTypeEnum{
	"private_endpoint": OutboundConnectionOutboundConnectionTypePrivateEndpoint,
	"none":             OutboundConnectionOutboundConnectionTypeNone,
}

// GetOutboundConnectionOutboundConnectionTypeEnumValues Enumerates the set of values for OutboundConnectionOutboundConnectionTypeEnum
func GetOutboundConnectionOutboundConnectionTypeEnumValues() []OutboundConnectionOutboundConnectionTypeEnum {
	values := make([]OutboundConnectionOutboundConnectionTypeEnum, 0)
	for _, v := range mappingOutboundConnectionOutboundConnectionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOutboundConnectionOutboundConnectionTypeEnumStringValues Enumerates the set of values in String for OutboundConnectionOutboundConnectionTypeEnum
func GetOutboundConnectionOutboundConnectionTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT",
		"NONE",
	}
}

// GetMappingOutboundConnectionOutboundConnectionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutboundConnectionOutboundConnectionTypeEnum(val string) (OutboundConnectionOutboundConnectionTypeEnum, bool) {
	enum, ok := mappingOutboundConnectionOutboundConnectionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
