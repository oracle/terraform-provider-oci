// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CircuitBreakingThresholdConfiguration Thresholds to fence the application can be configured with below parameters.
type CircuitBreakingThresholdConfiguration interface {
}

type circuitbreakingthresholdconfiguration struct {
	JsonData []byte
	Protocol string `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *circuitbreakingthresholdconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercircuitbreakingthresholdconfiguration circuitbreakingthresholdconfiguration
	s := struct {
		Model Unmarshalercircuitbreakingthresholdconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *circuitbreakingthresholdconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "HTTP":
		mm := HttpCircuitBreakingThresholdConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TCP":
		mm := TcpCircuitBreakingThresholdConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CircuitBreakingThresholdConfiguration: %s.", m.Protocol)
		return *m, nil
	}
}

func (m circuitbreakingthresholdconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m circuitbreakingthresholdconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CircuitBreakingThresholdConfigurationProtocolEnum Enum with underlying type: string
type CircuitBreakingThresholdConfigurationProtocolEnum string

// Set of constants representing the allowable values for CircuitBreakingThresholdConfigurationProtocolEnum
const (
	CircuitBreakingThresholdConfigurationProtocolHttp CircuitBreakingThresholdConfigurationProtocolEnum = "HTTP"
	CircuitBreakingThresholdConfigurationProtocolTcp  CircuitBreakingThresholdConfigurationProtocolEnum = "TCP"
)

var mappingCircuitBreakingThresholdConfigurationProtocolEnum = map[string]CircuitBreakingThresholdConfigurationProtocolEnum{
	"HTTP": CircuitBreakingThresholdConfigurationProtocolHttp,
	"TCP":  CircuitBreakingThresholdConfigurationProtocolTcp,
}

var mappingCircuitBreakingThresholdConfigurationProtocolEnumLowerCase = map[string]CircuitBreakingThresholdConfigurationProtocolEnum{
	"http": CircuitBreakingThresholdConfigurationProtocolHttp,
	"tcp":  CircuitBreakingThresholdConfigurationProtocolTcp,
}

// GetCircuitBreakingThresholdConfigurationProtocolEnumValues Enumerates the set of values for CircuitBreakingThresholdConfigurationProtocolEnum
func GetCircuitBreakingThresholdConfigurationProtocolEnumValues() []CircuitBreakingThresholdConfigurationProtocolEnum {
	values := make([]CircuitBreakingThresholdConfigurationProtocolEnum, 0)
	for _, v := range mappingCircuitBreakingThresholdConfigurationProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetCircuitBreakingThresholdConfigurationProtocolEnumStringValues Enumerates the set of values in String for CircuitBreakingThresholdConfigurationProtocolEnum
func GetCircuitBreakingThresholdConfigurationProtocolEnumStringValues() []string {
	return []string{
		"HTTP",
		"TCP",
	}
}

// GetMappingCircuitBreakingThresholdConfigurationProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCircuitBreakingThresholdConfigurationProtocolEnum(val string) (CircuitBreakingThresholdConfigurationProtocolEnum, bool) {
	enum, ok := mappingCircuitBreakingThresholdConfigurationProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
