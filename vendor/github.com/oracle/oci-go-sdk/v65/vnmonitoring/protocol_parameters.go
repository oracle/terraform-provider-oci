// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProtocolParameters Defines the IP protocol parameters for a `PathAnalyzerTest` resource.
type ProtocolParameters interface {
}

type protocolparameters struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *protocolparameters) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerprotocolparameters protocolparameters
	s := struct {
		Model Unmarshalerprotocolparameters
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *protocolparameters) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "UDP":
		mm := UdpProtocolParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCP":
		mm := TcpProtocolParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ICMP":
		mm := IcmpProtocolParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ProtocolParameters: %s.", m.Type)
		return *m, nil
	}
}

func (m protocolparameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m protocolparameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ProtocolParametersTypeEnum Enum with underlying type: string
type ProtocolParametersTypeEnum string

// Set of constants representing the allowable values for ProtocolParametersTypeEnum
const (
	ProtocolParametersTypeTcp  ProtocolParametersTypeEnum = "TCP"
	ProtocolParametersTypeUdp  ProtocolParametersTypeEnum = "UDP"
	ProtocolParametersTypeIcmp ProtocolParametersTypeEnum = "ICMP"
)

var mappingProtocolParametersTypeEnum = map[string]ProtocolParametersTypeEnum{
	"TCP":  ProtocolParametersTypeTcp,
	"UDP":  ProtocolParametersTypeUdp,
	"ICMP": ProtocolParametersTypeIcmp,
}

var mappingProtocolParametersTypeEnumLowerCase = map[string]ProtocolParametersTypeEnum{
	"tcp":  ProtocolParametersTypeTcp,
	"udp":  ProtocolParametersTypeUdp,
	"icmp": ProtocolParametersTypeIcmp,
}

// GetProtocolParametersTypeEnumValues Enumerates the set of values for ProtocolParametersTypeEnum
func GetProtocolParametersTypeEnumValues() []ProtocolParametersTypeEnum {
	values := make([]ProtocolParametersTypeEnum, 0)
	for _, v := range mappingProtocolParametersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProtocolParametersTypeEnumStringValues Enumerates the set of values in String for ProtocolParametersTypeEnum
func GetProtocolParametersTypeEnumStringValues() []string {
	return []string{
		"TCP",
		"UDP",
		"ICMP",
	}
}

// GetMappingProtocolParametersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtocolParametersTypeEnum(val string) (ProtocolParametersTypeEnum, bool) {
	enum, ok := mappingProtocolParametersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
