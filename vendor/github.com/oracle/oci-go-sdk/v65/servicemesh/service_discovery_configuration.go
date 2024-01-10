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

// ServiceDiscoveryConfiguration Service Discovery configuration for virtual deployments.
type ServiceDiscoveryConfiguration interface {
}

type servicediscoveryconfiguration struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *servicediscoveryconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerservicediscoveryconfiguration servicediscoveryconfiguration
	s := struct {
		Model Unmarshalerservicediscoveryconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *servicediscoveryconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DNS":
		mm := DnsServiceDiscoveryConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DISABLED":
		mm := DisabledServiceDiscoveryConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ServiceDiscoveryConfiguration: %s.", m.Type)
		return *m, nil
	}
}

func (m servicediscoveryconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m servicediscoveryconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceDiscoveryConfigurationTypeEnum Enum with underlying type: string
type ServiceDiscoveryConfigurationTypeEnum string

// Set of constants representing the allowable values for ServiceDiscoveryConfigurationTypeEnum
const (
	ServiceDiscoveryConfigurationTypeDns      ServiceDiscoveryConfigurationTypeEnum = "DNS"
	ServiceDiscoveryConfigurationTypeDisabled ServiceDiscoveryConfigurationTypeEnum = "DISABLED"
)

var mappingServiceDiscoveryConfigurationTypeEnum = map[string]ServiceDiscoveryConfigurationTypeEnum{
	"DNS":      ServiceDiscoveryConfigurationTypeDns,
	"DISABLED": ServiceDiscoveryConfigurationTypeDisabled,
}

var mappingServiceDiscoveryConfigurationTypeEnumLowerCase = map[string]ServiceDiscoveryConfigurationTypeEnum{
	"dns":      ServiceDiscoveryConfigurationTypeDns,
	"disabled": ServiceDiscoveryConfigurationTypeDisabled,
}

// GetServiceDiscoveryConfigurationTypeEnumValues Enumerates the set of values for ServiceDiscoveryConfigurationTypeEnum
func GetServiceDiscoveryConfigurationTypeEnumValues() []ServiceDiscoveryConfigurationTypeEnum {
	values := make([]ServiceDiscoveryConfigurationTypeEnum, 0)
	for _, v := range mappingServiceDiscoveryConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceDiscoveryConfigurationTypeEnumStringValues Enumerates the set of values in String for ServiceDiscoveryConfigurationTypeEnum
func GetServiceDiscoveryConfigurationTypeEnumStringValues() []string {
	return []string{
		"DNS",
		"DISABLED",
	}
}

// GetMappingServiceDiscoveryConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceDiscoveryConfigurationTypeEnum(val string) (ServiceDiscoveryConfigurationTypeEnum, bool) {
	enum, ok := mappingServiceDiscoveryConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
