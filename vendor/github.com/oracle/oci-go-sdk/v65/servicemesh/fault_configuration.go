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

// FaultConfiguration Fault injection configuration which is used to introduce faults in the incoming requests to the listener.
type FaultConfiguration interface {
}

type faultconfiguration struct {
	JsonData []byte
	Protocol string `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *faultconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfaultconfiguration faultconfiguration
	s := struct {
		Model Unmarshalerfaultconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *faultconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "GRPC":
		mm := GrpcFaultConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpFaultConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for FaultConfiguration: %s.", m.Protocol)
		return *m, nil
	}
}

func (m faultconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m faultconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FaultConfigurationProtocolEnum Enum with underlying type: string
type FaultConfigurationProtocolEnum string

// Set of constants representing the allowable values for FaultConfigurationProtocolEnum
const (
	FaultConfigurationProtocolHttp FaultConfigurationProtocolEnum = "HTTP"
	FaultConfigurationProtocolGrpc FaultConfigurationProtocolEnum = "GRPC"
)

var mappingFaultConfigurationProtocolEnum = map[string]FaultConfigurationProtocolEnum{
	"HTTP": FaultConfigurationProtocolHttp,
	"GRPC": FaultConfigurationProtocolGrpc,
}

var mappingFaultConfigurationProtocolEnumLowerCase = map[string]FaultConfigurationProtocolEnum{
	"http": FaultConfigurationProtocolHttp,
	"grpc": FaultConfigurationProtocolGrpc,
}

// GetFaultConfigurationProtocolEnumValues Enumerates the set of values for FaultConfigurationProtocolEnum
func GetFaultConfigurationProtocolEnumValues() []FaultConfigurationProtocolEnum {
	values := make([]FaultConfigurationProtocolEnum, 0)
	for _, v := range mappingFaultConfigurationProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetFaultConfigurationProtocolEnumStringValues Enumerates the set of values in String for FaultConfigurationProtocolEnum
func GetFaultConfigurationProtocolEnumStringValues() []string {
	return []string{
		"HTTP",
		"GRPC",
	}
}

// GetMappingFaultConfigurationProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFaultConfigurationProtocolEnum(val string) (FaultConfigurationProtocolEnum, bool) {
	enum, ok := mappingFaultConfigurationProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
