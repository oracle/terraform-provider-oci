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

// TrafficProtocolParameters Defines the traffic protocol parameters for the traffic in a `PathAnalysisResult`.
type TrafficProtocolParameters interface {
}

type trafficprotocolparameters struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *trafficprotocolparameters) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrafficprotocolparameters trafficprotocolparameters
	s := struct {
		Model Unmarshalertrafficprotocolparameters
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *trafficprotocolparameters) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ICMP":
		mm := IcmpTrafficProtocolParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "UDP":
		mm := UdpTrafficProtocolParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCP":
		mm := TcpTrafficProtocolParameters{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TrafficProtocolParameters: %s.", m.Type)
		return *m, nil
	}
}

func (m trafficprotocolparameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m trafficprotocolparameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TrafficProtocolParametersTypeEnum Enum with underlying type: string
type TrafficProtocolParametersTypeEnum string

// Set of constants representing the allowable values for TrafficProtocolParametersTypeEnum
const (
	TrafficProtocolParametersTypeTcp  TrafficProtocolParametersTypeEnum = "TCP"
	TrafficProtocolParametersTypeUdp  TrafficProtocolParametersTypeEnum = "UDP"
	TrafficProtocolParametersTypeIcmp TrafficProtocolParametersTypeEnum = "ICMP"
)

var mappingTrafficProtocolParametersTypeEnum = map[string]TrafficProtocolParametersTypeEnum{
	"TCP":  TrafficProtocolParametersTypeTcp,
	"UDP":  TrafficProtocolParametersTypeUdp,
	"ICMP": TrafficProtocolParametersTypeIcmp,
}

var mappingTrafficProtocolParametersTypeEnumLowerCase = map[string]TrafficProtocolParametersTypeEnum{
	"tcp":  TrafficProtocolParametersTypeTcp,
	"udp":  TrafficProtocolParametersTypeUdp,
	"icmp": TrafficProtocolParametersTypeIcmp,
}

// GetTrafficProtocolParametersTypeEnumValues Enumerates the set of values for TrafficProtocolParametersTypeEnum
func GetTrafficProtocolParametersTypeEnumValues() []TrafficProtocolParametersTypeEnum {
	values := make([]TrafficProtocolParametersTypeEnum, 0)
	for _, v := range mappingTrafficProtocolParametersTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTrafficProtocolParametersTypeEnumStringValues Enumerates the set of values in String for TrafficProtocolParametersTypeEnum
func GetTrafficProtocolParametersTypeEnumStringValues() []string {
	return []string{
		"TCP",
		"UDP",
		"ICMP",
	}
}

// GetMappingTrafficProtocolParametersTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTrafficProtocolParametersTypeEnum(val string) (TrafficProtocolParametersTypeEnum, bool) {
	enum, ok := mappingTrafficProtocolParametersTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
